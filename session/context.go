package session

import (
	"context"
	"net/http"

	"github.com/fox-one/mixin-sdk-go"
	"gorm.io/gorm"

	"github.com/unrolled/render"
)

type contextValueKey int

const (
	keyRequest           contextValueKey = 0
	keyDatabase          contextValueKey = 1
	keyRender            contextValueKey = 3
	keyLimiter           contextValueKey = 4
	keyMixinClient       contextValueKey = 5
	keyRemoteAddress     contextValueKey = 11
	keyAuthorizationInfo contextValueKey = 12
	keyRequestBody       contextValueKey = 13
)

func DB(ctx context.Context) *gorm.DB {
	v, _ := ctx.Value(keyDatabase).(*gorm.DB)
	return v
}

func Render(ctx context.Context) *render.Render {
	v, _ := ctx.Value(keyRender).(*render.Render)
	return v
}
func MixinClient(ctx context.Context) *mixin.Client {
	v, _ := ctx.Value(keyMixinClient).(*mixin.Client)
	return v
}

func Request(ctx context.Context) *http.Request {
	v, _ := ctx.Value(keyRequest).(*http.Request)
	return v
}

func RemoteAddress(ctx context.Context) string {
	v, _ := ctx.Value(keyRemoteAddress).(string)
	return v
}

func WithMixinClient(ctx context.Context, client *mixin.Client) context.Context {
	return context.WithValue(ctx, keyMixinClient, client)
}

func WithRender(ctx context.Context, render *render.Render) context.Context {
	return context.WithValue(ctx, keyRender, render)
}

func WithDatabase(ctx context.Context, database *gorm.DB) context.Context {
	return context.WithValue(ctx, keyDatabase, database)
}

func WithRequest(ctx context.Context, r *http.Request) context.Context {
	rCopy := new(http.Request)
	*rCopy = *r
	return context.WithValue(ctx, keyRequest, rCopy)
}

func WithRemoteAddress(ctx context.Context, remoteAddr string) context.Context {
	return context.WithValue(ctx, keyRemoteAddress, remoteAddr)
}
