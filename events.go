package harmony

import (
	// internals
	"strings"
	// externals
	"github.com/bwmarrin/discordgo"
)

// MessageCreate handles the MessageCreate event of discordgo
//
// Parameters:
//  s: the discordgo Session object
// 	m: the discordgo MessageCreate event object
func (h *Harmony) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	splitMessage := strings.Split(m.Content, " ")

	if m.Author.Bot && h.IgnoreBots {

		return

	}

	if len(splitMessage[0]) < (len(h.Prefix) + 1) {

		return

	}

	if splitMessage[0][0:len(h.Prefix)] != h.Prefix {

		return

	}

	h.ExecCommand(strings.ToLower(splitMessage[0][len(h.Prefix):]), m)

}

// MessageUpdate handles the MessageUpdate event of discordgo
//
// Parameters:
// 	s: the discordgo Session object
// 	m: the discordgo MessageUpdate event object
func (h.Harmony) MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {

	// do something
	return

}
