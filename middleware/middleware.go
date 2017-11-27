package middleware

import (
	"fmt"
	"net/http"
)

func TestMiddle(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "MAN IN THE MIDDLE")
		f.ServeHTTP(w, r)
	})
}

func SecondMiddle(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.ServeHTTP(w, r)
		fmt.Fprintln(w, "SECONDO")
	})
}
