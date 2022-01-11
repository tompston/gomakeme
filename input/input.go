package input

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// credits --> https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp
// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Project, error) {
	// Create config structure
	config := &Project{}
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Init new YAML decode
	d := yaml.NewDecoder(file)
	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	// if the ProjectConfig struc in the yml file is empty / not used,
	// assign them to be some predefined values
	if (ProjectConfig{} == config.ProjectConfig) {
		config.ProjectConfig = placeholder_config
	}
	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string
	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	// Actually parse the flags
	flag.Parse()
	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}
	// Return the configuration path
	return configPath, nil
}
