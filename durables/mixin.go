package durables

import (
	"github.com/fox-one/mixin-sdk-go"
	"github.com/lzm-cli/gin-web-server-template/config"
)

var client *mixin.Client

func GetMixinClient() *mixin.Client {
	var err error
	client, err = mixin.NewFromKeystore(&mixin.Keystore{
		ClientID:   config.C.Mixin.ClientID,
		SessionID:  config.C.Mixin.SessionID,
		PinToken:   config.C.Mixin.PinToken,
		PrivateKey: config.C.Mixin.PrivateKey,
	})
	if err != nil {
		panic(err)
	}

	return client
}
