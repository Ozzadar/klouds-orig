package controllers

import (
	"net/http"

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
	c.HTML(rw, http.StatusOK, "index", nil)
}
