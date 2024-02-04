package common

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"text/template"
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

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		return err
	}
	switch data.(type) {
	case error:
		_ = tmpl.Execute(w, struct {
			Error string
		}{data.(error).Error()})
	default:
		err = tmpl.Execute(w, data)
		if err != nil {
			_ = tmpl.Execute(w, struct {
				Error string
			}{err.Error()})
		}
	}

	return nil
}

func MapFormToStruct(values url.Values, v interface{}) error {
	// Get the type of the struct
	t := reflect.TypeOf(v).Elem()
	// Get the value of the struct
	val := reflect.ValueOf(v).Elem()
	// Loop through the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		// Get the field
		field := t.Field(i)
		// Get the tag value
		tag := field.Tag.Get("json")
		// Get the value from the form
		value := values.Get(tag)
		// Set the value of the field
		val.Field(i).SetString(value)
	}
	return nil
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
