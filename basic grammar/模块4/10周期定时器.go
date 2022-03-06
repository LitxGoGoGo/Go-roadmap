package main

import (
	"fmt"
	"time"
)

func main1001() {
	//time.Ticker{}
	ticker := time.NewTicker(time.Second * 2)

	go func() {
		i := 0
		for {
			<-ticker.C
			i++
			fmt.Println("定时器触发", i, "次")
			if i == 5 {
				ticker.Stop()
				break
			}
		}
	}()

	time.Sleep(time.Second * 15)
}
