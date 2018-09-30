/*

Package harmony provides an easy to use command handler for discordgo

*/
package harmony

import (
	// internals
	"strings"
	// externals
	"github.com/bwmarrin/discordgo"
)

// New creates a new CommandHandler
func New(prefix string, ignoreBots bool) *CommandHandler {

	return &CommandHandler{
		Prefix:     prefix,
		Commands:   map[string]*Command{},
		IgnoreBots: ignoreBots,
	}

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

	if h.OnMessageHandler != nil {

		go h.OnMessageHandler(s, m)

	}

	if m.Author.Bot && h.IgnoreBots {

		return

	}

	if len(m.Content) <= len(h.Prefix) || m.Content[0:len(h.Prefix)] != h.Prefix {

		return

	}

	if h.PreCommandHandler != nil {

		if !h.PreCommandHandler(s, m) {

			return

		}

	}

	postPrefix := strings.Split(m.Content[len(h.Prefix):], " ")
	if command, ok := h.Commands[strings.ToLower(postPrefix[0])]; ok {

		command.Run(s, m, postPrefix[1:])
		if command.SingleUse == true {

			h.RemoveCommand(strings.ToLower(postPrefix[0]))

		}

	}

}
