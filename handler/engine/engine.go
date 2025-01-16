package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/GarimaUttam/Car_management_system/models"
	"github.com/GarimaUttam/Car_management_system/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		service: service,
	}
}

func (e *EngineHandler) GetEngineByID(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := e.service.GetEngineByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("Error Writing Response: ", err)
	}
}

func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		log.Println("Error unmarsalling engine request body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdEngine,err := e.service.CreateEngine(ctx, &engineReq)
	if err != nil {
		log.Println("Error while creating Engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, err := json.Marshal(createdEngine)
	if err != nil {
		log.Println("Error marshalling response body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	_,_ = w.Write(resBody) 
}

func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]


	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		log.Println("Error unmarsalling engine request body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedEngine, err := e.service.UpdateEngine(ctx, id, &engineReq)
	if err != nil {
		log.Println("Error while updating Engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, err := json.Marshal(updatedEngine)
	if err != nil {
		log.Println("Error marshalling response body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	_,_ = w.Write(resBody) 
}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deletedEngine, err := e.service.DeletedEngine(ctx, id)
	if err != nil {
		log.Println("Error while deleting Engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string {"error": "Invalid ID or Engine not Found"}
		jsonResponse, _ := json.Marshal(response)
		_,_ = w.Write(jsonResponse)
		return
	}

	// check if the engine was deleted succesfully
	if deletedEngine.EngineID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Engine not found"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return
	}

	jsonResponse, err := json.Marshal(deletedEngine)
	if err != nil {
		log.Println("Error while marshalling deleted engine response:",err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error":"Internal server error"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}