package routes

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/<%= organization %>/<%= repo %>/middlewares"
	"github.com/<%= organization %>/<%= repo %>/models"
	"github.com/<%= organization %>/<%= repo %>/session"
	"github.com/<%= organization %>/<%= repo %>/views"
)

func registerUser(router *httptreemux.TreeMux) {
	impl := &usersImpl{}

	router.GET("/auth", impl.authenticate)
	router.GET("/me", impl.me)
}

type usersImpl struct{}

func (impl *usersImpl) authenticate(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	if err := r.ParseForm(); err != nil {
		views.RenderErrorResponse(w, r, session.BadDataError(r.Context()))
		return
	}
	code := r.Form.Get("code")
	if token, err := models.AuthenticateUserByOAuth(r.Context(), code); err != nil {
		views.RenderErrorResponse(w, r, err)
	} else {
		views.RenderDataResponse(w, r, token)
	}
}

func (impl *usersImpl) me(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	views.RenderDataResponse(w, r, middlewares.CurrentUser(r))
}

// POST demo
// var body models.User
// if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
// 	views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
// } else if err := user.UpdateUserProfile(r.Context(), middlewares.CurrentUser(r), body.Biography, body.Website); err != nil {
// 	views.RenderErrorResponse(w, r, err)
// } else {
// 	views.RenderDataResponse(w, r, "success")
// }
