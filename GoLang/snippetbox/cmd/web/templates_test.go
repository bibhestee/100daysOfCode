package main

import (
  "testing"
  "time"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/assert"
)


func TestHumanDate(t *testing.T) {
  tests := []struct{
    name string
    tm time.Time
    want string
  }{
    {
      name: "UTC",
      tm: time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
      want: "17 Mar 2024 at 10:15am",
    },
    {
      name: "Empty",
      tm: time.Time{},
      want: "",
    },
    {
      name: "CET",
      tm: time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
      want: "17 Mar 2024 at 9:15am",
    },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      hd := humanDate(tt.tm)
      assert.Equal(t, hd, tt.want)
    })
  }
}
