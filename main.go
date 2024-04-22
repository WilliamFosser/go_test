package main

import (
	"context"
	"fmt"

	"go.einride.tech/can/pkg/socketcan"
)

func main() {
	// Error handling omitted to keep example simple
	conn, _ := socketcan.DialContext(context.Background(), "vcan", "vcan0")

	recv := socketcan.NewReceiver(conn)
	for recv.Receive() {
		frame := recv.Frame()
		fmt.Println(frame.String())
	}
}
