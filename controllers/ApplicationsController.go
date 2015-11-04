package controllers

import (
	"net/http"
	"github.com/superordinate/klouds2.0/models"
	"gopkg.in/unrolled/render.v1"
	"github.com/julienschmidt/httprouter"
	"strings"
	"fmt"
)

type ApplicationsController struct {
	AppController
	*render.Render
}

func (c *ApplicationsController) AppAdmin(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

		if r.Method == "GET" {
		
		var user *models.User
		

		if getUserName(r) != "" || user != nil {
			user = GetUserByUsername(getUserName(r))
			
			if NotAdministrator(user, c, rw) {
				return
			}


			c.HTML(rw, http.StatusOK, "apps/admin/admin", user)
			
		} else {

			c.HTML(rw, http.StatusOK, "user/login", nil)
		}

	}
		
}

func (c *ApplicationsController) CreateApplication(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if r.Method == "GET" {
		
		var user *models.User
		//var newapp *models.Application
		if getUserName(r) != "" || user != nil{
			user = GetUserByUsername(getUserName(r))

			if NotAdministrator(user, c, rw) {
				return
			}

			c.HTML(rw, http.StatusOK, "apps/admin/newapp", nil)
			
		} else {
			c.HTML(rw, http.StatusOK, "user/login", nil)
		}

	} else if r.Method == "POST" {
		var user *models.User
		if getUserName(r) != "" || user != nil{

			user = GetUserByUsername(getUserName(r))

			if NotAdministrator(user, c, rw) {
				return
			}

			r.ParseForm();
			//Load Application from form

			appname := strings.ToLower(r.FormValue("appname"))
			image := r.FormValue("dockerimage")
			depends := r.FormValue("dependencies")
			//split them by commas
			dependencies := strings.Split(depends,",")

			//remove white space
			for i:=0; i< len(dependencies); i++ {
				dependencies[i] = stripSpaces(dependencies[i])
				fmt.Println(dependencies[i])
			}

			//Create structure array
			dependecystruct := make ([]models.Dependency, len(dependencies))

			for i:=0; i< len(dependencies); i++ {
				dependecystruct[i] = models.Dependency{Dependency: dependencies[i]}
			}

			envvars := r.FormValue("environmentvariables")
			environmentvariables := strings.Split(envvars,",")

			//remove white space
			for i:=0; i< len(environmentvariables); i++ {
				environmentvariables[i] = stripSpaces(environmentvariables[i])
				fmt.Println(environmentvariables[i])
			}

			//Create structure array
			envvarsstruct := make ([]models.EnvironmentVariable, len(environmentvariables))
			fmt.Println(environmentvariables)

			for i:=0; i< len(environmentvariables); i++ {
				split := strings.Split(environmentvariables[i], ":")
				envvarsstruct[i] = models.EnvironmentVariable{	Key: string(split[0]),
																Value: string(split[1])}
			}

			logo := r.FormValue("logo")
			internalport := r.FormValue("internalport")
			protocol := r.FormValue("protocol")
			description := r.FormValue("description")

			newapp := models.Application{
				Name: 					appname,
				DockerImage:			image,
				Dependencies: 			dependecystruct,
				EnvironmentVariables:	envvarsstruct,
				Logo:					logo,
				Description:			description,
				InternalPort: 			internalport,
				Protocol: 				protocol,
				Username:				getUserName(r)}
			
			//Validate fields
			newapp.ValidateApplication()

			//Check against application database for dependencies

			for i:=0; i< len(newapp.Dependencies); i++ {
				
				if !CheckApplicationExists(newapp.Dependencies[i].Dependency) {
					newapp.Message = newapp.Message + "Dependency " + newapp.Dependencies[i].Dependency + " does not exist in the database."
				}
			}

			if (newapp.Message != "") {
				//Validation failed
				c.HTML(rw, http.StatusOK, "apps/admin/newapp", newapp)
				return
			}

			//Create Application

			CreateApplication(&newapp)
			newapp.Message = "Application " + newapp.Name + " created successfully."
			c.HTML(rw, http.StatusOK, "apps/admin/newapp", newapp)
		} else {
			c.HTML(rw, http.StatusOK, "user/login", nil)
		}
	}
		
}

