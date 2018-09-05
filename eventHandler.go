package harmony

import (
	// internals
	// externals
	"github.com/bwmarrin/discordgo"
)

//
func (h *Harmony) AddEventHandler(singleUse bool, eventHandler func(s *discordgo.Session, e *Event)) {

	h.EventHandlers = append(h.EventHandlers, 
