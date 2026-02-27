package loglinter

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func extractLogCall(call *ast.CallExpr, pass *analysis.Pass) *LogCall {
	// Проверяем вызовы slog
	if isSlogCall(call, pass) {
		return extractFromSlog(call, pass)
	}

	// Проверяем вызовы zap
	if isZapCall(call, pass) {
		return extractFromZap(call, pass)
	}

	// Стандартная библиотека log
	if isStdLogCall(call, pass) {
		return extractFromStdLog(call, pass)
	}

	return nil
}

// Проверка на то, что вызов является логированием slog
func isSlogCall(call *ast.CallExpr, pass *analysis.Pass) bool {
	if len(call.Args) < 1 {
		return false
	}

	selExpr, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	ident, ok := selExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	if ident.Name != "slog" {
		return false
	}

	method := selExpr.Sel.Name
	return method == "Info" || method == "Error" || method == "Warn" || method == "Debug"
}

// Проверка на то, что вызов является логированием zap
func isZapCall(call *ast.CallExpr, pass *analysis.Pass) bool {
	selExpr, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	method := selExpr.Sel.Name
	return method == "Info" || method == "Error" || method == "Warn" || method == "Debug" || method == "Fatal"
}

// Проверка на то, что вызов является логированием стандартной библиотеки log
func isStdLogCall(call *ast.CallExpr, pass *analysis.Pass) bool {
	if len(call.Args) < 1 {
		return false
	}

	selExpr, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	ident, ok := selExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	if ident.Name != "log" {
		return false
	}

	method := selExpr.Sel.Name
	// common log methods
	return method == "Print" || method == "Println" || method == "Printf" ||
		method == "Fatal" || method == "Fatalf" || method == "Fatalln" ||
		method == "Panic" || method == "Panicf" || method == "Panicln"
}

// Извлечение информации из вызова slog
func extractFromSlog(call *ast.CallExpr, pass *analysis.Pass) *LogCall {
	selExpr := call.Fun.(*ast.SelectorExpr)
	method := selExpr.Sel.Name

	// Для slog первый аргумент - это сообщение
	if len(call.Args) > 0 {
		if msgExpr := getStringLiteral(call.Args[0]); msgExpr != "" {
			return &LogCall{
				Pos:      call.Pos(),
				Message:  msgExpr,
				Function: method,
			}
		}
	}

	return nil
}

// Извлечение информации из вызова zap
func extractFromZap(call *ast.CallExpr, pass *analysis.Pass) *LogCall {
	selExpr := call.Fun.(*ast.SelectorExpr)
	method := selExpr.Sel.Name

	if len(call.Args) > 0 {
		if msgExpr := getStringLiteral(call.Args[0]); msgExpr != "" {
			return &LogCall{
				Pos:      call.Pos(),
				Message:  msgExpr,
				Function: method,
			}
		}
	}

	return nil
}

// Извлечение информации из вызова стандартной библиотеки log
func extractFromStdLog(call *ast.CallExpr, pass *analysis.Pass) *LogCall {
	selExpr := call.Fun.(*ast.SelectorExpr)
	method := selExpr.Sel.Name

	if len(call.Args) > 0 {
		if msgExpr := getStringLiteral(call.Args[0]); msgExpr != "" {
			return &LogCall{
				Pos:      call.Pos(),
				Message:  msgExpr,
				Function: method,
			}
		}
	}

	return nil
}

// Получение строкового литерала из выражения
func getStringLiteral(expr ast.Expr) string {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return ""
	}

	// Удаляем кавычки
	return strings.Trim(lit.Value, `"`)
}
