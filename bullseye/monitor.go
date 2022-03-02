package bullseye

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Monitor struct {
	ApiKey 				string
	DecryptionString 	string
	Scopes 				[]string
	Verbose 			bool
	NotificationHandler func(product BaseProduct)

	connection      	*websocket.Conn
	isPinging 			bool
}

func (monitor *Monitor) Connect() error {
	ws := url.URL{Scheme: "ws", Host: "api.bullseye.pw", Path: "/v1/ws/"}
	headers := http.Header{
		"Authorization": {monitor.ApiKey},
		"scopes":	 	 {fmt.Sprintf("[%v]", strings.Join(monitor.Scopes[:], ","))},
	}
	connection, _, err := websocket.DefaultDialer.Dial(ws.String(), headers)
	if err != nil {
		if monitor.Verbose {
			log.Println("failed to connect to monitor:", err)
		}

		return err
	}
	monitor.connection = connection

	if monitor.Verbose {
		log.Println("connected to bullseye monitor")
	}

	go monitor.handleMessages()
	if !monitor.isPinging {
		go monitor.startPingInterval()
	}

	return nil
}

func (monitor *Monitor) startPingInterval() {
	monitor.isPinging = true
	for {
		err := monitor.ping()
		if err != nil {
			if monitor.Verbose {
				log.Println("failed to ping monitor:", err)
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func (monitor *Monitor) ping() error {
	err := monitor.connection.WriteMessage(websocket.TextMessage, []byte("PING_MONITOR"))
	if err != nil {
		return err
	}

	return nil
}

func (monitor *Monitor) reconnect() {
	if monitor.Verbose {
		log.Println("reconnecting to monitor")
	}
	if err := monitor.Connect(); err != nil {
		if monitor.Verbose {
			log.Println("failed to reconnect to monitor:", err)
		}
		time.Sleep(5 * time.Second)
		monitor.reconnect()
	}
}

func (monitor *Monitor) decryptMessage(message string) string {
	decryptedBytes, _ := base64.StdEncoding.DecodeString(message)
	decrypted := string(decryptedBytes)

	output := make([]byte, len(decrypted))
	for i := 0; i < len(decrypted); i++ {
		output[i] = decrypted[i] ^ monitor.DecryptionString[i % len(monitor.DecryptionString)]
	}

	return string(output)
}

func (monitor *Monitor) parseProduct(message string) BaseProduct {
	var product BaseProduct
	err := json.Unmarshal([]byte(message), &product)
	if err != nil {
		if monitor.Verbose {
			log.Println("failed to parse message:", err)
		}

		return BaseProduct{}
	}

	return product
}

func (monitor *Monitor) handleMessages() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := monitor.connection.ReadMessage()
			if err != nil {
				if monitor.Verbose {
					log.Println("failed to read message from monitor:", err)
				}
				monitor.reconnect()
				return
			}

			parsedProduct := monitor.parseProduct(monitor.decryptMessage(string(message)))
			if parsedProduct.StoreURL == "" {
				continue
			}

			go monitor.NotificationHandler(parsedProduct)
		}
	}()
}