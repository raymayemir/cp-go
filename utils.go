package cp_go

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ComputeHmac256(message, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(message)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
