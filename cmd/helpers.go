package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func exitOnErr(msg interface{}) {
	if msg != nil {
		fmt.Println("Error:", msg)
		os.Exit(1)
	}
}

func prettyJSON(obj interface{}) {
	prettyJSON, err := json.MarshalIndent(obj, "", "\t")
	exitOnErr(err)
	fmt.Println(string(prettyJSON))
}
