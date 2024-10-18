package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

// Визначаємо структуру User
type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Server - структура для HTTP-сервера
type Server struct{}

// Обробник реєстрації користувача
func (s *Server) handleUserRegistration() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Перевіряємо, що метод запиту - POST
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        // Створюємо канал для передачі користувача
        userChan := make(chan User)

        // Декодуємо JSON-дані з тіла запиту
        var newUser User
        if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        // Створюємо горутину для обробки реєстрації користувача
        go func(u User) {
            time.Sleep(2 * time.Second) // Імітація тривалої операції (напр., запис у базу даних)
            userChan <- u                // Передаємо користувача через канал після обробки
        }(newUser)

        // Очікуємо на результат від горутини
        registeredUser := <-userChan

        // Відправляємо відповідь з даними зареєстрованого користувача
        w.Header().Set("Content-Type", "application/json")
        response := fmt.Sprintf(`{"message": "User %s has been successfully registered", "email": "%s"}`, registeredUser.Username, registeredUser.Email)
        w.Write([]byte(response))
    }
}

func main() {
    server := &Server{}

    // Реєструємо обробник на маршрут /register
    http.HandleFunc("/register", server.handleUserRegistration())

    fmt.Println("Server is listening on port 8080")
    http.ListenAndServe(":8080", nil)
}

