package config

import (
	"errors"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/eapache/go-resiliency/breaker"
)

var (
	// circuitbreaker
	cb *breaker.Breaker
)

// CaptureNegroniHandler handle panic on negroni handler.
func CaptureNegroniHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	request, _ := httputil.DumpRequest(r, true)
	defer func() {
		if !recoveryBreak() {
			r := panicRecover(recover())
			if r != nil {
				log.Println(errors.New(r.Error()), string(request))
				http.Error(w, r.Error(), http.StatusInternalServerError)
			}
		}
	}()
	next(w, r)
}

func recoveryBreak() bool {
	if cb == nil {
		return false
	}

	if err := cb.Run(func() error {
		return nil
	}); err == breaker.ErrBreakerOpen {
		return true
	}
	return false
}

func panicRecover(rc interface{}) error {
	if cb != nil {
		r := cb.Run(func() error {
			return recovery(rc)
		})
		return r
	}
	return recovery(rc)
}

func recovery(r interface{}) error {
	var err error
	if r != nil {
		switch t := r.(type) {
		case string:
			err = errors.New(t)
		case error:
			err = t
		default:
			err = errors.New("Unknown error")
		}
	}
	return err
}
