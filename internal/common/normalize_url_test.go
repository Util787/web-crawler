package common

import "testing"

func TestNormalizeUrl(t *testing.T) {
	tests := []struct {
		name    string
		rawUrl  string
		want    string
		wantErr bool
	}{
		{
			name:    "Https with slash",
			rawUrl:  "https://someurl.com/path/",
			want:    "someurl.com/path",
			wantErr: false,
		},
		{
			name:    "Https with subpath",
			rawUrl:  "https://someurl.org/path",
			want:    "someurl.org/path",
			wantErr: false,
		},
		{
			name:    "Http with slash",
			rawUrl:  "http://someurl.com/path/",
			want:    "someurl.com/path",
			wantErr: false,
		},
		{
			name:    "Http with subpath",
			rawUrl:  "http://someurl.org/path",
			want:    "someurl.org/path",
			wantErr: false,
		},
		{
			name:    "With query",
			rawUrl:  "https://someurl.com/path/to/resource?query=123#section",
			want:    "someurl.com/path/to/resource",
			wantErr: false,
		},
		{
			name:    "Not url",
			rawUrl:  "not a url",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Empty string",
			rawUrl:  "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Whitespace only",
			rawUrl:  "   ",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Missing host",
			rawUrl:  "https:///path",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid host",
			rawUrl:  "https://someurl/path",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Missing path",
			rawUrl:  "https://someurl.com",
			want:    "someurl.com",
			wantErr: false,
		},
		{
			name:    "Unsupported scheme",
			rawUrl:  "tttt://someurl.com/path",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Missing scheme",
			rawUrl:  "someurl.com/path",
			want:    "",
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NormalizeUrl(testCase.rawUrl)
			if (err != nil) != testCase.wantErr {
				t.Errorf("NormalizeUrl() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if got != testCase.want {
				t.Errorf("NormalizeUrl() = %v, want %v", got, testCase.want)
			}
		})
	}
}
