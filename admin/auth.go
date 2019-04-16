package admin

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

type AuthUser struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      *Role
}

type AdminAuth struct {
	Users []AuthUser
}

func (admin Admin) ServeAuth() {
	admin.router.POST(path.Join(admin.ApiPrefix, "login"), func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var user []byte

		for _, i := range admin.Auth.Users {
			if i.Nickname == "iegor" {
				user, _ = json.Marshal(i)
				break
			}
		}

		if user != nil {
			w.Write([]byte(user))
		} else {
			http.Error(w, "Bad credentials", http.StatusUnauthorized)
		}
	})
}
