# hello-gui

[![CI](https://github.com/anitouchq/hello-gui/actions/workflows/ci.yml/badge.svg)](https://github.com/anitouchq/hello-gui/actions)

Учебное desktop-приложение на Go с библиотекой [Fyne](https://fyne.io/). GitHub Actions проверяет код и тесты, а по тегу версии собирает приложения под Linux, macOS и Windows.

## Возможности приложения

- Ввод имени и вывод приветствия.
- Отображение версии, переданной из Git-тега при release-сборке.
- Демонстрация тестируемой функции суммы диапазона.
- Завершение приложения кнопкой Exit.

## CI checks

- `gofmt` — проверка форматирования;
- `go vet ./...` — статический анализ;
- `go test ./... -v` — unit-тесты логики;
- `go build` — проверочная сборка GUI.

Fyne использует CGO и системные графические библиотеки. Поэтому релизные бинарники собираются на нативных runner'ах под каждую ОС.

## Release matrix

| Платформа | Runner | Артефакт |
|---|---|---|
| Linux x64 | Ubuntu | `hello-gui-linux-x64` |
| macOS | macOS | `hello-gui-macos` |
| Windows x64 | Windows + MinGW | `hello-gui-windows-x64.exe` |

## Как создать Release

```bash
git tag v1.0.0
git push origin v1.0.0
```

После CI три runner'а параллельно собирают приложения, загружают бинарники как artifacts, а отдельный job создаёт [GitHub Release](https://github.com/anitouchq/hello-gui/releases).

## Локальная сборка

Требуются Go 1.27+, GCC и системные графические библиотеки Fyne.

Для Ubuntu/Debian:

```bash
sudo apt-get install gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev
go test ./... -v
go run .
```

## Структура проекта

| Файл или каталог | Назначение |
|---|---|
| `.github/workflows/ci.yml` | Проверки, матричная сборка и публикация Release |
| `main.go` | Создание окна Fyne и обработчики интерфейса |
| `logic.go` | Логика приветствия и вычислений, отделённая от UI |
| `logic_test.go` | Unit-тесты прикладной логики |
| `go.mod` | Имя Go-модуля и прямые зависимости |
| `go.sum` | Контрольные суммы зависимостей |
| `.gitignore` | Локальные файлы, исключаемые из Git |

## Результаты

- [GitHub Actions](https://github.com/anitouchq/hello-gui/actions)
- [Опубликованные релизы](https://github.com/anitouchq/hello-gui/releases)


## 📸 Скриншоты

### 1. Работающее приложение

![Окно Hello GUI](screenshots/application.png)

### 2. GitHub Actions — CI/CD успешно прошёл

![GitHub Actions — успешная сборка](screenshots/github-actions-release.png)

### 3. GitHub Release v1.0.3 — 3 бинарника

![GitHub Release с бинарными файлами](screenshots/github-release.png)

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
