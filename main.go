package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token string

func init() {
	flag.StringVar(&Token, "t", os.Getenv("TOKEN"), "Bpt Token")
	flag.Parse()
}

func main() {
	bot, err := discordgo.New("my_token")
	if err != nil {
		fmt.Println("Bot login errored out: ", err)
		return
	}

	err = bot.Open()
	if err != nil {
		fmt.Println("It errored out again: ", err)
		return
	}

	fmt.Println("Bot is now running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = bot.Close()
	if err != nil {
		fmt.Println("Something failed when closing the bot: ", err)
	}
}
