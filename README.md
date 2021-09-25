# vaptcha-go
 This is a third-party golang SDK for [Vaptcha](https://www.vaptcha.com/).

### Quick Usage
1. get this module by command line

`go get github.com/Bisstocuz/vaptcha-go`

2. import this module

`import "github.com/Bisstocuz/vaptcha-go"`

3. code with this module

```
	result := vaptcha.RequestAndVerify(&vaptcha.CaptchaRequest{
		VID:      "your_captcha_unit_vid",
		Key:      "your_captcha_unit_key",
		Server:   "verify_url_from_frontend",
		Scene:    0,
		Token:    "token_from_frontend",
		ClientIP: "client_ip_address",
	})

	fmt.Println(result)
```

### More Usages
Check the docs: [pkg.go.dev](https://pkg.go.dev/github.com/Bisstocuz/vaptcha-go)
