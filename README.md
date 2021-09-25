# vaptcha-go
 This is a third-party golang SDK for [Vaptcha](https://www.vaptcha.com/).

### Quick Usage
1. get this module by command line

`go get github.com/Bisstocuz/vaptcha-go`

2. code with this module

```
	request := &CaptchaRequest{
		VID:      "your_captcha_unit_vid",
		Key:      "your_captcha_unit_key",
		Server:   "verify_server_link_from_frontend",
		Scene:    0,
		Token:    "token_from_frontend",
		ClientIP: "client_ip_address",
	}
	
	fmt.Println(RequestAndVerify(request))
```


