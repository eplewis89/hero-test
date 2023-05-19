package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strateos-test/state"
	"syscall"
	"time"
)

// variable to keep track of the current running state
var (
	currentState = state.NewGameState()
)

// func to run the program indefinitely
func runProgram() {
	// start the observer
	currentState.WaitGroup.Add(1)
	go func() {
		defer currentState.WaitGroup.Done()
		currentState.RunObserver()
	}()

	// start random intervalling of enemy spawn
	currentState.WaitGroup.Add(1)
	go func() {
		defer currentState.WaitGroup.Done()
		currentState.RunArachnidSense()
	}()

	// start hero service
	currentState.WaitGroup.Add(1)
	go func() {
		defer currentState.WaitGroup.Done()
		currentState.RunProtectServers()
	}()

	currentState.WaitGroup.Wait()

	fmt.Println("Total enemies defeated", currentState.GetTotalDefeated())
}

func main() {
	// seed random with pseudorandomness
	rand.Seed(time.Now().UnixNano())

	// graceful exit of main thread
	var stop = make(chan os.Signal, 1)

	// runs program indefinitely, until ctrl+c command is given
	runProgram()

	// notify ctrl-c
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// block main thread until stop is called
	<-stop
}
