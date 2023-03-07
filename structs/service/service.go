package service

type Service interface {
	Init()
	Start()
	Name() Name
	Version() string
}
