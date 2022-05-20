package main

import (
	"CivetDevicePluginCenter"
	PluginsManagement "CivetDevicePluginCenter/PluginsManagement/CivetDevicePluginCenter"
	"Civets/CivetPluginsFrameWork"
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

func (imp *DeviceProcessImp) AnalyzeDevice(tarsCtx context.Context, ProductID string, Data string, res *int32) (ret int32, err error) {
	var checkProductExistResult int32
	var checkProductResult []PluginsManagement.PluginBind
	var checkProductResultCount int32
	_, err = PluginWithProduct.QueryBindWithProduct(0, 0, ProductID, &checkProductResult, &checkProductResultCount, &checkProductExistResult)
	if err != nil {
		*res = -1
		return 0, errors.New("cannot Find Product")
	}
	arg, err := CivetFWIns.CallFuncWith1ArgWithReturn(ProductID, "Analyst", Data)
	if err != nil {
		*res = -1
		fmt.Println("Analyst Fail", err)
		return 0, err
	}
	if arg == true {
		*res = 1
		fmt.Println("Analyst success")
		return 0, nil
	} else {
		*res = -1
		return 0, err
	}
}

func (imp *DeviceProcessImp) SetDeviceProperties(tarsCtx context.Context, DeviceID string, jsons string, res *int32) (ret int32, err error) {
	var checkProductExistResult int32
	var checkProductResult []PluginsManagement.PluginBind
	var checkProductResultCount int32
	var DeviceModel CivetDevicePluginCenter.Device
	if TarDB.QueryRowsWithCondition(&DeviceModel, "deivce_id", DeviceID) {
		_, err = PluginWithProduct.QueryBindWithProduct(0, 0, DeviceModel.ProductID, &checkProductResult, &checkProductResultCount, &checkProductExistResult)
		if err != nil {
			*res = -1
			return 0, errors.New("cannot Find Product")
		}
		arg, err := CivetFWIns.CallFuncWith3ArgWithReturn(DeviceModel.ProductID, "Edit", "DeviceID", "'"+DeviceID+"'", jsons)
		if err != nil {
			*res = -1
			fmt.Println("Edit Fail", err)
			return 0, err
		}
		if arg == true {
			*res = 1
			fmt.Println("Edit success")
			return 0, nil
		} else {
			*res = -1
			return 0, err
		}
	} else {
		*res = -1
		fmt.Println(DeviceModel)
		return 0, errors.New("ID is not found")
	}
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
	var checkProductExistResult int32
	var checkProductResult []PluginsManagement.PluginBind
	var checkProductResultCount int32
	_, err = PluginWithProduct.QueryBindWithProduct(0, 0, device.ProductID, &checkProductResult, &checkProductResultCount, &checkProductExistResult)
	if err != nil {
		return 0, err
	}
	if checkProductResultCount > 0 {
		if TarDB.CheckTableExist("devices") {
			if TarDB.CreateRow(device) {
				fmt.Println(checkProductResult)
				arg, err := CivetFWIns.CallFuncWith1ArgWithReturn(checkProductResult[0].ProductID, "Create", "'"+device.DeivceID+"'")
				if err != nil {
					*res = -1
					fmt.Println("CreateDevice Fail")
					return 0, err
				}
				if arg == true {
					*res = 1
					fmt.Println("CreateDevice success")
					return 0, nil
				} else {
					*res = -1
					return 0, err
				}
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
	} else {
		*res = -1
		fmt.Println("CreateDevice Fail")
		return 0, err
	}

}

func (imp *DeviceProcessImp) RemoveDevice(tarsCtx context.Context, ID int32, res *int32) (ret int32, err error) {
	var checkProductExistResult int32
	var checkProductResult []PluginsManagement.PluginBind
	var checkProductResultCount int32
	var DeviceModel CivetDevicePluginCenter.Device
	if TarDB.QueryRowWithID(&DeviceModel, ID) {
		_, err = PluginWithProduct.QueryBindWithProduct(0, 0, DeviceModel.ProductID, &checkProductResult, &checkProductResultCount, &checkProductExistResult)
		if err != nil {
			*res = -1
			return 0, errors.New("cannot Find Product")
		}
		arg, err := CivetFWIns.CallFuncWith2ArgWithReturn(DeviceModel.ProductID, "Remove", "DeviceID", "'"+DeviceModel.DeivceID+"'")
		if err != nil {
			*res = -1
			fmt.Println("Del Device Fail")
			return 0, errors.New("delete Device failed From Plugins")
		}
		if arg == true {
			if TarDB.DelRowByCondition(&CivetDevicePluginCenter.Device{}, "id", ID) {
				*res = 1
				fmt.Println("Del Device success")
				return 0, nil
			} else {
				*res = -1
				return 0, errors.New("delete Device failed From deviceDatabase")
			}

		} else {
			*res = -1
			return 0, err
		}
	} else {
		*res = -1
		fmt.Println(DeviceModel)
		return 0, errors.New("ID is not found")
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

var comm *tars.Communicator
var PluginWithProduct *PluginsManagement.ManProcess

func initPluginsManagement() {
	comm = tars.NewCommunicator()
	//comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 172.25.0.3 -t 60000 -p 17890")
	//obj := "CivetDevicePluginCenter.DeviceManagerServer.DeviceProcessObj@tcp -h 127.0.0.1 -p 10015 -t 60000"
	obj := "CivetDevicePluginCenter.PluginsManagement.ManProcessObj"
	PluginWithProduct = new(PluginsManagement.ManProcess)
	comm.StringToProxy(obj, PluginWithProduct)
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

func loadBinds() bool {
	var BindList []PluginsManagement.PluginBind
	var ListCount int32
	var result int32
	_, err := PluginWithProduct.QueryBindAll(-1, -1, &BindList, &ListCount, &result)
	if err != nil {
		return false
	}
	if result == 1 && ListCount > 0 {
		for _, c := range BindList {
			fmt.Println(c)
			CivetFWIns.BindPlug(c.ProductID, PluginListWithID[c.ID])
		}
		return true
	} else if ListCount == 0 {
		tars.TLOG.Error("No Plugins Bind found")
		return true
	} else {
		tars.TLOG.Error("Plugin system may has error")

		return false
	}
}

var PluginListWithID map[int32]string

func LoadPlugins() bool {
	var PlugList []PluginsManagement.Plugin
	var ListCount int32
	var result int32
	PluginListWithID = make(map[int32]string)
	_, err := PluginWithProduct.QueryPlugins(-1, -1, &PlugList, &ListCount, &result)
	if err != nil {
		return false
	}
	if result == 1 && ListCount > 0 {
		for _, c := range PlugList {
			fmt.Println(c)
			CivetFWIns.RegPlug(c.Name, c.Path)
			PluginListWithID[c.ID] = c.Name
		}
		fmt.Println(CivetFWIns)
		return true
	} else if ListCount == 0 {
		tars.TLOG.Error("No Plugins found")
		return true
	} else {
		tars.TLOG.Error("Plugin system may has error")
		return false
	}
}

var CivetFWIns *CivetPluginsFrameWork.Plugins

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
	initPluginsManagement()
	CivetFWIns = CivetPluginsFrameWork.GetIns()

	if LoadPlugins() == false {
		tars.TLOG.Error("Plugin system may has error")
		os.Exit(-1)
	}

	if loadBinds() == false {
		tars.TLOG.Error("Plugin system may has error")
		os.Exit(-1)
	}
	return nil
}

// Destroy servant destory
func (imp *DeviceProcessImp) Destroy() {
	//destroy servant here:
	//...
}
