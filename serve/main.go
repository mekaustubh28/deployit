package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const BASE_URL = "http://localhost:8000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname := r.Host
		subdomain := strings.Split(hostname, ".")[0]

		resolveTo := BASE_URL + "/" + subdomain

		target, err := url.Parse(resolveTo)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ModifyResponse = func(response *http.Response) error {
			return nil
		}

		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Listening at :3000")
	http.ListenAndServe(":3000", nil)
}
