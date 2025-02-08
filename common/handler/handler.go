package handler

import (
	"encoding/json"
	"net/http"
	"reflect"

	log "github.com/sirupsen/logrus"
)

// todo: setup logger for package

type CoreService struct {
	R       *http.Request
	W       http.ResponseWriter
	Service any // for any services
}

func handleError(r *http.Request, w http.ResponseWriter, err error) {
	log.WithFields(log.Fields{"method": r.Method, "url": r.URL.Path}).Error(err)

	w.Header().Set("Content-Type", "application/json")

	if httpError, ok := err.(HttpCodeError); ok {
		w.WriteHeader(httpError.Code)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	b, err := json.Marshal(map[string]string{"error": err.Error()})
	if err != nil {
		log.Errorf("error on convert error to JSON - %s", err)
		return
	}

	if _, err := w.Write([]byte(b)); err != nil {
		log.Errorf("error on write error to response - %s", err)
	}
}

func handleResult(w http.ResponseWriter, result any) {
	if result == nil || reflect.ValueOf(result).IsNil() {
		w.WriteHeader(http.StatusOK)
		return
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Errorf("error on convert result to JSON - %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(b)
	if err != nil {
		log.Errorf("error on write result to response - %s", err)
		return
	}

}

// CreateHandler - приводит функцию к http.HanlderFunc, позволяя в хендлерах использовать CoreService.
// Так же служит единной точкой обратки респонсов.
func CreateHandler(fn func(CoreService) (any, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := fn(CoreService{R: r, W: w})
		if err != nil {
			handleError(r, w, err)
		} else {
			handleResult(w, result)
		}
	}
}
