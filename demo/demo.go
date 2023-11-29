package main

import (
	"fmt"
	"time"
)

func main() {
	//var age byte
	//fmt.Scanf("%d", &age)
	var age interface{}
	t, e := age.(string)
	fmt.Println(t)
	fmt.Println(e)

	//==============
	//done := make(chan struct{})
	//go func() {
	//	defer close(done)
	//	time.Sleep(time.Second)
	//}()
	//<-done
	//fmt.Println("go done")
	//==============
	done := make(chan struct{})
	timer := time.NewTimer(1 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println(time.Now())
				timer.Reset(1 * time.Second)
			case <-done:
				return
			}
		}
	}()
	<-time.After(5*time.Second + time.Millisecond*100)
	done <- struct{}{}
}
