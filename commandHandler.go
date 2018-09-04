package harmony

import (
	// internals
	"strings"
	// externals
	"github.com/bwmarrin/discordgo"
)

// AddCommand adds a command to the handler
//
// Parameters:
// 	name: the command name
// 	singleUse: a boolean telling if this command can only be invoked once
// 	commandHandler: a function to run when this command is called
func (h *Harmony) AddCommand(name string, singleUse bool, commandHandler func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)) {

	h.Commands[name] = &Command{
		Run:       commandHandler,
		SingleUse: singleUse,
	}

}

// RemoveCommand removes a command from the handler
//
// Parameters:
// 	name: the name of the command to remove
func (h *Harmony) RemoveCommand(name string) {

	delete(h.Commands, name)

}

// CommandExists tells you if a command is registered under a name in the handler
//
// Parameters:
// 	name: the name of the command to check existence for
//
// Returns:
// 	a boolean telling you if it exists or not
func (h *Harmony) CommandExists(name string) bool {

	_, ok := h.Commands[name]
	return ok

}

// ExecCommand executes a command on the handler
//
// Parameters:
// 	name: the name of the command to execute
// 	message: the message to run it with
func (h *Harmony) ExecCommand(name string, message *discordgo.MessageCreate) {

	if command, ok := h.Commands[name]; ok {

		args := strings.Split(message.Content, " ")[1:]

		command.Run(h.Session, message, args)
		if command.SingleUse == true {

			h.RemoveCommand(name)

		}

	}

}
