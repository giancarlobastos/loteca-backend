package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/giancarlobastos/loteca-backend/service"
	"github.com/gorilla/mux"
)

type Router struct {
	apiService    *service.ApiService
	updateService *service.UpdateService
}

func NewRouter(as *service.ApiService, us *service.UpdateService) *Router {
	return &Router{
		apiService:    as,
		updateService: us,
	}
}

func (router *Router) Start(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/lotteries/current", router.getCurrentLottery).Methods("GET")
	r.HandleFunc("/manager/{country}/teams", router.getTeams).Methods("GET")
	r.HandleFunc("/manager/{country}/teams", router.importTeams).Methods("POST")
	r.HandleFunc("/manager/{country}/competitions/{year}", router.getCompetitions).Methods("GET")
	r.HandleFunc("/manager/{country}/competitions/{competitionId}/{year}/matches", router.getMatches).Methods("GET")
	r.HandleFunc("/manager/{country}/competitions", router.importCompetitions).Methods("POST")
	r.HandleFunc("/manager/{country}/{year}/matches", router.importMatches).Methods("POST")
	log.Fatal(http.ListenAndServe(addr, r))
}

func (router *Router) getCurrentLottery(w http.ResponseWriter, r *http.Request) {
	lottery, err := router.apiService.GetCurrentLottery()

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, lottery)
}

func (router *Router) getTeams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["country"]
	teams, err := router.updateService.GetTeams(country)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, teams)
}

func (router *Router) importTeams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["country"]
	err := router.updateService.ImportTeams(country)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func (router *Router) getCompetitions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["country"]
	year, _ := strconv.ParseUint(vars["year"], 10, 32)
	competitions, err := router.updateService.GetCompetitions(country, uint(year))

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, competitions)
}

func (router *Router) importCompetitions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["country"]
	err := router.updateService.ImportCompetitions(country)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func (router *Router) getMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competitionId, _ := strconv.ParseUint(vars["competitionId"], 10, 32)
	year, _ := strconv.ParseUint(vars["year"], 10, 32)
	matches, err := router.updateService.GetMatches(uint32(competitionId), uint(year))

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, matches)
}

func (router *Router) importMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["country"]
	year, _ := strconv.ParseUint(vars["year"], 10, 32)
	err := router.updateService.ImportMatches(country, uint(year), false)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
