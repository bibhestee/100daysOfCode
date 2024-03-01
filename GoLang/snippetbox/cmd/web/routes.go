package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
  "github.com/justinas/alice"
)


func (app *application) routes() http.Handler {
  router := httprouter.New()

  router.NotFound = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    app.notFound(res)
  })

  router.MethodNotAllowed = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    app.clientError(res, http.StatusMethodNotAllowed)
  })

  fileServer := http.FileServer(http.Dir("./ui/static"))

  router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

  router.HandlerFunc(http.MethodGet, "/", app.home)
  router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
  router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
  router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

  return alice.New(app.recoverPanic, app.logRequest, secureHeader).Then(router)
}
