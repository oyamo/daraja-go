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
			}{consumerKey: "bidibidi", consumerSecret: "babababa"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct secrets",
			args: struct {
				consumerKey    string
				consumerSecret string
			}{consumerKey: "E22yMhsfYOZgAFQJY5N8aRkt4gQbvETC", consumerSecret: "zAFGe5cWKv3U1HQ7"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newAuthorization(tt.args.consumerKey, tt.args.consumerSecret, ENVIRONMENT_SANDBOX)
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("newAuthorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (err == nil) && len(got.AccessToken) == 0 {
				t.Errorf("newAuthorization() got = %v, want %v", got, tt.want)
			}
		})
	}
}
