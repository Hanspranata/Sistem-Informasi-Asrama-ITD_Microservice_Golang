package main

import (
	"log"

	"github.com/username/student-service/internal/http"
	"github.com/username/student-service/internal/student/repository"
	"github.com/username/student-service/internal/student/usecase"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	studentRepo := repository.NewStudentRepository(db)
	studentUsecase := usecase.NewStudentUsecase(studentRepo)
	handler := http.NewStudentHandler(studentUsecase)

	http.RunServer(handler)
}
