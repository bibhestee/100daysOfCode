package main

import (
  "bytes"
  "io"
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/assert"
)

func TestSecureHeaders(t *testing.T) {
  resp := httptest.NewRecorder()

  req, err := http.NewRequest(http.MethodGet, "/", nil)
  if err != nil {
    t.Fatal(err)
  }

  next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
  })

  secureHeader(next).ServeHTTP(resp, req)

  // Extract response data using Result method
  rs := resp.Result()
  // Check that middleware set content security policy
  expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
  assert.Equal(t, rs.Header.Get("Content-Security-Policy"), expectedValue)

  // Check that middleware set referrer policy
  expectedValue = "origin-when-cross-origin"
  assert.Equal(t, rs.Header.Get("Referrer-Policy"), expectedValue)

  // Check that middleware set X-Content-Type-Options
  expectedValue = "no-sniff"
  assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), expectedValue)

  // Check that middleware set X-Frame-Options
  expectedValue = "deny"
  assert.Equal(t, rs.Header.Get("X-Frame-Options"), expectedValue)

  // Check that middleware set X-XSS-Protection
  expectedValue = "0"
  assert.Equal(t, rs.Header.Get("X-XSS-Protection"), expectedValue)

  // Check that middleware has correctly called the next handler
  assert.Equal(t, rs.StatusCode, http.StatusOK)

  defer rs.Body.Close()
  body, err := io.ReadAll(rs.Body)
  if err != nil {
    t.Fatal(err)
  }
  bytes.TrimSpace(body)
  assert.Equal(t, string(body), "OK")
  }


func TestMethodNotAllowed(t *testing.T) {
  app := newTestApplication(t)

  ts := newTestServer(t, app.routes())

  defer ts.Close()

  code, _, body := ts.postForm(t, "/", nil)
  // Check if status code is equal to http.StatusMethodNotAllowed
  assert.Equal(t, code, http.StatusMethodNotAllowed)
  // Check if response is Method Not Allowed
  assert.Equal(t, string(body), "Method Not Allowed\n")
}
