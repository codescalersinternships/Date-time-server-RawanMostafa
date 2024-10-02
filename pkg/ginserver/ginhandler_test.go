package ginhandler

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

	formattedTime := time.Now().Format(time.ANSIC)
	timeJson, err := json.Marshal(formattedTime)
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
			expected:     formattedTime,
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
			testcaseName: "unsupported content type",
			method:       "GET",
			url:          "/datetime",
			statusCode:   http.StatusUnsupportedMediaType,
			expected:     http.StatusText(http.StatusUnsupportedMediaType),
			contentType:  "text/javascript; charset=utf-8",
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
			req.Header.Add("content-type", testcase.contentType)

			if err != nil {
				t.Errorf("Error in new request %v", err)
			}
			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error reading response body %v", err)
			}
			if testcase.contentType == "application/json" {
				assertEquality(t, testcase.expected, resBody)
			} else {
				assertEquality(t, testcase.expected, string(resBody))
			}
			assertEquality(t, testcase.statusCode, res.Code)

		})
	}

}
