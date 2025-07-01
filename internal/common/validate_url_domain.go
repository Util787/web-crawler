package common

import (
	"errors"
	"net/url"
)

// This function makes sure the currentUrl is on the same domain as the baseUrl
func ValidateURLDomain(baseUrl string, currentUrl string) error {
	parsedCurrentURL, err := url.Parse(currentUrl)
	if err != nil {
		return err
	}

	parsedBaseURL, err := url.Parse(baseUrl)
	if err != nil {
		return err
	}

	if parsedCurrentURL.Host != parsedBaseURL.Host {
		return errors.New("url is not on the provided domain")
	}
	return nil
}
