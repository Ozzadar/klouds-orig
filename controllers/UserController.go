package controllers

import (
	"net/http"
	"github.com/superordinate/klouds2.0/models"
	"gopkg.in/unrolled/render.v1"
	"github.com/julienschmidt/httprouter"
	"regexp"
	//"fmt"
)

type UserController struct {
	AppController
	*render.Render
}




func (c *UserController) Index(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		
		newuser:= models.User{	Username:	"ozzadar", 
								Email: 		"ozzadar@ozzadar.com", 
								FirstName:	"Paul",
								Surname:	"Mauviel",
								Password: 	"diamond11",
								Role:		"admin"}

		CreateUser(&newuser)

		c.HTML(rw, http.StatusOK, "index", nil)
		
}

func (c *UserController) Register(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	


	if r.Method == "GET" {
		
		c.HTML(rw, http.StatusOK, "user/register", nil)

	} else if r.Method == "POST" {
		r.ParseForm();

		reg, err := regexp.Compile(`\W`)

		if err != nil {
			panic(err)
		}
		errorstring := "";
		username := r.FormValue("username")
		email := r.FormValue("email")
		firstname := r.FormValue("firstname")
		surname := r.FormValue("lastname")
	//	password := r.FormValue("password")
	//	confirmpassword := r.FormValue("confirmpassword")

		if (reg.MatchString(string(username))) {
			errorstring = errorstring + "Username is invalid. A-Za-z0-9 only. -- "
		}
		if (reg.MatchString(string(firstname))) {
			errorstring = errorstring + "First Name is invalid. A-Za-z0-9 only. -- "
		}
		if (reg.MatchString(string(surname))) {
			errorstring = errorstring + "Last Name is invalid. A-Za-z0-9 only. -- "
		}
	
		reg , err = regexp.Compile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`)

		if err != nil {
			panic(err)
		}

		if (!(reg.MatchString(string(email)))) {
			errorstring = errorstring + "Email is invalid.\n"
		}

		if (errorstring != "") {
			error := ErrorMessage{Message: errorstring}

			c.HTML(rw, http.StatusOK, "user/register", error)
		}

	}

	
}