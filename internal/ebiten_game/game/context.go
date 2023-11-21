package game

import (
	"github.com/zq-xu/go-game/internal/ebiten_game/listener"
	"github.com/zq-xu/go-game/internal/ebiten_game/resource"
)

type Context struct {
	Resource *resource.Resource
	Listener listener.Listener
}

func NewContext() *Context {
	return &Context{
		Resource: resource.NewResource(),
		Listener: listener.NewListener(),
	}
}
