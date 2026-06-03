# 03 — директива replace

`replace` перенаправляет модуль на **другой источник**: локальную папку, форк
или конкретную версию. Частые сценарии:

- одновременная разработка приложения и его библиотеки;
- временный патч/форк сторонней зависимости;
- работа без доступа к интернету.

## Структура примера

```
03_replace-directive/
├── greeting/          # отдельный локальный модуль-библиотека
│   ├── go.mod         #   module example.com/greeting
│   └── greeting.go
└── app/               # приложение, использующее greeting
    ├── go.mod         #   require + replace => ../greeting
    └── main.go
```

Ключевая строка в `app/go.mod`:

```
require example.com/greeting v0.0.0
replace example.com/greeting => ../greeting
```

Модуль `example.com/greeting` нигде не опубликован — `replace` заставляет Go
брать его прямо из соседней папки.

## Запуск

```sh
cd 28_modules/03_replace-directive/app
go run .
```

Вывод:

```
Привет, мир!
```

## Формы replace

```
replace example.com/lib => ../lib                 # локальный путь
replace example.com/lib => example.com/fork v1.2.3 # другой модуль/версия
replace example.com/lib v1.0.0 => example.com/lib v1.0.1 # подмена версии
```

> С появлением воркспейсов (`go.work`, см. пример 04) для локальной
> мультимодульной разработки чаще используют их, а `replace` оставляют для
> форков и точечных подмен версий.
