package main

import (
	f "fmt"
	i "io/ioutil"
	l "log"
	h "net/http"
	o "os"
	s "strings"
)

type Route struct {
	URL  string
	File string
}

func m(i ...interface{}) []interface{} { return i }

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

	csvBytes, _ := i.ReadFile("routes.ssv")
	csvString := string(csvBytes)

	for _, line := range s.Split(csvString, "\n") {
		parts := s.Split(line, " ")
		if len(parts) >= 2 && parts[0][0] != '#' {
			if parts[0][len(parts[0])-1] == '*' {
				path := string([]rune(parts[0])[:len(parts[0])-1])
				h.Handle(path, h.StripPrefix(path, h.FileServer(h.Dir(parts[1]))))
			} else {
				route := &Route{URL: parts[0], File: parts[1]}
				h.HandleFunc(route.URL, func(w h.ResponseWriter, r *h.Request) {
					h.ServeFile(w, r, route.File)
				})
			}
		}
	}

	if err := h.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
