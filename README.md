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
Прежде всего в корне проекта надо создать файл ```.golangci.yml``` следующего содержания
```
version: "2"

plugins:
  enable:
    - loglinter
  local:
    - path: ./plugin/loglinter.so

linters:
  default: none
  enable:
    - loglinter

rules:
  lower-case: true
  spec-symbols: true
  check-english: true
  sensitive-data: true

patterns:
  sensitive-keywords:
    - password
    - passwd
    - pwd
    - api_key
    - apikey
    - api-key
    - secret
    - credential
    - credentials
    - private_key
    - privatekey
    - private-key
    - access_token
    - accesstoken
    - refresh_token
    - refreshtoken
    - bearer
    - aws_secret
    - db_password
    - database_password
  patterns:
    - (?i)(token|password|passwd|pwd|api[_-]?key|secret|credential|access[_-]?token|refresh[_-]?token|bearer|private[_-]?key)\s*[:=]
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
В ту же папку скопировать папку проекта ```/loglinter``` и ```/testdata``` и добавить файл конфигурации ```.golangci.yml```

В свойства функции ```New()``` в ```mylinter.go``` вписать следующий модуль

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
golangci-lint cache clean

go run ./cmd/golangci-lint/ run --no-config --default=none --enable=loglinter ./pkg/golinters/loglinter/testdata/example.go
```
После всех проделанных шагов должен появиться вывод в консоль
```bash
pkg/golinters/loglinter/testdata/example.go:16:14: Make first letter to lower case (loglinter)
        log.Println("This is a log")
                    ^
pkg/golinters/loglinter/testdata/example.go:17:14: Make first letter to lower case (loglinter)
        log.Println("Bad message!")
                    ^
pkg/golinters/loglinter/testdata/example.go:18:13: Make first letter to lower case (loglinter)
        log.Fatalf("Русский лог с заглавной")
                   ^
pkg/golinters/loglinter/testdata/example.go:19:2: Error log: "пример лога!" --Message should not contain special symbols-- (loglinter)
        log.Println("пример лога!")
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
pkg/golinters/loglinter/testdata/example.go:39:2: Error log: "server started! 🚀" --Message should not contain special symbols-- (loglinter)
        logger.Debug("server started! 🚀")
        ^
pkg/golinters/loglinter/testdata/example.go:40:2: Error log: "connection failed!!!" --Message should not contain special symbols-- (loglinter)
        logger.Info("connection failed!!!")
        ^
pkg/golinters/loglinter/testdata/example.go:41:2: Error log: "warning: something went wrong..." --Message should not contain special symbols-- (loglinter)
        logger.Debug("warning: something went wrong...")
        ^
10 issues:
* loglinter: 10
exit status 1
```
Как-то так :З
## Бонусные задания
1. Конфигурация: Добавить возможность настройки правил через
конфигурационный файл ✅

Был добавлен в линтер пакет конфигурации ```config```, в котором есть функция, 
считывающая данные из файла ```.golangci.yml```. Загрузка осуществляется в ```loglinter.go```. Теперь есть возможность выбора активных правил.

2. Авто-исправление: Реализовать SuggestedFixes для
автоматического исправления ошибок ✅

После применения следующей команды в файле с тестовыми примерами будут внесены изменения в сам файл, буквы 
в логах станут строчными. Автоисправление работает только для 1 правила.

При использовании кастомных настроек стоит сбрасывать кэш после каждого изменения конфигурационного файла командой
```bash
golangci-lint cache clean  
```
Затем запустить команду
```bash
go run ./cmd/golangci-lint/ run --no-config --fix --default=none --enable=loglinter ./pkg/golinters/loglinter/testdata/example.go
```
Изменёния будут зафиксированы в ```/testdata/example.go```

3. Кастомные паттерны: Добавить возможность указывать свои
паттерны для проверки чувствительных данных ✅

Добавлена новая структура в конфигурационный файл, паттерны и ключевые слова для проверки чувствительных 
данных можно записывать в конфигурационный файл ```.golangci.yml```

4. CI/CD: Подготовить CI Gitlab/GitHub для автоматической сборки и
тестирования ✅

Добавлен файл ```.github/workflows/ci.yml``` с автотестами и бидом перед коммитом