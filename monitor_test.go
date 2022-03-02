package _go

import (
	"fmt"
	"github.com/BullseyeMonitors/bullseye-go/monitor"
	"testing"
)

func TestMonitor(t *testing.T) {
	bullseye := monitor.Monitor{
		ApiKey:           "KEY",
		DecryptionString: "DECRYPTION_KEY",
		Scopes:           []string{"amazon"},
		Verbose:          true,
		NotificationHandler: NotificationHandler,
	}
	err := bullseye.Connect()
	if err != nil {
		return
	}
	for {

	}
}

func NotificationHandler(product monitor.BaseProduct) {
	fmt.Println(product)
}