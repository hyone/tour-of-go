package main

import (
  "fmt"
  "net/http"
)

type String string

func (s String) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  fmt.Fprint(w, s)
}


type Struct struct {
  Greeting string
  Punct string
  Who string
}

func (s *Struct) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  fmt.Fprintf(
    w, "%s%s %s",
    s.Greeting, s.Punct, s.Who,
  )
}


func main () {
  http.Handle("/string", String("I'm a frayed knot."))
  http.Handle("/struct", &Struct {"Hello", ":", "Gphers!"})
  http.ListenAndServe("localhost:4000", nil)
}
