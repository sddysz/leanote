package controllers

import (
	"github.com/revel/revel"
	//	"encoding/json"
	//	"gopkg.in/mgo.v2/bson"
	. "github.com/sddysz/leanote/app/lea"
	"github.com/sddysz/leanote/app/lea/captcha"
	//	"github.com/sddysz/leanote/app/types"
	//	"io/ioutil"
	//	"fmt"
	//	"math"
	//	"os"
	//	"path"
	//	"strconv"
	"net/http"
)

// 验证码服务
type Captcha struct {
	BaseController
}

type Ca string

func (r Ca) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "image/png")
}

func (c Captcha) Get() revel.Result {
	c.Response.ContentType = "image/png"
	image, str := captcha.Fetch()
	image.WriteTo(c.Response.Out)

	sessionId := c.Session["_ID"]
	//	LogJ(c.Session)
	//	Log("------")
	//	Log(str)
	//	Log(sessionId)
	Log("..")
	sessionService.SetCaptcha(sessionId, str)

	return c.Render()
}
