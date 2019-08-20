package models

// CallableCommand is the default interface for all Commands.
// It needs to be implemented in all commands from all plugins and build-in functions
type CallableCommand interface {
	OnCall()
}

func a() {

}

type Command struct {
	name string
}
