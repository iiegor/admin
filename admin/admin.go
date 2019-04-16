package admin

import (
	"net/http"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type AdminConfig struct {
	Prefix string
	ApiPrefix string
	UI     bool
	DB     *xorm.Engine
	Debug  bool
	Auth 	 *AdminAuth
}

type Admin struct {
	*AdminConfig

	router    *httprouter.Router
	resources []*Resource
	startTime time.Time
}

func New(config *AdminConfig) *Admin {
	admin := Admin{
		router:      httprouter.New(),
		AdminConfig: config,
	}

	if admin.AdminConfig.Prefix == "" {
		admin.AdminConfig.Prefix = "/"
	}

	if admin.AdminConfig.ApiPrefix == "" {
		admin.AdminConfig.ApiPrefix = "/api"
	}

	if admin.AdminConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	return &admin
}

func (admin *Admin) AddResource(val interface{}, config ResourceConfig) {
	resource := CreateResource(val, config)
	resource.admin = admin

	admin.resources = append(admin.resources, resource)

	resource.RegisterRoutes()
}

func (admin *Admin) MountTo(mux *http.ServeMux) {
	muxHandler := admin.MuxHandler()

	if len(admin.Prefix) > 0 {
		mux.Handle(admin.Prefix, muxHandler)
	}

	mux.Handle(admin.Prefix+"/", muxHandler)

	admin.startTime = time.Now()
}

func (admin *Admin) MuxHandler() http.Handler {
	router := admin.router

	admin.ServeAuth()
	admin.ServeUIMeta()

	if admin.AdminConfig.UI {
		admin.ServeUI()
	}

	return router
}
