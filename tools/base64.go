package tools

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
)

func Base64Decode(str string) []byte {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		decodeBytes, err := base64.RawURLEncoding.DecodeString(str)
		if err != nil {
			return make([]byte, 0)
		}
		return decodeBytes
	}
	return decodeBytes
}

func Base64Encode(str []byte) string {
	return base64.RawURLEncoding.EncodeToString(str)
}

func SafeBase64Encode(s string) string {
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "=", "")
	return s
}

func GetBase64ByURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(body), nil
}

func GetBytesByURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
