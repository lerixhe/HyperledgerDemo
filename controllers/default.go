package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// get方法：请求key对应的value
func (c *MainController) GetValue() {
	// 从请求中获取"key"对于的key
	key := c.GetString("key")
	// 处理未获取到key
	if key == "" {
		// 做出失败回复
		handleResponse(c, 400, "Request parameter key can't be empty!")
		return
	}
	logs.Info("key:" + key)

	// TODO，与区块链交互，取得返回值，写入response

	var response []byte
	// 做出成功回复
	handleResponse(c, 200, response)

}

/*
定义处理数据回复的方法
*/
func handleResponse(c *MainController, code int, msg interface{}) {
	//  判断状态码
	if code >= 400 {
		logs.Error(msg)
	} else {
		logs.Info(msg)
	}
	// 创建状态码对应的响应头
	c.Ctx.ResponseWriter.WriteHeader(code)
	//  将信息也写入响应，注意需要断言信息类型：错误响应是string，正确响应是字节数组，而写入响应的需要是字节数据
	b, ok := msg.([]byte)
	if ok {
		c.Ctx.ResponseWriter.Write(b)
	} else {
		s := msg.(string)
		c.Ctx.ResponseWriter.Write([]byte(s))
	}

}
