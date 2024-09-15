package config

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func createTempConfigFile(t *testing.T, content map[string]string) string {
	configPath, err := ConstructConfigLocation()
	if err != nil {
		t.Fatalf("ConstructConfigLocation returned an error: %v", err)
	}

	dir := filepath.Dir(configPath)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory %s: %v", dir, err)
	}

	fileData, err := json.Marshal(content)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	err = os.WriteFile(configPath, fileData, 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	return configPath
}

func TestReadConfigFile_Success(t *testing.T) {
	expectedConfig := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}

	configPath := createTempConfigFile(t, expectedConfig)
	defer os.Remove(configPath)

	actualConfig, err := ReadConfigFile()
	if err != nil {
		t.Fatalf("ReadConfigFile returned an error: %v", err)
	}

	if len(expectedConfig) != len(actualConfig) {
		t.Fatalf("Expected %d config entries, got %d", len(expectedConfig), len(actualConfig))
	}
	for key, expectedValue := range expectedConfig {
		if actualValue, ok := actualConfig[key]; !ok {
			t.Errorf("Key %s not found in actual config", key)
		} else if actualValue != expectedValue {
			t.Errorf("For key %s, expected value %s, got %s", key, expectedValue, actualValue)
		}
	}
}

func TestReadConfigFile_FileNotFound(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS") == "1" {
		configPath, _ := ConstructConfigLocation()
		os.Remove(configPath)

		ReadConfigFile()

		t.Fatalf("Expected ReadConfigFile to call log.Fatalln and exit")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestReadConfigFile_FileNotFound")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS=1")
	err := cmd.Run()

	if exitError, ok := err.(*exec.ExitError); ok && !exitError.Success() {
	} else {
		t.Fatalf("Expected subprocess to exit with non-zero status")
	}
}

func TestReadConfigFile_InvalidJSON(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS") == "1" {
		configPath, err := ConstructConfigLocation()
		if err != nil {
			t.Fatalf("ConstructConfigLocation returned an error: %v", err)
		}
		dir := filepath.Dir(configPath)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
		err = os.WriteFile(configPath, []byte("{invalid json}"), 0644)
		if err != nil {
			t.Fatalf("Failed to write invalid config file: %v", err)
		}

		ReadConfigFile()

		t.Fatalf("Expected ReadConfigFile to call log.Fatalln and exit due to invalid JSON")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestReadConfigFile_InvalidJSON")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS=1")
	err := cmd.Run()

	if exitError, ok := err.(*exec.ExitError); ok && !exitError.Success() {
	} else {
		t.Fatalf("Expected subprocess to exit with non-zero status")
	}
}
