package driver

import "github.com/zengjiangbin/driver/structs/service"

func Drive(s service.Service) {
	s.Init()
	s.Start()
}
