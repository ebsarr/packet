package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Configure prompts the user to configure a default API key
func Configure() error {
	return nil
}

// GetAPIKey return either the default configured key or the one passed through the cli,
// which has highest priority
func GetAPIKey() (string, error) {

	apiKey := RootCmd.Flag("key").Value.String()
	if apiKey == "" {
		// TODO: Get and assign the configure default key
	}

	if apiKey == "" {
		// API Key was neither configured, neither passed through the cli
		return apiKey, errors.New("API key is missing\nConfigure the API key, or specify default key with the --key flag")
	}

	return apiKey, nil
}

// MarshallAndPrint pretty-prints any object as a JSON string
func MarshallAndPrint(v interface{}) error {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(b))
	return nil
}
