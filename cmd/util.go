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

	"github.com/spf13/cobra"
)

// ConfigDir is the location of the config file under user's $HOME dir
const ConfigDir = ".packet"

// ConfigFile is the filename of the config file
const ConfigFile = "config"

// Config represent default configurations
type Config struct {
	APIKEY           string `json:"APIKEY"`
	DefaultProjectID string `json:"DEFAULT_PROJECT_ID"`
}

// Configure prompts the user to configure a default API key
func Configure() error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	// create directory to save configs
	dirPath := filepath.Join(u.HomeDir, ConfigDir)
	filePath := filepath.Join(dirPath, ConfigFile)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, 0755)
	}

	// Modify to true if user changes the config
	var hasChanged bool

	// Declare values for user prompt
	var newKey string
	var currentKey string
	var keySuffix string
	var currentProjectID string
	var projectID string

	// Get the current key, create a suffix for hint.
	// the current key shall not be displayed on the console
	conf, _ := ReadConfigs()
	if conf != nil {
		currentKey = conf.APIKEY
		currentProjectID = conf.DefaultProjectID
	} else {
		conf = &Config{}
	}

	if currentKey != "" && len(currentKey) > 5 {
		keySuffix = currentKey[len(currentKey)-4:]
	}

	if keySuffix != "" {
		keySuffix = "*****" + keySuffix
	}

	// Get API key and default project ID from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your API key [ %s ]: ", keySuffix)
	newKey, _ = reader.ReadString('\n')
	newKey = strings.TrimSpace(newKey)
	fmt.Printf("Enter your default project ID [ %s ]: ", currentProjectID)
	projectID, _ = reader.ReadString('\n')
	projectID = strings.TrimSpace(projectID)

	// For debug:
	// fmt.Printf("New key is: %s\n", newKey)
	// fmt.Printf("New PID: %s\n", projectID)

	if newKey != "" {
		conf.APIKEY = newKey
		hasChanged = true
	}

	if projectID != "" {
		conf.DefaultProjectID = projectID
		hasChanged = true
	}

	if hasChanged {
		// Write to config file
		c, err := json.MarshalIndent(conf, "", "\t")
		if err != nil {
			return err
		}
		e := ioutil.WriteFile(filePath, c, 0755)
		return e
	}

	return nil
}

// ReadConfigs reads the current config file and returns a Config object
func ReadConfigs() (*Config, error) {
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

// GetAPIKey returns either the default configured key or the one passed through the cli,
// which has highest priority
func GetAPIKey() (string, error) {

	apiKey := RootCmd.Flag("key").Value.String()
	if apiKey == "" {
		config, _ := ReadConfigs()
		apiKey = config.APIKEY
	}

	if apiKey == "" {
		// API Key was neither configured, neither passed through the cli
		return apiKey, errors.New("API key is missing\nConfigure a default one with \"packet configure\", or specify with the --key flag")
	}

	return apiKey, nil
}

// GetProjectID returns the project ID passed to the CLI, otherwise the configure default ID.
func GetProjectID(cmd *cobra.Command) string {
	// The flag to pass a project ID shall always be "--project-id"
	projectID := cmd.Flag("project-id").Value.String()
	if projectID == "" {
		config, _ := ReadConfigs()
		projectID = config.DefaultProjectID
	}

	return projectID

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
