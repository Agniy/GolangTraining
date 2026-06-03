module example.com/app

go 1.21

// Обратите внимание: здесь есть require, но НЕТ директивы replace.
// Локальную привязку к ../greeting обеспечивает go.work уровнем выше.
// Если убрать go.work, для сборки понадобится либо replace, либо публикация
// модуля greeting.
require example.com/greeting v0.0.0
