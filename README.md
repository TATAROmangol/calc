# Calc

В данном проекте реализован калькулятор с приоритетами и скобками, который работает на сервере. Тесты все калькулятора (pkg/calc/calc_test.go), все тесты сервера (pkg/server/server_test.go).

## Работа калькулятора:
Калькулятор принимает строку в которой написано математическое выражение, и считает. При возникновении неполадок выдает ошибки : 
- Неправильный формат
- Отсутвие открывающей или закрывающей скобок
- деление на ноль.

## Работа сервера:

1) Обращение к серверу:
    Сервер принимает только команды **POST**, в виде JSON в котором есть поле "expression" а в нем математическое выражение, по урл localhost/api/v1/calculate.
    
    Пример использования:
    ```bash
    curl --location 'localhost/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
    "expression": "2+2*2"
    }'
    ```
    Ответ:
    ```bash
    {
        "result": 6
    }  
    ```

2) Ответы сервера:
    1. Статус код 200 - все успешно, ответ в графе result:

    Пример запроса:
    ```bash
    curl -i --location 'localhost/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
    "expression": "2+2*2"
    }'
    ```

    Ответ:
    ```bash
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Sat, 21 Dec 2024 11:27:35 GMT
    Content-Length: 18

    {
        "result": 6
    }   
    ```

    2. Статус код 422 - входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций пользователь ввёл символ английского алфавита:

    Пример запроса:
    ```bash
    curl -i --location 'localhost/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
    "expression": "2+2*"
    }'
    ```

    Ответ:
    ```bash
    HTTP/1.1 422 Unprocessable Entity
    Content-Type: application/json
    Date: Sat, 21 Dec 2024 11:32:55 GMT
    Content-Length: 55

    {
        "error": "Неправильный формат"
    }
    ```

    3. Статус код 500 - непедвиденная ошибка, к примеру неправильная форма в запросе:

    Пример запроса:
    ```bash
    curl -i --location 'localhost/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
    "p": "2+2*2"
    }'
    ```

    Ответ:
    ```bash
    HTTP/1.1 500 Internal Server Error
    Content-Type: application/json
    Date: Sat, 21 Dec 2024 11:35:14 GMT
    Content-Length: 53

    {
        "error": "Что-то пошло не так"
    } 
    ```  

    4.Запуск сервера:
    Команда в терминал из папки проекта: 
    ```bash
        go run ./cmd/main.go
    ```


