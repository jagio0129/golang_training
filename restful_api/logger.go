package main

import "time"
import "log"
import "github.com/julienschmidt/httprouter"
import "net/http"

var logger = func(method, uri, name string, start time.Time) {
	log.Printf("\"method\":%q \"uri\":%q \"name\":%q \"time\":%q", method, uri, name, time.Since(start))
}

func Logging(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		logger(r.Method, r.URL.Path, name, start)
	}
}
