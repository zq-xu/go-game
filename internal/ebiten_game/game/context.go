package game

import "github.com/zq-xu/2d-game/internal/ebiten_game/resource"

type Context struct {
	Resource *resource.Resource
}

func NewContext() *Context {
	return &Context{
		Resource: resource.NewResource(),
	}
}
