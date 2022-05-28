package db

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	date, err := time.Parse("2006-01-02", "1999-07-22")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(date.Format("2006-01-02"))
}

func TestClient_Add(t *testing.T) {
	client := []Client{{
		PhoneNum:  "7777",
		Firstname: "Vlad",
		Lastname:  "Mikhin",
		Birthday:  "1999-07-22",
		Orders: []Order{
			{
				PhoneNum: "7777",
				Status:   0,
			}, {
				PhoneNum: "7777",
				Status:   0,
			},
		},
	}, {
		PhoneNum:  "6666",
		Firstname: "test",
		Lastname:  "test",
		Birthday:  "1999-07-22",
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
		PhoneNum:  "7777",
		Firstname: "update",
		Lastname:  "update",
		Birthday:  "update",
		Orders:    nil,
	}, {
		PhoneNum:  "6666",
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
		PhoneNum: "7777",
	}, {
		PhoneNum: "6666",
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
		PhoneNum: "7777",
	}, {
		PhoneNum: "6666",
	}}
	for i := range client {
		t.Run(fmt.Sprintf("Test-%v", i+1), func(t *testing.T) {
			if err := client[i].Del(); err != nil {
				t.Error(err)
			}
		})
	}
}
