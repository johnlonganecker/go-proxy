// helpful resource with additional content:
// http://stackoverflow.com/questions/21270945/how-to-read-the-response-from-a-newsinglehostreverseproxy
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

const BaseUrl = "http://localhost:8000"
const ListeningPort = "8888"

func main() {
	// intercept call
	http.HandleFunc("/sayblah", SayBlah)

	// all other traffic pass on
	http.HandleFunc("/", ProxyFunc)
	http.ListenAndServe(":"+ListeningPort, nil)
}

func ProxyFunc(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(BaseUrl)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

func SayBlah(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BLAH"))
}
