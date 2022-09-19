package checker

import (
	"github.com/ahmedkhaeld/bookings-api/internals/data"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Response struct {
	Error   bool              `json:"error"`
	Errors  map[string]string `json:"errors"`
	Message string            `json:"message"`
}

func InvalidCredentials(w http.ResponseWriter) error {
	var res Response

	res.Error = true
	res.Message = "Invalid Auth Credentials"

	err := data.WriteJSON(w, http.StatusUnauthorized, res, "response")
	if err != nil {
		return err
	}
	return nil
}

// PasswordMatches validate user password, takes 2 args that will be compared against each other, using the bcrypt pkg
//hash is what is pulled out of the db, and the password
// user entered on the input field,
func PasswordMatches(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

//FailedValidation call when validation fails
func FailedValidation(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	var res Response

	res.Error = true
	res.Message = "failed validation"
	res.Errors = errors
	data.WriteJSON(w, http.StatusUnprocessableEntity, res, "response")
}
