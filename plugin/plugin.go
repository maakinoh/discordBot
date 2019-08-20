package plugin

import (
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"

	"github.com/robertkrimen/otto"
)

// Plugin represent a loadable extention to the bot
type Plugin struct {
	// Name of the plugin
	Name string `json:"name"`

	// Maintainer of the plugin
	Maintainer string `json:"maintainer"`

	// Path of the JavaScript file
	ScriptPath string `json:"scriptPath"`

	BotCommands []string `json:"botcommands"`

	discordSession *discordgo.Session `json:"-"`

	ottoVM *otto.Otto `json:"-"`
}

func NewPlugin(session *discordgo.Session, scriptPath string) Plugin {
	return Plugin{discordSession: session, BotCommands: make([]string, 0), ScriptPath: scriptPath}
}

func (plugin *Plugin) InitPlugin(session *discordgo.Session) {
	plugin.discordSession = session
}

// ExecuteCommand is called when a user type a bot command that is associated with this Plugin.
// This function will start a new otto JS VM and call the OnCommand function from the script
func (plugin *Plugin) ExecuteCommand(msg *discordgo.MessageCreate) {
	vm := otto.New()

	plugin.ottoVM = vm

	content, fileErr := ioutil.ReadFile(".bot/plugins/" + plugin.ScriptPath)

	if fileErr != nil {
		fmt.Println(fileErr.Error())
	}

	res := string(content)
	vm.Run(res)

	val, errFunc := vm.Get("OnCommand")
	if errFunc != nil {
		fmt.Println("Error with command handler on plugin: " + plugin.Name + " error: " + errFunc.Error())
	}

	// Check if the function exists and if there where no errors
	if val.IsFunction() || errFunc != nil {

	}
	msga := NewMessge(*plugin, msg)

	val.Call(val, msga)
}
