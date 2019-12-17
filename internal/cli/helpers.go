package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

func prettyJSON(obj interface{}) error {
	prettyJSON, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(prettyJSON))

	return nil
}

func validURL(val interface{}) error {
	_, err := url.ParseRequestURI(val.(string))
	if err != nil {
		return errors.New("the value is not a valid URL")
	}
	return nil
}
