package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/giancarlobastos/loteca-backend/domain"
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
	r.HandleFunc("/authenticate", router.authenticate).Methods("POST")
	r.HandleFunc("/lotteries/current", router.getCurrentLottery).Methods("GET")
	r.HandleFunc("/lotteries/{number}", router.getLottery).Methods("GET")
	r.HandleFunc("/manager/{country}/teams", router.getTeams).Methods("GET")
	r.HandleFunc("/manager/{country}/teams", router.importTeams).Methods("POST")
	r.HandleFunc("/manager/{country}/competitions/{year}", router.getCompetitions).Methods("GET")
	r.HandleFunc("/manager/{country}/competitions/{competitionId}/{year}/matches", router.getMatches).Methods("GET")
	r.HandleFunc("/manager/{country}/competitions", router.importCompetitions).Methods("POST")
	r.HandleFunc("/manager/{country}/competitions/{competitionId}/{year}/matches", router.importMatches).Methods("POST")
	r.HandleFunc("/manager/lotteries", router.createLottery).Methods("POST")
	r.HandleFunc("/manager/odds/{matchId}", router.importOdds).Methods("POST")
	log.Fatal(http.ListenAndServe(addr, r))
}

func (router *Router) authenticate(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	token := r.Header.Get("token")
	authenticatedUser, err := router.apiService.Authenticate(user, token)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, authenticatedUser)
}

func (router *Router) createLottery(w http.ResponseWriter, r *http.Request) {
	var lottery domain.Lottery
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&lottery); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	l, err := router.apiService.CreateLottery(lottery)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	router.respondCreated(w, r, fmt.Sprintf("/lotteries/%d", l.Id))
}

func (router *Router) getLottery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, _ := strconv.Atoi(vars["number"])

	lottery, err := router.apiService.GetLottery(number)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, lottery)
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
	year, _ := strconv.Atoi(vars["year"])
	competitions, err := router.updateService.GetCompetitions(country, year)

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
	competitionId, _ := strconv.Atoi(vars["competitionId"])
	year, _ := strconv.Atoi(vars["year"])
	matches, err := router.updateService.GetMatches(competitionId, year)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, matches)
}

func (router *Router) importMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competitionId, _ := strconv.Atoi(vars["competitionId"])
	year, _ := strconv.Atoi(vars["year"])
	err := router.updateService.ImportMatches(competitionId, year)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func (router *Router) importOdds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchId, _ := strconv.Atoi(vars["matchId"])
	err := router.updateService.ImportOdds(matchId)

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

func (router *Router) respondCreated(w http.ResponseWriter, r *http.Request, path string) {
	w.Header().Set("Path", r.Host+path)
	w.WriteHeader(http.StatusCreated)
}
