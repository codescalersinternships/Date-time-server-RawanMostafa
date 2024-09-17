package internal

import (
	"io"
	"log"
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
	req := httptest.NewRequest("GET","http://localhost:8080", nil)

	w := httptest.NewRecorder()
	HttpHome(w, req)
	resp := w.Result()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body %v", err)
	}

	expected := "Welcome to my datetime server!"

	assertEquality(t, expected, string(resBody))
	assertEquality(t, 200, resp.StatusCode)
}

func TestHttpHandler(t *testing.T) {

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
			url:          "http://localhost:8080/datetime",
			statusCode:   200,
			expected:     truncateToSec(time.Now()).String(),
		},
		{
			testcaseName: "wrong method",
			method:       "POST",
			url:          "http://localhost:8080/datetime",
			statusCode:   405,
			expected:     http.StatusText(http.StatusMethodNotAllowed) + "\n",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.testcaseName, func(t *testing.T) {

			req := httptest.NewRequest(testcase.method, testcase.url, nil)

			w := httptest.NewRecorder()
			HttpHandler(w, req)
			resp := w.Result()
			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Error reading response body %v", err)
			}

			assertEquality(t, testcase.expected, string(resBody))
			assertEquality(t, testcase.statusCode, resp.StatusCode)

		})
	}

}

