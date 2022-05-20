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
	imp := new(PluginDeviceControllerImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("PluginDeviceControllerImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(CivetDevicePluginCenter.PluginDeviceController)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".PluginDeviceControllerObj")

	// Run application
	tars.Run()
}
