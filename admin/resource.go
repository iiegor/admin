package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type ResourceConfig struct {
	TableName string
	Methods   []string
	Parent    *Resource
}

type Resource struct {
	Name      string
	TableName string

	value   interface{}
	methods []string
	admin   *Admin
	mounted bool

	metrics *ResourceMetrics
}

type ResourceMetrics struct {
	hits int64
}

func CreateResource(val interface{}, config ResourceConfig) *Resource {
	res := Resource{
		Name:      GetName(val),
		TableName: config.TableName,
		methods:   config.Methods,
		metrics:   new(ResourceMetrics),
		value:     val,
	}

	if len(config.TableName) == 0 {
		res.TableName = inflection.Plural(ToParamString(res.Name))
	}

	if len(config.Methods) == 0 {
		res.methods = []string{"read"}
	}

	return &res
}

type ResourceQuery struct {
	Limit  int
	Offset int
}

// TODO: Move routing logic to route.go
func (res *Resource) RegisterRoutes() {
	var rootPath = path.Join(res.RoutePrefix(), inflection.Plural(res.ToParam()))
	res.GetAdmin().router.GET(rootPath, ChainMiddlewares(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		query := res.parseQuery(r.URL)

		fetchFromDatabase := res.NewSlice()
		err = res.GetAdmin().DB.Table(res.TableName).Limit(query.Limit, query.Offset).Find(fetchFromDatabase)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"resource": res.Name,
				"method": "read",
				"url": r.URL,
				"error": err.Error(),
			}).Error()

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(fetchFromDatabase)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(data))
	}, MetricsMiddleware(res)))

	logrus.WithFields(logrus.Fields{
		"resource": res.Name,
		"method":   "read",
		"path":     rootPath,
	}).Debug("GET route registered")

	// FIXME: Logs will output method=delete since it's the last iteration done.
	//  A solution could be storing the current method inside another variable and use that instead.
	for _, method := range res.methods {
		switch strings.ToLower(method) {
		case "read":			
			var resourcePath = path.Join(rootPath, res.ParamID())
			res.GetAdmin().router.GET(resourcePath, ChainMiddlewares(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				var resourceID = p.ByName(res.ParamID()[1:])

				fetchFromDatabase := res.NewStruct()
				has, err := res.GetAdmin().DB.Table(res.TableName).Where("id = ?", resourceID).Get(fetchFromDatabase)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"resource": res.Name,
						"method": method,
						"url": r.URL,
						"error": err.Error(),
					}).Error()

					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				if has {
					data, _ := json.Marshal(fetchFromDatabase)
					w.Header().Add("Content-Type", "application/json")
					w.Write([]byte(data))
				}
			}, MetricsMiddleware(res)))

			logrus.WithFields(logrus.Fields{
				"resource": res.Name,
				"method":   method,
				"path":     resourcePath,
			}).Debug("GET route registered")
		case "create":
			var resourcePath = path.Join(rootPath)
			res.GetAdmin().router.PUT(resourcePath, ChainMiddlewares(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				w.Write([]byte("Hello, World!"))
			}, MetricsMiddleware(res)))

			logrus.WithFields(logrus.Fields{
				"resource": res.Name,
				"method":   method,
				"path":     resourcePath,
			}).Debug("PUT route registered")
		case "delete":
			var resourcePath = path.Join(rootPath, res.ParamID())
			res.GetAdmin().router.DELETE(resourcePath, ChainMiddlewares(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				w.Write([]byte("Hello, World!"))
			}, MetricsMiddleware(res)))

			logrus.WithFields(logrus.Fields{
				"resource": res.Name,
				"method":  	method,
				"path":     resourcePath,
			}).Debug("DELETE route registered")
		case "update":
			var resourcePath = path.Join(rootPath, res.ParamID())
			res.GetAdmin().router.PATCH(resourcePath, ChainMiddlewares(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				w.Write([]byte("Hello, World!"))
			}, MetricsMiddleware(res)))

			logrus.WithFields(logrus.Fields{
				"resource": res.Name,
				"method":   method,
				"path":     resourcePath,
			}).Debug("PATCH route registered")
		default:
			logrus.WithFields(logrus.Fields{
				"resource": res.Name,
				"method":   method,
			}).Warn("Unknown method provided")
		}
	}

	res.mounted = true
}

func (res *Resource) parseQuery(url *url.URL) ResourceQuery {
	query := new(ResourceQuery)

	limit, err := strconv.Atoi(url.Query().Get("limit"))
	if err == nil && limit <= 50 {
		query.Limit = limit
	} else {
		query.Limit = 10
	}

	offset, err := strconv.Atoi(url.Query().Get("offset"))
	if err == nil {
		query.Offset = offset
	} else {
		query.Offset = 0
	}

	return *query
}

func (res *Resource) NewStruct() interface{} {
	if res.value == nil {
		return nil
	}
	return reflect.New(Indirect(reflect.ValueOf(res.value)).Type()).Interface()
}

func (res *Resource) NewSlice() interface{} {
	if res.value == nil {
		return nil
	}
	sliceType := reflect.SliceOf(reflect.TypeOf(res.value))
	slice := reflect.MakeSlice(sliceType, 0, 0)
	slicePtr := reflect.New(sliceType)
	slicePtr.Elem().Set(slice)
	return slicePtr.Interface()
}

// ToParam e.g. "MyModel" -> "mymodel"
func (res Resource) ToParam() string {
	return ToParamString(res.Name)
}

// ParamId e.g. ":course_id"
func (res Resource) ParamID() string {
	return fmt.Sprintf(":%v_id", inflection.Singular(res.ToParam()))
}

func (res Resource) GetAdmin() *Admin {
	return res.admin
}

// RoutePrefix returns e.g. "/admin/api"
func (res Resource) RoutePrefix() string {
	return path.Join(res.admin.Prefix, res.GetAdmin().ApiPrefix)
}
