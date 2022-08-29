package data

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

//Response create a type check the response is ok
type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// ReadJSON  allows receiving generic json data from request body,
// provide clean way to read any kind of json from a request, convert it to go data structure
//assuming that request body has only a single json value
func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// maxBytes define the max size of request allowed
	//prevents clients from accidentally or maliciously sending a large
	//request and wasting server resources.
	maxBytes := 1048756
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	//read the request body
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data) //decode body into data
	if err != nil {
		return err
	}

	// assume to decode a json file that has one entry
	//decode body into data
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

//WriteJSON output data as a json to the browser,
//takes in w so where to write response to,
// status code which set the status code for the response,
//the data that will be converted to json, and
// headers to be sent
func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {

	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	json, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(json)
	if err != nil {
		return err
	}
	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	WriteJSON(w, statusCode, theError, "error")
}
