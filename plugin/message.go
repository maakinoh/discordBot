package plugin

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Message struct {
	// The actual message
	Content string

	// The author of the message
	Author User

	// The Guild (discord server) the message come frome
	Guild Guild

	plugin         Plugin
	discordMessage *discordgo.MessageCreate
}

func NewMessge(plug Plugin, discordmsg *discordgo.MessageCreate) Message {
	return Message{Content: discordmsg.Content, plugin: plug, discordMessage: discordmsg, Author: NewUser(plug, discordmsg.Author)}
}

// Answer can be called from a JS Pluin script file to answer to a message.
// This function takes one string as parameter which will be send back to the
// same channel or to the same user where the original message came frome.
func (msg Message) Answer(content string) {
	s := msg.plugin.discordSession
	fmt.Println(msg.discordMessage.ChannelID)
	s.ChannelMessageSend(msg.discordMessage.ChannelID, content)
}
