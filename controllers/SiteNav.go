package controllers

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
	"github.com/julienschmidt/httprouter"
)

type SiteNav struct {
	AppController
	*render.Render
}

func (c *SiteNav) ServeFull(rw http.ResponseWriter, r *http.Request, p httprouter.Params, pageName string) {

	
}

func (c *SiteNav) Index(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c.HTML(rw, http.StatusOK, "index", nil)
}
