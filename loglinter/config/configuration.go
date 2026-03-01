package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Rules    RulesConfig    `yaml:"rules"`
	Patterns PatternsConfig `yaml:"patterns"`
}

// Структура для хранения конфигурации правил линтера
type RulesConfig struct {
	LowerCase     bool `yaml:"lower-case"`
	SpecSymbols   bool `yaml:"spec-symbols"`
	CheckEnglish  bool `yaml:"check-english"`
	SensitiveData bool `yaml:"sensitive-data"`
}

// Структура для хранения конфигурации шаблонов (например, чувствительных ключевых слов)
type PatternsConfig struct {
	SensitiveKeywords []string `yaml:"sensitive-keywords"`
	Patterns          []string `yaml:"patterns"`
}

// Хранит текущую конфигурацию правил
var rules = RulesConfig{}
var patterns = PatternsConfig{}

// LoadConfig читает .golangci.yml и обновляет глобальную конфигурацию
func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}
	rules = cfg.Rules
	patterns = cfg.Patterns
	return nil
}

// Проверка правил
func CheckRulesConfig(ruleName string) bool {
	switch ruleName {
	case "lower-case":
		return rules.LowerCase
	case "spec-symbols":
		return rules.SpecSymbols
	case "check-english":
		return rules.CheckEnglish
	case "sensitive-data":
		return rules.SensitiveData
	default:
		return false
	}
}

// Выгрузка ключевых слов из конфигурации
func GetSensitiveKeywords() []string {
	return patterns.SensitiveKeywords
}

// Выгрузка шаблонов из конфигурации
func GetPatterns() []string {
	return patterns.Patterns
}
