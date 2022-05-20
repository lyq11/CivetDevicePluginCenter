package main

import (
	instance "CivetRedis/CivetRedis/BaseInstance"
	"Civets/CivetPluginsFrameWork"
	"Civets/CivetRedis"
	"Civets/CivetRedis/BaseInstance"
	"context"
	"encoding/json"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"plugin"
)

// ProcessImp servant implementation
type ProcessImp struct {
}

func (imp *ProcessImp) Reboot(tarsCtx context.Context, c *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *ProcessImp) Shutdown(tarsCtx context.Context, c *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func running(ctx context.Context, consumerName string) {
	for true {
		result, err := newConsumer.ReadFromGroupQueue(ctx, consumerName, 1, 0)
		fmt.Println(err)
		fmt.Println("the stream is", result[0].Stream)
		fmt.Println("the msg is", result[0].Messages)
		fmt.Println("the id is", result[0].Messages[0].ID)
		value := result[0].Messages[0].Values
		for s, i := range value {
			//fmt.Println("the value is", s)
			//fmt.Println("the value is", i.(string))
			_, err := PFW.Cal(s, "Analyst", i.(string))
			if err != nil {
				tars.TLOG.Error(err)
			}
		}

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

// Init servant init
var PFW *CivetPluginsFrameWork.Plugins
var newConsumer *CivetRedis.Consumer
var PlugsMapper map[string]plugin.Plugin
var Plugs map[string]string

func (imp *ProcessImp) Init() error {
	var TarsConfig = tars.NewRConf("CivetDevicePluginCenter", "PluginManagerServer", tars.GetServerConfig().BasePath)
	config, _ := TarsConfig.GetConfig("PluginManagerServer.conf")
	var mc Config
	json.Unmarshal([]byte(config), &mc)

	PFW = CivetPluginsFrameWork.GetIns()

	redisconf := BaseInstance.RedisBaseConfig{
		Host:     "127.0.0.1",
		Port:     "6379",
		Password: "",
		UserName: "",
		Db:       0,
		Size:     20,
	}
	ctx := context.Background()
	newConsumer = CivetRedis.CreateConsumer((*instance.RedisBaseConfig)(&redisconf), "CivetDevicePluginCenter", "MessageQueue")
	newConsumer.CreateGroup(ctx, "$")
	for i := 0; i < 2; i++ {
		newConsumer.CreateGroupConsumer(ctx, fmt.Sprintf("Process%d", i))
		go running(ctx, fmt.Sprintf("Process%d", i))
	}
	return nil
}

// Destroy servant destroy
func (imp *ProcessImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *ProcessImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *ProcessImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
