# Go JSON Crypto Service

Сервис на Go, предназначенный для генерации случайных JSON-файлов, их шифрования с использованием алгоритма Fernet, расшифровки и записи нужных данных в PostgreSQL.

---

## Архитектура

```text
go-json-crypto-service/
├── cmd/                          # Основная логика приложения
│   ├── crypto/                   # Шифрование и расшифровка JSON (Fernet)
│   │   └── crypto.go
│   │
│   ├── db/                       # Подключение к PostgreSQL, создание таблиц, операции с БД
│   │   └── db.go
│   │
│   ├── generator/                # Генерация случайных JSON-объектов
│   │   ├── generator.go
│   │   └── generator_test.go
│   │
│   ├── parser/                   # Парсинг JSON-файлов
│   │   ├── parser.go
│   │   └── parser_test.go
│
├── .env.example                  # Пример конфигурации
├── .gitignore                    # Список игнорируемых файлов/папок для Git
├── go.mod                        # Модули Go
├── go.sum                        # Контрольные суммы зависимостей Go
├── main.go                       # Точка входа, выбор режима, запуск функций
└── README.md                     # Общая документация проекта
```

---

## Установка и запуск без Docker

1. Клонируем проект

```bash
git clone https://github.com/adil-cpu/go-json-crypto-service.git
cd go-json-crypto-service
go mod tidy
```

2. Создаем файл `.env`, скопировав `.env.example`

```bash
cp .env.example .env
```

Сгенерировать `FERNET_KEY` можно так:

```go
fernetKey := fernet.Key{}
fernetKey.Generate()
fmt.Println(fernetKey.Encode())
```


### Генерация и шифрование

```env
MODE=ENCRYPTION
```

```bash
go run main.go
```

### Расшифровка и сохранение в PostgreSQL

```env
MODE=DECRYPTION
```

```bash
go run main.go
```

### Структура БД

```sql
-- Список нужных ключей
CREATE TABLE key_list (
    id SERIAL PRIMARY KEY,
    key_name TEXT UNIQUE NOT NULL
);

-- Таблица для хранения отфильтрованных данных
CREATE TABLE json_data (
    id SERIAL PRIMARY KEY,
    key_name TEXT NOT NULL,
    value TEXT
);
```

### Тесты

Модульные тесты находятся в папках `cmd/generator` и `cmd/parser`

```bash
go test ./...
```

## Установка и запуск с Docker

1. Установить зависимости

Docker: https://docs.docker.com/get-started/get-docker/

Docker Compose (обычно встроет в Docker Desktop): https://docs.docker.com/compose/install/

2. Клонируем проект

```bash
git clone https://github.com/adil-cpu/go-json-crypto-service.git
cd go-json-crypto-service
```

3. Создаем файл `.env`, скопировав `.env.example`

```bash
cp .env.example .env
```

4. Запуск проекта

```bash
docker-compose up --build
```

5. Смена режима (ENCRYPTION ↔ DECRYPTION)

* Открываем `.env` или `docker-compose.yml`

* Меняем
```env
MODE=DECRYPTION
```

* Перезапускаем 
```bash
docker-compose up --build
```