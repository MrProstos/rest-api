package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

const (
	username string = "VLADISALV"
	password string = "1111"
)

type testServer struct {
	client *http.Client
	server *httptest.Server
}

func NewTestServer() *testServer {
	return &testServer{}
}

func (ts *testServer) Init() *testServer {
	ts.server = httptest.NewServer(NewServer().Init())
	ts.client = ts.server.Client()
	return ts
}

func (ts *testServer) getClient() *http.Client {
	return ts.client
}

func (ts *testServer) getServer() *httptest.Server {
	return ts.server
}

func (ts *testServer) Auth() error {
	req, err := http.NewRequest("POST", ts.server.URL+"/auth/", nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)

	resp, err := ts.client.Do(req)
	if err != nil {
		return err
	}

	ts.client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return err
	}
	req.URL.Path = "/"
	ts.client.Jar.SetCookies(req.URL, resp.Cookies())
	return nil
}

func TestRegistration(t *testing.T) {
	svr := NewTestServer().Init()

	req, err := http.NewRequest("POST", svr.server.URL+"/registration/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(username, password)

	resp, err := svr.client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestAuth(t *testing.T) {
	svr := NewTestServer().Init()

	req, err := http.NewRequest("POST", svr.server.URL+"/auth/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(username, password)

	resp, err := svr.client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp.Cookies())
}

func TestAddClient(t *testing.T) {
	svr := NewTestServer().Init()
	if err := svr.Auth(); err != nil {
		t.Fatal(err)
	}
	users := map[int]map[string]any{
		0: {
			"phone_num": "666",
			"firstname": "vlad",
			"lastname":  "mikhin",
			"birthday":  "1999-07-22",
		},
		1: {
			"phone_num": "777",
			"firstname": "vlad",
			"lastname":  "mikhin",
			"birthday":  "1999-07-22",
			"orders": []map[string]any{
				{
					"phone_num": "777",
					"to":        "sergey",
					"body":      "text",
					"status":    1,
				},
			},
		},
	}
	for _, i := range users {
		b, err := json.Marshal(i)
		if err != nil {
			t.Error(err)
		}
		req, err := http.NewRequest("POST", svr.server.URL+"/client/add/", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		for _, i := range svr.client.Jar.Cookies(req.URL) {
			req.AddCookie(i)
		}

		resp, err := svr.client.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		b, _ = ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	}
}

func TestShowClients(t *testing.T) {
	svr := NewTestServer().Init()
	if err := svr.Auth(); err != nil {
		t.Fatal(err)
	}

	phone_num := []string{"777", "123456789", "666"}
	for _, i := range phone_num {
		req, err := http.NewRequest("GET", svr.server.URL+fmt.Sprintf("/client/%v", i), nil)
		if err != nil {
			t.Fatal(err)
		}
		for _, i := range svr.client.Jar.Cookies(req.URL) {
			req.AddCookie(i)
		}

		resp, err := svr.client.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	}
}

func TestUpdateClient(t *testing.T) {
	svr := NewTestServer().Init()
	if err := svr.Auth(); err != nil {
		t.Fatal(err)
	}

	users := map[int]map[string]any{
		0: {
			"phone_num": "666",
			"firstname": "UPDATE",
			"lastname":  "UPDATE",
			"birthday":  "1999-07-22",
		},
		1: {
			"phone_num": "777",
			"firstname": "UPDATE",
			"lastname":  "UPDATE",
			"birthday":  "1999-07-22",
			"orders": []map[string]any{
				{
					"phone_num": "777",
					"to":        "UPDATE",
					"body":      "UPDATE",
					"status":    1,
				},
			},
		},
	}

	for _, i := range users {
		b, _ := json.Marshal(i)
		req, err := http.NewRequest("PUT", svr.server.URL+"/client/update/", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		for _, i := range svr.client.Jar.Cookies(req.URL) {
			req.AddCookie(i)
		}

		resp, err := svr.client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(resp)
	}
}

func TestDelClient(t *testing.T) {
	svr := NewTestServer().Init()
	if err := svr.Auth(); err != nil {
		t.Fatal(err)
	}

	users := map[int]map[string]any{
		0: {
			"phone_num": "666",
		},
		1: {
			"phone_num": "777",
		},
	}
	for _, i := range users {
		b, _ := json.Marshal(i)
		req, err := http.NewRequest("DELETE", svr.server.URL+"/client/delete/", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		for _, i := range svr.client.Jar.Cookies(req.URL) {
			req.AddCookie(i)
		}

		resp, err := svr.client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(resp)
	}
}

func TestAddOrder(t *testing.T) {
	svr := NewTestServer().Init()
	if err := svr.Auth(); err != nil {
		t.Fatal(err)
	}
	order := map[int]map[string]any{
		0: {
			"phone_num": "777",
			"to":        "sergey",
			"body":      "text",
			"status":    1,
		},
		1: {
			"phone_num": "666",
			"to":        "sergey1",
			"body":      "text1",
			"status":    1,
		},
		2: {
			"phone_num": "777",
			"to":        "sergey2",
			"body":      "text2",
			"status":    1,
		},
	}
	for _, i := range order {
		b, err := json.Marshal(i)
		if err != nil {
			t.Error(err)
		}
		req, err := http.NewRequest("POST", svr.server.URL+"/order/add/", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		for _, i := range svr.client.Jar.Cookies(req.URL) {
			req.AddCookie(i)
		}

		resp, err := svr.client.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		b, _ = ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	}
}

func TestShowOrder(t *testing.T) {
	svr := NewTestServer().Init()
	if err := svr.Auth(); err != nil {
		t.Fatal(err)
	}

	phone_num := []string{"777", "123456789", "666"}
	for _, i := range phone_num {
		req, err := http.NewRequest("GET", svr.server.URL+fmt.Sprintf("/order/%v", i), nil)
		if err != nil {
			t.Fatal(err)
		}
		for _, i := range svr.client.Jar.Cookies(req.URL) {
			req.AddCookie(i)
		}

		resp, err := svr.client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	}
}

func TestUpdateOrder(t *testing.T) {
	svr := NewTestServer().Init()
	if err := svr.Auth(); err != nil {
		t.Fatal(err)
	}

	order := map[int]map[string]any{
		0: {
			"phone_num": "777",
			"to":        "UPDATE",
			"body":      "UPDATE",
			"status":    1,
		},
		1: {
			"phone_num": "666",
			"to":        "UPDATE",
			"body":      "UPDATE",
			"status":    1,
		},
	}

	for _, i := range order {
		b, _ := json.Marshal(i)
		req, err := http.NewRequest("PUT", svr.server.URL+"/order/update/", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		for _, i := range svr.client.Jar.Cookies(req.URL) {
			req.AddCookie(i)
		}

		resp, err := svr.client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(resp)
	}
}

func TestDelOrder(t *testing.T) {
	svr := NewTestServer().Init()
	if err := svr.Auth(); err != nil {
		t.Fatal(err)
	}

	order := map[int]map[string]any{
		0: {
			"phone_num": "777",
		},
		1: {
			"phone_num": "666",
		},
	}

	for _, i := range order {
		b, _ := json.Marshal(i)
		req, err := http.NewRequest("DELETE", svr.server.URL+"/client/delete/", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		for _, i := range svr.client.Jar.Cookies(req.URL) {
			req.AddCookie(i)
		}

		resp, err := svr.client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(resp)
	}
}
