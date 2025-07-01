package common

import "testing"

func TestValidateURLDomain(t *testing.T) {
	tests := []struct {
		name       string
		baseUrl    string
		currentUrl string
		wantErr    bool
	}{
		{
			name:       "OK",
			baseUrl:    "https://example.com",
			currentUrl: "https://example.com",
			wantErr:    false,
		},
		{
			name:       "Not on the same domain",
			baseUrl:    "https://example.com",
			currentUrl: "https://example1.com",
			wantErr:    true,
		},
		{
			name:       "Not on the same TLD",
			baseUrl:    "https://example.com",
			currentUrl: "https://example.org",
			wantErr:    true,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if err := ValidateURLDomain(testCase.baseUrl, testCase.currentUrl); (err != nil) != testCase.wantErr {
				t.Errorf("ValidateURLDomain() error = %v, wantErr %v", err, testCase.wantErr)
			}
		})
	}
}
