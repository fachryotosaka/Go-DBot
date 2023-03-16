package main

import (
	"fmt"

	"github.com/fachryotosoka/Go-DBot/bot"
	"github.com/fachryotosoka/Go-DBot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.StartBot()
	<-make(chan struct{})
	return
}
