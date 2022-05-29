package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var token *http.Cookie

func getToken(t *testing.T) {
	client := http.Client{}
	auth, err := http.NewRequest("GET", "http://localhost:2000/auth/", nil)
	if err != nil {
		t.Fatal(err)
	}
	auth.SetBasicAuth("test", "test")
	authResponse, _ := client.Do(auth)
	if authResponse.StatusCode != 200 {
		t.Fatal(authResponse)
	}

	token = authResponse.Cookies()[0]

	err = authResponse.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

}

func TestShowClients(t *testing.T) {
	getToken(t)
	client := http.Client{}
	showclient, err := http.NewRequest("GET", "http://localhost:2000/clietns/7777", nil)
	if err != nil {
		t.Fatal(err)
	}
	showclient.AddCookie(token)
	showclientResponse, _ := client.Do(showclient)
	if showclientResponse.StatusCode != 200 {
		t.Fatal(showclientResponse)
	}
	b, _ := ioutil.ReadAll(showclientResponse.Body)
	fmt.Println(string(b))
}

func TestAddUpdateDelClient(t *testing.T) {
	getToken(t)
	httpClient := http.Client{}
	t.Run("Test AddClient", func(t *testing.T) {
		body := map[string]interface{}{
			"phone_num": "98765431",
			"firstname": "TestAddClient",
			"lastname":  "TestAddClient",
			"birthday":  "1999-07-22",
			"orders": []map[string]interface{}{
				{"phone_num": "98765431",
					"to":     "vlad",
					"body":   "vlad",
					"status": 1},
			},
		}
		jsonStr, _ := json.Marshal(body)

		addclientRequest, _ := http.NewRequest("POST", "http://localhost:2000/client/add/", bytes.NewReader(jsonStr))
		addclientRequest.AddCookie(token)

		addclientResponse, _ := httpClient.Do(addclientRequest)
		if addclientResponse.StatusCode != 200 {
			t.Fatal(addclientResponse)
		}

		b, _ := ioutil.ReadAll(addclientResponse.Body)
		fmt.Println(string(b))
		addclientResponse.Body.Close()
	})

	t.Run("Test UpdateClient", func(t *testing.T) {
		body := map[string]interface{}{
			"phone_num": "98765431",
			"firstname": "TestUpdateClient",
			"lastname":  "TestUpdateClient",
			"birthday":  "1999-09-09",
		}
		jsonStr, _ := json.Marshal(body)

		updateclientRequest, _ := http.NewRequest("PUT", "http://localhost:2000/client/update/", bytes.NewReader(jsonStr))
		updateclientRequest.AddCookie(token)

		updateclientResponse, _ := httpClient.Do(updateclientRequest)
		if updateclientResponse.StatusCode != 200 {
			t.Fatal(updateclientResponse)
		}

		b, _ := ioutil.ReadAll(updateclientResponse.Body)
		fmt.Println(string(b))
		updateclientResponse.Body.Close()
	})

	t.Run("Test DelClient", func(t *testing.T) {
		body := map[string]interface{}{
			"phone_num": "98765431",
		}
		jsonStr, _ := json.Marshal(body)

		updateclientRequest, _ := http.NewRequest("DELETE", "http://localhost:2000/client/delete/", bytes.NewReader(jsonStr))
		updateclientRequest.AddCookie(token)

		updateclientResponse, _ := httpClient.Do(updateclientRequest)
		if updateclientResponse.StatusCode != 200 {
			t.Fatal(updateclientResponse)
		}

		b, _ := ioutil.ReadAll(updateclientResponse.Body)
		fmt.Println(string(b))

		updateclientResponse.Body.Close()
	})

}
