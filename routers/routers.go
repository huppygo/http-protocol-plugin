package routers

import (
	"errors"
	"http-procotol-plugin/controller"
	"http-procotol-plugin/global"
	"http-procotol-plugin/service"
	"strings"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.Use(IsAuthDevice())
	c := router.Group("api")
	tp := &controller.TpController{}
	{
		//提供给tp平台的接口
		c.GET("form/config", tp.Config)                 //获取插件表单配置
		c.POST("device/config/update", tp.UpdateDevice) //修改设备表单配置
		c.POST("device/config/add", tp.AddDevice)       //新增网关子设备
		c.POST("device/config/delete", tp.DeleteDevice) //删除设备配
		//提供给设备的接口
		c.POST("device/attributes", tp.Attributes)      //属性上报
		c.POST("device/event", tp.Event)                //事件上报
		c.POST("device/command/reply", tp.CommandReply) //命令执行结果上报

	}
}

func IsAuthDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		if !strings.Contains(url, "config") {
			//accesstoken := c.Param("accesstoken")
			body, _ := ioutil.ReadAll(c.Request.Body)
			bodyJson := make(map[string]interface{})
			if err := json.Unmarshal(body, &bodyJson); err != nil {
				log.Println("json转换失败", err)
			//	return err
			}
			accesstoken := bodyJson["imei"].(string)
			if _, ok := global.DevicesMap.Load(accesstoken); !ok {
				if err := service.TpDeviceAccessToken(accesstoken); err != nil {
					c.AbortWithError(401, errors.New("device is unauth"))
				}
			}
		}
	}
}
