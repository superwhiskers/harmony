/*

Package harmony provides a simple and easy to use framework for discordgo bots

*/
package harmony

import (
	// internals
	"strings"
	// externals
	"github.com/bwmarrin/discordgo"
)

// New creates a new CommandHandler
func New(prefix string, ignoreBots bool, session *discordgo.Session) *CommandHandler {

	return &CommandHandler{
		Prefix:     prefix,
		Session:    session,
		Commands:   map[string]*Command{},
		IgnoreBots: ignoreBots,
	}

}

// Init initiates the event handling features of the library
func (h *CommandHandler) Init() {

	h.Session.AddHandler(h.onMessage)
	
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

// onMessage handles the onMessage event of discordgo
func (h *CommandHandler) onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	splitMessage := strings.Split(m.Content, " ")

	if h.OnMessageHandler != nil {

		go h.OnMessageHandler(s, m)

	}

	if m.Author.Bot && h.IgnoreBots {

		return

	}

	if len(splitMessage[0]) < (len(h.Prefix) + 1) {

		return

	}

	if h.PreCommandHandler != nil {

		if !h.PreCommandHandler(s, m) {

			return

		}

	}

	if splitMessage[0][0:len(h.Prefix)] != h.Prefix {

		return

	}

	if command, ok := h.Commands[strings.ToLower(splitMessage[0][len(h.Prefix):])]; ok {

		command.Run(s, m, splitMessage[1:])
		if command.SingleUse == true {

			h.RemoveCommand(strings.ToLower(splitMessage[0][len(h.Prefix):]))

		}

	}

}
