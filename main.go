package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/suconghou/goexcel/route"
	"github.com/suconghou/goexcel/util"
)

func main() {
	var (
		port int
		host string
	)
	flag.IntVar(&port, "p", 6060, "listen port")
	flag.StringVar(&host, "h", "", "bind address")
	flag.Parse()
	if err := serve(host, port); err != nil {
		util.Log.Print(err)
	}
}

func serve(host string, port int) error {
	http.HandleFunc("/", routeMatch)
	util.Log.Printf("Starting up on port %d", port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}

func routeMatch(w http.ResponseWriter, r *http.Request) {
	found := false
	for _, p := range route.Rules {
		if p.Reg.MatchString(r.URL.Path) {
			found = true
			if err := p.Handler(w, r, p.Reg.FindStringSubmatch(r.URL.Path)); err != nil {
				util.Log.Print(err)
			}
			break
		}
	}
	if !found {
		fallback(w, r)
	}
}

func fallback(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}