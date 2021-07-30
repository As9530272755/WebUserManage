package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"webUser/user"
)

var (
	Addr string
)

func Web() {
	Nusers := user.NewUser()

	Addr = ":8888"
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("template/userlist.html")
		if err != nil {
			log.Println(err)
		}
		tpl.ExecuteTemplate(rw, "userlist.html", Nusers)
	})

	http.HandleFunc("/create/", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tpl, err := template.ParseFiles("template/create.html")
			if err != nil {
				log.Printf("create err: %v", err)
			}
			tpl.ExecuteTemplate(rw, "create.html", nil)
		} else {
			name := r.FormValue("name")
			sex := r.FormValue("sex") == "1"
			age := r.FormValue("age")
			addr := r.FormValue("addr")
			user.Adduser(rw, r, name, age, addr, sex)
		}
		http.Redirect(rw, r, "/", 302)
	})

	http.HandleFunc("/delete/", func(rw http.ResponseWriter, r *http.Request) {
		if Id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err == nil {
			user.DelUser(rw, r, Id)
		}
	})

	http.HandleFunc("/usererr/", func(rw http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("template/usererr.html")
		if err != nil {
			log.Println(err)
			return
		}
		tpl.ExecuteTemplate(rw, "usererr.html", Nusers)
	})

	http.HandleFunc("/edit/", func(rw http.ResponseWriter, r *http.Request) {

		id := 0
		if Id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err == nil {
			user.JudgeId(Id)
			id = int(Id)
		}
		// fmt.Println(open)

		if r.Method == "GET" {
			tpl, err := template.ParseFiles("template/edit.html")
			if err != nil {
				log.Printf("create err: %v", err)
			}
			tpl.ExecuteTemplate(rw, "edit.html", nil)
		} else {
			fmt.Println("opst")
			name := r.FormValue("name")
			age := r.FormValue("age")
			addr := r.FormValue("addr")
			sex := r.FormValue("sex") == "1"
			user.Edit(rw, r, name, age, addr, sex, &id)
		}

	})

	http.ListenAndServe(Addr, nil)

}
