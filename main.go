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

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    l.Fatal(err)
  }
  h.HandleFunc("/", func(w h.ResponseWriter, r *h.Request) {
    h.ServeFile(w, r, "app/index.html")
  })
  h.Handle("/css/", h.StripPrefix("/css/", h.FileServer(h.Dir("app/styles/"))))
  if err := h.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
