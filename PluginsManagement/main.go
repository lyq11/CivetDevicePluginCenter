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
	imp := new(ManProcessImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("ManProcessImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(CivetDevicePluginCenter.ManProcess)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".ManProcessObj")

	// Run application
	tars.Run()
}
