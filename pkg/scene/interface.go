package scene

type Scene interface {
	Name() string
	OnEnter()
	OnExit()
	Update(dt float32)
	Draw()
}
