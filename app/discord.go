package app

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/maakinoh/discordBot/utils"

	"github.com/bwmarrin/discordgo"
)

var botInstance discordgo.Session

func StartBot(token, botDir string) {

	if token == "" {
		fmt.Println("No Discord token set")
		return
	}

	botInstance, err := discordgo.New("Bot " + token)

	if err != nil {

	}

	botInstance.AddHandler(handle)

	err = botInstance.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	botInstance.Close()
}

func handle(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	cfg := utils.GetConfigurationInstance()

	// Check if the sended message starts with the configurated command prefix
	if strings.HasPrefix(m.Content, cfg.CommandPrefix) {
		trimmed := strings.TrimLeft(m.Content, cfg.CommandPrefix)
		for _, p := range cfg.Plugins {
			for _, c := range p.BotCommands {
				if c == trimmed {
					p.InitPlugin(s)
					p.ExecuteCommand(m)
				}
			}
		}

	} else {
		fmt.Println("Not a command")
	}
}
