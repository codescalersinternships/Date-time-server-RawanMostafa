package internal

import (
	"io"
	"log"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func assertEquality(t *testing.T, obj1 any, obj2 any) {
	t.Helper()
	if reflect.TypeOf(obj1) != reflect.TypeOf(obj2) {
		t.Errorf("Error! type mismatch, wanted %t got %t", reflect.TypeOf(obj1), reflect.TypeOf(obj2))
	}
	if !reflect.DeepEqual(obj1, obj2) {
		t.Errorf("Error! values mismatch, wanted %v got %v", obj1, obj2)
	}
}

func TestHome(t *testing.T) {

}

func TestGetDate(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:8080", nil)

	w := httptest.NewRecorder()
	GetDate(w, req)
	resp := w.Result()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body %v", err)
	}

	currentTime := time.Now()
	want := truncateToSec(currentTime).String()

	assertEquality(t, want, string(resBody))

}
