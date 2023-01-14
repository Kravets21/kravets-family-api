package rest

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	baseErrors "kravets-family-api/pkg/errors"
	"net/http"
	"reflect"
)

type structResponse struct {
	Data  interface{} `json:"data"`
	Extra interface{} `json:"extra,omitempty"`
}

func StructResponse(w http.ResponseWriter, s interface{}, extra interface{}) {
	response(w, s, extra, http.StatusOK)
}

func response(w http.ResponseWriter, s interface{}, extra interface{}, code int) {
	if s == nil || (reflect.ValueOf(s).Kind() == reflect.Ptr && reflect.ValueOf(s).IsNil()) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(structResponse{
		Data:  s,
		Extra: extra,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}

func InternalServerErrorResponse(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func BadRequestResponse(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func ErrResponse(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	errCause := errors.Cause(err)
	if errCause == baseErrors.ErrDBNotOneRowsAffected || errCause == baseErrors.ErrDBRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	InternalServerErrorResponse(w, err)
}
