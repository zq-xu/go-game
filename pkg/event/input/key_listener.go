package input

type InputListener interface {
	Update()
	Reload()

	Listen(e *KeyEvent)
	Idle(fn func())
	Blocks(fn ...func() bool)
}

type inputListener struct {
	blocks    []func() bool
	keyEvents []*KeyEvent
	idle      func()
}

func NewInputListener() InputListener {
	return &inputListener{}
}

func (il *inputListener) Update() {
	var hasKeyPressed bool

	for _, fn := range il.blocks {
		if fn() {
			return
		}
	}

	for _, v := range il.keyEvents {
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

	if il.idle != nil {
		il.idle()
	}
}

func (il *inputListener) Listen(e *KeyEvent) { il.keyEvents = append(il.keyEvents, e) }

func (il *inputListener) Idle(fn func()) { il.idle = fn }

func (il *inputListener) Reload() {}

func (il *inputListener) Blocks(fn ...func() bool) { il.blocks = append(il.blocks, fn...) }
