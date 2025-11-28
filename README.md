🎓 Fullstack LMS — Learning Management System

Простой, но расширяемый менеджер курсов: создавайте, редактируйте, удаляйте курсы через красивый интерфейс.  
**Backend**: Go (Chi, GORM, PostgreSQL)  
**Frontend**: React (JavaScript, Tailwind CSS)  
**Сборка и запуск**: Docker / Docker Compose  
**Особенность**: фронтенд без nginx — только `serve` или `react-scripts`.

---

## 📁 Структура проекта

```
FULLSTACK_LMS/
├── backend/          # Go-бэкенд
│   ├── app/          # main.go
│   ├── internal/     # api, config, model, repository, service
│   ├── migrations/   # SQL-миграции (Goose)
│   ├── pkg/          # logger, validator
│   └── Dockerfile
├── frontend/         # React-фронтенд
│   ├── src/          # компоненты: AddCourseForm, CourseCard, App.jsx
│   ├── public/
│   ├── package.json
│   └── Dockerfile
├── docker-compose.yml      
└── README.md   
```

# Настройка окружения

Необходимо создать два .env файла для запуска проекта, вот их содержание:

postgres.env - для PostgreSQL:

```
POSTGRES_USER=ваш_пользователь
POSTGRES_PASSWORD=ваш_пароль
POSTGRES_DB=ваша_база_данных
```

app.env - для приложения (укажите те же данные, что и для PostgreSQL):

```
DB_USER=ваш_пользователь
DB_PASS=ваш_пароль
DB_NAME=ваша_база_данных
DB_HOST=ваш хост
DB_PORT=ваш порт(обычно 5432)
```

# Запуск через Docker

``` 
docker compose up --build 
```

# Просмотр логов отдельных сервисов

Выполните команду:
```
docker-compose logs backend
docker-compose logs frontend
```
# Остановка всех сервисов

Выполните: 
```
docker-compose down
``` 

🧪 Ручное тестирование API

📥 Создать курс 

```
curl -X POST http://localhost:8081/api/courses \
  -H "Content-Type: application/json" \
  -d '{"title":"Go для продакшена","description":"Курс по надёжным сервисам","author":"Артём"}'
```

📤 Получить все курсы 

```
curl http://localhost:8081/api/courses
```

🗑 Удалить курс 

```
curl -X DELETE http://localhost:8081/api/courses/ВАШ_UUID
# Пример UUID: f47ac10b-58cc-4372-a567-0e02b2c3d479
```

🧑‍💻 Локальная разработка 

Backend (Go) 

```
cd backend
go mod tidy
go run app/main.go
# → Слушает :8081
# → Подключается к локальному PostgreSQL (настройки в env)
```

Frontend (React)

```
cd frontend
npm install
npm start
# → http://localhost:3000
```

🧰 Миграции (Goose) 

Файл миграции: backend/migrations/000001_create_courses_table.sql 

```
goose -dir ./migrations postgres "postgres://DB_USER:DB_PASSWORD@localhost:port/DB_NAME?sslmode=disable" up
```