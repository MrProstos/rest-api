package myldap

import (
	"reflect"
	"testing"
)

func TestNewOperator(t *testing.T) {

	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *operator
	}{
		{
			name: "",
			args: args{
				username: "vlad",
				password: "vlad",
			},
			want: &operator{
				Username: "vlad",
				Password: "vlad",
			},
		}, {
			name: "",
			args: args{
				username: "asdw1",
				password: "asdw1",
			},
			want: &operator{
				Username: "asdw1",
				Password: "asdw1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOperator(tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operator_AddUser(t *testing.T) {
	conf := newConf(Url, Bind, Password)
	err := conf.Connect()
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				Username: "Haruk",
				Password: "Haruk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operator := &operator{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			if err := operator.AddUser(); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_operator_Search(t *testing.T) {
	conf := newConf(Url, Bind, Password)
	err := conf.Connect()
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test - 1",
			fields: fields{
				Username: "Haruk",
				Password: "Haruk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operator := &operator{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			if err := operator.Search(); (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
