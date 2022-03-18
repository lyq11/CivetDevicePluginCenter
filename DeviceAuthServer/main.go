package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"CivetDevicePluginCenter"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(DeviceAuthProcessImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("DeviceAuthProcessImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(CivetDevicePluginCenter.DeviceAuthProcess)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".DeviceAuthProcessObj")

	// Run application
	tars.Run()
}
