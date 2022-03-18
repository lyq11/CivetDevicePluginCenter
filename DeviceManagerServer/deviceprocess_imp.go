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

// DeviceProcessImp servant implementation
type DeviceProcessImp struct {
}

func (imp *DeviceProcessImp) QueryAllDevice(tarsCtx context.Context, offset int32, limit int32, devicelist *[]CivetDevicePluginCenter.Device, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, devicelist) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *DeviceProcessImp) CreateDevice(tarsCtx context.Context, device *CivetDevicePluginCenter.Device, res *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("devices") {
		if TarDB.CreateRow(device) {
			*res = 1
			return 0, nil
		} else {
			*res = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(CivetDevicePluginCenter.Device{}) {
			if TarDB.CreateRow(device) {
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

func (imp *DeviceProcessImp) RemoveDevice(tarsCtx context.Context, DeviceID int32, res *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&CivetDevicePluginCenter.Device{}, "id", DeviceID) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *DeviceProcessImp) EditDevice(tarsCtx context.Context, info *CivetDevicePluginCenter.Device, searchKey string, searchValue string, res *int32) (ret int32, err error) {
	if TarDB.EditRowByConditionbyModel(info, searchKey, searchValue) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

// TarDB Init servant init
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

func (imp *DeviceProcessImp) Init() error {
	var TarsConfig = tars.NewRConf("CivetDevicePluginCenter", "DeviceManagerServer", tars.GetServerConfig().BasePath)
	config, _ := TarsConfig.GetConfig("DeviceManagerServer.conf")
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
func (imp *DeviceProcessImp) Destroy() {
	//destroy servant here:
	//...
}
