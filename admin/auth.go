package admin

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (admin Admin) Auth() {
	admin.router.POST("/api/login", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("{\"firstName\": \"Iegor\", \"lastName\": \"Azuaga\", \"nickName\": \"iegor\" }"))
	})
}
