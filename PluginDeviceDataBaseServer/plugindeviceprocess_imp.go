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

// PluginDeviceProcessImp servant implementation
type PluginDeviceProcessImp struct {
}

func (imp *PluginDeviceProcessImp) QueryDeviceWithCondition(tarsCtx context.Context, deviceNameJson string, Result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceProcessImp) QueryDeviceCondition(tarsCtx context.Context, deviceNameJson string, SearchKey string, SearchValue string, col *[]CivetDevicePluginCenter.TarsWarpPacket, Result *int32) (ret int32, err error) {
	model := CivetTarsDataBase.DBModel{}
	err = json.Unmarshal([]byte(deviceNameJson), &model)
	if err != nil {
		fmt.Println(err)
		*Result = -1
		return 0, err
	} else {
		ins.ConnectDataBaseByJson([]byte(deviceNameJson))
		var result []*CivetTarsDataBase.WarpPacket
		res := ins.QueryRowsConditionWithOutModel(model.TableName, &result, SearchKey, SearchValue)
		if res == true {
			*Result = 1
			var sf []CivetDevicePluginCenter.TarsWarpPacket
			for _, packet := range result {
				fmt.Println("the packet is ", packet)
				temp := CivetDevicePluginCenter.TarsWarpPacket{Key: packet.Keys, Base: packet.Base}
				sf = append(sf, temp)
			}
			*col = sf
			*Result = 1
			fmt.Println("the result is ", sf)
			return 0, nil
		} else {
			*Result = -1
			return 0, errors.New("remove fail")
		}
	}
}

func (imp *PluginDeviceProcessImp) QueryDeviceAll(tarsCtx context.Context, deviceNameJson string, offset int32, limit int32, col *[]CivetDevicePluginCenter.TarsWarpPacket, count *int32, Result *int32) (ret int32, err error) {
	model := CivetTarsDataBase.DBModel{}
	err = json.Unmarshal([]byte(deviceNameJson), &model)
	if err != nil {
		fmt.Println(err)
		*Result = -1
		return 0, err
	} else {
		ins.ConnectDataBaseByJson([]byte(deviceNameJson))
		var result []*CivetTarsDataBase.WarpPacket
		res := ins.QueryRowsAllWithOutModelOSLM(model.TableName, offset, limit, &result)
		if res == true {
			*Result = 1
			var sf []CivetDevicePluginCenter.TarsWarpPacket
			for _, packet := range result {
				fmt.Println("the packet is ", packet)
				temp := CivetDevicePluginCenter.TarsWarpPacket{Key: packet.Keys, Base: packet.Base}
				sf = append(sf, temp)
			}
			*col = sf
			*Result = 1
			fmt.Println("the result is ", sf)
			return 0, nil
		} else {
			*Result = -1
			return 0, errors.New("remove fail")
		}
	}
}

func (imp *PluginDeviceProcessImp) QueryBindDeviceWithCondition(tarsCtx context.Context, deviceNameJson string, companyID string, Result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceProcessImp) QueryBindDeviceCondition(tarsCtx context.Context, deviceNameJson string, companyID string, SearchKey string, SearchValue string, col *[]CivetDevicePluginCenter.TarsWarpPacket, Result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceProcessImp) QueryBindDevicecAll(tarsCtx context.Context, deviceNameJson string, offset int32, limit int32, companyID string, col *[]CivetDevicePluginCenter.TarsWarpPacket, count *int32, Result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceProcessImp) CreateDevice(tarsCtx context.Context, deviceNameJson string, col []CivetDevicePluginCenter.Column, Result *int32) (ret int32, err error) {

	model := CivetTarsDataBase.DBModel{}
	ins.ConnectDataBaseByJson([]byte(deviceNameJson))
	res, err := ins.CreateTableByJson([]byte(deviceNameJson), &model, "")
	if err != nil {
		*Result = -1
		return -1, err
	}
	fmt.Printf("run here")
	var coverlist []CivetTarsDataBase.Column
	for _, column := range col {
		cover := CivetTarsDataBase.Column{
			Key:   column.Key,
			Value: column.Value,
		}
		coverlist = append(coverlist, cover)
	}
	rs := ins.CreateRowWithOutModel(res, coverlist)
	if rs != true {
		*Result = -1
		return -2, errors.New("something error")
	} else {
		*Result = 1
		return 0, nil
	}
}

func (imp *PluginDeviceProcessImp) RemoveDevice(tarsCtx context.Context, deviceNameJson string, SearchKey string, SearchValue string, Result *int32) (ret int32, err error) {
	model := CivetTarsDataBase.DBModel{}
	err = json.Unmarshal([]byte(deviceNameJson), &model)
	if err != nil {
		fmt.Println(err)
		*Result = -1
		return 0, err
	} else {
		ins.ConnectDataBaseByJson([]byte(deviceNameJson))
		res := ins.DeleteRowWithOutModel(model.TableName, SearchKey, SearchValue)
		if res == true {
			*Result = 1
			return 0, nil
		} else {
			*Result = -1
			return 0, errors.New("remove fail")
		}
	}
}

func (imp *PluginDeviceProcessImp) BindDeviceCompany(tarsCtx context.Context, deviceNameJson string, companyID string, col []CivetDevicePluginCenter.Column, Result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceProcessImp) UnbindDeviceCompany(tarsCtx context.Context, deviceNameJson string, companyID string, SearchKey string, SearchValue string, Result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

var ins *CivetTarsDataBase.CivetData

//func (imp *PluginDeviceProcessImp) QueryDeviceWithCondition(tarsCtx context.Context, deviceNameJson string, companyID string, Result *int32) (ret int32, err error) {
//	//TODO implement me
//	panic("implement me")
//}

//func (imp *PluginDeviceProcessImp) QueryDeviceCondition(tarsCtx context.Context, deviceNameJson string, companyID string, SearchKey string, SearchValue string, col *[]CivetDevicePluginCenter.TarsWarpPacket, Result *int32) (ret int32, err error) {
//
//}

//func (imp *PluginDeviceProcessImp) QueryDeviceAll(tarsCtx context.Context, deviceNameJson string, offset int32, limit int32, companyID string, col *[]CivetDevicePluginCenter.TarsWarpPacket, count *int32, Result *int32) (ret int32, err error) {
//
//}

//func (imp *PluginDeviceProcessImp) CreateDevice(tarsCtx context.Context, deviceNameJson string, companyID string, col []CivetDevicePluginCenter.Column, Result *int32) (ret int32, err error) {
//
//}

//func (imp *PluginDeviceProcessImp) RemoveDevice(tarsCtx context.Context, deviceNameJson string, companyID string, SearchKey string, SearchValue string, Result *int32) (ret int32, err error) {
//
//}

func (imp *PluginDeviceProcessImp) EditDevice(tarsCtx context.Context, deviceNameJson string, SearchKey string, SearchValue string, col []CivetDevicePluginCenter.Column, Result *int32) (ret int32, err error) {
	model := CivetTarsDataBase.DBModel{}
	err = json.Unmarshal([]byte(deviceNameJson), &model)
	if err != nil {
		fmt.Println(err)
		*Result = -1
		return 0, err
	} else {
		ins.ConnectDataBaseByJson([]byte(deviceNameJson))
		var coverlist []CivetTarsDataBase.Column
		for _, column := range col {
			cover := CivetTarsDataBase.Column{
				Key:   column.Key,
				Value: column.Value,
			}
			coverlist = append(coverlist, cover)
		}
		res, errs := ins.EditRowWithOutModel(model.TableName, SearchKey, SearchValue, coverlist)
		if res {
			*Result = 1
			return 0, nil
		} else {
			*Result = -1
			return 0, errs
		}
	}
}

func (imp *PluginDeviceProcessImp) InitDevice(tarsCtx context.Context, deviceNameJson string, Result *int32) (ret int32, err error) {
	//TODO implement me
	var Model CivetTarsDataBase.DBModel
	ins.CreateDataBaseByJson([]byte(deviceNameJson), &Model)
	*Result = 1
	return 0, nil
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

// Init servant init
func (imp *PluginDeviceProcessImp) Init() error {
	var casd = tars.NewRConf("CivetDevicePluginCenter", "PluginDeviceDataBaseServer", tars.GetServerConfig().BasePath)
	config, _ := casd.GetConfig("PluginDeviceDataBaseServer.conf")
	var mc Config
	json.Unmarshal([]byte(config), &mc)
	if mc.DBName == "" || mc.DBHost == "" || mc.DBUserName == "" {
		fmt.Println("The Mysql Config error")
		os.Exit(1)
	} else {
		fmt.Println("the config is", config, mc.DBHost, mc.DBPassword, mc.DBUserName, mc.DBName)
	}
	ins = CivetTarsDataBase.CreateCivetData(mc.DBUserName, mc.DBPassword, mc.DBHost, mc.DBPort)
	ins.ConnectOrCreateDataBase(mc.DBName)

	return nil
}

// Destroy servant destory
func (imp *PluginDeviceProcessImp) Destroy() {
	//destroy servant here:
	//...
}
