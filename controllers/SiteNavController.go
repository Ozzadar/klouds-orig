package controllers

import (
	"net/http"
	"github.com/superordinate/klouds2.0/models"
	"gopkg.in/unrolled/render.v1"
	"github.com/julienschmidt/httprouter"

)

type SiteNavController struct {
	AppController
	*render.Render
}

func (c *SiteNavController) ServeFull(rw http.ResponseWriter, r *http.Request, p httprouter.Params, pageName string) {

	
}

func (c *SiteNavController) Index(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var user *models.User

		if getUserName(r) != "" {
			user = GetUserByUsername(getUserName(r))
			c.HTML(rw, http.StatusOK, "index", user)
			
		} else {
			c.HTML(rw, http.StatusOK, "index", nil)
		}
}
