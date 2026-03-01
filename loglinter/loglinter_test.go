package loglinter

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/FYNJYserty/log-linter-go/loglinter/config"
)

func TestLogLinter(t *testing.T) {
	// Загружаем конфигурацию из существующего файла .golangci.yml в корневой директории
	cfgPath := "../.golangci.yml"
	if err := config.LoadConfig(cfgPath); err != nil {
		t.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	// Копируем конфигурацию в корень для запускаемого линтера
	cmd := exec.Command("cp", cfgPath, ".golangci.yml")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to copy config: %v\n%s", err, out)
	}
	defer os.Remove(".golangci.yml")

	cmd = exec.Command("go", "build", "-o", "../build/mylinter", "../cmd")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build linter: %v\n%s", err, out)
	}

	cmd = exec.Command("../build/mylinter", "../testdata/example.go")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("expected non-zero exit code, got output: %s", out)
	}

	outStr := string(out)

	if !strings.Contains(outStr, "lowercase letter") && config.CheckRulesConfig("lower-case") {
		t.Errorf("expected at least one lowercase-letter error, got:\n%s", outStr)
	}
	if !strings.Contains(outStr, "sensitive data") && config.CheckRulesConfig("sensitive-data") {
		t.Errorf("expected at least one sensitive-data error, got:\n%s", outStr)
	}
}
