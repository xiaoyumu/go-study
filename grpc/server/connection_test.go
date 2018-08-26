package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConncurrency(t *testing.T) {

	var pool chan int

	pool = make(chan int, 10)
	defer func() { close(pool) }()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop() // Make sure the ticker get stopped when this function finish

	start := time.Now()

	//pool <- 88

	var actual int
	select {
	case x := <-pool:
		actual = x
	case <-ticker.C:
		elapsed := time.Now().Sub(start)
		fmt.Printf("Timed out after %s", elapsed.String())
		actual = -1
	}

	assert.Equal(t, 88, actual)
}
