package rules

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"go.yaml.in/yaml/v4"
)

func LoadRules(path string) []Rule {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("read rules directory %w", err)
	}
	result := make([]Rule, 0)
	for _, v := range entries {
		if v.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(v.Name()))
		if ext != ".yaml" && ext != ".yml" {
			log.Printf("file is not yaml %s", v.Name())
			continue
		}
		data, err := os.ReadFile(filepath.Join(path, v.Name()))
		if err != nil {
			log.Printf("cannot read file %s", filepath.Join(path, v.Name()))
			continue
		}

		var rule Rule
		err = yaml.Unmarshal(data, &rule)
		if err != nil {
			log.Printf("cannot read unmarshal %s", filepath.Join(path, v.Name()))
			continue
		}
		result = append(result, rule)
	}
	return result
}
