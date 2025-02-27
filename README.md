# Запуск

Сервис имеет одну команду с одним флагом:
 * go run main serve --config="config.toml"

В докере:
 * docker-compose up

# Коментарии по ТЗ
Функциональные требования:
- ~~Создание зашифрованной строки~~ - пакет hasher
- ~~Выдача зашифрованной строки~~ - [POST] /hash
- ~~Слой кеширования~~ - редис + мок (для тестов + возможности поднять без кеша)
- ~~Обработка ошибок~~ - логаем ошибки 
- ~~Серивис и редис обернуты в докер контейнеры и запуск происходит с них же~~
- ~~Конфигурационные файлы~~ - я бы предпочел ENV т.к при контеринизации c ними проше работать.

Нефункциональные требования: 
- ~~Документирование кода и архитектуры приложения для лучшей поддерживаемости и расширяемости~~ - есть коменты к некоторым сложным(на мой взгляд функциям), остальные считаю понятны по неймингу.

# API 
 * [POST] 0.0.0.0:8090/hash
```JSON
{
    "text": "sss",
    "alg" : "sha256"
}
```
Возможные алгоритмы:
 * SHA256
 * MD5

Возможные коды и респонсы
 * 200
 * 400/500

при 200:
```JSON
{
    "text": "a871c47a7f48a12b38a994e48a9659fab5d6376f3dbce37559bcb617efe8662d"
}
```
при 400/500:
```JSON
{
    "error": "текст ошибки"
}
```

# Тесты
 * go test ./...
