package models

import (
	"strconv"
	"fmt"
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/golyu/sql-build-example/models/build"
)

type User struct {
	Id         int `json:"id" orm:"column(id)" insert:"id"`
	Username   string `json:"username" orm:"column(username)" insert:"username"`
	Sex        string `json:"sex" orm:"column(sex)" insert:"sex"`
	Password   string `json:"password" orm:"column(password)" insert:"password"`
	Name       string `json:"name" orm:"column(name)" insert:"name"`
	AddTime    string `json:"add_time" orm:"column(add_time)" insert:"add_time"`
	State      int `json:"state" orm:"column(state)" insert:"state"`
	Money      float64 `json:"money" orm:"column(money)" insert:"money"`
	RegAddress string `json:"reg_address" orm:"column(reg_address)" insert:"reg_address"`
	Qq         string `json:"qq" orm:"column(qq)" insert:"qq"`
	Email      string `json:"email" orm:"column(email)" insert:"email"`
}

func CreateUser() User {
	return User{
		Id:         -1,
		Username:   "",
		Sex:        "",
		Password:   "",
		Name:       "",
		AddTime:    "",
		State:      -1,
		Money:      -1,
		RegAddress: "",
		Qq:         "",
		Email:      "",
	}
}

//主要为了对比下面的UpdateUser方法
func UpdateUserTemp(id int, username, sex, name string, state int, money float64, qq, email string) error {
	if id > 0 {
		set := ""
		if len(username) > 0 {
			set += " name = '" + username + "'"
		}
		if len(sex) > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " sex = '" + sex + "'"
		}
		if len(name) > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " name = '" + name + "'"
		}
		if state > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " state = " + strconv.Itoa(state)
		}
		if money > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " money = " + fmt.Sprintf("%g", money)
		}
		if len(qq) > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " qq = '" + qq + "'"
		}
		if len(email) > 0 {
			if len(set) > 0 {
				set += ","
			}
			set += " email = '" + email + "'"
		}
		if len(set) > 0 {
			set = " SET" + set
		} else {
			return errors.New("not need update column")
		}
		sql := fmt.Sprintf("UPDATE user "+set+" WHERE id = %d",
			id)
		o := orm.NewOrm()
		_, err := o.Raw(sql).Exec()
		return err
	}
	return errors.New("Id can not <=0")
}

func (m *User) Insert() error {
	sql, err := build.NewInsert("user").
		Value(m).String()
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.Raw(sql).Exec()
	return err
}

func UpdateUser(id int, username, sex, name string, state int, money float64, qq, email string) error {
	sql,err:=build.NewUpdate("user").
	Where_(id,"id").
	Set(username,"username").
	Set(sex,"sex").
	Set(name,"name").
	Set(state,"state").
	Set(money,"money").
	Set(qq,"qq").
	Set(email,"email").String()
	if err != nil {
		return err
	}
	o:=orm.NewOrm()
	_,err = o.Raw(sql).Exec()
	return err
}
func SelectUser(username, sex, name string, state int, money float64, qq, email string) (User, error) {
	sql, err := build.NewSelect("user").
		Where_(username, "username").
		Where(sex, "sex").
		Where(name, "name").
		Where(state, "state").
		Where(money, "money").
		Where(qq, "qq").
		Where(email, "email").
		String()
	if err != nil {
		return User{}, err
	}
	o := orm.NewOrm()
	var user User
	err = o.Raw(sql).QueryRow(&user)
	return user, err
}