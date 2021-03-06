package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	bot, err := discordgo.New(os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("Bot login errored out: ", err)
		return
	}

	bot.Open()

	fmt.Println("Bot is now running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
	if err != nil {
		fmt.Println("Something failed when closing the bot: ", err)
	}
}
