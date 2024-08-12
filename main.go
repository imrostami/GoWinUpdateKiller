package main

import (
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"time"
)

func main() {

	appTimer := time.NewTicker(3 * time.Second)
	defer appTimer.Stop()

	for range appTimer.C {
		HandleWindowsUpdateTask()
	}

}

func HandleWindowsUpdateTask() {
	const windowsUpdateServiceName = "wuauserv"
	serviceManager, err := mgr.Connect()

	if err != nil {
		log.Fatalf("windows service manager not started , %v", err)
	}
	defer serviceManager.Disconnect()

	targetService, err := serviceManager.OpenService(windowsUpdateServiceName)

	if err != nil {
		log.Fatalf("failed to load windows update service %v", err)
	}
	serviceStatus, err := targetService.Query()

	if err != nil {
		log.Fatalf("failed to run query %v", err)
	}

	if serviceStatus.State == svc.Running {
		targetService.Control(svc.Stop)
	}
}
