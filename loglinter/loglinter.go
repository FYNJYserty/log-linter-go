package loglinter

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/FYNJYserty/log-linter-go/loglinter/config"
	"github.com/FYNJYserty/log-linter-go/loglinter/rules"
	"golang.org/x/tools/go/analysis"
)

// Структура вывода информации о логе
type LogCall struct {
	Pos      token.Pos
	Message  string
	Function string // "Info", "Error", etc.
	LitPos   token.Pos
	LitEnd   token.Pos
}

// Основная функция анализа, которая будет вызываться при запуске линтера
func run(pass *analysis.Pass) (interface{}, error) {
	// Загружаем конфигурацию, ищем .golangci.yml вверх по директориям
	if len(pass.Files) > 0 {
		currentDir := filepath.Dir(pass.Fset.Position(pass.Files[0].Pos()).Filename)
		
		for filepath.Dir(currentDir) != currentDir {
			if config.LoadConfig(filepath.Join(currentDir, ".golangci.yml")) == nil {
				break
			}
			currentDir = filepath.Dir(currentDir)
		}
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			node, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			logCall := extractLogCall(node, pass)
			if logCall == nil {
				return true
			}

			// Применяем все правила к найденному лог-вызову
			checkAllRules(pass, logCall)
			return true
		})
	}
	return nil, nil
}

// Функция проверки всех правил
func checkAllRules(pass *analysis.Pass, call *LogCall) {
	if !rules.CheckLowerCase(call.Message) && config.CheckRulesConfig("lower-case") {
		startPos := call.LitPos + 1
		origFirst := []byte(string([]rune(call.Message)[0]))
		endPos := startPos + token.Pos(len(origFirst))

		newFirst := []byte(strings.ToLower(string([]rune(call.Message)[0])))

		edit := analysis.TextEdit{
			Pos:     startPos,
			End:     endPos,
			NewText: newFirst,
		}

		fix := analysis.SuggestedFix{
			Message:   "Make first letter to lower case",
			TextEdits: []analysis.TextEdit{edit},
		}

		pass.Report(analysis.Diagnostic{
			Pos:            call.LitPos,
			Message:        "Make first letter to lower case",
			SuggestedFixes: []analysis.SuggestedFix{fix},
		})
		pass.Reportf(call.LitPos, "\nError log: \"%s\" --Message should start with a lowercase letter--", call.Message)
	}
	if rules.ContainsSensitiveData(call.Message) && config.CheckRulesConfig("sensitive-data") {
		pass.Reportf(call.Pos, "\nError log: \"%s\" --Message contains sensitive data (password, token, api_key, etc)--", call.Message)
	}
	if !rules.IsNoSpecSymbols(call.Message) && config.CheckRulesConfig("spec-symbols") {
		pass.Reportf(call.Pos, "\nError log: \"%s\" --Message should not contain special symbols--", call.Message)
	}
	if !rules.IsEnglish(call.Message) && config.CheckRulesConfig("check-english") {
		pass.Reportf(call.Pos, "\nError log: \"%s\" --Message should be in English--", call.Message)
	}
}
