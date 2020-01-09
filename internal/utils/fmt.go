package utils

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

func PrettyJSON(obj interface{}) error {
	prettyJSON, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(prettyJSON))

	return nil
}

// TODO: use me ;)
func PrettyYAML(obj interface{}) error {
	prettyYAML, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	fmt.Println(string(prettyYAML))

	return nil
}
