# План: добавление примеров возможностей Go начиная с 2015 года

Репозиторий — классический курс Todd McLeod (GOPATH-стиль, без `go.mod`),
покрывает Go примерно до версии 1.5 (2015). Цель — добавить примеры всего,
что появилось в языке с тех пор. Нумерация разделов продолжается с 28.

Текущая версия Go в системе: go1.26.

## Что уже есть / чего нет

Курс заканчивается на: горутины/каналы, error handling, тестирование, пакеты
стандартной библиотеки (01–27). Отсутствует всё, что добавлено с 2015:
модули, context, дженерики, обёртка ошибок, новые встроенные функции,
итераторы, embed и т.д.

---

## Приоритет 1 — фундаментальные изменения

### `28_modules/` — Go Modules (Go 1.11, 2018) ✅ ГОТОВО
- `01_go-mod-init` — создание модуля, `go.mod`/`go.sum` ✅
- `02_adding-dependency` — `go get`, версионирование, `go.sum` ✅
- `03_replace-directive` — `replace` на локальный модуль ✅
- `04_workspaces` — `go.work` (Go 1.18) ✅
- Решение по GOPATH: каждый пример сделан отдельным самостоятельным модулем со
  своим `go.mod`, старая структура 01–27 не затронута. Все примеры проходят
  `go vet` и собираются.

### `29_context/` — пакет `context` (Go 1.7, 2016)
- `01_background-todo`
- `02_with-cancel`
- `03_with-timeout`
- `04_with-deadline`
- `05_with-value`
- `06_http-context` — отмена запросов

### `30_error-wrapping/` — обёртка ошибок (Go 1.13, 2019)
- `01_fmt-errorf-w` — глагол `%w`
- `02_errors-is`
- `03_errors-as`
- `04_errors-unwrap`
- `05_errors-join` (Go 1.20)
- Расширяет существующий `23_error-handling`.

### `31_generics/` — дженерики (Go 1.18, 2022) ⭐ самое важное
- `01_type-parameters` — синтаксис `[T any]`
- `02_constraints` — `comparable`, `constraints.Ordered`
- `03_generic-functions` — Map/Filter/Reduce
- `04_generic-types` — generic-стек/очередь
- `05_type-inference`
- `06_any-vs-interface`
- `07_generic-aliases` (Go 1.24)

---

## Приоритет 2 — стандартная библиотека

### `32_slices-maps-cmp/` — новые пакеты (Go 1.21, 2023)
- `01_slices-package` — `slices.Sort/Contains/Index`
- `02_maps-package` — `maps.Keys/Values/Clone`
- `03_cmp-package`

### `33_builtins/` — новые встроенные функции
- `01_min-max` (Go 1.21)
- `02_clear` (Go 1.21)

### `34_embed/` — `//go:embed` (Go 1.16, 2021)
- `01_embed-string`
- `02_embed-bytes`
- `03_embed-fs`

### `35_slog/` — структурированное логирование `log/slog` (Go 1.21)

---

## Приоритет 3 — более новые фичи

### `36_iterators/` — range-over-func итераторы (Go 1.23, 2024)
- `01_iter-package`
- `02_range-over-func`
- `03_seq-seq2`

### `37_range-improvements/` — улучшения цикла `for` (Go 1.22, 2024)
- `01_range-over-int` — `for i := range 10`
- `02_loop-var-per-iteration` — изменение семантики (частая ошибка с
  горутинами; расширить `22_go-routines/15`)

### `38_number-literals/` — литералы чисел (Go 1.13)
- `01_binary-octal`
- `02_digit-separators` (`1_000_000`)
- `03_hex-floats`

### `39_modern-stdlib/` — прочее
- `01_math-rand-v2` (Go 1.22)
- `02_unique-package` (Go 1.23)
- `03_sync-map` (Go 1.9)
- `04_type-aliases` (Go 1.9)
- `05_json-omitzero` (Go 1.24)

### `40_fuzzing/` — fuzz-тесты (Go 1.18) — расширяет `24_testing`

---

## Порядок реализации

1. **Модули** (`28`) — без этого современные примеры не запускаются как раньше.
2. **Дженерики** (`31`), **context** (`29`), **error wrapping** (`30`) — самые
   востребованные.
3. Стандартная библиотека (`32`–`35`).
4. Новые синтаксические возможности (`36`–`39`).

## Открытые вопросы
- Как поступить с миграцией на модули, не ломая GOPATH-структуру остального
  репозитория (отдельный go.mod в новых папках vs. модуль на весь репо).