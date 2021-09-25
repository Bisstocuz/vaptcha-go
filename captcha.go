package vaptcha

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrIllegalServer = errors.New("illegal server")
	ErrWrongUserID   = errors.New("userid error")
	ErrEmptyID       = errors.New("id empty")
	ErrWrongID       = errors.New("id error")
	ErrWrongScene    = errors.New("scene error")
	ErrWrongToken    = errors.New("token error")
	ErrExpiredToken  = errors.New("token expired")
	ErrOverrun       = errors.New("frequency overrun")
	ErrBadRequest    = errors.New("bad request")
)

type CaptchaRequest struct {
	VID      string `json:"id"`
	Key      string `json:"secretkey"`
	Server   string `json:"-"`
	Scene    int    `json:"scene"`
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

type CaptchaResponse struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

// Request will send request to verify server and get response, if server is illegal, returns ErrIllegalServer,
// you can use this function alone to handle error details
func (request *CaptchaRequest) Request() (*CaptchaResponse, error) {
	// validate server
	err := request.validateServer()
	if err != nil {
		return nil, err
	}
	// post request
	buf := bytes.NewReader(mustMarshalToJson(request))
	resp, err := http.Post(request.Server, "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// bind data to response struct
	var response CaptchaResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Verify will verify the response and return nil if pass, otherwise errors
func (response *CaptchaResponse) Verify() error {
	if response.Success == 1 {
		return nil
	} else {
		return errors.New(response.Msg)
	}
}

// RequestAndVerify will request and verify captcha info, return true if pass, otherwise false
func RequestAndVerify(request *CaptchaRequest) bool {
	response, err := request.Request()
	if err != nil {
		return false
	}
	err = response.Verify()
	if err != nil {
		return false
	}
	return true
}
