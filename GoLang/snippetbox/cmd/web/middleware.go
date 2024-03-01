package main

import (
  "net/http"
)


func secureHeader(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
    res.Header().Set("Referrer-Policy", "origin-when-cross-origin")
    res.Header().Set("X-Content-Type-Options", "no-sniff")
    res.Header().Set("X-Frame-Options", "deny")
    res.Header().Set("X-XSS-Protection", "0")

    next.ServeHTTP(res, req)
  })
}


func (app *application) logRequest(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
      app.infoLog.Printf("%s - %s %s %s", req.RemoteAddr, req.Proto, req.Method, req.URL.RequestURI())

    next.ServeHTTP(res, req)
  })
}
