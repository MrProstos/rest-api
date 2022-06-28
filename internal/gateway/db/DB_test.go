package db_test

import (
	"fmt"
	"github.com/MrProstos/rest-api/internal/gateway/db"
	"reflect"
	"testing"
)

func Test_Insert(t *testing.T) {
	err := db.NewDataBaseConfig().Connect()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  db.Tables
		wantErr bool
	}{
		{
			name: "Test Insert Client and Order - 1",
			fields: db.NewClient().SetClient(
				"777",
				"VLad",
				"Mikhin",
				"1999-07-22",
				db.SetArrayOrder(db.NewOrder().SetOrder("777", "", "", 0))),
			wantErr: false,
		},
		{
			name:    "Test Insert Client and Order- 2",
			fields:  db.NewClient().SetClient("", "", "", "", nil),
			wantErr: true,
		},
		{
			name:    "Test Insert Order - 3",
			fields:  db.NewOrder().SetOrder("777", "Test", "Test", 0),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insert, err := db.Tables.Insert(tt.fields)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}

			_select, err := db.Tables.Select(tt.fields)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}

			if !reflect.DeepEqual(insert, _select) {
				t.Error("the structures are not equal")
			}
		})
	}
}

func Test_Select(t *testing.T) {
	err := db.NewDataBaseConfig().Connect()
	if err != nil {
		t.Fatal(err)
	}

	test := []struct {
		name   string
		fields db.Tables
	}{
		{
			name:   "Test Select Client and Order - 1",
			fields: db.NewClient().SetClient("777", "", "", "", nil),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			_select, err := db.Tables.Select(tt.fields)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(tt.fields, _select) {
				t.Error("the structures are not equal")
			}
			fmt.Println(_select)
		})
	}

}

func Test_Update(t *testing.T) {
	err := db.NewDataBaseConfig().Connect()
	if err != nil {
		t.Fatal(err)
	}

	test := []struct {
		name    string
		fields  db.Tables
		wantErr bool
	}{
		{
			name: "Test Update Client and Order - 1",
			fields: db.NewClient().SetClient(
				"777",
				"UPDATE",
				"UPDATE",
				"",
				db.SetArrayOrder(db.NewOrder().SetOrder("777", "UPDATE", "UPDATE", 0))),
			wantErr: false,
		},
		{
			name: "Test Update Client and Order - 2 ",
			fields: db.NewClient().SetClient(
				"777",
				"UPDATE2",
				"UPDATE2",
				"",
				db.SetArrayOrder(db.NewOrder().SetOrder("777", "UPDATE2", "UPDATE2", 0))),
			wantErr: false,
		},
		{
			name: "Test Update Client and Order - 3 ",
			fields: db.NewClient().SetClient(
				"",
				"",
				"",
				"",
				nil),
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			update, err := db.Tables.Update(tt.fields)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}

			_select, err := db.Tables.Select(tt.fields)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}

			if !reflect.DeepEqual(update, _select) {
				t.Error("the structures are not equal")
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	err := db.NewDataBaseConfig().Connect()
	if err != nil {
		t.Fatal(err)
	}

	test := []struct {
		name   string
		fields db.Tables
	}{
		{
			name:   "Test Select Client and Order - 1",
			fields: db.NewClient().SetClient("777", "", "", "", nil),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			_, err := db.Tables.Delete(tt.fields)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
