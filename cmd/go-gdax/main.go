package main

import (
	"os"
	"os/signal"
	"syscall"

	gdax "github.com/banaio/go-gdax"
	"github.com/banaio/go-gdax/websocket"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		DisableSorting:  false,
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	passphrase := os.Getenv("COINBASE_PASSPHRASE")
	isSandbox := false

	client := gdax.NewClient(secret, key, passphrase)

	productIds := []string{
		"BTC-EUR",
		// "BTC-USD",
	}
	ws := websocket.New(client, productIds, isSandbox)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	exitChan := make(chan int)
	go func() {
		for {
			<-signalChan
			exitChan <- 1

			// s := <-signalChan
			// switch s {
			// // kill -SIGHUP XXXX
			// case syscall.SIGHUP:
			// 	fmt.Println("hungup")

			// // kill -SIGINT XXXX or Ctrl+c
			// case syscall.SIGINT:
			// 	fmt.Println("Warikomi")

			// // kill -SIGTERM XXXX
			// case syscall.SIGTERM:
			// 	fmt.Println("force stop")
			// 	exit_chan <- 0

			// // kill -SIGQUIT XXXX
			// case syscall.SIGQUIT:
			// 	fmt.Println("stop and core dump")
			// 	exit_chan <- 0

			// default:
			// 	fmt.Println("Unknown signal.")
			// 	exit_chan <- 1
			// }
		}
	}()

	go func() {
		// ws.Run(secret, key, passphrase)
		ws.Run("", "", "")
	}()

	code := <-exitChan
	os.Exit(code)
}
