package common

import (
	"reflect"
	"testing"
)

func Test_getURLsFromHTML(t *testing.T) {
	tests := []struct {
		name       string
		htmlBody   string
		rawBaseURL string
		want       []string
		wantErr    bool
	}{
		{
			name:       "OK absolute url",
			htmlBody:   "<a href='https://example.com'>Example</a><a href='https://example.com/foo'>Foo</a>",
			rawBaseURL: "https://example.com",
			want:       []string{"https://example.com", "https://example.com/foo"},
			wantErr:    false,
		},
		{
			name:       "OK relative url",
			htmlBody:   `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`,
			rawBaseURL: "https://example.com",
			want:       []string{"https://example.com/foo", "https://example.com/bar/baz"},
			wantErr:    false,
		},
		{
			name:       "absolute and relative URLs",
			rawBaseURL: "https://example.com",
			htmlBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Example</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Example</span>
				</a>
			</body>
		</html>
		`,
			want:    []string{"https://example.com/path/one", "https://other.com/path/one"},
			wantErr: false,
		},
		{
			name:       "OK with multiple equal links",
			rawBaseURL: "https://example.com",
			htmlBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Example</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Example</span>
				</a>
                <a href="https://other.com/path/one">
					<span>Example</span>
				</a>
                <a href="https://other.com/path/one">
					<span>Example</span>
				</a>
				<a href="/path/one">
					<span>Example</span>
				</a>
			</body>
		</html>
		`,
			want:    []string{"https://example.com/path/one", "https://other.com/path/one", "https://other.com/path/one", "https://other.com/path/one", "https://example.com/path/one"},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := GetURLsFromHTML(testCase.htmlBody, testCase.rawBaseURL)
			if (err != nil) != testCase.wantErr {
				t.Errorf("getURLsFromHTML() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("getURLsFromHTML() = %v, want %v", got, testCase.want)
			}
		})
	}
}
