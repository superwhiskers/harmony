/*

Package harmony provides an easy to use command handler for discordgo

*/
package harmony

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// New creates a new CommandHandler
func New(prefix string) *CommandHandler {

	return &CommandHandler{prefix, []Command{}}

}

// AddCommand adds a command to the handler
func (h *CommandHandler) AddCommand(name string, singleUse bool, commandHandler func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)) {

	h.Commands[name] = &Command{
		Run:       commandHandler,
		SingleUse: singleUse,
	}

}

// RemoveCommand removes a command from the handler
func (h *CommandHandler) RemoveCommand(name string) {

	delete(h.Commands, name)

}

// OnMessage handles the onMessage event of discordgo
func (h *CommandHandler) OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	splitMessage := strings.Split(m.Content, " ")

	if m.Author.Bot && h.IgnoreBots {

		return

	}

	if h.OnMessage != nil {

		go h.OnMessage(s, m)

	}

	if command, ok := h.Commands[strings.ToLower(splitMessage[0][len(h.Prefix):])]; ok {

		command.Run(s, m, splitMessage[1:])
		if command.SingleUse == true {

			h.RemoveCommand(strings.ToLower(splitMessage[0][len(h.Prefix):]))

		}

	}

}