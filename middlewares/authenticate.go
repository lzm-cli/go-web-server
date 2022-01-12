package middlewares

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/<%= organization %>/<%= repo %>/models"
	"github.com/<%= organization %>/<%= repo %>/session"
	"github.com/<%= organization %>/<%= repo %>/views"
)

var whitelist = [][2]string{
	{"GET", "^/auth$"},
}

type contextValueKey struct{ int }

var keyCurrentUser = contextValueKey{1000}

func CurrentUser(r *http.Request) *models.User {
	u, _ := r.Context().Value(keyCurrentUser).(*models.User)
	return u
}

func Authenticate(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			handleUnauthorized(handler, w, r)
			return
		}
		u, err := models.AuthenticateUserByToken(r.Context(), header[7:])
		if models.CheckEmptyError(err) != nil {
			views.RenderErrorResponse(w, r, err)
		} else if u == nil {
			handleUnauthorized(handler, w, r)
		} else {
			ctx := context.WithValue(r.Context(), keyCurrentUser, u)
			handler.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func handleUnauthorized(handler http.Handler, w http.ResponseWriter, r *http.Request) {
	for _, pp := range whitelist {
		if pp[0] != r.Method {
			continue
		}
		if matched, err := regexp.MatchString(pp[1], r.URL.Path); err == nil && matched {
			handler.ServeHTTP(w, r)
			return
		}
	}

	views.RenderErrorResponse(w, r, session.AuthorizationError(r.Context()))
}
