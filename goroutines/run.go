package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)

func Go() {
	for i := 1; i < 1000; i++ {
		go concurrent(i)
	}	
	time.Sleep(2 * time.Second)
}

func concurrent(i int) {
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Printf("Goroutine %d finished after %d ms.\n", i, n)
}

