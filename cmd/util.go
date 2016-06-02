package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// ConfigDir is the location of the config file under user's $HOME dir
const ConfigDir = ".packet"

// ConfigFile is the filename of the config file
const ConfigFile = "config"

// Config represent default configurations
type Config struct {
	APIKEY string `json:"APIKEY"`
}

// Configure prompts the user to configure a default API key
func Configure() error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	dirPath := filepath.Join(u.HomeDir, ConfigDir)
	filePath := filepath.Join(dirPath, ConfigFile)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, 0755)
	}

	// Declare values for user prompt
	var newKey string
	var currentKey string
	var keySuffix string

	// Get the current key
	conf, _ := ReadKey()
	if conf != nil {
		currentKey = conf.APIKEY
	}

	if currentKey != "" && len(currentKey) > 5 {
		keySuffix = currentKey[len(currentKey)-4:]
	}

	if keySuffix != "" {
		keySuffix = "*****" + keySuffix
	}

	// Get API from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your API key [ %s ]: ", keySuffix)
	newKey, _ = reader.ReadString('\n')
	newKey = strings.TrimSpace(newKey)

	// For debug:
	// fmt.Printf("New key is: %s\n", newKey)

	if newKey == "" {
		// No change.
		return nil
	}

	var newConf = &Config{
		APIKEY: newKey,
	}

	// Write to config file
	c, err := json.Marshal(newConf)
	e := ioutil.WriteFile(filePath, c, 0755)

	return e
}

// ReadKey reads the APIKEY from the config file
func ReadKey() (*Config, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(u.HomeDir, ConfigDir, ConfigFile)

	conf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	e := json.Unmarshal(conf, &c)
	if e != nil {
		return nil, e
	}

	return &c, nil
}

// GetAPIKey return either the default configured key or the one passed through the cli,
// which has highest priority
func GetAPIKey() (string, error) {

	apiKey := RootCmd.Flag("key").Value.String()
	if apiKey == "" {
		config, err := ReadKey()
		if err != nil {
			apiKey = ""
		} else {
			apiKey = config.APIKEY
		}
	}

	if apiKey == "" {
		// API Key was neither configured, neither passed through the cli
		return apiKey, errors.New("API key is missing\nConfigure a default one with \"packet configure\", or specify with the --key flag")
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
