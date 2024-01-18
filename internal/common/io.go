package common

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReaderToStruct(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)
	return nil
}

func RenderJSON(w http.ResponseWriter, v interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	switch v.(type) {
	case error:
		err := json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{v.(error).Error()})
		if err != nil {
			return
		}
	default:
		err := json.NewEncoder(w).Encode(v)
		if err != nil {
			return
		}
	case string:
		err := json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{v.(string)})
		if err != nil {
			return
		}
	case nil:
	}

}
