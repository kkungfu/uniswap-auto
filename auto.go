package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hirokimoto/uniswap-auto/services"
	"github.com/hirokimoto/uniswap-auto/utils"
)

func ethQuery(target chan string) {
	for {
		utils.ETHQuery(target)
		time.Sleep(time.Second * 2)
	}
}

func xiQuery(target chan string) {
	for {
		utils.XIQuery(target)
		time.Sleep(time.Second * 3)
	}
}

func tradesQuery(target chan string) {
	for {
		utils.TradesQuery(target)
		time.Sleep(time.Second * 3)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go ethQuery(c1)
	go xiQuery(c2)
	go tradesQuery(c3)

	var eth utils.Crypto
	var xi utils.Tokens
	var swaps utils.Swaps

	go func() {
		for {
			select {
			case msg1 := <-c1:
				json.Unmarshal([]byte(msg1), &eth)
				services.Price(eth, xi)
			case msg2 := <-c2:
				json.Unmarshal([]byte(msg2), &xi)
				services.Price(eth, xi)
			case msg3 := <-c3:
				json.Unmarshal([]byte(msg3), &swaps)
				min, max := services.MinAndMax(swaps)
				fmt.Println("Min: ", min)
				fmt.Println("Max: ", max)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
