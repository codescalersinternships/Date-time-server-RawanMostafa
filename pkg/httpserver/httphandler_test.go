package httphandler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
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
			statusCode:   http.StatusOK,
			expected:     formattedTime,
			contentType:  "text/plain",
		},
		{
			testcaseName: "correct method and url, json type",
			method:       "GET",
			url:          "/datetime",
			statusCode:   http.StatusOK,
			expected:     timeJson,
			contentType:  "application/json",
		},
		{
			testcaseName: "unsupported content type",
			method:       "GET",
			url:          "/datetime",
			statusCode:   http.StatusUnsupportedMediaType,
			expected:     http.StatusText(http.StatusUnsupportedMediaType) + "\n",
			contentType:  "text/javascript; charset=utf-8",
		},
		{
			testcaseName: "wrong method",
			method:       "POST",
			url:          "/datetime",
			statusCode:   http.StatusMethodNotAllowed,
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
