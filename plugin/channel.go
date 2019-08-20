package plugin

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Channel struct {
	Name string

	plugin Plugin

	channel *discordgo.Channel
}

func NewChannel(plugin Plugin, channel *discordgo.Channel) Channel {
	return Channel{plugin: plugin, channel: channel, Name: channel.Name}
}

func (channel Channel) GetAllMenbers() {
	ret := make([]*User, 0)
	for _, i := range channel.channel.Recipients {
		user := NewUser(channel.plugin, i)
		ret = append(ret, &user)
	}
}

func (channel Channel) IsNSFW() bool {
	return channel.channel.NSFW
}

// SendMassage can be called from JS to send a new message to the channel
func (channel Channel) SendMessage(msg string) {
	_, err := channel.plugin.discordSession.ChannelMessageSend(channel.channel.ID, msg)
	if err != nil {
		fmt.Print(err.Error())
	}
}
