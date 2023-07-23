package controllers

import (
	"encoding/json"
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type Reselfts struct {
	reselft []string
}

type UploadImageController struct {
	beego.Controller
}

type GetImagePathController struct {
	beego.Controller
}

type GetImagePathInUseController struct {
	beego.Controller
}

type UseHeadImageController struct {
	beego.Controller
}

func (c *UploadImageController) Post() {
	id, _ := c.GetInt("userid")
	_, head, err := c.GetFile("uploadfile")
	username := c.GetString("username")
	if err != nil {
		fmt.Println("read file err   :", err, "\n")
		c.Ctx.WriteString("{'code':400,'massage':'failed'}")
	}
	fileExt := path.Ext(head.Filename)
	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg" {
		c.Ctx.WriteString("{'code':400,'massage':'file is not png,jpg or jpeg'}")
	}
	if head.Size > 5000000 {
		c.Ctx.WriteString("{'code':400,'massage':'file is too large'}")
	}
	fileName := time.Now().Unix()
	err = c.SaveToFile("uploadfile", "static/img/"+strconv.Itoa(int(fileName))+fileExt)
	if err != nil && uploadImageData(id, username, strconv.Itoa(int(fileName))+fileExt) {
		fmt.Println("save err   :", err, "\n")
		c.Ctx.WriteString("{'code':400,'massage':'failed'}")
	} else {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	}
}

func (c *GetImagePathController) Get() {
	var reselfts Reselfts
	id, _ := c.GetInt("userid")
	re := getImagePathByUserId(id)
	reselfts.reselft = re
	jsonByte, _ := json.Marshal(reselfts.reselft)
	jsonStr := string(jsonByte)
	if re != nil {
		c.Ctx.WriteString("{'code':200,'massage':" + jsonStr + "}")
	} else {
		c.Ctx.WriteString("{'code':400,'massage':'failed or not img'}")
	}
}

func (c *GetImagePathInUseController) Get() {
	var reselfts Reselfts
	id, _ := c.GetInt("userid")
	re := getImagePathByUserIdInUse(id)
	reselfts.reselft = re
	jsonByte, _ := json.Marshal(reselfts.reselft)
	jsonStr := string(jsonByte)
	if re != nil {
		c.Ctx.WriteString("{'code':200,'massage':" + jsonStr + "}")
	} else {
		c.Ctx.WriteString("{'code':400,'massage':'failed or not img'}")
	}
}

func (c *UseHeadImageController) Put() {
	imgPath := c.GetString("imgpath")
	userid, _ := c.GetInt("userid")
	if useHeadImage(imgPath, userid) {
		c.Ctx.WriteString("{'code':200,'massage':'successfully'}")
	} else {
		c.Ctx.WriteString("{'code':400,'massage':'failed'}")
	}
}
