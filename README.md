# Тестовое задание по направлению «Backend-разработка. Golang (Selectel)»
## Линтер для проверки лог-записей

Описание задания:

Необходимо разработать линтер для Go, совместимый с golangci-lint,
который будет анализировать лог-записи в коде и проверять их
соответствие установленным правилам.

## Технические требования:
1. Язык разработки: Go 1.22+
2. Совместимость: Линтер должен работать как плагин для
golangci-lint
3. Поддерживаемые логгеры:
- log/slog
- go.uber.org/zap

## Порядок проверки программы
Копирование репозитория:
```bash
git clone https://github.com/FYNJYserty/log-linter-go
cd <место расположения репозитория>
```
Сборка программы:
```bash
go build -o ./build/mylinter ./cmd/main.go
```
Проверка работы программы на основе тестовой программы с логами:
```bash
./build/mylinter ./testdata/example.go
```
## Интеграция с golangci-lint
Прежде всего в папке ```/testdata``` надо создать файл ```.golangci.yml``` следующего содержания
```
version: "2"

linters:
  default: none
  enable:
    - loglinter
```
Далее скопировать golangci-lint в проект по команде
```bash
git clone https://github.com/golangci/golangci-lint
```
Создать папку линтера по пути 
```/pkg/golinters/loglinter/mylinter.go```
В файл Go вписать следующее:
```
package loglinter

import (
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/loglinter/loglinter"
)

func New() *goanalysis.Linter {

	return goanalysis.
		NewLinter("loglinter", "My description", []*analysis.Analyzer{
			loglinter.Analyzer,
		}, nil).WithLoadMode(goanalysis.LoadModeSyntax)
}
```
В ту же папку скопировать папку проекта ```/loglinter``` и ```/testdata```

В свойства функции ```New()``` вписать следующий модуль

```
linter.NewConfig(loglinter.New()).
			WithLoadForGoAnalysis(),
```
Перейти в репозиторий линтера и осуществить сборку по командам
```bash
cd golangci-lint
make build
```
Проверить наличие линтера
```bash
./golangci-lint linters | grep loglinter
```
Проверить работу линта по команде на основе тестовых логов с прошлого раздела
```bash
go run ./cmd/golangci-lint/ run --no-config --default=none --enable=loglinter ./pkg/golinters/loglinter/testdata/example.go
```
После всех проделанных шагов должен появиться вывод в консоль
```bash
pkg/golinters/loglinter/testdata/example.go:16:2: Error log: "This is a log" --Message should start with a lowercase letter-- (loglinter)
        log.Println("This is a log")
        ^
pkg/golinters/loglinter/testdata/example.go:17:2: Error log: "Bad message!" --Message should start with a lowercase letter-- (loglinter)
        log.Println("Bad message!")
        ^
pkg/golinters/loglinter/testdata/example.go:18:2: Error log: "Русский лог с заглавной" --Message should start with a lowercase letter-- (loglinter)
        log.Fatalf("Русский лог с заглавной")
        ^
pkg/golinters/loglinter/testdata/example.go:19:2: Error log: "пример лога!" --Message should not contain special symbols-- (loglinter)
        log.Println("пример лога!")
        ^
pkg/golinters/loglinter/testdata/example.go:27:2: Error log: "This is an error message slog" --Message should start with a lowercase letter-- (loglinter)
        slog.Info("This is an error message slog")
        ^
pkg/golinters/loglinter/testdata/example.go:28:2: Error log: "api_key=sk_live_abcd1234" --Message contains sensitive data (password, token, api_key, etc)-- (loglinter)
        slog.Debug("api_key=sk_live_abcd1234")
        ^
pkg/golinters/loglinter/testdata/example.go:29:2: Error log: "user password: password123" --Message contains sensitive data (password, token, api_key, etc)-- (loglinter)
        slog.Info("user password: password123")
        ^
pkg/golinters/loglinter/testdata/example.go:30:2: Error log: "token: 1233qwee" --Message contains sensitive data (password, token, api_key, etc)-- (loglinter)
        slog.Debug("token: 1233qwee")
        ^
pkg/golinters/loglinter/testdata/example.go:38:2: Error log: "This is standard library logging" --Message should start with a lowercase letter-- (loglinter)
        logger.Info("This is standard library logging")
        ^
pkg/golinters/loglinter/testdata/example.go:39:2: Error log: "server started! 🚀" --Message should not contain special symbols-- (loglinter)
        logger.Debug("server started! 🚀")
        ^
pkg/golinters/loglinter/testdata/example.go:40:2: Error log: "connection failed!!!" --Message should not contain special symbols-- (loglinter)
        logger.Info("connection failed!!!")
        ^
pkg/golinters/loglinter/testdata/example.go:41:2: Error log: "warning: something went wrong..." --Message should not contain special symbols-- (loglinter)
        logger.Debug("warning: something went wrong...")
        ^
12 issues:
* loglinter: 12
exit status 1
```
Как-то так :З
## Бонусные задания
1. Конфигурация: Добавить возможность настройки правил через
конфигурационный файл ❌
2. Авто-исправление: Реализовать SuggestedFixes для
автоматического исправления ошибок ❌
3. Кастомные паттерны: Добавить возможность указывать свои
паттерны для проверки чувствительных данных ❌
4. CI/CD: Подготовить CI Gitlab/GitHub для автоматической сборки и
тестирования ✅
