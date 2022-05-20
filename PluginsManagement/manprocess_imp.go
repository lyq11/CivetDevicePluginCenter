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

// ManProcessImp servant implementation
type ManProcessImp struct {
}

func (imp *ManProcessImp) QueryBindWithProduct(tarsCtx context.Context, offset int32, limit int32, ProductID string, result *[]CivetDevicePluginCenter.PluginBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), result, count, "product_id", ProductID) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *ManProcessImp) QueryBindWithProductKey(tarsCtx context.Context, offset int32, limit int32, ProductID int32, result *[]CivetDevicePluginCenter.PluginBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), result, count, "id", ProductID) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *ManProcessImp) QueryBindWithPluginID(tarsCtx context.Context, offset int32, limit int32, PluginID int32, result *[]CivetDevicePluginCenter.PluginBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), result, count, "plugin_id", PluginID) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *ManProcessImp) EditPlugin(tarsCtx context.Context, plugin *CivetDevicePluginCenter.Plugin, searchKey string, searchValue string, res *int32) (ret int32, err error) {
	if TarDB.EditRowByConditionbyModel(plugin, searchKey, searchValue) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *ManProcessImp) UnBindWithID(tarsCtx context.Context, relation *CivetDevicePluginCenter.PluginBind, res *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&CivetDevicePluginCenter.PluginBind{}, "id", relation.ID) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("unbind Fail")
	}
}

func (imp *ManProcessImp) CreatePlugin(tarsCtx context.Context, plugin *CivetDevicePluginCenter.Plugin, res *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("plugins") {
		if TarDB.CreateRow(plugin) {
			*res = 1
			return 0, nil
		} else {
			*res = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(CivetDevicePluginCenter.Plugin{}) {
			if TarDB.CreateRow(plugin) {
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

func (imp *ManProcessImp) DeletePlugin(tarsCtx context.Context, pluginID int32, res *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&CivetDevicePluginCenter.Plugin{}, "id", pluginID) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *ManProcessImp) EditPluginAddress(tarsCtx context.Context, plugin *CivetDevicePluginCenter.Plugin, searchKey string, searchValue string, res *int32) (ret int32, err error) {
	if TarDB.EditRowByConditionbyModel(plugin, searchKey, searchValue) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *ManProcessImp) QueryPlugins(tarsCtx context.Context, offset int32, limit int32, Pluginlist *[]CivetDevicePluginCenter.Plugin, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, Pluginlist) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *ManProcessImp) Bind(tarsCtx context.Context, relation *CivetDevicePluginCenter.PluginBind, res *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("plugin_binds") {
		if TarDB.CreateRow(relation) {
			*res = 1
			return 0, nil
		} else {
			*res = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(CivetDevicePluginCenter.PluginBind{}) {
			if TarDB.CreateRow(relation) {
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

func (imp *ManProcessImp) QueryBindAll(tarsCtx context.Context, offset int32, limit int32, result *[]CivetDevicePluginCenter.PluginBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, result) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

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

var TarDB *CivetTarsDataBase.CivetData

func (imp *ManProcessImp) Init() error {
	var TarsConfig = tars.NewRConf("CivetDevicePluginCenter", "PluginsManagement", tars.GetServerConfig().BasePath)
	config, _ := TarsConfig.GetConfig("PluginsManagement.conf")
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
func (imp *ManProcessImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *ManProcessImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *ManProcessImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
