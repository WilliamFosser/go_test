package main

import (
	"context"
	"fmt"

	"go.einride.tech/can/pkg/candevice"
	"go.einride.tech/can/pkg/socketcan"
)

//Fungerende fra dette tidspunkt.

func main() {
	// Set up CAN interface
	d, err := candevice.New("vcan0")
	if err != nil {
		fmt.Println("Failed to create CAN device:", err)
		return
	}

	/*
		err = d.SetBitrate(500000)
		if err != nil {
			fmt.Println("Failed to set bitrate:", err)
			return
		}

		err = d.SetUp()
		if err != nil {
			fmt.Println("Failed to set up CAN device:", err)
			return
		}
	*/

	defer d.SetDown()

	// Receive CAN frames
	conn, err := socketcan.DialContext(context.Background(), "can", "vcan0")
	if err != nil {
		fmt.Println("Failed to dial socketcan:", err)
		return
	}

	recv := socketcan.NewReceiver(conn)
	var count int = 0
	for recv.Receive() {
		frame := recv.Frame()
		fmt.Println(frame.String())
		count++
		fmt.Println(count)
	}

}
