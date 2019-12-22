package main
import (
  l "log"
  f "fmt"
  h "net/http"
  o "os"
)

func determineListenAddress() (string, error) {
  port := o.Getenv("PORT")
  if port == "" {
    return "", f.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func hello(w h.ResponseWriter, r *h.Request) {
  f.Fprintln(w, "Hello World")
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    l.Fatal(err)
  }
  h.HandleFunc("/", hello)
  l.Printf("Listening on %s...\n", addr)
  if err := h.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
