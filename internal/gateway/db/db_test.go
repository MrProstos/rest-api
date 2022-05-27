package db

import (
	"fmt"
	"testing"
)

func TestClient_Add(t *testing.T) {
	client := []Client{{
		Phone_num: "7777",
		Firstname: "Vlad",
		Lastname:  "Mikhin",
		Birthday:  "22-07-1999",
		Orders:    nil,
	}, {
		Phone_num: "6666",
		Firstname: "test",
		Lastname:  "test",
		Birthday:  "test",
		Orders:    nil,
	}}
	for i := range client {
		t.Run(fmt.Sprintf("Test-%v", i+1), func(t *testing.T) {
			if err := client[i].Add(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestClient_Update(t *testing.T) {
	client := []Client{{
		Phone_num: "7777",
		Firstname: "update",
		Lastname:  "update",
		Birthday:  "update",
		Orders:    nil,
	}, {
		Phone_num: "6666",
		Firstname: "update",
		Lastname:  "update",
		Birthday:  "update",
		Orders:    nil,
	}}
	for i := range client {
		t.Run(fmt.Sprintf("Test-%v", i+1), func(t *testing.T) {
			if err := client[i].Update(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestClient_Show(t *testing.T) {
	client := []Client{{
		Phone_num: "7777",
	}, {
		Phone_num: "6666",
	}}
	for i := range client {
		t.Run(fmt.Sprintf("Test-%v", i+1), func(t *testing.T) {
			if err := client[i].Show(); err != nil {
				t.Error(err)
			}
			fmt.Println(client[i])
		})
	}
}

func TestClient_Del(t *testing.T) {
	client := []Client{{
		ID: 1,
	}, {
		ID: 2,
	}}
	for i := range client {
		t.Run(fmt.Sprintf("Test-%v", i+1), func(t *testing.T) {
			if err := client[i].Del(); err != nil {
				t.Error(err)
			}
		})
	}
}
