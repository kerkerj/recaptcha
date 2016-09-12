package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// R type represents an object of Recaptcha and has public property Secret,
// which is secret obtained from google recaptcha tool admin interface
type R struct {
	Secret    string
	lastError []string
}

// Struct for parsing json in google's response
type googleResponse struct {
	Success    bool
	ErrorCodes []string `json:"error-codes"`
}

// url to post submitted re-captcha response to
var postURL = "https://www.google.com/recaptcha/api/siteverify"

func doPostRequest(postUrl string, strToVerify string, r *R) (*http.Response, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.PostForm(postUrl, url.Values{"secret": {r.Secret}, "response": {strToVerify}})

	return resp, err
}

// Verify method, verifies if the response string is valid re-captcha response and returns true or false
// This method also records any errors in validation.
// These errors can be received by calling LastError() method.
func (r *R) Verify(response string) bool {
	r.lastError = make([]string, 1)

	resp, err := doPostRequest(postURL, response, r)
	if err != nil {
		r.lastError = append(r.lastError, err.Error())
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.lastError = append(r.lastError, err.Error())
		return false
	}

	gr := new(googleResponse)
	if err = json.Unmarshal(body, gr); err != nil {
		r.lastError = append(r.lastError, err.Error())
		return false
	}

	if !gr.Success {
		r.lastError = append(r.lastError, gr.ErrorCodes...)
	}
	return gr.Success
}

// LastError returns errors occurred in last re-captcha validation attempt
func (r R) LastError() []string {
	return r.lastError
}
