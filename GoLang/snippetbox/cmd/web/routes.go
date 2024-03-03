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

  dynamic := alice.New(app.sessionManager.LoadAndSave)

  router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
  router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
  router.Handler(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
  router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

  standard := alice.New(app.recoverPanic, app.logRequest, secureHeader)

  return standard.Then(router)
}
