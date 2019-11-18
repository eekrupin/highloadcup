package main

import (
	"github.com/eekrupin/hlc-travels/api"
)

//func DumbHandler(writer http.ResponseWriter, request *http.Request) {
//	writer.Write([]byte("Hello, I'm Web Server!"))
//}

func main() {
	api.Run()
	//http.HandleFunc("/", DumbHandler)
	//http.ListenAndServe(":80", nil)

}

//GET /<entity>/<id> для получения данных о сущности
//GET /users/<id>/visits для получения списка посещений пользователем
//GET /locations/<id>/avg для получения средней оценки достопримечательности
//POST /<entity>/<id> на обновление
//POST /<entity>/new на создание

//docker build -t dumb .
//docker run --rm -p 8080:80 -t dumb

//docker-compose up --build --abort-on-container-exit

//docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -d mysql:5.7
//--CREATE SCHEMA IF NOT EXISTS`travels` DEFAULT CHARACTER SET utf8 ;
