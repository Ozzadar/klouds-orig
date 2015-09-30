package controllers

import (
	"fmt"
	"io/ioutil"
	"github.com/astaxie/beego"
	"net/http"
	"bytes"
)

type MainController struct {
	beego.Controller
}

type LaunchController struct {
	beego.Controller
}

func (this *MainController) activeContent(view string) {
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Sidebar"] = "sidebar.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplNames = view + ".tpl"

	sess := this.GetSession("acme")
	if sess != nil {
		this.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		this.Data["First"] = m["first"]
	}
}

func (this *MainController) Get() {
	this.activeContent("index")

	//******** This page requires login
	sess := this.GetSession("acme")
	if sess == nil {
		this.Redirect("/user/login/home", 302)
		return
	}
	m := sess.(map[string]interface{})
	fmt.Println("username is", m["username"])
	fmt.Println("logged in at", m["timestamp"])
}

func (this *MainController) Notice() {
	this.activeContent("notice")

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = n
	}
}

func (this *LaunchController) activeContent(view string) {
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Sidebar"] = "sidebar.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	fileselect := "<select name='filelist' form='form1'>"

	files, err := ioutil.ReadDir("./restcalls")
	if err != nil {

		fmt.Println(err)
	}
    for _, f := range files {
    		fileselect= fileselect + "<option>" + f.Name() + "</option>"
    }

    fileselect = fileselect + "</select>"

 

	this.TplNames = view + ".tpl"
    this.Data["FileSelect"] = fileselect
	sess := this.GetSession("acme")
	if sess != nil {
		this.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		this.Data["First"] = m["first"]
	}
}

func (this *LaunchController) Get() {
	this.activeContent("appLaunch")

	//******** This page requires login
	sess := this.GetSession("acme")
	if sess == nil {
		this.Redirect("/user/login/home", 302)
		return
	}
	m := sess.(map[string]interface{})
	fmt.Println("username is", m["username"])
	fmt.Println("logged in at", m["timestamp"])
}

func (this *LaunchController) Notice() {
	this.activeContent("appLaunch")

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = n
	}
}

func (this *LaunchController) Post() {
	this.activeContent("appLaunch")

	file := this.GetString("filelist")

	thethings, err := ioutil.ReadFile("./restcalls/" + file)

	if err != nil {
		panic(err)
	}
	url := "http://localhost:8001/api/v1/namespaces/default/pods"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(thethings))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
        panic(err) //Something is wrong while sending request
    }

    if res.StatusCode != 201 {
        fmt.Printf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
    }

    fmt.Println(res.StatusCode)

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = n
	}
}
