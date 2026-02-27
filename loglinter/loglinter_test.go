package loglinter

import (
	"os/exec"
	"strings"
	"testing"
)

func TestLogLinter(t *testing.T) {
	cmd := exec.Command("go", "build", "-o", "../build/mylinter", "../cmd")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build linter: %v\n%s", err, out)
	}

	cmd = exec.Command("../build/mylinter", "../testdata/example.go")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("expected non-zero exit code, got output: %s", out)
	}

	outStr := string(out)

	if !strings.Contains(outStr, "lowercase letter") {
		t.Errorf("expected at least one lowercase-letter error, got:\n%s", outStr)
	}
	if !strings.Contains(outStr, "sensitive data") {
		t.Errorf("expected at least one sensitive-data error, got:\n%s", outStr)
	}
}
