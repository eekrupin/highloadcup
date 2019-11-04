package main

import "net/http"

func DumbHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, I'm Web Server!"))
}

func main() {
	api.Run()
	http.HandleFunc("/", DumbHandler)
	http.ListenAndServe(":80", nil)

}

//GET /<entity>/<id> для получения данных о сущности
//GET /users/<id>/visits для получения списка посещений пользователем
//GET /locations/<id>/avg для получения средней оценки достопримечательности
//POST /<entity>/<id> на обновление
//POST /<entity>/new на создание

//docker build -t dumb .
//docker run --rm -p 8080:80 -t dumb
