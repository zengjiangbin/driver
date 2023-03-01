package driver

import "github.com/zengjiangbin/driver/structs/service"

func Driver(s service.Service) {
	s.Init()
	s.Start()
}
