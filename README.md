# README.md (Русский)

# Task_API

Это руководство описывает процесс запуска и использования Task_API, простого API для управления задачами. Он позволяет добавлять, просматривать, удалять задачи и отслеживать их статус.

## 🚀 Запуск API

### Предварительные требования
*   **Git**: Для клонирования репозитория.
*   **Go**: Компилятор Go (версия 1.18 или выше) для запуска проекта из исходного кода.

### Шаги по запуску
1.  **Клонируйте репозиторий:**
    ```bash
    git clone https://github.com/ваш_пользователь/ваш_репозиторий.git # Замените на актуальный URL
    ```
2.  **Перейдите в корневую директорию проекта:**
    ```bash
    cd ваш_репозиторий # Замените на имя папки проекта
    ```

### Способы запуска
#### Для Windows (использование исполняемого файла)
Если у вас есть скомпилированный исполняемый файл `app.exe`:
```bash
.\cmd\app\app.exe
```

#### Для любой ОС (с помощью Go компилятора)
Убедитесь, что у вас установлен Go:
```bash
go run cmd/app/main.go
```

## ⚙️ Конфигурация (config.yaml)

Файл `config.yaml` находится в корневой директории проекта. Вы можете изменить порт, на котором будет работать сервер, отредактировав значение `localhost:8080` (например, на `localhost:8081`).

Пример `config.yaml`:
```yaml
server_address: "localhost:8080" # Измените на желаемый адрес/порт
```

При изменении порта, сервер будет слушать соединения на новом порте (например, `localhost:8081` вместо `localhost:8080`).

## 🤝 Взаимодействие с API

Рекомендуется использовать такие инструменты, как [Postman](https://www.postman.com/downloads/), [Insomnia](https://insomnia.rest/download) или `curl` для взаимодействия с API.

**Базовый URL:** `http://127.0.0.1:8080` (или любой другой порт, указанный в `config.yaml`).

### 1. Проверка статуса сервера

*   **URL:** `/`
*   **Метод:** `GET`
*   **Пример запроса:**
    ```
    GET http://127.0.0.1:8080
    ```
*   **Ожидаемый ответ:**
    ```json
    {
        "status": "connect"
    }
    ```

### 2. Добавление новой задачи

*   **URL:** `/tasks/add`
*   **Метод:** `POST`
*   **Тело запроса (JSON):**
    ```json
    {
        "name": "{название_задачи}"
    }
    ```
*   **Пример запроса (Postman/curl):**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"name": "Купить молоко"}' http://127.0.0.1:8080/tasks/add
    ```
*   **Ожидаемые ответы:**
    *   Если задача успешно добавлена:
        ```json
        {
            "status": "add task"
        }
        ```
    *   Если задача с таким именем уже существует:
        ```json
        {
            "status": "task already exists"
        }
        ```

### 3. Получение информации о задаче

*   **URL:** `/tasks/{название_задачи}`
*   **Метод:** `GET`
*   **Пример запроса:**
    ```
    GET http://127.0.0.1:8080/tasks/Купить%20молоко
    ```
*   **Ожидаемые ответы:**
    *   Если задача не найдена:
        ```json
        {
            "status": "task not found"
        }
        ```
    *   Если задача найдена:
        ```json
        {
            "name": "Купить молоко",          # Название задачи
            "status": "in progress",          # Статус задачи: "completed" (выполнена) | "in progress" (в процессе выполнения)
            "lead_time_min": 5,               # Примерное время выполнения задачи (от 3 до 5 минут)
            "end_time": "2025-06-14T15:28:50.1430089+03:00" # Время начала обработки задачи
        }
        ```

### 4. Получение списка всех задач

*   **URL:** `/tasks`
*   **Метод:** `GET`
*   **Пример запроса:**
    ```
    GET http://127.0.0.1:8080/tasks
    ```
*   **Ожидаемый ответ:**
    Список всех добавленных и не удаленных задач. Формат ответа будет массивом JSON-объектов, аналогичных тем, что возвращаются при получении информации о конкретной задаче.
    ```json
    [
        {
            "name": "Купить молоко",
            "status": "in progress",
            "lead_time_min": 5,
            "end_time": "2025-06-14T15:28:50.1430089+03:00"
        },
        {
            "name": "Погулять с собакой",
            "status": "completed",
            "lead_time_min": 3,
            "end_time": "2025-06-14T15:20:00.0000000+03:00"
        }
    ]
    ```

### 5. Удаление задачи

*   **URL:** `/tasks/delete`
*   **Метод:** `POST`
*   **Тело запроса (JSON):**
    ```json
    {
        "name": "{название_задачи}"
    }
    ```
*   **Пример запроса (Postman/curl):**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"name": "Купить молоко"}' http://127.0.0.1:8080/tasks/delete
    ```
*   **Ожидаемые ответы:**
    *   Если задача успешно удалена:
        ```json
        {
            "status": "delete task"
        }
        ```
    *   Если задача не найдена:
        ```json
        {
            "status": "task not found"
        }
        ```

## 🛑 Завершение работы

Для остановки работающего сервера API просто нажмите `Ctrl + C` в консоли, где был запущен сервер.

---

# README.md (English)

# Task_API

This guide describes the process of running and using the Task_API, a simple API for managing tasks. It allows you to add, view, delete tasks, and track their status.

## 🚀 Running the API

### Prerequisites
*   **Git**: For cloning the repository.
*   **Go**: Go compiler (version 1.18 or higher) for running the project from source.

### Launch Steps
1.  **Clone the Repository:**
    ```bash
    git clone https://github.com/your_user/your_repository.git # Replace with the actual URL
    ```
2.  **Navigate to the Project Root Directory:**
    ```bash
    cd your_repository # Replace with your project folder name
    ```

### Ways to Run
#### For Windows (Using the Executable)
If you have a pre-compiled `app.exe` executable:
```bash
.\cmd\app\app.exe
```

#### For Any OS (Using Go Compiler)
Ensure you have Go installed:
```bash
go run cmd/app/main.go
```

## ⚙️ Configuration (config.yaml)

The `config.yaml` file is located in the project's root directory. You can change the port on which the server will run by editing the `localhost:8080` value (e.g., to `localhost:8081`).

Example `config.yaml`:
```yaml
server_address: "localhost:8080" # Change to your desired address/port
```

By changing the port, the server will listen for connections on the new port (e.g., `localhost:8081` instead of `localhost:8080`).

## 🤝 API Usage

It is recommended to use tools like [Postman](https://www.postman.com/downloads/), [Insomnia](https://insomnia.rest/download), or `curl` for interacting with the API.

**Base URL:** `http://127.0.0.1:8080` (or any other port specified in `config.yaml`).

### 1. Server Status Check

*   **URL:** `/`
*   **Method:** `GET`
*   **Example Request:**
    ```
    GET http://127.0.0.1:8080
    ```
*   **Expected Response:**
    ```json
    {
        "status": "connect"
    }
    ```

### 2. Add New Task

*   **URL:** `/tasks/add`
*   **Method:** `POST`
*   **Request Body (JSON):**
    ```json
    {
        "name": "{task_name}"
    }
    ```
*   **Example Request (Postman/curl):**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"name": "Buy milk"}' http://127.0.0.1:8080/tasks/add
    ```
*   **Expected Responses:**
    *   If the task is added successfully:
        ```json
        {
            "status": "add task"
        }
        ```
    *   If a task with the same name already exists:
        ```json
        {
            "status": "task already exists"
        }
        ```

### 3. Get Task Information

*   **URL:** `/tasks/{task_name}`
*   **Method:** `GET`
*   **Example Request:**
    ```
    GET http://127.0.0.1:8080/tasks/Buy%20milk
    ```
*   **Expected Responses:**
    *   If the task is not found:
        ```json
        {
            "status": "task not found"
        }
        ```
    *   If the task is found:
        ```json
        {
            "name": "Buy milk",              // Task name
            "status": "in progress",         // Task status: "completed" | "in progress"
            "lead_time_min": 5,              // Estimated task completion time (3 to 5 minutes)
            "end_time": "2025-06-14T15:28:50.1430089+03:00" // Task processing start time
        }
        ```

### 4. Get All Tasks

*   **URL:** `/tasks`
*   **Method:** `GET`
*   **Example Request:**
    ```
    GET http://127.0.0.1:8080/tasks
    ```
*   **Expected Response:**
    A list of all added and non-deleted tasks. The response format will be an array of JSON objects, similar to those returned when getting information for a specific task.
    ```json
    [
        {
            "name": "Buy milk",
            "status": "in progress",
            "lead_time_min": 5,
            "end_time": "2025-06-14T15:28:50.1430089+03:00"
        },
        {
            "name": "Walk the dog",
            "status": "completed",
            "lead_time_min": 3,
            "end_time": "2025-06-14T15:20:00.0000000+03:00"
        }
    ]
    ```

### 5. Delete Task

*   **URL:** `/tasks/delete`
*   **Method:** `POST`
*   **Request Body (JSON):**
    ```json
    {
        "name": "{task_name}"
    }
    ```
*   **Example Request (Postman/curl):**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"name": "Buy milk"}' http://127.0.0.1:8080/tasks/delete
    ```
*   **Expected Responses:**
    *   If the task is deleted successfully:
        ```json
        {
            "status": "delete task"
        }
        ```
    *   If the task is not found:
        ```json
        {
            "status": "task not found"
        }
        ```

## 🛑 Stopping the API

To stop the running API server, simply press `Ctrl + C` in the console where the server was launched.
