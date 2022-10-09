package durables

import (
	"context"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/<%= organization %>/<%= repo %>/config"
	"github.com/<%= organization %>/<%= repo %>/session"
)

var MixinCtx context.Context

func GetMixinClient() *mixin.Client {
	client, err := mixin.NewFromKeystore(&mixin.Keystore{
		ClientID:   config.C.Mixin.ClientID,
		SessionID:  config.C.Mixin.SessionID,
		PinToken:   config.C.Mixin.PinToken,
		PrivateKey: config.C.Mixin.PrivateKey,
	})
	if err != nil {
		panic(err)
	}

	MixinCtx = session.WithMixinClient(context.Background(), client)
	return client
}
