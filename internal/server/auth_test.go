package server

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRegistration(t *testing.T) {
	r := httptest.NewRequest("GET", "http://localhost:2000/registration/", nil)
	r.SetBasicAuth("test1", "test1")
	w := httptest.NewRecorder()
	Registration(w, r)
	if w.Result().StatusCode != 200 {
		t.Errorf("%v", w)
	}
	fmt.Println(w)
}

func TestAuth(t *testing.T) {
	r := httptest.NewRequest("GET", "http://localhost:2000/auth/", nil)
	r.SetBasicAuth("test1", "test1")
	w := httptest.NewRecorder()
	Auth(w, r)
	if w.Result().StatusCode != 200 {
		t.Errorf("%v", w)
	}
	fmt.Println(w.Result())
}
