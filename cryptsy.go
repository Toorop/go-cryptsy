package cryptsy

import (
	"fmt"
)

const (
	API_BASE_PUB               = "http://pubapi.cryptsy.com/api.php"
	API_BASE_PRIV              = "https://api.cryptsy.com/api"
	PUSHER_ENDPOINT            = ""
	DEFAULT_HTTPCLIENT_TIMEOUT = 30
)

type cryptsy struct {
	client *client
}

func New(pubKey, privKey string) {
	return &cryptsy{NewClient(pubKey, privKey)}
}
