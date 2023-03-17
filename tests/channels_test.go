package tests

import (
	"fmt"
	"testing"
	"time"
)

func acceptChannel(ch <-chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("finished work")
			return
		default:
			fmt.Println("do work")
			time.Sleep(time.Second)
		}
	}
}

func TestChannels(t *testing.T) {
	ch := make(chan bool, 1)
	go acceptChannel(ch)
	time.Sleep(time.Second * 5)
	close(ch)
	time.Sleep(time.Second)
}
