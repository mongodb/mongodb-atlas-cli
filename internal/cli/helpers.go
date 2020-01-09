package cli

import (
	"errors"
	"net/url"
)

func validURL(val interface{}) error {
	_, err := url.ParseRequestURI(val.(string))
	if err != nil {
		return errors.New("the value is not a valid URL")
	}
	return nil
}
