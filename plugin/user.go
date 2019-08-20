package plugin

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type User struct {
	Username string
	Bot      bool
	Verified bool

	dcUser *discordgo.User
	plugin Plugin
}

func NewUser(plugin Plugin, user *discordgo.User) User {
	return User{dcUser: user, plugin: plugin, Username: user.Username}
}

// SendPrivateMessage can be called from JS to send a private message to the user
func (user User) SendPrivateMessage(msg string) {
	ch, err := user.plugin.discordSession.UserChannelCreate(user.dcUser.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.plugin.discordSession.ChannelMessageSend(ch.ID, msg)
}
