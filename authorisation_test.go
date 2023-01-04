package darajago

import (
	"reflect"
	"testing"
)

func TestNewAuthorization(t *testing.T) {
	type args struct {
		consumerKey    string
		consumerSecret string
	}
	tests := []struct {
		name    string
		args    args
		want    *Authorization
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuthorization(tt.args.consumerKey, tt.args.consumerSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthorization() got = %v, want %v", got, tt.want)
			}
		})
	}
}