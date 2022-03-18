package main

import (
	"CivetDevicePluginCenter"
	"Civets/CivetTarsDataBase"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"os"
)

// DeviceClassProcessImp servant implementation
type DeviceClassProcessImp struct {
}

func (imp *DeviceClassProcessImp) CreateDeviceClass(tarsCtx context.Context, info *CivetDevicePluginCenter.DeviceClass, res *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("device_classes") {
		if TarDB.CreateRow(info) {
			*res = 1
			return 0, nil
		} else {
			*res = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(CivetDevicePluginCenter.DeviceClass{}) {
			if TarDB.CreateRow(info) {
				*res = 1
				return 0, nil
			} else {
				*res = -1
				return 0, err
			}
		} else {
			return 0, err
		}
	}
}

func (imp *DeviceClassProcessImp) RemoveDeviceClass(tarsCtx context.Context, DeviceClassID int32, res *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&CivetDevicePluginCenter.DeviceClass{}, "id", DeviceClassID) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *DeviceClassProcessImp) EditDeviceClass(tarsCtx context.Context, info *CivetDevicePluginCenter.DeviceClass, searchKey string, searchValue string, res *int32) (ret int32, err error) {
	if TarDB.EditRowByConditionbyModel(info, searchKey, searchValue) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *DeviceClassProcessImp) QueryAllDevice(tarsCtx context.Context, offset int32, limit int32, deviceClasslist *[]CivetDevicePluginCenter.DeviceClass, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, deviceClasslist) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

// Init servant init
var TarDB *CivetTarsDataBase.CivetData

type Config struct {
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBHost     string `json:"DBHost"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
	RDUserName string `json:"RDUserName"`
	RDPassword string `json:"RDPassword"`
	RDHost     string `json:"RDHost"`
	RDPort     string `json:"RDPort"`
	RDSize     int    `json:"RDSize"`
}

func (imp *DeviceClassProcessImp) Init() error {
	var TarsConfig = tars.NewRConf("CivetDevicePluginCenter", "DeviceClassManagerServer", tars.GetServerConfig().BasePath)
	config, _ := TarsConfig.GetConfig("DeviceClassManagerServer.conf")
	var mc Config
	json.Unmarshal([]byte(config), &mc)
	if mc.DBName == "" || mc.DBHost == "" || mc.DBUserName == "" {
		fmt.Println("The Mysql Config error")
		os.Exit(1)
	} else {
		fmt.Println("the config is", config, mc.DBHost, mc.DBPassword, mc.DBUserName, mc.DBName)
	}
	TarDB = CivetTarsDataBase.CreateCivetData(mc.DBUserName, mc.DBPassword, mc.DBHost, mc.DBPort)
	TarDB.ConnectOrCreateDataBase(mc.DBName)
	return nil
}

// Destroy servant destory
func (imp *DeviceClassProcessImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *DeviceClassProcessImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *DeviceClassProcessImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
