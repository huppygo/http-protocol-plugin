package controller

import (
	"http-procotol-plugin/service"
	"http-procotol-plugin/utils"
	"http-procotol-plugin/global"
	"io/ioutil"
	"net/http"
	"time"
	"log"
	"errors"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//相关视图，跳转至services里对应服务
type TpController struct{}

var Response = utils.Response

//获取表单配置
func (w *TpController) Config(ctx *gin.Context) {
	if data, err := service.TpSer.Config(); err != nil || data == "" {
		Response.Failed(ctx)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "404",
			"ts":   time.Now().UnixMicro(),
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "200",
			"data": data,
			"ts":   time.Now().UnixMicro(),
		})
	}
}

//更新设备
func (w *TpController) UpdateDevice(ctx *gin.Context) {
	var device utils.Device
	_ = ctx.ShouldBindJSON(&device)
	if err := service.TpSer.UpdateDevice(device); err != nil {
		Response.Failed(ctx)
	} else {
		Response.OK(ctx)
	}
}

//新增设备
func (w *TpController) AddDevice(ctx *gin.Context) {
	var device utils.Device
	err := ctx.ShouldBindJSON(&device)
	if err != nil {
		Response.Failed(ctx)
	}
	if err := service.TpSer.AddDevice(device); err != nil {
		Response.Failed(ctx)
	} else {
		Response.OK(ctx)
	}
}

//删除设备
func (w *TpController) DeleteDevice(ctx *gin.Context) {
	var device utils.Device
	_ = ctx.ShouldBindJSON(&device)
	if err := service.TpSer.DeleteDevice(device); err != nil {
		Response.Failed(ctx)
	} else {
		Response.OK(ctx)
	}

}

//接收属性
func (w *TpController) Attributes(ctx *gin.Context) {
	//accesstoken := ctx.Param("accesstoken")
	//body, _ := ioutil.ReadAll(ctx.Request.Body)
	//bodyJson := make([]map[string]interface{})
	var body = `[{"cho":"0.0","dia":"0.0","imei":"864383063591361","glu":"11.1"}, {"cho":"1.0","dia":"097","imei":"864383063591361","glu":"0.0"}, {"cho":"2.0","dia":"097","imei":"864383063591361","glu":"0.0"}]`
	var bodyJson []map[string]interface{}
	if err := json.Unmarshal([]byte(body), &bodyJson); err != nil {
		log.Println("json转换失败", err)
	//	return err
	}
	log.Println("body:",body)
	log.Println("bodyJson:",bodyJson)

	bodyData := make(map[string]interface{})
	bodyObject_counts := len(bodyJson)
	for i := 0; i < bodyObject_counts; i++{
	    log.Println("bodyJson[i]:",bodyJson[i])
	    if imei, exist := bodyJson[i]["imei"]; exist {
	        bodyData["imei"] = imei.(string)
	    }
	    if local, exist := bodyJson[i]["local"]; exist {
	        bodyData["local"] = local.(string)
	    }
	    if recTime, exist := bodyJson[i]["recTime"]; exist {
	        bodyData["recTime"] = recTime.(string)
	    }
	    if male, exist := bodyJson[i]["male"]; exist {
            bodyData["male"] = male.(string)
	    }
	    if sys, exist := bodyJson[i]["sys"]; exist {
	        if "000" != sys.(string) {
	            bodyData["sys"] = sys.(string)
	        }
	    }
	    if dia, exist := bodyJson[i]["dia"]; exist {
	        if "000" != dia.(string) {
	            bodyData["dia"] = dia.(string)
	        }
	    }
	    if pul, exist := bodyJson[i]["pul"]; exist {
	        if "000" != pul.(string) {
	            bodyData["pul"] = pul.(string)
	        }
	    }
	    if eat, exist := bodyJson[i]["eat"]; exist {
	        bodyData["eat"] = eat.(string)
	    }
	    if glu, exist := bodyJson[i]["glu"]; exist {
	        if "0.0" != glu.(string) {
	            bodyData["glu"] = glu.(string)
	        }
	    }
	    if cho, exist := bodyJson[i]["cho"]; exist {
	        if "0.0" != cho.(string) {
	            bodyData["cho"] = cho.(string)
	        }
	    }
	    if tri, exist := bodyJson[i]["tri"]; exist {
	        if "0.0" != tri.(string) {
	            bodyData["tri"] = tri.(string)
	        }
	    }
	    if uri, exist := bodyJson[i]["uri"]; exist {
	        if "000" != uri.(string) {
	            bodyData["uri"] = uri.(string)
	        }
	    }
	    if xy, exist := bodyJson[i]["xy"]; exist {
	        if "000" != xy.(string) {
	            bodyData["xy"] = xy.(string)
	        }
	    }
	    if hr, exist := bodyJson[i]["hr"]; exist {
	        if "000" != hr.(string) {
	            bodyData["hr"] = hr.(string)
	        }
	    }
	    if tw, exist := bodyJson[i]["tw"]; exist {
	        if "000" != tw.(string) {
	            bodyData["tw"] = tw.(string)
	        }
	    }
	}

	accesstoken := bodyData["imei"].(string)
	bodyOut, err := json.Marshal(bodyData)
	if err == nil{
	}else{
	    //输出错误
	    log.Println("json转换失败", err)
	}
	log.Println("bodyOut:",bodyOut)
	//  查找是否注册设备
	if _, ok := global.DevicesMap.Load(accesstoken); ok {
		if err := service.TpSer.Attributes(accesstoken, bodyOut); err != nil {
			Response.Failed(ctx)
		} else {
			Response.OK(ctx)
		}
	} else {
		if err := service.TpDeviceAccessToken(accesstoken); err != nil {
			ctx.AbortWithError(401, errors.New("device is unauth"))
		}
	}
}

//接收事件数据
func (w *TpController) Event(ctx *gin.Context) {
	//accesstoken := ctx.Param("accesstoken")
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	//bodyJson := make([]map[string]interface{})
	var bodyJson []map[string]interface{}
	if err := json.Unmarshal([]byte(body), &bodyJson); err != nil {
		log.Println("json转换失败", err)
	//	return err
	}

	bodyData := make(map[string]interface{})
	bodyObject_counts := len(bodyJson)
	for i := 0; i < bodyObject_counts; i++{
	    
	    if imei, exist := bodyJson[i]["imei"]; exist {
	        bodyData["imei"] = imei.(string)
	    }
	    if local, exist := bodyJson[i]["local"]; exist {
	        bodyData["local"] = local.(string)
	    }
	    if recTime, exist := bodyJson[i]["recTime"]; exist {
	        bodyData["recTime"] = recTime.(string)
	    }
	    if male, exist := bodyJson[i]["male"]; exist {
            bodyData["male"] = male.(string)
	    }
	    if sys, exist := bodyJson[i]["sys"]; exist {
	        if "000" != sys.(string) {
	            bodyData["sys"] = sys.(string)
	        }
	    }
	    if dia, exist := bodyJson[i]["dia"]; exist {
	        if "000" != dia.(string) {
	            bodyData["dia"] = dia.(string)
	        }
	    }
	    if pul, exist := bodyJson[i]["pul"]; exist {
	        if "000" != pul.(string) {
	            bodyData["pul"] = pul.(string)
	        }
	    }
	    if eat, exist := bodyJson[i]["eat"]; exist {
	        bodyData["eat"] = eat.(string)
	    }
	    if glu, exist := bodyJson[i]["glu"]; exist {
	        if "0.0" != glu.(string) {
	            bodyData["glu"] = glu.(string)
	        }
	    }
	    if cho, exist := bodyJson[i]["cho"]; exist {
	        if "0.0" != cho.(string) {
	            bodyData["cho"] = cho.(string)
	        }
	    }
	    if tri, exist := bodyJson[i]["tri"]; exist {
	        if "0.0" != tri.(string) {
	            bodyData["tri"] = tri.(string)
	        }
	    }
	    if uri, exist := bodyJson[i]["uri"]; exist {
	        if "000" != uri.(string) {
	            bodyData["uri"] = uri.(string)
	        }
	    }
	    if xy, exist := bodyJson[i]["xy"]; exist {
	        if "000" != xy.(string) {
	            bodyData["xy"] = xy.(string)
	        }
	    }
	    if hr, exist := bodyJson[i]["hr"]; exist {
	        if "000" != hr.(string) {
	            bodyData["hr"] = hr.(string)
	        }
	    }
	    if tw, exist := bodyJson[i]["tw"]; exist {
	        if "000" != tw.(string) {
	            bodyData["tw"] = tw.(string)
	        }
	    }
	}

	accesstoken := bodyData["imei"].(string)
	bodyOut, err := json.Marshal(bodyData)
	if err == nil{
	}else{
	    //输出错误
	    log.Println("json转换失败", err)
	}

	//  查找是否注册设备
	if _, ok := global.DevicesMap.Load(accesstoken); ok {
		if err := service.TpSer.Event(accesstoken, bodyOut); err != nil {
			Response.Failed(ctx)
		} else {
			Response.OK(ctx)
		}
	} else {
		if err := service.TpDeviceAccessToken(accesstoken); err != nil {
			ctx.AbortWithError(401, errors.New("device is unauth"))
		}
	}
}

//接收命令响应数据
func (w *TpController) CommandReply(ctx *gin.Context) {
	//accesstoken := ctx.Param("accesstoken")
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	//bodyJson := make([]map[string]interface{})
	var bodyJson []map[string]interface{}
	if err := json.Unmarshal([]byte(body), &bodyJson); err != nil {
		log.Println("json转换失败", err)
	//	return err
	}
	
	bodyData := make(map[string]interface{})
	bodyObject_counts := len(bodyJson)
	for i := 0; i < bodyObject_counts; i++{
	    
	    if imei, exist := bodyJson[i]["imei"]; exist {
	        bodyData["imei"] = imei.(string)
	    }
	    if local, exist := bodyJson[i]["local"]; exist {
	        bodyData["local"] = local.(string)
	    }
	    if recTime, exist := bodyJson[i]["recTime"]; exist {
	        bodyData["recTime"] = recTime.(string)
	    }
	    if male, exist := bodyJson[i]["male"]; exist {
            bodyData["male"] = male.(string)
	    }
	    if sys, exist := bodyJson[i]["sys"]; exist {
	        if "000" != sys.(string) {
	            bodyData["sys"] = sys.(string)
	        }
	    }
	    if dia, exist := bodyJson[i]["dia"]; exist {
	        if "000" != dia.(string) {
	            bodyData["dia"] = dia.(string)
	        }
	    }
	    if pul, exist := bodyJson[i]["pul"]; exist {
	        if "000" != pul.(string) {
	            bodyData["pul"] = pul.(string)
	        }
	    }
	    if eat, exist := bodyJson[i]["eat"]; exist {
	        bodyData["eat"] = eat.(string)
	    }
	    if glu, exist := bodyJson[i]["glu"]; exist {
	        if "0.0" != glu.(string) {
	            bodyData["glu"] = glu.(string)
	        }
	    }
	    if cho, exist := bodyJson[i]["cho"]; exist {
	        if "0.0" != cho.(string) {
	            bodyData["cho"] = cho.(string)
	        }
	    }
	    if tri, exist := bodyJson[i]["tri"]; exist {
	        if "0.0" != tri.(string) {
	            bodyData["tri"] = tri.(string)
	        }
	    }
	    if uri, exist := bodyJson[i]["uri"]; exist {
	        if "000" != uri.(string) {
	            bodyData["uri"] = uri.(string)
	        }
	    }
	    if xy, exist := bodyJson[i]["xy"]; exist {
	        if "000" != xy.(string) {
	            bodyData["xy"] = xy.(string)
	        }
	    }
	    if hr, exist := bodyJson[i]["hr"]; exist {
	        if "000" != hr.(string) {
	            bodyData["hr"] = hr.(string)
	        }
	    }
	    if tw, exist := bodyJson[i]["tw"]; exist {
	        if "000" != tw.(string) {
	            bodyData["tw"] = tw.(string)
	        }
	    }
	}

	accesstoken := bodyData["imei"].(string)
	bodyOut, err := json.Marshal(bodyData)
	if err == nil{
	}else{
	    //输出错误
	    log.Println("json转换失败", err)
	}
	
	//  查找是否注册设备
	if _, ok := global.DevicesMap.Load(accesstoken); ok {
		if err := service.TpSer.CommandReply(accesstoken, bodyOut); err != nil {
			Response.Failed(ctx)
		} else {
			Response.OK(ctx)
		}
	} else {
		if err := service.TpDeviceAccessToken(accesstoken); err != nil {
			ctx.AbortWithError(401, errors.New("device is unauth"))
		}
	}
}
