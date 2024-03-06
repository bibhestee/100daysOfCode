package main

import (
  "bytes"
  "errors"
  "fmt"
  "net/http"
  "runtime/debug"
  "time"
  "github.com/go-playground/form/v4"
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

  buf := new(bytes.Buffer)

  err := ts.ExecuteTemplate(buf, "base", data)
  if err != nil {
    app.serverError(res, err)
  }

  res.WriteHeader(status)

  buf.WriteTo(res)

}


func (app *application) newTemplateData(req *http.Request) *templateData {
  return &templateData{
    CurrentYear: time.Now().Year(),
    Flash: app.sessionManager.PopString(req.Context(), "flash"),
    isAuthenticated: app.isAuthenticated(req),
  }
}

func (app *application) decodePostForm(req *http.Request, dst any) error {
  err := req.ParseForm()
  if err != nil {
    return err
  }

  err = app.formDecoder.Decode(dst, req.PostForm)
  if err != nil {
    var InvalidDecoderError *form.InvalidDecoderError

    if errors.As(err, &InvalidDecoderError) {
      panic(err)
    }
    return err
  }

  return nil
}

func (app *application) isAuthenticated(req *http.Request) bool {
  return app.sessionManager.Exists(req.Context(), "authenticatedUserID")
}
