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

// DeviceAuthProcessImp servant implementation
type DeviceAuthProcessImp struct {
}

func (imp *DeviceAuthProcessImp) QueryCompanyDeviceCBindAll(tarsCtx context.Context, offset int32, limit int32, result *[]CivetDevicePluginCenter.CompanyDeivceBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, result) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *DeviceAuthProcessImp) QueryCompanyBindWithDeviceC(tarsCtx context.Context, offset int32, limit int32, DeviceClassID int32, result *[]CivetDevicePluginCenter.CompanyDeivceBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), result, count, "deivce_class_id", DeviceClassID) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *DeviceAuthProcessImp) QueryDeviceCBindWithCompanyID(tarsCtx context.Context, offset int32, limit int32, CompanyID int32, result *[]CivetDevicePluginCenter.CompanyDeivceBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), result, count, "company_id", CompanyID) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *DeviceAuthProcessImp) BindProductCompany(tarsCtx context.Context, relation *CivetDevicePluginCenter.CompanyProductBind, res *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("company_product_binds") {
		if TarDB.CreateRow(relation) {
			*res = 1
			return 0, nil
		} else {
			*res = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(CivetDevicePluginCenter.CompanyProductBind{}) {
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

func (imp *DeviceAuthProcessImp) UnBindProductCompany(tarsCtx context.Context, relation *CivetDevicePluginCenter.CompanyProductBind, res *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&CivetDevicePluginCenter.CompanyProductBind{}, "id", relation.ID) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *DeviceAuthProcessImp) QueryProductCompanyBindAll(tarsCtx context.Context, offset int32, limit int32, result *[]CivetDevicePluginCenter.CompanyProductBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, result) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *DeviceAuthProcessImp) QueryCompanyBindWithProductID(tarsCtx context.Context, offset int32, limit int32, ProductClassID int32, result *[]CivetDevicePluginCenter.CompanyProductBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), result, count, "product_id", ProductClassID) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *DeviceAuthProcessImp) QueryProductBindWithCompanyID(tarsCtx context.Context, offset int32, limit int32, CompanyID int32, result *[]CivetDevicePluginCenter.CompanyProductBind, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), result, count, "company_id", CompanyID) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *DeviceAuthProcessImp) BindDeviceCompany(tarsCtx context.Context, relation *CivetDevicePluginCenter.CompanyDeivceBind, res *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("company_deivce_binds") {
		if TarDB.CreateRow(relation) {
			*res = 1
			return 0, nil
		} else {
			*res = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(CivetDevicePluginCenter.CompanyDeivceBind{}) {
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

func (imp *DeviceAuthProcessImp) UnBindDeviceCompany(tarsCtx context.Context, relation *CivetDevicePluginCenter.CompanyDeivceBind, res *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&CivetDevicePluginCenter.CompanyDeivceBind{}, "id", relation.ID) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

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

func (imp *DeviceAuthProcessImp) Init() error {
	var TarsConfig = tars.NewRConf("CivetDevicePluginCenter", "DeviceAuthServer", tars.GetServerConfig().BasePath)
	config, _ := TarsConfig.GetConfig("DeviceAuthServer.conf")
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
func (imp *DeviceAuthProcessImp) Destroy() {
	//destroy servant here:
	//...
}
