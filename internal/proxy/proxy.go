package proxy

import (
    "io"
    "log"
    "net/http"
    "net/url"
)

// Backend server list (can add more for load balancing)
var backends = []string{
    "http://localhost:9000",
    "http://localhost:9001",
}

// Round-robin index
var current = 0

// pickBackend selects backend in round-robin fashion
func pickBackend() string {
    backend := backends[current%len(backends)]
    current++
    return backend
}

// Start launches the reverse proxy
func Start(addr string) {
    handler := func(w http.ResponseWriter, r *http.Request) {
        backendURL := pickBackend()

        target, err := url.Parse(backendURL)
        if err != nil {
            http.Error(w, "Bad backend URL", http.StatusInternalServerError)
            return
        }

        // Forward request to backend
        resp, err := http.Get(target.String() + r.URL.Path)
        if err != nil {
            http.Error(w, "Backend not reachable", http.StatusBadGateway)
            return
        }
        defer resp.Body.Close()

        // Copy response headers
        for k, v := range resp.Header {
            w.Header()[k] = v
        }
        w.WriteHeader(resp.StatusCode)

        // Copy response body
        io.Copy(w, resp.Body)
    }

    log.Println("Mirro running on", addr)
    http.ListenAndServe(addr, http.HandlerFunc(handler))
}