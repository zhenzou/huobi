package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func NewSigner(key string) *Signer {
	return &Signer{
		key: []byte(key),
	}
}

type Signer struct {
	key []byte
}

func (p *Signer) Sign(method string, host string, path string, parameters string) string {
	if method == "" || host == "" || path == "" || parameters == "" {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString("\n")
	sb.WriteString(host)
	sb.WriteString("\n")
	sb.WriteString(path)
	sb.WriteString("\n")
	sb.WriteString(parameters)

	return p.sign(sb.String())
}

func (p *Signer) sign(payload string) string {
	hash := hmac.New(sha256.New, p.key)
	_, _ = hash.Write([]byte(payload))
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}
