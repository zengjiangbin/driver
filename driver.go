package driver

import (
	"fmt"
	"github.com/zengjiangbin/driver/cmd"
	"github.com/zengjiangbin/driver/structs/service"
	"os"
	"runtime/debug"
	"time"
)

func Drive(s service.Service) {
	defer func() {
		if r := recover(); r != nil {
			now := time.Now().Format("2006-01-02 15:04:05")
			err := fmt.Sprintf(
				"\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n"+
					"[%s] main panic err: %v\n", now, r)
			fmt.Fprint(os.Stderr, err)
			debug.PrintStack()
			fmt.Fprint(os.Stderr, "================================================================================\n")
			os.Exit(-1)
		}
	}()
	cmd.Execute(s)
}
