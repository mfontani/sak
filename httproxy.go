package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// HTTPProxy handles "httproxy LISTEN_HOST LISTEN_PORT FWD_HOST FWD_PORT" command.
func HTTPProxy(args []string) {
	RequireArgs(4, args)
	listenOn := fmt.Sprintf("%s:%s", args[0], args[1])
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         listenOn,
	}
	hostPort := fmt.Sprintf("%s:%s", args[2], args[3])
	httpHostPort := fmt.Sprintf("http://%s", hostPort)
	remote, err := url.Parse(httpHostPort)
	if err != nil {
		panic(err)
	}
	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			origHost := r.Host
			r.Host = hostPort
			p.ServeHTTP(w, r)
			end := time.Now()
			elapsed := end.Sub(start)
			dump, err := httputil.DumpRequest(r, true)
			if err == nil {
				fmt.Printf("%s from %s in %s: [Host %q] %q\n",
					end.Format(time.RFC3339Nano), r.RemoteAddr, elapsed, origHost, dump)
			}
		}
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
