package test

import (
	"testing"
	"net"
	"github.com/fifman/surlane/src/surlane"
	"sync"
	"context"
)

var (
	address = "45.78.12.43:1190"
	//testAddress = "45.78.12.43:1191"
	testAddress = "127.0.0.1:1191"
	background = context.Background()
)

func TestDial(t *testing.T) {
	multiDial(1000, dial)
}

func multiDial(n int, dialFunc func(*surlane.LocalContext)bool) {
	ctx := surlane.NewContext(&surlane.RootContext, "test dial")
	wg := sync.WaitGroup{}
	wg.Add(n)
	errCount := 0
	for i:=0; i<n; i++ {
		go func() {
			if dialFunc(ctx) {
				errCount ++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	println(errCount)
}

func dial(ctx *surlane.LocalContext) bool {
	conn, err := (&net.Dialer{}).DialContext(background, "tcp", address)
	if err != nil {
		ctx.LogError(err, "dial error!")
		return true
	}
	conn.Close()
	return false
}

func TestDemoDial(t *testing.T) {
	multiDial(1000, dialDemo)
}

func dialDemo(ctx *surlane.LocalContext) bool {
	conn, err := (&net.Dialer{}).DialContext(background, "tcp", testAddress)
	if err != nil {
		ctx.LogError(err, "dial error!")
		return true
	}
	defer conn.Close()
	_, err = conn.Write([]byte{1,2,3,4})
	/*
	if err != nil {
		ctx.LogError(err, "write error!")
		return true
	}
	buffer := make([]byte, 10)
	n, err := conn.Read(buffer)
	if n > 0 {
		fmt.Println("bytes:", buffer[:n])
	}
	if err != nil && err != io.EOF {
		ctx.LogError(err, "read error!")
		return true
	}
	*/
	return false
}

