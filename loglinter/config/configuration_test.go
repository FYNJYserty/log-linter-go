package config

import (
    "os"
    "testing"
)

func TestCheckRulesConfig(t *testing.T) {
	// Тестирование включения правил
	rules = RulesConfig{
		LowerCase:     true,
		SpecSymbols:   false,
		CheckEnglish:  false,
		SensitiveData: false,
	}

	if !CheckRulesConfig("lower-case") {
		t.Error("Expected lower-case rule to be enabled")
	}
	if CheckRulesConfig("spec-symbols") {
		t.Error("Expected spec-symbols rule to be disabled")
	}
	if CheckRulesConfig("check-english") {
		t.Error("Expected check-english rule to be disabled")
	}
	if CheckRulesConfig("sensitive-data") {
		t.Error("Expected sensitive-data rule to be disabled")
	}
}

func TestLoadConfig(t *testing.T) {
	tmp := t.TempDir()
	path := tmp + "/.golangci.yml"
	content := []byte("rules:\n  lower-case: true\n  sensitive-data: false\n")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
	if err := LoadConfig(path); err != nil {
		t.Fatalf("LoadConfig returned error: %v", err)
	}
	if !CheckRulesConfig("lower-case") {
		t.Error("lower-case should be enabled after loading config")
	}
	if CheckRulesConfig("sensitive-data") {
		t.Error("sensitive-data should be disabled after loading config")
	}
}
