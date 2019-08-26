package godev

import (
	"strings"

	"gopkg.in/yaml.v3"
)

func YAML(input interface{}) string {
	out, _ := yaml.Marshal(input)
	return strings.TrimSpace(string(out))
}
