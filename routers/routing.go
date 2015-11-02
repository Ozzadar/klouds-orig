package routers

import (
	"net/http"
	"gopkg.in/unrolled/render.v1"
	"github.com/julienschmidt/httprouter"
	"github.com/superordinate/klouds2.0/controllers"
)

type Routing struct {

	Render *render.Render
	Mux *httprouter.Router

}

func (r *Routing) Init() {


	r.Render = render.New(render.Options{Directory: "views"})
	r.Mux = httprouter.New()

	c := &controllers.SiteNav{Render: r.Render}

	
	r.Mux.GET("/", c.Index)
	r.Mux.NotFound = http.FileServer(http.Dir("public"))
}