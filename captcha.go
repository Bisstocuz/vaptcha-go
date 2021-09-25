package vaptcha

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrIllegalServer = errors.New("illegal server")
)

type CaptchaRequest struct {
	VID      string `json:"id"`
	Key      string `json:"secretkey"`
	Server   string `json:"-"`
	Scene    string `json:"scene"`
	Token    string `json:"token"`
	ClientIP string `json:"ip"`
	UserID   string `json:"userid,omitempty"`
}

// validateServer will validate verify server's url, if illegal, return error, otherwise nil
func (request *CaptchaRequest) validateServer() error {
	u, err := url.Parse(request.Server)
	if err != nil {
		return ErrIllegalServer
	}
	host := u.Hostname()
	if strings.HasSuffix(host, ".vaptcha.com") || strings.HasSuffix(host, ".vaptcha.net") {
		return nil
	}
	return ErrIllegalServer
}
