package config

import "github.com/BurntSushi/toml"

// Contains configuration values needed for the server
type Config struct {
  static_root struct {
    static_views_directory string
    static_template_name string
  }
}

// Loads a configuration from a TOML file
func GenerateConfigFromFile(file_path string) (*Config, error) {
  var config Config

  _, err := toml.DecodeFile(file_path, &config)
  if err != nil {
    return nil, err
  }

  return &config, nil
}
