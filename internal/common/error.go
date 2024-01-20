package common

type ValidationError struct {
	Message string `json:"message"`
}

func (v ValidationError) Error() string {
	return v.Message
}
