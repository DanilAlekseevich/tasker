tasker/
├── cmd/
│   └── app/
│       └── main.go                 # Точка входа
├── internal/
│   ├── config/
│   │   ├── config.go              # Конфигурация приложения
│   │   └── database.go            # Настройки БД
│   ├── domain/
│   │   ├── task.go                # Доменная модель Task
│   │   ├── user.go                # Доменная модель User  
│   │   └── project.go             # Доменная модель Project
│   ├── service/
│   │   ├── task_service.go        # Бизнес-логика задач
│   │   ├── user_service.go        # Бизнес-логика пользователей
│   │   └── auth_service.go        # Логика аутентификации
│   ├── repository/
│   │   ├── interfaces/
│   │   │   ├── task_repository.go # Интерфейсы репозиториев
│   │   │   └── user_repository.go
│   │   └── postgres/
│   │       ├── task_repository.go # Реализация для PostgreSQL
│   │       └── user_repository.go
│   ├── handler/
│   │   ├── task_handler.go        # HTTP обработчики задач
│   │   ├── user_handler.go        # HTTP обработчики пользователей
│   │   └── auth_handler.go        # Обработчики аутентификации
│   ├── middleware/
│   │   ├── auth.go                # JWT middleware
│   │   ├── cors.go                # CORS middleware  
│   │   ├── logging.go             # Логирование запросов
│   │   └── error.go               # Обработка ошибок
│   └── router/
│       └── router.go              # Настройка маршрутов
├── migrations/
│   ├── 001_create_users_table.up.sql
│   ├── 001_create_users_table.down.sql
│   ├── 002_create_projects_table.up.sql
│   ├── 002_create_projects_table.down.sql
│   ├── 003_create_tasks_table.up.sql
│   └── 003_create_tasks_table.down.sql
├── configs/
│   ├── config.yaml                # Конфигурация для разработки
│   └── config.prod.yaml           # Конфигурация для продакшена
├── api/
│   └── openapi.yaml               # OpenAPI/Swagger спецификация  
├── tests/
│   ├── integration/
│   │   └── task
