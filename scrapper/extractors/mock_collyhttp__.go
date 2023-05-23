package extractor

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
<a class="top-card-layout__title">This is a test page</p>
<p class="topcard__org-name-link">This is a test paragraph</p>
<span class="topcard__flavor topcard__flavor--bullet">loca</span>
<p class="posted-time-ago__text topcard__flavor--metadata">age</p>
<div class="description__text description__text--rich">
<div class="show-more-less-html__markup">
desc
</div>
</div>
<a class="topcard__org-name-link topcard__flavor--black-link">link</a>

<ul class="description__job-criteria-list">
<li> 
<span class="description__job-criteria-subheader">Seniority level</span>
<span class="description__job-criteria-text description__job-criteria-text--criteria">asd</span>
</li>
<li> 
<span class="description__job-criteria-subheader">Employment type</span>
<span class="description__job-criteria-text description__job-criteria-text--criteria">asd</span>
</li>
<li> 
<span class="description__job-criteria-subheader">Job function</span>
<span class="description__job-criteria-text description__job-criteria-text--criteria">adasd</span>
</li>
<li> 
<span class="description__job-criteria-subheader">Industries</span>
<span class="description__job-criteria-text description__job-criteria-text--criteria">asdsad</span>
</li>
</ul>
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
