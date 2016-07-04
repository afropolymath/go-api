package main

import (
  "log"
  "net/http"
  "time"
)

func Logger(inner http.Handler, name String) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, *r http.Request) {
    start := time.Now()
    inner
  })
}