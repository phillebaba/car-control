package main

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatal("Could not initiate rpio")
	}

	defer rpio.Close()

	in1 := rpio.Pin(23)
	in1.Output()
	in2 := rpio.Pin(24)
	in2.Output()
	//ena := rpio.Pin(25)

	in1.PullUp()
}
