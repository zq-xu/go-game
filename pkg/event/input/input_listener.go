package input

type InputListener interface {
	Update()
	Reload()

	Listen(e *KeyEvent)
	Idle(fn func())
}

type inputListener struct {
	idle      func()
	keyEvents []*KeyEvent
}

func NewInputListener() InputListener {
	return &inputListener{}
}

func (g *inputListener) Update() {
	var hasKeyPressed bool

	for _, v := range g.keyEvents {
		isPressed, exclusive := v.listern()
		if isPressed {
			hasKeyPressed = true
		}

		if exclusive {
			break
		}
	}

	if hasKeyPressed {
		return
	}

	if g.idle != nil {
		g.idle()
	}
}

func (g *inputListener) Listen(e *KeyEvent) { g.keyEvents = append(g.keyEvents, e) }
func (g *inputListener) Idle(fn func())     { g.idle = fn }
func (g *inputListener) Reload()            {}
