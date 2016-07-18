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

// Configs represents a set of profiles
type Configs struct {
	Profiles map[string]*Config `json:"profiles"`
}

// Config represent default configurations
type Config struct {
	APIKEY           string `json:"APIKEY"`
	DefaultProjectID string `json:"DEFAULT_PROJECT_ID"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%-32s\t%s", c.APIKEY, c.DefaultProjectID)
}

// Configure prompts the user to configure a default API key
func Configure() error {

	// Modify to true if user changes the config
	var hasChanged bool

	// Declare values for user prompt
	var newKey string
	var currentKey string
	var keySuffix string
	var currentProjectID string
	var projectID string

	// Get the profile name from CLI
	profile := getProfile()

	// Get the current key, create a suffix for hint.
	// the current key shall not be displayed on the console
	confs, _ := ReadConfigs()
	var conf *Config
	if confs != nil {
		if _, ok := confs.Profiles[profile]; ok {
			conf = confs.Profiles[profile]
			currentKey = conf.APIKEY
			currentProjectID = conf.DefaultProjectID
		} else {
			conf = &Config{}
		}
	} else {
		confs = &Configs{}
		confs.Profiles = make(map[string]*Config)
		conf = &Config{}
	}

	if currentKey != "" && len(currentKey) > 5 {
		keySuffix = currentKey[len(currentKey)-4:]
	}

	if keySuffix != "" {
		keySuffix = "****************************" + keySuffix
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

	confs.Profiles[profile] = conf

	if hasChanged {
		// Write to config file
		writeConfigFile(confs)
	}

	return nil
}

// ReadConfigs reads the current config file and returns a Config object
func ReadConfigs() (*Configs, error) {
	c, err := readConfigFile()
	if err != nil {
		return nil, err
	}

	// For backward compatibility. This is not essential but will
	// avoid errors after  upgrades from v1.1 to v1.2 and up
	// If it happens that the c.Profiles is empty, we try to read the configs
	// assuming the old format, and automatically make a new defualt profile
	if len(c.Profiles) == 0 {
		defaultConf, err := readOldConfigFile()
		if err != nil {
			return nil, err
		}
		if defaultConf.APIKEY != "" {
			c.Profiles = make(map[string]*Config)
			c.Profiles["default"] = defaultConf
			// Now write the profile in new format
			err := writeConfigFile(c)
			if err != nil {
				return nil, err
			}
		}
	}

	return c, nil
}

// GetAPIKey returns either the default configured key or the one passed through the cli,
// which has highest priority
func GetAPIKey() (string, error) {

	apiKey := RootCmd.Flag("key").Value.String()
	if apiKey == "" {
		profile := getProfile()
		configs, err := ReadConfigs()
		if err != nil {
			return apiKey, err
		}
		if _, found := configs.Profiles[profile]; found {
			apiKey = configs.Profiles[profile].APIKEY
		}
	}

	if apiKey == "" {
		// API Key was neither configured, neither passed through the cli
		return apiKey, errors.New("API key is missing\nConfigure your credentials with `packet configure`, or use the `--key` flag")
	}

	return apiKey, nil
}

// GetProjectID returns the project ID passed to the CLI, otherwise the configure default ID.
func GetProjectID(cmd *cobra.Command) string {
	// The flag to pass a project ID shall always be "--project-id"
	projectID := cmd.Flag("project-id").Value.String()
	if projectID == "" {
		profile := getProfile()
		configs, err := ReadConfigs()
		if err != nil {
			return projectID
		}
		if _, found := configs.Profiles[profile]; found {
			projectID = configs.Profiles[profile].DefaultProjectID
		}
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

// Helper functions

func getProfile() string {
	profile := RootCmd.Flag("profile").Value.String()
	if profile == "" {
		profile = "default"
	}
	return profile
}

func readConfigFile() (*Configs, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(u.HomeDir, ConfigDir, ConfigFile)

	confs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("API key is missing\nConfigure your credentials with `packet configure`, or use the `--key` flag")
	}

	var c Configs
	err = json.Unmarshal(confs, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func readOldConfigFile() (*Config, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(u.HomeDir, ConfigDir, ConfigFile)

	confs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	err = json.Unmarshal(confs, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func writeConfigFile(c *Configs) error {
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

	cBytes, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, cBytes, 0755)
	return err
}
