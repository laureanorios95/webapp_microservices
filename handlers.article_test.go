// handlers.article_test.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)

	for _, article := range articleList {
		req, _ := http.NewRequest("GET", "/article/view/"+fmt.Sprint(article.ID), nil)

		testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
			statusOK := w.Code == http.StatusOK

			p, err := ioutil.ReadAll(w.Body)
			pageOK := err == nil && strings.Index(string(p), "<h1>"+article.Title+"</h2>") > 0 &&
				strings.Index(string(p), "<p>"+article.Content+"</p>") > 0

			return statusOK && pageOK
		})
	}
}

func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)
	r.GET("/", showIndexPage)

	var should = make([]string, 0, len(articleList))

	for _, article := range articleList {
		should = append(should, "{"+"\"id\":\""+fmt.Sprint(article.ID)+
			"\",\"title\":\""+article.Title+"\",\"content\":\""+article.Content+"\"}")
	}

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p),
			strings.Join(should, ",")) > 0

		return statusOK && pageOK
	})
}

func TestArticleXML(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)

	for _, article := range articleList {
		req, _ := http.NewRequest("GET", "/article/view/"+fmt.Sprint(article.ID), nil)
		req.Header.Set("Accept", "application/json")

		testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
			statusOK := w.Code == http.StatusOK

			p, err := ioutil.ReadAll(w.Body)
			pageOK := err == nil && strings.Index(string(p),
				"{ id:"+fmt.Sprint(article.ID)+", title:"+article.Title+
					", content:"+article.Content+" }") > 0

			return statusOK && pageOK
		})
	}
}
