package admin

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     Role
}

type AdminAuth struct {
	Users []AuthUser
}

type UIAuth struct {
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}

func (admin *Admin) ServeAuth() {
	admin.router.POST(admin.ApiPrefix+"/auth", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var authUser AuthUser
		for _, user := range admin.Auth.Users {
			if user.Username == "iegora" && user.Password == "2008a" {
				authUser = user
				break
			}
		}

		if authUser.Username != "" && authUser.Password != "" {
			data, _ := json.Marshal(UIAuth{
				Error: false,
				Token: "test-123",
			})

			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(data))
		} else {
			data, _ := json.Marshal(UIAuth{
				Error:   true,
				Message: "Invalid credentials",
			})

			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(data))
		}
	})
}
