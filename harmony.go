/*

Package harmony provides a simple and easy to use framework for discordgo bots

*/
package harmony

import (
	// externals
	"github.com/bwmarrin/discordgo"
)

// New creates a new Harmony object
//
// Parameters:
// 	prefix: the command prefix for the bot
// 	ignoreBots: a boolean telling if the command handler should ignore commands invoked by bots
// session: the discordgo Session object
//
// Returns:
// 	a Harmony object with which you can use to make bots with
func New(prefix string, ignoreBots bool, session *discordgo.Session) *Harmony {

	return &Harmony{
		Prefix:     prefix,
		Session:    session,
		Commands:   map[string]*Command{},
		IgnoreBots: ignoreBots,
	}

}

// Init initiates the event handling features of the library
func (h *Harmony) Init() {

	h.Session.AddHandler(h.MessageCreate)

}
