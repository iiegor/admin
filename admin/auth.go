package admin

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

func (admin Admin) Auth() {
	admin.router.POST(path.Join(admin.ApiPrefix, "login"), func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("{\"firstName\": \"Iegor\", \"lastName\": \"Azuaga\", \"nickName\": \"iegor\" }"))
	})
}
