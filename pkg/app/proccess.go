package app

type Proccess interface {
	Init()
	Tick()
	Stop()
}
