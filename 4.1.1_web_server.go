package main

import (
	"fmt"
	"net/http"
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RawQuery: ", r.URL.String())           // URL с параметрами
	fmt.Println("Name: ", r.URL.Query().Get("name"))    // значение параметра
	fmt.Println("IsExist: ", r.URL.Query().Has("name")) // существует ли такой параметр
	w.Write([]byte("Привет, Stepik!"))
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/", handler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
