# dish-hub

## Просмотр данных в MongoDB
### Вариант 1: MongoDB Compass (рекомендуется)
Скачайте MongoDB Compass

Подключитесь с настройками:

```bash
Hostname: localhost
Port: 27017
```
Откройте базу **dish_hub** → коллекцию **recipes**

### Вариант 2: Командная строка
```bash
mongosh
use dish_hub
db.recipes.find().pretty()
```

## Установка и запуск

### 1. Клонирование репозитория
```bash
git clone https://github.com/bulbosaur/dish-hub
cd dish-hub
```

### 2. Запуск MongoDB
Убедитесь, что MongoDB запущен локально:

```bash
# Для Linux/macOS
sudo systemctl start mongod

# Для Windows (запустите из папки установки)
mongod.exe
```

### 3. Установите зависимости
```bash
go mod download
```
### 4. Запустите приложение
``` bash
go run ./cmd/main.go
```