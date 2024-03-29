package admin

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/julienschmidt/httprouter"
)

type UIMeta struct {
	Resources []UIMetaResource `json:"resources"`
}

// TODO: Necesitaríamos enviar una definición de los tipos de datos
type UIMetaResource struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Methods []string          `json:"methods"`
	Fields  map[string]string `json:"fields"`
}

func (admin Admin) BuildUIMeta() UIMeta {
	meta := UIMeta{}

	for _, res := range admin.resources {
		metaRes := UIMetaResource{
			Id:      inflection.Plural(res.ToParam()),
			Name:    inflection.Plural(res.Name),
			Methods: res.methods,
			Fields:  res.GetFields(),
		}
		meta.Resources = append(meta.Resources, metaRes)
	}

	return meta
}

type UIMetricsResource struct {
	Name string `json:"name"`
	Hits int64  `json:"hits"`
}

type UIMetrics struct {
	Resources []UIMetricsResource `json:"resources"`
}

func (admin Admin) BuildUIMetrics() UIMetrics {
	meta := UIMetrics{}
	for _, res := range admin.resources {
		metricsRes := UIMetricsResource{
			Name: inflection.Plural(res.ToParam()),
			Hits: res.metrics.hits,
		}
		meta.Resources = append(meta.Resources, metricsRes)
	}

	return meta
}

func (admin Admin) ServeUIMeta() {
	meta := admin.BuildUIMeta()

	admin.router.GET(path.Join(admin.AdminConfig.Prefix, "api", "ui", "meta"), func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data, err := json.Marshal(meta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(data))
	})

	admin.router.GET(path.Join(admin.AdminConfig.Prefix, "api", "ui", "metrics"), func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		metrics := admin.BuildUIMetrics()

		data, err := json.Marshal(metrics)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(data))
	})
}

func (admin Admin) ServeUI() {
	admin.router.GET("/ui/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var url = r.URL.String()

		if strings.Contains(url, "/static/") {
			http.ServeFile(w, r, path.Join("./", strings.Replace(url, "/ui/static", "/ui/dist/static", -1)))
		} else {
			http.ServeFile(w, r, "./ui/dist/index.html")
		}
	})
}
