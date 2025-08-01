package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequest(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
