package main

import (
	"CivetDevicePluginCenter"
	"Civets/CivetTarsDataBase"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
)

var ins *CivetTarsDataBase.CivetData

// DataBaseManagerImp servant implementation
type DataBaseManagerImp struct {
}

func (imp *DataBaseManagerImp) QueryDeviceCondition(tarsCtx context.Context, deviceNameJson string, companyID string, SearchKey string, SearchValue string, col *[]CivetDevicePluginCenter.TarsWarpPacket, Result *int32) (ret int32, err error) {
	model := CivetTarsDataBase.DBModel{}
	err = json.Unmarshal([]byte(deviceNameJson), &model)
	if err != nil {
		fmt.Println(err)
		*Result = -1
		return 0, err
	} else {
		ins.ConnectDataBaseByJson([]byte(deviceNameJson))
		var result []*CivetTarsDataBase.WarpPacket
		res := ins.QueryRowsConditionWithOutModel(model.TableName+companyID, &result, SearchKey, SearchValue)
		if res == true {
			*Result = 1
			var sf []CivetDevicePluginCenter.TarsWarpPacket
			for _, packet := range result {
				temp := CivetDevicePluginCenter.TarsWarpPacket{Key: packet.Keys, Base: packet.Base}
				sf = append(sf, temp)
			}
			*col = sf
			return 0, nil
		} else {
			*Result = -1
			return 0, errors.New("remove fail")
		}
	}
}

func (imp *DataBaseManagerImp) EditDevice(tarsCtx context.Context, deviceNameJson string, companyID string, SearchKey string, SearchValue string, col []CivetDevicePluginCenter.Column, Result *int32) (ret int32, err error) {
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
		res, errs := ins.EditRowWithOutModel(model.TableName+companyID, SearchKey, SearchValue, coverlist)
		if res {
			*Result = 1
			return 0, nil
		} else {
			*Result = -1
			return 0, errs
		}
	}
}

func (imp *DataBaseManagerImp) CreateDevice(tarsCtx context.Context, deviceNameJson string, companyID string, col []CivetDevicePluginCenter.Column, Result *int32) (ret int32, err error) {
	model := CivetTarsDataBase.DBModel{}
	ins.ConnectDataBaseByJson([]byte(deviceNameJson))
	res, err := ins.CreateTableByJson([]byte(deviceNameJson), &model, companyID)
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

type MysqlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Port     string `json:"port"`
}

func (imp *DataBaseManagerImp) QueryDeviceWithCondition(tarsCtx context.Context, deviceNameJson string, companyID string, Result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *DataBaseManagerImp) QueryDeviceAll(tarsCtx context.Context, deviceNameJson string, companyID string, col *[]CivetDevicePluginCenter.TarsWarpPacket, Result *int32) (ret int32, err error) {
	model := CivetTarsDataBase.DBModel{}
	err = json.Unmarshal([]byte(deviceNameJson), &model)
	if err != nil {
		fmt.Println(err)
		*Result = -1
		return 0, err
	} else {
		ins.ConnectDataBaseByJson([]byte(deviceNameJson))
		var result []*CivetTarsDataBase.WarpPacket
		res := ins.QueryRowsAllWithOutModel(model.TableName+companyID, &result)
		if res == true {
			*Result = 1
			var sf []CivetDevicePluginCenter.TarsWarpPacket
			for _, packet := range result {
				temp := CivetDevicePluginCenter.TarsWarpPacket{Key: packet.Keys, Base: packet.Base}
				sf = append(sf, temp)
			}
			*col = sf
			return 0, nil
		} else {
			*Result = -1
			return 0, errors.New("remove fail")
		}
	}
}

func (imp *DataBaseManagerImp) RemoveDevice(tarsCtx context.Context, deviceNameJson string, companyID string, SearchKey string, SearchValue string, Result *int32) (ret int32, err error) {
	model := CivetTarsDataBase.DBModel{}
	err = json.Unmarshal([]byte(deviceNameJson), &model)
	if err != nil {
		fmt.Println(err)
		*Result = -1
		return 0, err
	} else {
		ins.ConnectDataBaseByJson([]byte(deviceNameJson))
		res := ins.DeleteRowWithOutModel(model.TableName+companyID, SearchKey, SearchValue)
		if res == true {
			*Result = 1
			return 0, nil
		} else {
			*Result = -1
			return 0, errors.New("remove fail")
		}
	}

}

func (imp *DataBaseManagerImp) InitDevice(tarsCtx context.Context, deviceNameJson string, Result *int32) (ret int32, err error) {
	//TODO implement me
	var Model CivetTarsDataBase.DBModel
	ins.CreateDataBaseByJson([]byte(deviceNameJson), &Model)
	*Result = 1
	return 0, nil
}

// Init servant init
func (imp *DataBaseManagerImp) Init() error {
	var casd = tars.NewRConf("civetUserManagement", "AdminDataProcessServer", tars.GetServerConfig().BasePath)
	config, _ := casd.GetConfig("AdminDataProcessServer.conf")
	var mc MysqlConfig
	json.Unmarshal([]byte(config), &mc)
	ins = CivetTarsDataBase.CreateCivetData("root", "LYQsy921023", "127.0.0.1", "3306")

	return nil
}

// Destroy servant destory
func (imp *DataBaseManagerImp) Destroy() {
	//destroy servant here:
	//...
}
