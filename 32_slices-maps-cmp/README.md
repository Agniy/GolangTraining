# 32 — slices, maps, cmp

Три обобщённых пакета стандартной библиотеки, добавленные в **Go 1.21 (2023)**.
Все реализованы на дженериках (раздел 31), поэтому работают с любыми типами и
проверяются на этапе компиляции. Раньше эти функции писали руками или брали из
`golang.org/x/exp`.

## Содержание

| Пример | Тема | Версия Go |
|---|---|---|
| [01_slices-package](01_slices-package/) | `Sort`, `Contains`, `Index`, `BinarySearch`, `Min/Max`, `Reverse`, `Equal`, `SortFunc` | 1.21 |
| [02_maps-package](02_maps-package/) | `Clone`, `Equal`, `Copy`, `DeleteFunc`; получение ключей/значений | 1.21 |
| [03_cmp-package](03_cmp-package/) | `cmp.Compare`, `cmp.Less`, `cmp.Or`, `cmp.Ordered` | 1.21 / 1.22 |

## Запуск

Раздел — один модуль (`go.mod`, `go 1.22`). Каждый пример — пакет `main`:

```sh
cd 32_slices-maps-cmp
go run ./01_slices-package
go run ./02_maps-package
go run ./03_cmp-package
```

> `go.mod` требует `go 1.22` ради `cmp.Or` (пример 03). Всё остальное работает
> начиная с Go 1.21.

## Шпаргалка

### slices
```go
slices.Sort(s)                       // сортировка на месте
slices.SortFunc(s, cmpFn)            // сортировка по компаратору (-1/0/+1)
slices.Contains(s, v)                // есть ли элемент
slices.Index(s, v)                   // позиция или -1
slices.BinarySearch(s, v)            // (pos, found) в отсортированном срезе
slices.Min(s) / slices.Max(s)        // паникуют на пустом срезе
slices.Reverse(s)                    // разворот на месте
slices.Equal(a, b)                   // поэлементное равенство
```

### maps
```go
maps.Clone(m)                        // поверхностная копия
maps.Equal(a, b)                     // поэлементное равенство
maps.Copy(dst, src)                  // скопировать пары (перезаписывая)
maps.DeleteFunc(m, fn)               // удалить пары по условию
```

> `maps.Keys` / `maps.Values` в Go 1.23 стали **итераторами** (`iter.Seq`,
> раздел 36). Современный способ получить отсортированные ключи:
> `slices.Sorted(maps.Keys(m))`. В примере 02 для совместимости с go 1.22
> ключи собираются циклом `range`.

### cmp
```go
cmp.Compare(a, b)   // -1, 0, +1 — готовый компаратор для slices.SortFunc
cmp.Less(a, b)      // a < b
cmp.Or(vals...)     // первое НЕнулевое значение (Go 1.22)
cmp.Ordered         // ограничение для дженериков: числа и строки
```

Типичный приём — многоуровневая сортировка:
```go
slices.SortFunc(people, func(a, b Person) int {
    return cmp.Or(
        cmp.Compare(a.Age, b.Age),   // сначала по возрасту
        cmp.Compare(a.Name, b.Name), // при равенстве — по имени
    )
})
```
