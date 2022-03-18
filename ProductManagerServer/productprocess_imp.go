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

// ProductProcessImp servant implementation
type ProductProcessImp struct {
}

func (imp *ProductProcessImp) EditProduct(tarsCtx context.Context, ID *CivetDevicePluginCenter.Product, Key string, Value string, res *int32) (ret int32, err error) {
	if TarDB.EditRowByConditionbyModel(ID, Key, Value) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *ProductProcessImp) QueryAllProduct(tarsCtx context.Context, offset int32, limit int32, result *[]CivetDevicePluginCenter.Product, count *int32, res *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, result) {
		*res = 1
		return 0, nil
	} else {
		*res = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *ProductProcessImp) CreateProduct(tarsCtx context.Context, product *CivetDevicePluginCenter.Product, res *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("products") {
		if TarDB.CreateRow(product) {
			*res = 1
			return 0, nil
		} else {
			*res = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(CivetDevicePluginCenter.Product{}) {
			if TarDB.CreateRow(product) {
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

func (imp *ProductProcessImp) DeleteProduct(tarsCtx context.Context, ID int32, res *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&CivetDevicePluginCenter.Product{}, "id", ID) {
		*res = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

// Init servant init
func (imp *ProductProcessImp) Init() error {
	var TarsConfig = tars.NewRConf("CivetDevicePluginCenter", "ProductManagerServer", tars.GetServerConfig().BasePath)
	config, _ := TarsConfig.GetConfig("ProductManagerServer.conf")
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

func (imp *ProductProcessImp) Destroy() {

}

func (imp *ProductProcessImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *ProductProcessImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
