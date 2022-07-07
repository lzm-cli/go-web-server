package middlewares

import (
	"net/http"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/<%= organization %>/<%= repo %>/session"
	"github.com/unrolled/render"
	"gorm.io/gorm"
)

func Context(handler http.Handler, db *gorm.DB, client *mixin.Client, render *render.Render) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := session.WithRequest(r.Context(), r)
		ctx = session.WithMixinClient(ctx, client)
		ctx = session.WithDatabase(ctx, db)
		ctx = session.WithRender(ctx, render)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
