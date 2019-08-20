package plugin

import "github.com/bwmarrin/discordgo"

type Guild struct {
	// The Name of the Guild
	Name string

	guild *discordgo.Guild

	plugin Plugin
}

func NewGuild(plugin Plugin, dcGuild *discordgo.Guild) Guild {
	return Guild{Name: dcGuild.Name, plugin: plugin, guild: dcGuild}
}

func (guild Guild) GetChannelByName(name string) {

}

func (guild Guild) GetGuildMenbers() []User {
	ret := make([]User, 0)
	for _, m := range guild.guild.Members {
		ret = append(ret, NewUser(guild.plugin, m.User))
	}
	return ret
}
