package db

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDataBaseConfig(t *testing.T) {
	tests := []struct {
		name string
		want *dataBaseConfig
	}{
		{
			name: "Test - 1",
			want: &dataBaseConfig{
				dbName: "",
				dbPass: "",
				dbUser: "",
				dbHost: "",
			},
		}, {
			name: "Test - 2",
			want: &dataBaseConfig{
				dbName: "",
				dbPass: "",
				dbUser: "",
				dbHost: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataBaseConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDataBaseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrder(t *testing.T) {
	tests := []struct {
		name string
		want *order
	}{
		{
			name: "Test - 1",
			want: &order{
				ID:       0,
				PhoneNum: "",
				To:       "",
				Body:     "",
				Status:   0,
			},
		}, {
			name: "Test - 2",
			want: &order{
				ID:       0,
				PhoneNum: "",
				To:       "",
				Body:     "",
				Status:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataBaseConfig_Connect(t *testing.T) {
	type fields struct {
		dbName string
		dbPass string
		dbUser string
		dbHost string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test - 1",
			fields: fields{
				dbName: "postgres",
				dbPass: "Zz123456",
				dbUser: "postgres",
				dbHost: "localhost",
			},
			wantErr: false,
		}, {
			name: "Test - 2",
			fields: fields{
				dbName: "postgres",
				dbPass: "",
				dbUser: "postgres",
				dbHost: "localhost",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := &dataBaseConfig{
				dbName: tt.fields.dbName,
				dbPass: tt.fields.dbPass,
				dbUser: tt.fields.dbUser,
				dbHost: tt.fields.dbHost,
			}
			if err := database.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dataBaseConfig_SetConnect(t *testing.T) {
	type args struct {
		dbName string
		dbPass string
		dbUser string
		dbHost string
	}
	tests := []struct {
		name string
		args args
		want *dataBaseConfig
	}{
		{
			name: "Test - 1",
			args: args{
				dbName: "postgres",
				dbPass: "",
				dbUser: "postgres",
				dbHost: "localhost",
			},
			want: &dataBaseConfig{
				dbName: "postgres",
				dbPass: "",
				dbUser: "postgres",
				dbHost: "localhost",
			},
		}, {
			name: "Test - 2",
			args: args{
				dbName: "test",
				dbPass: "test",
				dbUser: "test",
				dbHost: "test",
			},
			want: &dataBaseConfig{
				dbName: "test",
				dbPass: "test",
				dbUser: "test",
				dbHost: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := &dataBaseConfig{}
			if got := database.SetConnect(tt.args.dbName, tt.args.dbPass, tt.args.dbUser, tt.args.dbHost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConnect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_isValid(t *testing.T) {
	type fields struct {
		PhoneNum  string
		Firstname string
		Lastname  string
		Birthday  string
		Orders    []order
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test - 1",
			fields: fields{
				PhoneNum:  "",
				Firstname: "",
				Lastname:  "",
				Birthday:  "",
				Orders:    nil,
			},
			wantErr: true,
		},
		{
			name: "Test - 2",
			fields: fields{
				PhoneNum:  "89537648822",
				Firstname: "Vlad",
				Lastname:  "Mikhin",
				Birthday:  "1999-07-22",
				Orders:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &client{
				PhoneNum:  tt.fields.PhoneNum,
				Firstname: tt.fields.Firstname,
				Lastname:  tt.fields.Lastname,
				Birthday:  tt.fields.Birthday,
				Orders:    tt.fields.Orders,
			}
			if err := client.isValid(); (err != nil) != tt.wantErr {
				t.Errorf("isValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Tables_Insert(t *testing.T) {
	conf := NewDataBaseConfig().SetConnect("postgres", "Zz123456", "postgres", "localhost").Connect()
	if conf != nil {
		t.Fatal(conf)
	}

	tests := []struct {
		name    string
		fileds  Tables
		wantErr bool
	}{
		{
			name: "Test - 1",
			fileds: &client{
				PhoneNum:  "89537648822",
				Firstname: "Vlad",
				Lastname:  "Mikhin",
				Birthday:  "1999-07-22",
				Orders: []order{{
					PhoneNum: "89537648822",
					To:       "test",
					Body:     "test",
					Status:   0,
				}},
			},
			wantErr: false,
		}, {
			name: "Test - 2",
			fileds: &client{
				PhoneNum:  "",
				Firstname: "test",
				Lastname:  "test",
				Birthday:  "",
				Orders:    nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Tables.Insert(tt.fileds)
			if err != nil && !tt.wantErr {
				t.Errorf("Tables.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Tables_Select(t *testing.T) {
	conf := NewDataBaseConfig().SetConnect("postgres", "Zz123456", "postgres", "localhost").Connect()
	if conf != nil {
		t.Fatal(conf)
	}

	tests := []struct {
		name    string
		fileds  Tables
		wantErr bool
	}{
		{
			name: "Test - 1",
			fileds: &client{
				PhoneNum: "89537648822",
			},
			wantErr: false,
		}, {
			name: "Test - 2",
			fileds: &order{
				PhoneNum: "89537648822",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := Tables.Select(tt.fileds)
			if err != nil {
				t.Errorf("Tables.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(msg)
		})
	}
}

func Test_Tables_Update(t *testing.T) {
	conf := NewDataBaseConfig().SetConnect("postgres", "Zz123456", "postgres", "localhost").Connect()
	if conf != nil {
		t.Fatal(conf)
	}

	tests := []struct {
		name    string
		fileds  Tables
		wantErr bool
	}{
		{
			name: "Test - 1",
			fileds: &client{
				PhoneNum:  "89537648822",
				Firstname: "UPDATE1",
				Lastname:  "UPDATE",
				Birthday:  "2022-11-2",
				Orders:    nil,
			},
			wantErr: false,
		}, {
			name: "Test - 2",
			fileds: &order{
				PhoneNum: "89537648822",
				To:       "UPDATE",
				Body:     "UPDATE",
				Status:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := Tables.Update(tt.fileds)
			if err != nil {
				t.Errorf("Tables.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(msg)
		})
	}
}

func Test_Tables_Delete(t *testing.T) {
	conf := NewDataBaseConfig().SetConnect("postgres", "Zz123456", "postgres", "localhost").Connect()
	if conf != nil {
		t.Fatal(conf)
	}

	tests := []struct {
		name    string
		fileds  Tables
		wantErr bool
	}{
		{
			name: "Test - 1",
			fileds: &client{
				PhoneNum: "89537648822",
			},
			wantErr: false,
		}, {
			name: "Test - 2",
			fileds: &order{
				PhoneNum: "89537648822",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := Tables.Delete(tt.fileds)
			if err != nil {
				t.Errorf("Tables.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(msg)
		})
	}
}
