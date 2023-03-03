package service

type Service interface {
	Init()
	Start()
	Name() ServiceName
	Version() string
}
