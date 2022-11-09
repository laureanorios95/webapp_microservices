// handlers.article_test.go

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
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

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			log.Fatal(err)
		}
		dec := json.NewDecoder(strings.NewReader(string(p)))

		// read open bracket
		t, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v\n", t, t)

		var articleStruct article

		for i := 0; dec.More(); i++ {
			err = dec.Decode(&articleStruct)
			if err != nil {
				log.Fatal(err)
			}
			if articleStruct != articleList[i] {
				return false
			}
		}

		// pageOK := err == nil && strings.Index(string(p),
		// 	articleList) > 0

		return statusOK //&& pageOK
	})
}

func TestArticleXML(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)

	var articleStruct article

	for _, ar := range articleList {
		req, _ := http.NewRequest("GET", "/article/view/"+fmt.Sprint(ar.ID), nil)
		req.Header.Set("Accept", "application/xml")

		testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
			statusOK := w.Code == http.StatusOK

			p, err := ioutil.ReadAll(w.Body)
			if err != nil {
				log.Fatal(err)
			}
			if err := xml.Unmarshal(p, &articleStruct); err != nil {
				log.Fatal(err)
			}
			if articleStruct != articleList[ar.ID-1] {
				return false
			}

			return statusOK
		})
	}
}
