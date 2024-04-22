package main

import (
	"context"
	"fmt"

	"go.einride.tech/can/pkg/socketcan"
)

func main() {
	conn, err := socketcan.DialContext(context.Background(), "vcan", "vcan0")
	if err != nil {
		fmt.Println("Failed to dial socketcan:", err)
		return
	}

	recv := socketcan.NewReceiver(conn)
	for recv.Receive() {
		frame := recv.Frame()
		fmt.Println(frame.String())
	}
}
