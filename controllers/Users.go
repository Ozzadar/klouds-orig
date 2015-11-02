package controllers

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	AppController
	*render.Render
}

func (c *UserController) ServeFull(rw http.ResponseWriter, r *http.Request, p httprouter.Params, pageName string) {
	c.HTML(rw, http.StatusOK, pageName, nil)
}

func (c *UserController) Index(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c.ServeFull(rw,r,p,"example")
}
func (c *UserController) Book(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		
		c.ServeFull(rw,r,p,"index")
		
}