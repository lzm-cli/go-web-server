package routes

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/bugsnag/bugsnag-go/errors"
	"github.com/dimfeld/httptreemux"
	"github.com/<%= organization %>/<%= repo %>/config"
	"github.com/<%= organization %>/<%= repo %>/session"
	"github.com/<%= organization %>/<%= repo %>/views"
)

func RegisterHandlers(router *httptreemux.TreeMux) {
	router.MethodNotAllowedHandler = func(w http.ResponseWriter, r *http.Request, _ map[string]httptreemux.HandlerFunc) {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context(), ""))
	}
	router.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		views.RenderErrorResponse(w, r, session.NotFoundError(r.Context(), ""))
	}
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, rcv interface{}) {
		err := fmt.Errorf(string(errors.New(rcv, 2).Stack()))
		views.RenderErrorResponse(w, r, session.ServerError(r.Context(), err))
	}
}

func RegisterRoutes(router *httptreemux.TreeMux) {
	registerUser(router)
	router.GET("/", root)
	router.GET("/_hc", healthCheck)
}

func root(w http.ResponseWriter, r *http.Request, params map[string]string) {
	views.RenderDataResponse(w, r, map[string]string{
		"build":      config.BuildVersion + "-" + runtime.Version(),
		"developers": "https://snapshot-api.mixin.one",
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request, params map[string]string) {
	views.RenderBlankResponse(w, r)
}
