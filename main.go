package main

import (
	"fmt"
)

func main() {
	err := config.Readconfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
