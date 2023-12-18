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
	"fmt"
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

//解析函数
func processBodyJson(bodyJson []map[string]interface{}) map[string]interface{} { 
	bodyData := make(map[string]interface{}) 
	//log.Println("bodyJson Process", bodyJson)
	 for i := 0; i < len(bodyJson); i++{
	    //log.Println("bodyJson[i]:",bodyJson[i])
	    if imei, exist := bodyJson[i]["imei"].(string); exist {
	        bodyData["imei"] = imei
	    }
	    if local, exist := bodyJson[i]["local"].(string); exist {
	        bodyData["local"] = local
	    }
		if recTime, exist := bodyJson[i]["recTime"]; exist {  
			if recTimeFloat, ok := recTime.(float64); ok {  
				recTimeString := fmt.Sprintf("%.0f", recTimeFloat)  
				bodyData["recTime"] = recTimeString  
			} else {  
				// Handle error or log a warning  
			}  
		}
	    if male, exist := bodyJson[i]["male"].(string); exist {
            bodyData["male"] = male
	    }
	    if sys, exist := bodyJson[i]["sys"].(string); exist {
	        if "000" != sys {
	            bodyData["sys"] = sys
	        }else{
				if _, ok := bodyData["sys"]; !ok {  
					bodyData["sys"] = sys  
				}
			}
	    }
		if dia, exist := bodyJson[i]["dia"].(string); exist {
	        if "000" != dia {
	            bodyData["dia"] = dia
	        }else{
				if _, ok := bodyData["dia"]; !ok {  
					bodyData["dia"] = dia  
				}
			}
	    }
	    if pul, exist := bodyJson[i]["pul"].(string); exist {
	        if "000" != pul {
	            bodyData["pul"] = pul
			}else{
				if _, ok := bodyData["pul"]; !ok {  
					bodyData["pul"] = pul  
				}
			}
	    }
	    if eat, exist := bodyJson[i]["eat"].(string); exist {
	        bodyData["eat"] = eat
	    }
	    if glu, exist := bodyJson[i]["glu"].(string); exist {
	        if "0.0" != glu {
	            bodyData["glu"] = glu
			}else{
				if _, ok := bodyData["glu"]; !ok {  
					bodyData["glu"] = glu  
				}
			}
	    }
	    if cho, exist := bodyJson[i]["cho"].(string); exist {
	        if "0.0" != cho {
	            bodyData["cho"] = cho
			}else{
				if _, ok := bodyData["cho"]; !ok {  
					bodyData["cho"] = cho  
				}
			}
	    }
	    if tri, exist := bodyJson[i]["tri"].(string); exist {
	        if "0.00" != tri {
	            bodyData["tri"] = tri
			}else{
				if _, ok := bodyData["tri"]; !ok {  
					bodyData["tri"] = tri  
				}
			}
	    }
	    if uri, exist := bodyJson[i]["uri"].(string); exist {
	        if "000" != uri {
	            bodyData["uri"] = uri
			}else{
				if _, ok := bodyData["uri"]; !ok {  
					bodyData["uri"] = uri  
				}
			}
	    }
	    if xy, exist := bodyJson[i]["xy"].(string); exist {
	        if "000" != xy {
	            bodyData["xy"] = xy
			}else{
				if _, ok := bodyData["xy"]; !ok {  
					bodyData["xy"] = xy  
				}
			}
	    }
	    if hr, exist := bodyJson[i]["hr"].(string); exist {
	        if "000" != hr {
	            bodyData["hr"] = hr
			}else{
				if _, ok := bodyData["hr"]; !ok {  
					bodyData["hr"] = hr  
				}
			}
	    }
	    if tw, exist := bodyJson[i]["tw"].(string) ; exist {
	        if "000" != tw {
	            bodyData["tw"] = tw
			}else{
				if _, ok := bodyData["tw"]; !ok {  
					bodyData["tw"] = tw  
				}
			}
	    }
	} 
	return bodyData;
}

//接收属性
func (w *TpController) Attributes(ctx *gin.Context) {
	//accesstoken := ctx.Param("accesstoken")
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	log.Println("bodyReceive:",body)
	// 尝试解析为数组  
	var bodyJson []map[string]interface{}  
	DecodeBodyData := make(map[string]interface{})
	if err := json.Unmarshal([]byte(body), &bodyJson); err == nil {  
		//fmt.Println("body 是数组") 		
		//调用解析函数
		DecodeBodyData = processBodyJson(bodyJson)
	} else {  
		// 如果解析为数组失败，尝试解析为对象  
		var bodyMap map[string]interface{}  
		if err := json.Unmarshal([]byte(body), &bodyMap); err == nil {  
			DecodeBodyData = bodyMap
		} else {  
			log.Println("无法解析 body，可能不是有效的 JSON")  
			Response.Failed(ctx)
		}  
	}  

	accesstoken := DecodeBodyData["imei"].(string)
	log.Println("DecodeBodyData:",DecodeBodyData)
	bodyOut, err := json.Marshal(DecodeBodyData)
	if err == nil{
	}else{
		//输出错误
		log.Println("json转换失败", err)
		Response.Failed(ctx)
	}
	//log.Println("bodyOut:",bodyOut)
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

	// 尝试解析为数组  
	var bodyJson []map[string]interface{}  
	DecodeBodyData := make(map[string]interface{})
	if err := json.Unmarshal([]byte(body), &bodyJson); err == nil {  
		//fmt.Println("body 是数组") 		
		//调用解析函数
		DecodeBodyData = processBodyJson(bodyJson)
	} else {  
		// 如果解析为数组失败，尝试解析为对象  
		var bodyMap map[string]interface{}  
		if err := json.Unmarshal([]byte(body), &bodyMap); err == nil {  
			DecodeBodyData = bodyMap
		} else {  
			log.Println("无法解析 body，可能不是有效的 JSON")  
			Response.Failed(ctx)
		}  
	}  

	accesstoken := DecodeBodyData["imei"].(string)
	//log.Println("DecodeBodyData:",DecodeBodyData)
	bodyOut, err := json.Marshal(DecodeBodyData)
	if err == nil{
	}else{
	    //输出错误
	    log.Println("json转换失败", err)
		Response.Failed(ctx)
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

	// 尝试解析为数组  
	var bodyJson []map[string]interface{}  
	DecodeBodyData := make(map[string]interface{})
	if err := json.Unmarshal([]byte(body), &bodyJson); err == nil {  
		//fmt.Println("body 是数组") 		
		//调用解析函数
		DecodeBodyData = processBodyJson(bodyJson)
	} else {  
		// 如果解析为数组失败，尝试解析为对象  
		var bodyMap map[string]interface{}  
		if err := json.Unmarshal([]byte(body), &bodyMap); err == nil {  
			DecodeBodyData = bodyMap
		} else {  
			log.Println("无法解析 body，可能不是有效的 JSON")  
			Response.Failed(ctx)
		}  
	}  

	accesstoken := DecodeBodyData["imei"].(string)
	//log.Println("DecodeBodyData:",DecodeBodyData)
	bodyOut, err := json.Marshal(DecodeBodyData)
	if err == nil{
	}else{
	    //输出错误
	    log.Println("json转换失败", err)
		Response.Failed(ctx)
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
