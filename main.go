package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func requestDump(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request\n%s\n\n", strings.Repeat("#", len("Request")))
	fmt.Fprintf(w, "%v %v %v\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host: %v\n", r.Host)

	keys := make([]string, len(r.Header))
	for name := range r.Header {
		keys = append(keys, name)
	}

	sort.Strings(keys)

	for _, key := range keys {
		for _, h := range r.Header[key] {
			fmt.Fprintf(w, "%v: %v\n", key, h)
		}
	}

	fmt.Fprintf(w, "\n\nRemoteAddr\n%s\n\n", strings.Repeat("#", len("RemoteAddr")))
	fmt.Fprintln(w, r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", requestDump)
	http.ListenAndServe(":8080", nil)
}
