package input

import "github.com/hajimehoshi/ebiten/v2"

type InputListener interface {
	Update()
	Reload()

	Listen(e *KeyEvent)
	Idle(fn func())
}

type inputListener struct {
	lastPress ebiten.Key

	idle      func()
	keyEvents []*KeyEvent
}

func NewInputListener() InputListener {
	return &inputListener{}
}

func (g *inputListener) Update() {
	var isPressed bool

	for _, v := range g.keyEvents {
		pressed, exclusive := v.listern(g.lastPress)
		if pressed {
			isPressed = pressed
			g.lastPress = v.Key
		}

		if exclusive {
			break
		}
	}

	if isPressed {
		return
	}

	g.lastPress = 0

	if g.idle != nil {
		g.idle()
	}
}

func (g *inputListener) Listen(e *KeyEvent) { g.keyEvents = append(g.keyEvents, e) }
func (g *inputListener) Idle(fn func())     { g.idle = fn }
func (g *inputListener) Reload()            {}
