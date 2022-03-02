
# Bullseye Go
a go library for connecting to the bullseye monitor web socket.

# Installation
```
go get github.com/BullseyeMonitors/bullseye-go
```

# Usage
```go
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
```
