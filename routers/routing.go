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

	controllers.Init()
	r.Render = render.New(render.Options{Directory: "views" })
	r.Mux = httprouter.New()

	c := &controllers.SiteNavController{Render: r.Render}
	u := &controllers.UserController{Render: r.Render}
	
	r.Mux.GET("/", c.Index)
	r.Mux.GET("/user", u.Login)
	r.Mux.POST("/user/register", u.Register)
	r.Mux.GET("/user/register", u.Register)
	r.Mux.POST("/user/logout", u.Logout)
	r.Mux.POST("/user/login", u.Login)
	r.Mux.GET("/user/login", u.Login)
	r.Mux.GET("/user/profile", u.Profile)
	r.Mux.POST("/user/profile", u.Profile)
	
	r.Mux.NotFound = http.FileServer(http.Dir("public"))
}