package main

import (
  "errors"
  "fmt"
  "net/http"
  "strconv"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/validator"
  "github.com/julienschmidt/httprouter"
)

type snippetCreateForm struct {
  Title string `form:"title"`
  Content string `form:"content"`
  Expires int `form:"expires"`
  validator.Validator `form:"-"`
}

type userSignupForm struct {
  Name string `form:"name"`
  Email string `form:"email"`
  Password string `form:"password"`
  validator.Validator `form:"-"`
}

type userLoginForm struct {
  Email string `form:"email"`
  Password string `form:"password"`
  validator.Validator `form:"-"`
}


func (app *application) home(res http.ResponseWriter, req *http.Request) {

  // Query the database
  snippets, err := app.snippets.Latest()
  if err != nil {
    app.serverError(res, err)
    return
  }

  data := app.newTemplateData(req)
  data.Snippets = snippets

  app.render(res, http.StatusOK, "home.html", data)
}


func (app *application) snippetView(res http.ResponseWriter, req *http.Request) {
  params := httprouter.ParamsFromContext(req.Context())

  id, err := strconv.Atoi(params.ByName("id"))
  if err != nil || id < 1 {
    app.notFound(res)
    return
  }

  snippet, err := app.snippets.Get(id)
  if err != nil {
    if errors.Is(err, models.ErrNoRecord) {
      app.notFound(res)
    } else {
      app.serverError(res, err)
    }
    return
  }

  data := app.newTemplateData(req)
  data.Snippet = snippet

  app.render(res, http.StatusOK, "view.html", data)
}


func (app *application) snippetCreate(res http.ResponseWriter, req *http.Request) {
  data := app.newTemplateData(req)
  data.Form = snippetCreateForm{
    Expires: 365,
  }
  app.render(res, http.StatusOK, "create.html", data)
}


func (app *application) snippetCreatePost(res http.ResponseWriter, req *http.Request) {
  var form snippetCreateForm

  err := app.decodePostForm(req, &form)
  if err != nil {
    app.clientError(res, http.StatusBadRequest)
    return
  }

  // Validate data
  form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")

  form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long.")
  form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
  form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")


  if !form.Valid() {
    data := app.newTemplateData(req)
    data.Form = form
    app.render(res, http.StatusUnprocessableEntity, "create.html", data)
    return
  }

  // Add form data to database
  id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
  if err != nil {
    app.serverError(res, err)
  }
  app.sessionManager.Put(req.Context(), "flash", "Snippet successfully created!")

  http.Redirect(res, req, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}


func (app *application) userSignup(res http.ResponseWriter, req *http.Request) {
  data := app.newTemplateData(req)
  data.Form = userSignupForm{}
  app.render(res, http.StatusOK, "signup.html", data)
}

func (app *application) userSignupPost(res http.ResponseWriter, req *http.Request) {
  var form userSignupForm

  err := app.decodePostForm(req, &form)
  if err != nil {
    app.clientError(res, http.StatusBadRequest)
    return
  }

  // Validate
  form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
  form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
  form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
  form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
  form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")


  if !form.Valid() {
    data := app.newTemplateData(req)
    data.Form = form
    app.render(res, http.StatusUnprocessableEntity, "signup.html", data)
  }
  // Add form data to database
  err = app.users.Insert(form.Name, form.Email, form.Password)
  if err != nil {
    if errors.Is(err, models.ErrDuplicateEmail) {
      form.AddFieldError("email", "Email address is already in use")
      data := app.newTemplateData(req)
      data.Form = form
      app.render(res, http.StatusUnprocessableEntity, "signup.html", data)
    } else {
      app.serverError(res, err)
    }
    return
  }

  app.sessionManager.Put(req.Context(), "flash", "Your signup was successful. Please log in.")
  http.Redirect(res, req, "/user/login", http.StatusSeeOther)
}

func (app *application) userLogin(res http.ResponseWriter, req *http.Request) {
  data := app.newTemplateData(req)
  data.Form = userLoginForm{}
  app.render(res, http.StatusOK, "login.html", data)
}

func (app *application) userLoginPost(res http.ResponseWriter, req *http.Request) {
  var form userLoginForm

  err := app.decodePostForm(req, &form)
  if err != nil {
    app.clientError(res, http.StatusBadRequest)
    return
  }

  form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
  form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
  form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")

  if !form.Valid() {
    data := app.newTemplateData(req)
    data.Form = form
    app.render(res, http.StatusUnprocessableEntity, "login.html", data)
  }

  id, err := app.users.Authenticate(form.Email, form.Password)
  if err != nil {
    if errors.Is(err, models.ErrInvalidCredentials) {
      form.AddNonFieldError("Email or password is incorrect")
      data := app.newTemplateData(req)
      data.Form = form
      app.render(res, http.StatusUnprocessableEntity, "login.html", data)
    } else {
      app.serverError(res, err)
    }
    return
  }

  err = app.sessionManager.RenewToken(req.Context())
  if err != nil {
    app.serverError(res, err)
    return
  }

  app.sessionManager.Put(req.Context(), "authenticatedUserID", id)

  http.Redirect(res, req, "/snippet/create", http.StatusSeeOther)
}

func (app *application) userLogoutPost(res http.ResponseWriter, req *http.Request) {
  err := app.sessionManager.RenewToken(req.Context())
  if err != nil {
    app.serverError(res, err)
  }

  app.sessionManager.Remove(req.Context() ,"authenticatedUserID")
  app.sessionManager.Put(req.Context(), "flash", "You've been logged out successfully!")
  http.Redirect(res, req, "/", http.StatusSeeOther)
}

