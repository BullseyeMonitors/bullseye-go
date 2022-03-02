package _go

import (
	"bullsye-go/bullseye"
	"fmt"
	"testing"
)

func TestMonitor(t *testing.T) {
	monitor := bullseye.Monitor{
		ApiKey:           "KEY",
		DecryptionString: "DECRYPTION_KEY",
		Scopes:           []string{"amazon"},
		Verbose:          true,
		NotificationHandler: NotificationHandler,
	}
	err := monitor.Connect()
	if err != nil {
		return
	}
	for {

	}
}

func NotificationHandler(product bullseye.BaseProduct) {
	fmt.Println(product)
}