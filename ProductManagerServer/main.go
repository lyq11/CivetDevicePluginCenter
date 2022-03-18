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
	imp := new(ProductProcessImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("ProductProcessImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(CivetDevicePluginCenter.ProductProcess)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".ProductProcessObj")

	// Run application
	tars.Run()
}
