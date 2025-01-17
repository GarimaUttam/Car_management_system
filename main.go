package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GarimaUttam/Car_management_system/driver"
	carStore "github.com/GarimaUttam/Car_management_system/store/car"
	engineStore "github.com/GarimaUttam/Car_management_system/store/engine"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	carService "github.com/GarimaUttam/Car_management_system/service/car"
	engineService "github.com/GarimaUttam/Car_management_system/service/engine"

	carHandler "github.com/GarimaUttam/Car_management_system/handler/car"
	engineHandler "github.com/GarimaUttam/Car_management_system/handler/engine"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	carStore := carStore.New(db)
	carService := carService.NewCarService(carStore)

	engineStore := engineStore.New(db)
	engineService := engineService.NewEngineService(engineStore)

	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	router := mux.NewRouter()

	schemaFile := "store/schema.sql"
	if err := executeSchemaFile(db, schemaFile);err != nil {
		log.Fatal("Error while executing the schema file :", err)
	}

	router.HandleFunc("/cars/{id}", carHandler.GetCarByID).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarByID).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engine/{id}", engineHandler.GetEngineByID).Methods("GET")
	router.HandleFunc("/engine", engineHandler.GetEngineByID).Methods("POST")
	router.HandleFunc("/engine/{id}", engineHandler.GetEngineByID).Methods("PUT")
	router.HandleFunc("/engine/{id}", engineHandler.GetEngineByID).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}

func executeSchemaFile(db *sql.DB, fileName string) error {
	sqlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}