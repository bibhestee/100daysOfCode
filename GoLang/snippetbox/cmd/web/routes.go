package main

import (
  "net/http"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/ui"
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

  fileServer := http.FileServer(http.FS(ui.Files))

  router.Handler(http.MethodGet, "/static/*filepath", fileServer)

  router.HandlerFunc(http.MethodGet, "/ping", ping)

  dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

  router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
  router.Handler(http.MethodGet, "/about", dynamic.ThenFunc(app.about))
  protected := dynamic.Append(app.requireAuthentication)

  // Snippet Routes
  router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
  router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
  router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
  router.Handler(http.MethodGet, "/account/view", protected.ThenFunc(app.account))
  router.Handler(http.MethodGet, "/account/password/update", protected.ThenFunc(app.accountPasswordUpdate))
  router.Handler(http.MethodPost, "/account/password/update", protected.ThenFunc(app.accountPasswordUpdatePost))

  // User Routes
  router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
  router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
  router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
  router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))
  router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))

  standard := alice.New(app.recoverPanic, app.logRequest, secureHeader)

  return standard.Then(router)
}
