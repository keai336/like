package session

import (
	"sync"
	"testing"
)

func TestSeson(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	c := make(chan string)
	go A(c, &wait)
	c <- "1"
	wait.Wait()
	c <- "q"
	c <- "quit"

}
