package durables

import (
	"github.com/fox-one/mixin-sdk-go"
	"github.com/<%= organization %>/<%= repo %>/config"
)

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
	return client
}
