# Подключение `loglint` через `golangci-lint custom`

`loglint` подключается к `golangci-lint` как модульный плагин `github.com/irumako/loglint/plugin`.

## Что нужно заранее

- `golangci-lint` v2
- `go` и `git`
- доступ к модулю `github.com/irumako/loglint`
- тег релиза `loglint`, который будет указан в `plugins.version`

## 1. Конфиг для сборки кастомного бинаря

Создайте файл `.custom-gcl.yml`:

```yaml
version: v2.11.3

plugins:
  - module: github.com/irumako/loglint
    import: github.com/irumako/loglint/plugin
    version: <tag>
```

`<tag>` должен совпадать с опубликованным тегом репозитория `github.com/irumako/loglint`.

## 2. Конфиг проекта для запуска линтера

Добавьте `loglint` в обычный `.golangci.yml` проекта:

```yaml
version: "2"

linters:
  default: none
  enable:
    - loglint
  settings:
    custom:
      loglint:
        type: module
        description: Checks log messages.
```

Имя `loglint` должно совпадать с именем анализатора, который экспортирует плагин.

## 3. Сборка и запуск

Соберите кастомный бинарь:

```bash
golangci-lint custom
```

По умолчанию команда создаёт бинарь `custom-gcl` в текущей директории. Затем запустите его в проекте:

```bash
./custom-gcl run
```

На Windows:

```powershell
.\custom-gcl.exe run
```
