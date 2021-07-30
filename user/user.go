package user

import (
	"fmt"
	"net/http"
)

type user struct {
	Id   int64
	Name string
	Sex  bool
	Age  string
	Addr string
}

var inituser = []user{
	{1, "zz", true, "20", "重庆"},
	{2, "xx", false, "40", "天津"},
	{3, "hh", true, "31", "上海"},
	{4, "kk", false, "22", "北京"},
}

func NewUser() *[]user {
	return &inituser
}

func Adduser(rw http.ResponseWriter, r *http.Request, name, age, addr string, sex bool) {
	for i := 0; i < len(inituser); i++ {
		if inituser[i].Name == name {
			http.Redirect(rw, r, "/usererr/", 302)
			break
		}
		if len(inituser)-1 == i {
			inituser = append(inituser, user{
				int64(len(inituser) + 1),
				name,
				sex,
				age,
				addr,
			})
			break
		}

	}
}

func DelUser(rw http.ResponseWriter, r *http.Request, id int64) {
	NewUser := make([]user, 0)
	for _, users := range inituser {
		if users.Id == id {
			continue
		}
		NewUser = append(NewUser, users)
	}
	inituser = NewUser
	http.Redirect(rw, r, "/", 302)
}

func JudgeId(id int64) *int64 {
	var ID int64
	for i := 0; i < len(inituser); i++ {
		if inituser[i].Id == id {
			ID = id
			break
		}
	}
	return &ID
}

func Edit(rw http.ResponseWriter, r *http.Request, name, age, addr string, sex bool, Id *int) {

	fmt.Println(*Id)
	inituser[*Id].Name = name
	inituser[*Id].Age = age
	inituser[*Id].Addr = addr
	inituser[*Id].Sex = sex
	http.Redirect(rw, r, "/", 302)
}
