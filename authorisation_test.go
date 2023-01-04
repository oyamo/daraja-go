package darajago

import (
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
		{
			name: "wrong secrets",
			args: struct {
				consumerKey    string
				consumerSecret string
			}{consumerKey: "E22yMhsfYOZgAFQJY5N8aRkt4gQbvETC", consumerSecret: "zAFGe5cWKv3U1HQ7"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuthorization(tt.args.consumerKey, tt.args.consumerSecret)
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (err == nil) && len(got.AccessToken) == 0 {
				t.Errorf("NewAuthorization() got = %v, want %v", got, tt.want)
			}
		})
	}
}
