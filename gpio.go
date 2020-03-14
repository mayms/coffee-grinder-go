package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

func initGpio() {
	state, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Using drivers:\n")
	for _, driver := range state.Loaded {
		fmt.Printf("- %s\n", driver)
	}
	fmt.Printf("Drivers skipped:\n")
	for _, failure := range state.Skipped {
		fmt.Printf("- %s: %s\n", failure.D, failure.Err)
	}
	fmt.Printf("Drivers failed to load:\n")
	for _, failure := range state.Failed {
		fmt.Printf("- %s: %v\n", failure.D, failure.Err)
	}
}

func startCounter() {
	counter := make(chan int)
	go func() {
		button := rpi.P1_5
		if err := button.In(gpio.PullUp, gpio.FallingEdge); err != nil {
			log.Fatal(err)
		}

		i := 0
		for {
			button.WaitForEdge(-1)
			button.WaitForEdge(100 * time.Millisecond)
			button.WaitForEdge(100 * time.Millisecond)
			button.WaitForEdge(100 * time.Millisecond)
			i = i + 1
			counter <- i
		}
	}()

	led := gpio.Low
	setLed(led)

	var exit bool
	for !exit {
		select {
		case i := <-counter:
			exit = i > 100
			led = !led
			setLed(led)
			fmt.Println("Counter: ", i)
		}
	}
}

func setLed(level gpio.Level) {
	if err := rpi.P1_3.Out(level); err != nil {
		log.Fatal(err)
	}
}
