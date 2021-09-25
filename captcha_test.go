package vaptcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateServer(t *testing.T) {
	req := &CaptchaRequest{Server: "https://xsdr.vaptcha.net/verify"}
	assert.Equal(t, nil, req.validateServer(), "legal server not pass")

	req = &CaptchaRequest{Server: "https://x.cvaptcha.com/verify"}
	assert.Equal(t, ErrIllegalServer, req.validateServer(), "illegal server approve")
}
