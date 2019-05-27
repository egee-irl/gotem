package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

//func start() discordgo.Session{
//	bot, err := discordgo.New(os.Getenv("TOKEN"))
//	if err != nil {
//		fmt.Println("Bot login errored out: ", err)
//		return
//	}
//	return bot
//}

func test() {
	fmt.Println("Bot is now running!")
}

func main() {
	bot, err := discordgo.New(os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("Bot login errored out: ", err)
		return
	}

	err = bot.Open()
	if err != nil {
		fmt.Println("It errored out again: ", err)
		return
	}

	test()

	var sc = make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
	if err != nil {
		fmt.Println("Something failed when closing the bot: ", err)
	}
}
