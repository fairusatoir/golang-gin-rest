package utils

type Result struct {
	StatusCode int
	Err        error
	Data       interface{}
}

func NewResult(d interface{}, sc int, err error) *Result {
	return &Result{
		StatusCode: sc,
		Err:        err,
		Data:       d,
	}
}
