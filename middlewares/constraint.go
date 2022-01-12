package middlewares

import (
	"fmt"
	"net"
	"net/http"

	"github.com/<%= organization %>/<%= repo %>/session"
	"github.com/<%= organization %>/<%= repo %>/views"
)

func parseRemoteAddr(remoteAddress string) (string, error) {
	host, _, err := net.SplitHostPort(remoteAddress)
	if err == nil {
		remoteAddress = host
	}
	ip := net.ParseIP(remoteAddress)
	if ip == nil {
		return "", fmt.Errorf("invalid remote address %s", remoteAddress)
	}
	return ip.String(), nil
}

func Constraint(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddress, err := parseRemoteAddr(r.RemoteAddr)
		if err != nil {
			views.RenderBlankResponse(w, r)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization,Mixin-Conversation-ID")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,PUT,POST,DELETE")
		w.Header().Set("Access-Control-Max-Age", "600")
		if r.Method == "OPTIONS" {
			views.RenderBlankResponse(w, r)
		} else {
			ctx := session.WithRemoteAddress(r.Context(), remoteAddress)
			handler.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
