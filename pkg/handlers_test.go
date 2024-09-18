package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func assertEquality(t *testing.T, obj1 any, obj2 any) {
	t.Helper()
	if reflect.TypeOf(obj1) != reflect.TypeOf(obj2) {
		t.Errorf("Error! type mismatch, wanted: %t got: %t", reflect.TypeOf(obj1), reflect.TypeOf(obj2))
	}
	if !reflect.DeepEqual(obj1, obj2) {
		t.Errorf("Error! values mismatch, wanted: %v got: %v", obj1, obj2)
	}
}

func TestHttpHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	HttpHome(w, req)
	resp := w.Result()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body %v", err)
	}

	expected := "Welcome to my datetime server!"

	assertEquality(t, expected, string(resBody))
	assertEquality(t, 200, resp.StatusCode)
}

func TestHttpHandler(t *testing.T) {
	truncatedTime := truncateToSec(time.Now())
	timeJson, err := json.Marshal(truncatedTime)
	if err != nil {
		t.Errorf("error converting to json: %v", err)
	}
	testcases := []struct {
		testcaseName string
		method       string
		url          string
		statusCode   int
		expected     any
		contentType  string
	}{
		{
			testcaseName: "correct method and url, plain text type",
			method:       "GET",
			url:          "/datetime",
			statusCode:   200,
			expected:     truncatedTime.String(),
			contentType:  "text/plain",
		},
		{
			testcaseName: "correct method and url, json type",
			method:       "GET",
			url:          "/datetime",
			statusCode:   200,
			expected:     timeJson,
			contentType:  "application/json",
		},
		{
			testcaseName: "wrong method",
			method:       "POST",
			url:          "/datetime",
			statusCode:   405,
			expected:     http.StatusText(http.StatusMethodNotAllowed) + "\n",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.testcaseName, func(t *testing.T) {

			req := httptest.NewRequest(testcase.method, testcase.url, nil)
			req.Header.Add("content-type", testcase.contentType)

			w := httptest.NewRecorder()
			HttpHandler(w, req)
			resp := w.Result()
			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Error reading response body %v", err)
			}
			if testcase.contentType == "application/json" {
				assertEquality(t, testcase.expected, resBody)
			} else {
				assertEquality(t, testcase.expected, string(resBody))
			}
			assertEquality(t, testcase.statusCode, resp.StatusCode)

		})
	}

}

func TestGinHome(t *testing.T) {

	r := gin.Default()
	r.GET("/", GinHome)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	expected := "Welcome to my datetime server!"
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body %v", err)
	}
	assertEquality(t, expected, string(resBody))
	assertEquality(t, 200, res.Code)
}

func TestGinHandler(t *testing.T) {

	testcases := []struct {
		testcaseName string
		method       string
		url          string
		statusCode   int
		expected     string
	}{
		{
			testcaseName: "correct method and url",
			method:       "GET",
			url:          "/datetime",
			statusCode:   200,
			expected:     truncateToSec(time.Now()).String(),
		},
		{
			testcaseName: "wrong method",
			method:       "POST",
			url:          "/datetime",
			statusCode:   405,
			expected:     http.StatusText(http.StatusMethodNotAllowed),
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.testcaseName, func(t *testing.T) {

			r := gin.Default()
			if testcase.method == "GET" {

				r.GET(testcase.url, GinHandler)
			} else {
				r.POST(testcase.url, GinHandler)

			}

			req, err := http.NewRequest(testcase.method, testcase.url, nil)
			if err != nil {
				t.Errorf("Error in new request %v", err)
			}
			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error reading response body %v", err)
			}
			assertEquality(t, testcase.expected, string(resBody))
			assertEquality(t, testcase.statusCode, res.Code)

		})
	}

}
