package harmony

import "github.com/bwmarrin/discordgo"

// Harmony is a struct that contains data required for the framework to operate
//
// Fields:
// 	Prefix: the command prefix for the bot
// 	Session: the discordgo session object
// 	Commands: a map of command names to command objects
// 	EventHandlers: a list of functions used to handle events
// 	IgnoreBots: a boolean telling the command handler to ignore bots or not
type Harmony struct {
	Prefix        string
	Session       *discordgo.Session
	Commands      map[string]*Command
	EventHandlers []func(s *discordgo.Session, e *Event)
	IgnoreBots    bool
}

// Command is a structure that contains data that helps the CommandHandler execute commands
//
// Fields:
// 	Run: the function to call to run the command
// 	SingleUse: a boolean reperesenting if this command can only be run once
type Command struct {
	Run       func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
	SingleUse bool
}

// Event is a generic event struct that is sent as an argument to the event handler function
//
// Fields:
// 	EventObject: the object that the event was called with
// 	EventName: the name of the event
type Event struct {
	Event     interface{}
	EventName string
}

// PreCommandEvent is an event emitted by harmony before a command is run
//
// Fields:
// 	Message: the discordgo message object emitted
// 	CommandName: the name of the command to execute
type PreCommandEvent struct {
	Message     *discordgo.MessageCreate
	CommandName string
	Arguments   []string
}
