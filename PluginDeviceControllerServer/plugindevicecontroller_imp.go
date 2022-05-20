package main

import (
	PluginsManagement "CivetDevicePluginCenter/PluginsManagement/CivetDevicePluginCenter"
	"Civets/CivetPluginsFrameWork"
	"context"
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"os"
)

// PluginDeviceControllerImp servant implementation
type PluginDeviceControllerImp struct {
}

func (imp *PluginDeviceControllerImp) GetDeviceJson(tarsCtx context.Context, ProductID string, jsons *string, result *int32) (ret int32, err error) {
	res, err := CivetFWIns.CallFuncWithPString(ProductID, "GetJson", jsons)
	if err != nil {
		*result = -1
		return -1, err
	}
	if res == true {
		*result = 1
		return 1, nil
	} else {
		*result = -1
		return -1, errors.New("getDeviceJsonFail")
	}

}

func (imp *PluginDeviceControllerImp) SetDeviceProp(tarsCtx context.Context, companyID string, ProductID string, DeviceID string, jsons string, res *int32) (ret int32, err error) {
	arg, err := CivetFWIns.CallFuncWith4ArgWithReturn(ProductID, "Edit", companyID, "DeviceID", "'"+DeviceID+"'", jsons)
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
}

func (imp *PluginDeviceControllerImp) CreateDevice(tarsCtx context.Context, companyID string, ProductID string, DeviceID string, res *int32) (ret int32, err error) {
	arg, err := CivetFWIns.CallFuncWith2ArgWithReturn(ProductID, "Create", companyID, "'"+DeviceID+"'")
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
}

func (imp *PluginDeviceControllerImp) DelDevice(tarsCtx context.Context, companyID string, ProductID string, DeviceID string, res *int32) (ret int32, err error) {
	arg, err := CivetFWIns.CallFuncWith3Arg(ProductID, "Remove", companyID, "DeviceID", "'"+DeviceID+"'")
	if err != nil {
		*res = -1
		fmt.Println("DelDevice Fail")
		return 0, err
	}
	if arg == true {
		*res = 1
		fmt.Println("DelDevice success")
		return 0, nil
	} else {
		*res = -1
		return 0, err
	}
}

func (imp *PluginDeviceControllerImp) QueryDevice(tarsCtx context.Context, companyID string, ProductID string, DeviceID string, jsons *string, res *int32) (ret int32, err error) {
	fmt.Println("QueryDevice")
	arg, err := CivetFWIns.CallFuncWith4ArgWithPstring(ProductID, "QueryDetail", companyID, "DeviceID", "'"+DeviceID+"'", jsons)
	if err != nil {
		*res = -1
		fmt.Println("QueryDevice")
		return 0, err
	}

	if arg == true {
		*res = 1
		fmt.Println("QueryDevice")
		return 0, nil
	} else {
		*res = -1
		return 0, err
	}
}

func (imp *PluginDeviceControllerImp) SetDeviceDesiredProperty(tarsCtx context.Context) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceControllerImp) QueryDeviceDesiredProperty(tarsCtx context.Context) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceControllerImp) DeleteDeviceDesiredProperty(tarsCtx context.Context) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceControllerImp) CallDeviceService(tarsCtx context.Context) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceControllerImp) CallDeviceDesiredService(tarsCtx context.Context) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceControllerImp) QueryDeviceDesiredCallService(tarsCtx context.Context) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *PluginDeviceControllerImp) LoadPlugins(tarsCtx context.Context, c *int32) (ret int32, err error) {
	var PlugList []PluginsManagement.Plugin
	var ListCount int32
	var result int32
	_, err = pluginsManagement.QueryPlugins(-1, -1, &PlugList, &ListCount, &result)
	if err != nil {
		return 0, err
	}
	if result == 1 && ListCount > 0 {
		for _, c := range PlugList {
			fmt.Println(c)
			CivetFWIns.RegPlug(c.Name, c.Path)
		}
		fmt.Println(CivetFWIns)
		*c = 1
		return 0, nil
	} else {
		tars.TLOG.Error("Plugin system may has error")
		*c = -1
		return 0, errors.New("LoadPlugins Fail")
	}
}

func (imp *PluginDeviceControllerImp) LoadBinds(tarsCtx context.Context, c *int32) (ret int32, err error) {
	var BindList []PluginsManagement.PluginBind
	var ListCount int32
	var result int32
	_, err = pluginsManagement.QueryBindAll(-1, -1, &BindList, &ListCount, &result)
	if err != nil {
		return 0, err
	}
	if result == 1 && ListCount > 0 {
		for _, c := range BindList {
			fmt.Println(c)
			CivetFWIns.BindPlug(c.ProductID, "math")
		}
		*c = 1
		return 0, nil
	} else {
		tars.TLOG.Error("Plugin system may has error")
		*c = -1
		return 0, errors.New("LoadBinds Fail")
	}

}

var pluginsManagement *PluginsManagement.ManProcess
var CivetFWIns *CivetPluginsFrameWork.Plugins

// Init servant init
func (imp *PluginDeviceControllerImp) Init() error {
	var tarBackground = context.Background()
	var initPluginRes int32
	var initPluginBindRes int32
	comm := tars.NewCommunicator()
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 172.25.0.3 -t 60000 -p 17890")
	pluginsManagement = new(PluginsManagement.ManProcess)
	obj := "CivetDevicePluginCenter.PluginsManagement.ManProcessObj"
	comm.StringToProxy(obj, pluginsManagement)
	CivetFWIns = CivetPluginsFrameWork.GetIns()
	_, _ = imp.LoadPlugins(tarBackground, &initPluginRes)
	if initPluginRes != 1 {
		tars.TLOG.Error("Plugin system may has error")
		os.Exit(-1)
	}
	_, _ = imp.LoadBinds(tarBackground, &initPluginBindRes)
	if initPluginBindRes != 1 {
		tars.TLOG.Error("Plugin system may has error")
		os.Exit(-1)
	}

	return nil
}

// Destroy servant destory
func (imp *PluginDeviceControllerImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *PluginDeviceControllerImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *PluginDeviceControllerImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
