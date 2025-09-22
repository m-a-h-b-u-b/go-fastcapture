package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	packets := 0
	// TODO: wire capture.Backend and count packets
	fmt.Printf(\"Processed %d packets in %v\\n\", packets, time.Since(start))
}