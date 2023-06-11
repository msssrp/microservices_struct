package api

import "net/http"

func Authorization(next http.Handler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //somethin here to handler middleware
    next.ServeHTTP(w, r)
  })
}
