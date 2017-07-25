package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golyu/sql-build-example/models"
	"encoding/json"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (c *UserController) FindUser() {
	//username, sex, name, state, money, qq, email string
	username := c.GetString("username")
	sex := c.GetString("username")
	name := c.GetString("name")
	state, _ := c.GetInt("state",-1)
	money, _ := c.GetFloat("money",-1)
	qq := c.GetString("qq")
	email := c.GetString("emial")
	user, err := models.SelectUser(username, sex, name, state, money, qq, email)
	if err != nil {
		c.EchoJSON("1103", "find err", err.Error())
	} else {
		c.EchoJSON("1101", "success", user)
	}
}
func (c *UserController) AddUser() {
	user  := models.CreateUser()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.EchoJSON("1102", "params err", err.Error())
	}
	err = user.Insert()
	if err != nil {
		c.EchoJSON("1103", "insert err", err.Error())
	} else {
		c.EchoJSON("1101", "success", err.Error())
	}
}
func (c *UserController) UpdateUser() {
	id, _ := c.GetInt("id",-1)
	username := c.GetString("username")
	sex := c.GetString("username")
	name := c.GetString("name")
	state, _ := c.GetInt("state",-1)
	money, _ := c.GetFloat("money",-1)
	qq := c.GetString("qq")
	email := c.GetString("emial")
	err := models.UpdateUser(id, username, sex, name, state, money, qq, email)
	if err != nil {
		c.EchoJSON("1103", "update err", err.Error())
	} else {
		c.EchoJSON("1101", "success", "")
	}
}

func (c *UserController) EchoJSON(code, msg string, data interface{}) {
	var jsonData = make(map[string]interface{})
	jsonData["Code"] = code
	jsonData["Msg"] = msg
	jsonData["Data"] = data
	c.Data["json"] = jsonData
	c.ServeJSON()
}
