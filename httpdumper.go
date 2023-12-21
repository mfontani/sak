package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"sort"
	"time"
)

func sortedHeadersTable(r *http.Request) [][]string {
	var table [][]string
	var sortedHeaders []string
	for h := range r.Header {
		sortedHeaders = append(sortedHeaders, h)
	}
	sort.Strings(sortedHeaders)
	for _, h := range sortedHeaders {
		var hv = []string{h, r.Header.Get(h)}
		table = append(table, hv)
	}
	return table
}

// HTTPDumper handles "httpdump LISTEN_HOST LISTEN_PORT" command.
func HTTPDumper(args []string) {
	RequireArgs(2, args)
	listenOn := fmt.Sprintf("%s:%s", args[0], args[1])
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         listenOn,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Pragma", "no-cache")
		for _, h := range sortedHeadersTable(r) {
			_, err := fmt.Fprintf(w, "'%s': '%s'\n", h[0], h[1])
			if err != nil {
				fmt.Println("error:", err)
				return
			}
		}
		end := time.Now()
		elapsed := end.Sub(start)
		dump, err := httputil.DumpRequest(r, true)
		if err == nil {
			fmt.Printf("%s from %s in %s: %q\n",
				end.Format(time.RFC3339Nano), r.RemoteAddr, elapsed, dump)
		}
	}
	http.HandleFunc("/", handler)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
