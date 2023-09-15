package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func timer(d time.Duration) {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(d)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Printf("\r%s", t.Format("15:04:05"))
		}
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	work := 25 * time.Minute
	rest := 5 * time.Minute
	longRest := 15 * time.Minute

	gomos := 0
	for {
		fmt.Println("Work!")
		timer(work)
		gomos++

		if gomos%4 == 0 {
			fmt.Println("Long rest!")
			timer(longRest)
		} else {
			fmt.Println("Rest!")
			timer(rest)
		}
	}

}
