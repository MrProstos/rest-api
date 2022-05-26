package core

import (
	"testing"
)

func TestClient_IsValid(t *testing.T) {
	type fields struct {
		Phone_num string
		Firstname string
		Lastname  string
		Birthday  string
		Orders    []Order
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test 1 = False",
			fields: fields{
				Phone_num: "",
				Firstname: "",
				Lastname:  "",
				Birthday:  "",
				Orders:    nil,
			},
			wantErr: true,
		},
		{
			name: "Test 2 = True",
			fields: fields{
				Phone_num: "999",
				Firstname: "999",
				Lastname:  "999",
				Birthday:  "999",
				Orders:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				Phone_num: tt.fields.Phone_num,
				Firstname: tt.fields.Firstname,
				Lastname:  tt.fields.Lastname,
				Birthday:  tt.fields.Birthday,
				Orders:    tt.fields.Orders,
			}
			if err := client.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Add(t *testing.T) {
	type fields struct {
		Phone_num string
		Firstname string
		Lastname  string
		Birthday  string
		Orders    []Order
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test 1 = True",
			fields: fields{
				Phone_num: "123",
				Firstname: "Vlad",
				Lastname:  "Mikhin",
				Birthday:  "22-07-1999",
				Orders:    nil,
			},
			wantErr: false,
		}, {
			name: "Test 2 = True",
			fields: fields{
				Phone_num: "89531222390",
				Firstname: "Test",
				Lastname:  "Test",
				Birthday:  "777",
				Orders: []Order{{
					Client_id: 2,
					Title:     "Test",
					To:        "Test",
					Body:      "Test",
					Status:    1,
				}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				Phone_num: tt.fields.Phone_num,
				Firstname: tt.fields.Firstname,
				Lastname:  tt.fields.Lastname,
				Birthday:  tt.fields.Birthday,
				Orders:    tt.fields.Orders,
			}
			if err := client.Add(); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Update(t *testing.T) {
	type fields struct {
		ID        uint
		Phone_num string
		Firstname string
		Lastname  string
		Birthday  string
		Orders    []Order
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test 1 = True",
			fields: fields{
				Phone_num: "123",
				Firstname: "update",
				Lastname:  "update",
				Birthday:  "update",
				Orders:    nil,
			},
			wantErr: false,
		}, {
			name: "Test 2 = True", fields: fields{
				Phone_num: "89531222390",
				Firstname: "update2",
				Lastname:  "update2",
				Birthday:  "update2",
				Orders:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				Phone_num: tt.fields.Phone_num,
				Firstname: tt.fields.Firstname,
				Lastname:  tt.fields.Lastname,
				Birthday:  tt.fields.Birthday,
				Orders:    tt.fields.Orders,
			}
			if err := client.Update(); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Show(t *testing.T) {
	type fields struct {
		ID        uint
		Phone_num string
		Firstname string
		Lastname  string
		Birthday  string
		Orders    []Order
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				Phone_num: tt.fields.Phone_num,
				Firstname: tt.fields.Firstname,
				Lastname:  tt.fields.Lastname,
				Birthday:  tt.fields.Birthday,
				Orders:    tt.fields.Orders,
			}
			if err := client.Show(); (err != nil) != tt.wantErr {
				t.Errorf("Show() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Del(t *testing.T) {
	type fields struct {
		Phone_num string
		Firstname string
		Lastname  string
		Birthday  string
		Orders    []Order
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				Phone_num: tt.fields.Phone_num,
				Firstname: tt.fields.Firstname,
				Lastname:  tt.fields.Lastname,
				Birthday:  tt.fields.Birthday,
				Orders:    tt.fields.Orders,
			}
			if err := client.Del(); (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
