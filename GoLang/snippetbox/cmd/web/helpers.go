package main

import (
  "fmt"
  "net/http"
  "runtime/debug"
)

func (app *application) serverError(res http.ResponseWriter, err error) {
  trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
  app.errorLog.Output(2, trace)

  http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}


func (app *application) clientError(res http.ResponseWriter, status int) {
  http.Error(res, http.StatusText(status), status)
}


func (app *application) notFound(res http.ResponseWriter) {
  app.clientError(res, http.StatusNotFound)
}


func (app *application) render(res http.ResponseWriter, status int, page string, data *templateData) {
  ts, ok := app.templateCache[page]
  if !ok {
    err := fmt.Errorf("the template %s does not exist", page)
    app.serverError(res, err)
    return
  }

  res.WriteHeader(status)

  err := ts.ExecuteTemplate(res, "base", data)
  if err != nil {
    app.serverError(res, err)
  }

}
