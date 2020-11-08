package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// CloserHandler handles os interrupt & prints ending info
func CloserHandler() {
	start := time.Now()
	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rEnding load testing.")
		// outputLoadInto()
		elapsed := time.Since(start)
		fmt.Printf("Total Time: %s", elapsed)
		os.Exit(0)
	}()
}

// func outputLoadInto() {
// 	fmt.Println("Num Threads:", numWorker)
// }
