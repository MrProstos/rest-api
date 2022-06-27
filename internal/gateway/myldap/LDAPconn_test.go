package myldap

import (
	"reflect"
	"testing"
)

func Test_conf_Connect(t *testing.T) {
	type fields struct {
		url      string
		bind     string
		password string
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
			conf := &conf{
				url:      tt.fields.url,
				bind:     tt.fields.bind,
				password: tt.fields.password,
			}
			if err := conf.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newConf(t *testing.T) {
	type args struct {
		url      string
		bind     string
		password string
	}
	tests := []struct {
		name string
		args args
		want *conf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConf(tt.args.url, tt.args.bind, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
