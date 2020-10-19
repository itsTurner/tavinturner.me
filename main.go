package main

import (
	e "errors"
	i "io/ioutil"
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
		return "", e.New("$PORT not set")
	}
	return ":" + port, nil
}

func readFileAsString(path string) (string, error) {
	bytes, err := i.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func routesFromSSV(fileString string, delimiter string) {
	for _, line := range s.Split(fileString, "\n") {
		parts := s.Split(line, delimiter)
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
}

func main() {
	routes, err := readFileAsString("routes.ssv")
	if err != nil {
		panic(err)
	}

	routesFromSSV(routes, " ")

	addr, err := determineListenAddress()
	if err != nil {
		panic(err)
	}

	if err := h.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
