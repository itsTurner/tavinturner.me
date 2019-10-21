package main
import ."net/http"
func main() {
  HandleFunc("/", func(w ResponseWriter, r *Request) {
    ServeFile(w, r, "app/index.html")
  })
  Handle("/css/", StripPrefix("/css/", FileServer(Dir("app/styles/"))))
  ListenAndServe(":8000", nil)
}
