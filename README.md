This is google golang module for google re-captcha.


Installation
----------------

```
go get github.com/kerkerj/recaptcha

```


Usage
----------------

Usage example can be found in example/main.go file.



import

```go
import "github.com/kerkerj/recaptcha"
```

setup keys

```go
// Setup siteKey and secretKey
sitekey := "{Your site key here}"
re := recaptcha.R{
	Secret: "{Your secret here}",
}
```

verify

```go
// get recaptcha response, and then verify it.
challenge := r.FormValue("g-recaptcha-response")
isValid := re.Verify(challenge)
```



Documentation
----------------

Available on godocs: http://godoc.org/github.com/kerkerj/recaptcha

Google reCAPTCHA guides: https://developers.google.com/recaptcha/intro
