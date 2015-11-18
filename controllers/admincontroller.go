package controllers

import (
	"net/http"
	"github.com/superordinate/klouds/models"
	"gopkg.in/unrolled/render.v1"
	"github.com/julienschmidt/httprouter"
	"strings"
	"fmt"
)

type AdminController struct {
	AppController
	*render.Render
}

func (c *AdminController) AppAdmin(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

		if r.Method == "GET" {
		
		var user *models.User
		


		if getUserName(r) != "" || user != nil {
			user = GetUserByUsername(getUserName(r))
			
			if NotAdministrator(user, c, rw) {
				return
			}

					//Get application list for user
			runningapps := []models.RunningApplication{}

			GetRunningApplications(&runningapps)

			for i:=0; i<len(runningapps);i++ {
				runningapps[i].User = *user
			}

			if len(runningapps) == 0 {
				runningapps = []models.RunningApplication{models.RunningApplication{User: *user}}	
			}

			c.HTML(rw, http.StatusOK, "admin/admin", runningapps)
			
		} else {

			c.HTML(rw, http.StatusOK, "user/login", nil)
		}

	}
		
}

func (c *AdminController) CreateApplication(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if r.Method == "GET" {
		
		var user *models.User
		//var newapp *models.Application
		if getUserName(r) != "" || user != nil{
			user = GetUserByUsername(getUserName(r))

			if NotAdministrator(user, c, rw) {
				return
			}

			newapp := &models.Application{User: *user}
			c.HTML(rw, http.StatusOK, "admin/newapp", newapp)
			
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

			//Strips white space
			for i:=0; i< len(environmentvariables); i++ {
				environmentvariables[i] = stripSpaces(environmentvariables[i])
			}

			//Create structure array
			envvarsstruct := make ([]models.EnvironmentVariable, len(environmentvariables))
			fmt.Println("length! : ",len(environmentvariables))

			for i:=0; i < len(environmentvariables); i++ {
				split := strings.Split(environmentvariables[i], ":")
				if len(envvarsstruct) >= 1 {
					envvarsstruct[i] = models.EnvironmentVariable{}

					for j:= 0; j< len(split); j++ {
						if j == 0 {
							envvarsstruct[i].Key = split[j]
						} else if j == 1 {
							envvarsstruct[i].Value = split[j]
						} else {
							envvarsstruct[i].Value = envvarsstruct[i].Value + ":" + split[j]
						}
					}
					
				}

			}

			logo := r.FormValue("logo")
			if logo == "" {
				logo ="/images/noimage.png"
			}

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
				User:					*user}
			
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
				c.HTML(rw, http.StatusOK, "admin/newapp", newapp)
				return
			}

			//Create Application

			CreateApplication(&newapp)
			newapp.Message = "Application " + newapp.Name + " created successfully."
			c.HTML(rw, http.StatusOK, "admin/newapp", newapp)
		} else {
			c.HTML(rw, http.StatusOK, "user/login", nil)
		}
	}
		
}

func (c *AdminController) EditApplication(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if r.Method == "GET" {
		
		var user *models.User
		//var newapp *models.Application
		if getUserName(r) != "" || user != nil{
			user = GetUserByUsername(getUserName(r))

			if NotAdministrator(user, c, rw) {
				return
			}

			newapp := GetApplicationByName(p.ByName("appID"))
			c.HTML(rw, http.StatusOK, "admin/editapp", newapp)
			
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

			//Strips white space
			for i:=0; i< len(environmentvariables); i++ {
				environmentvariables[i] = stripSpaces(environmentvariables[i])
			}

			//Create structure array
			envvarsstruct := make ([]models.EnvironmentVariable, len(environmentvariables))
			fmt.Println("length! : ",len(environmentvariables))

			for i:=0; i < len(environmentvariables); i++ {
				split := strings.Split(environmentvariables[i], ":")
				if len(envvarsstruct) >= 1 {
					envvarsstruct[i] = models.EnvironmentVariable{}

					for j:= 0; j< len(split); j++ {
						if j == 0 {
							envvarsstruct[i].Key = split[j]
						} else if j == 1 {
							envvarsstruct[i].Value = split[j]
						} else {
							envvarsstruct[i].Value = envvarsstruct[i].Value + ":" + split[j]
						}
					}
					
				}

			}

			logo := r.FormValue("logo")
			if logo == "" {
				logo ="/images/noimage.png"
			}

			internalport := r.FormValue("internalport")
			protocol := r.FormValue("protocol")
			description := r.FormValue("description")
			newapp := GetApplicationByName(p.ByName("appID"))

			
				newapp.DockerImage =		image
				newapp.Dependencies =		dependecystruct
				newapp.EnvironmentVariables=envvarsstruct
				newapp.Logo =				logo
				newapp.Description =		description
				newapp.InternalPort =		internalport
				newapp.Protocol =			protocol
				newapp.User =				*user
			
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
				c.HTML(rw, http.StatusOK, "admin/editapp", newapp)
				return
			}

			//Update Application


			UpdateApplication(newapp)

			newapp.Message = "Application " + newapp.Name + " updated successfully."
			c.HTML(rw, http.StatusOK, "admin/editapp", newapp)
		} else {
			c.HTML(rw, http.StatusOK, "user/login", nil)
		}
	}
		
}


func (c *AdminController) DeleteApplication(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if r.Method == "GET" {
		
		var user *models.User
		//var newapp *models.Application
		if getUserName(r) != "" || user != nil{
			user = GetUserByUsername(getUserName(r))

			if NotAdministrator(user, c, rw) {
				return
			}

			newapp := GetApplicationByName(p.ByName("appID"))
			newapp.User = *user
			c.HTML(rw, http.StatusOK, "admin/delete", newapp)
			
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

			
			newapp := GetApplicationByName(p.ByName("appID"))
			newapp.User = *user

			//Delete Application


			DeleteApplication(newapp,user.Username)

			newapp.Message = "Application " + newapp.Name + " deleted successfully."
			c.HTML(rw, http.StatusOK, "admin/deleted", newapp)
		} else {
			c.HTML(rw, http.StatusOK, "user/login", nil)
		}
	}
		
}