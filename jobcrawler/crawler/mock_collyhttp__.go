package crawler

import (
	"net/http"
	"net/http/httptest"
)

func newUnstartedTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
<title>Test Page</title>
</head>
<body>
<h1>Hello World</h1>
<a href="https://linkedin.com/jobs/view/sad">This is a test page</p>
<p class="description">This is a test paragraph</p>
</body>
</html>
		`))
	})

	return httptest.NewUnstartedServer(mux)
}

func newTestServer() *httptest.Server {
	srv := newUnstartedTestServer()
	srv.Start()
	return srv
}

func newUnstartedTestServerWithError() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/json")
		w.WriteHeader(http.StatusNotFound)
	})

	return httptest.NewUnstartedServer(mux)
}

func newTestServerWithError() *httptest.Server {
	srv := newUnstartedTestServerWithError()
	srv.Start()
	return srv
}
