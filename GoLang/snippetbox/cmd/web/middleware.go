package main

import (
  "context"
  "fmt"
  "net/http"
  "github.com/justinas/nosurf"
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


func (app *application) recoverPanic(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    defer func() {
      if err := recover(); err != nil {
        res.Header().Set("Connection", "close")
        app.serverError(res, fmt.Errorf("%s", err))
      }

    }()
    next.ServeHTTP(res, req)
  })

}
func (app *application) requireAuthentication(next http.Handler) http.Handler {
    return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
      if !app.isAuthenticated(req) {
        app.sessionManager.Put(req.Context(), "urlPath", req.RequestURI)
        http.Redirect(res, req, "/user/login", http.StatusSeeOther)
        return
      }

      res.Header().Add("Cache-Control", "no-store")
      next.ServeHTTP(res, req)
    })
}

func noSurf(next http.Handler) http.Handler {
  csrfHandler := nosurf.New(next)
  csrfHandler.SetBaseCookie(http.Cookie{
    HttpOnly: true,
    Path: "/",
    Secure: true,
  })

  return csrfHandler
}

func (app *application) authenticate(next http.Handler) http.Handler {
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    id := app.sessionManager.GetInt(req.Context(), "authenticatedUserID")
    if id == 0 {
      next.ServeHTTP(res, req)
      return
    }

    exists, err := app.users.Exists(id)
    if err != nil {
      app.serverError(res, err)
      return
    }

    if exists {
      ctx := context.WithValue(req.Context(), isAuthenticatedContextKey, true)
      req = req.WithContext(ctx)
    }
    next.ServeHTTP(res, req)
  })
}
