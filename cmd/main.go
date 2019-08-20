package main

import (
	"flag"
	"fmt"

	"github.com/maakinoh/discordBot/app"
)

func main() {
	fmt.Println("Wellcome to this Discord Bot")

	token := flag.String("DISCORD-TOKEN", "", "discord Token")
	botDir := flag.String("BOT-DIR", "", "Dir of the bot")

	flag.Parse()
	//utils.GetConfigurationInstance()
	app.StartBot(*token, *botDir)
}
