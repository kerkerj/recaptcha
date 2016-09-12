package recaptcha

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestLastError(t *testing.T) {
	// Arrange
	testSecret := "test_recaptcha_secret"
	re := R{
		Secret: testSecret,
	}

	// Act
	lastError := re.LastError()

	// Assert
	if len(lastError) > 1 {
		t.Errorf("re.LastError() returned %s, expected %s", lastError, "[]")
	}

}

func TestVerify(t *testing.T) {
	// Arrange
	testSecret := "test_recaptcha_secret"
	testResponseChallenge := "test_recaptcha_response_challenge"

	re := R{
		Secret: testSecret,
	}

	// Act
	result := re.Verify(testResponseChallenge)

	// Assert
	if result != false {
		t.Errorf("re.Verify(%s) returned %b, expected %b", testResponseChallenge, result, false)
	}
}

func ExampleR_Verify() {
	sitekey := "{Your site key here}"
	re := R{
		Secret: "{Your secret here}",
	}

	form := fmt.Sprintf(`
		<html>
			<head>
				<script src='https://www.google.com/recaptcha/api.js'></script>
			</head>
			<body>
				<form action="/submit" method="post">
					<div class="g-recaptcha" data-sitekey="%s"></div>
					<input type="submit">
				</form>
			</body>
		</html>
	`, sitekey)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, form)
	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		response := r.FormValue("g-recaptcha-response")
		isValid := re.Verify(response)
		if isValid {
			fmt.Fprintf(w, "Valid")
		} else {
			fmt.Fprintf(w, "Invalid! These errors ocurred: %v", re.LastError())
		}
	})

	err := http.ListenAndServe(":8100", nil)

	if err != nil {
		log.Printf("Could not start server. %s", err)
	}
}
