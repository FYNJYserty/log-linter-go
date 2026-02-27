package loglinter

import (
	"go/ast"
	"go/token"

	"github.com/FYNJYserty/log-linter-go/loglinter/rules"
	"golang.org/x/tools/go/analysis"
)

// Структура вывода информации о логе
type LogCall struct {
	Pos      token.Pos
	Message  string
	Function string // "Info", "Error", etc.
}

// Основная функция анализа, которая будет вызываться при запуске линтера
func run(pass *analysis.Pass) (interface{}, error) {
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
	if !rules.CheckLowerCase(call.Message) {
		pass.Reportf(call.Pos, "\nError log: \"%s\" --Message should start with a lowercase letter--", call.Message)
	}
	if rules.ContainsSensitiveData(call.Message) {
		pass.Reportf(call.Pos, "\nError log: \"%s\" --Message contains sensitive data (password, token, api_key, etc)--", call.Message)
	}
	if !rules.IsNoSpecSymbols(call.Message) {
		pass.Reportf(call.Pos, "\nError log: \"%s\" --Message should not contain special symbols--", call.Message)
	}
	if !rules.IsEnglish(call.Message) {
		pass.Reportf(call.Pos, "\nError log: \"%s\" --Message should be in English--", call.Message)
	}
}
