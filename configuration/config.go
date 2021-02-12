package configuration

/*
 * Copied from https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp
 */
import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type ConfigFile struct {
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

type EnvConfig struct {
	Cert string
	Key  string
}

func ReadEnv() *EnvConfig {
	var envConf = &EnvConfig{}
	envConf.Cert = os.Getenv("state-store.cert-path")
	envConf.Key = os.Getenv("state-store.key-path")
	return envConf
}

// NewConfig returns a new decoded ConfigFile struct
func ReadConfig(configPath string) (*ConfigFile, error) {
	// Create config structure
	var config = &ConfigFile{}

	// Open config file
	var file, err = os.Open(configPath)
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
