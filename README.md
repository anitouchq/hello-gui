# 🎨 hello-gui — Go + Fyne → GitHub Releases

Учебное desktop-приложение на Go с библиотекой [Fyne](https://fyne.io/). GitHub Actions проверяет код и тесты, затем по тегу версии собирает нативные приложения под Linux, macOS и Windows.

## Что умеет приложение

- Ввод имени и вывод приветствия.
- Отображение версии, переданной из Git-тега при release-сборке.
- Демонстрация тестируемой функции суммы диапазона.
- Завершение приложения кнопкой Exit.

## CI checks

- `gofmt` — проверка форматирования;
- `go vet ./...` — статический анализ;
- `go test ./... -v` — unit-тесты логики;
- `go build` — smoke-build GUI приложения.

Для Linux runner workflow устанавливает системные OpenGL/X11/Wayland зависимости. Fyne использует CGO, поэтому релизные бинарники собираются на нативных runner'ах, а не обычной Go cross-compilation.

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

Workflow сначала запускает CI. Затем три нативных runner'а параллельно собирают приложение, загружают бинарники как artifacts, а отдельный job создаёт GitHub Release.

## Локальная сборка

Требуется Go 1.27+, GCC и системные графические библиотеки Fyne.

Ubuntu/Debian:

```bash
sudo apt-get install gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev
```

Далее:

```bash
go mod tidy
go test ./... -v
go run .
```

## Permissions

Только release job использует `contents: write` для публикации GitHub Release через встроенный `GITHUB_TOKEN`.

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
