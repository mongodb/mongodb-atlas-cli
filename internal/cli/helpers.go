package cli

import (
	"encoding/json"
	"fmt"
)

func prettyJSON(obj interface{}) error {
	prettyJSON, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(prettyJSON))

	return nil
}
