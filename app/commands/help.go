package commands

import (
	"github.com/maakinoh/discordBot/models"
)

// NewHelpCommand create a new instance of a HelpCommand
func NewHelpCommand() HelpCommand {
	return HelpCommand{}
}

type HelpCommand struct {
	models.Command
}
