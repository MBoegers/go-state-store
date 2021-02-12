package configuration

/*
 * Copied from https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp
 */
import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Edit struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"edit"`

	Read struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"read"`

	Event struct {
		Time time.Duration `yaml:"time"`
	} `yaml:"event"`
}

// NewConfig returns a new decoded Config struct
func ReadConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

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

	return config, nil
}
