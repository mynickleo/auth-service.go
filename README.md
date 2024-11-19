# ✨ Auth Service Backend

Это небольшой бэкенд-сервис для тестового задания


## 📦 Используемые технологии | Technologies Used
- Go
- PostgreSQL
- Sqlc

## 📂 Структура проекта | Project Structure
- `cmd/server/`
- `config/`
- `internal/`
- - `app/`
- - `controllers/`
- - `database/`
- - - `postgres/`
- - - - `migrations/`
- - - - `queries/`
- - `interfaces/`
- - `models/`
- - `module/`
- - `repository/`
- - `services/`
- - `utils/`
- `pkg/`
- - `sqlcqueries/`

## ⚙️ Установка и запуск | Installation and Launch

```bash
git clone https://github.com/mynickleo/auth-service.go.git
cd auth-service.go
```

```bash
docker-compose up --build
```

## 🔗 API

**`Аутентификация`**
- POST /api/auth/register - регистрация

```
{
  "email": "test@test.com",
  "password": "12345665",
  "full_name": "user"
}
```

- POST /api/auth/login - вход

```
{
  "email": "test@test.com",
  "password": "12345665"
}
```

**`Системный статус`**
- GET /api/ready

**`Управление пользователями`**
> Проверка на jwt токен
- POST /api/users
- GET /api/users
- GET /api/users/:id
- PUT /api/users/:id
- DELETE /api/users/:id

**`Управление поинтами`**
> Проверка на jwt токен
- PUT /api/points/add - дать очки себе
- GET /api/points/leaderboard - посмотреть на доску лидеров