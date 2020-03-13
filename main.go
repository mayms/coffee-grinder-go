package main

import (
	"fmt"
	"log"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
	"time"
)

func main() {
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

	if err := rpi.P1_33.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	var exit bool
	for !exit {
		select {
		case lala := <-done:
			exit = lala
			fmt.Println("Done!")
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
