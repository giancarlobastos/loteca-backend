package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/giancarlobastos/loteca-backend/domain"
	"github.com/giancarlobastos/loteca-backend/security"
	"github.com/giancarlobastos/loteca-backend/service"
	"github.com/gorilla/mux"
)

type Router struct {
	apiService               *service.ApiService
	updateService            *service.UpdateService
	authenticationMiddleware *security.AuthenticationMiddleware
}

func NewRouter(as *service.ApiService, us *service.UpdateService) *Router {
	amw := security.NewAuthenticationMiddleware(as)

	return &Router{
		apiService:               as,
		updateService:            us,
		authenticationMiddleware: amw,
	}
}

func (router *Router) Start(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/lotteries/current", router.getCurrentLottery).Methods("GET")
	r.HandleFunc("/lotteries/{number}", router.getLottery).Methods("GET")
	r.HandleFunc("/live/{lotteryId}", router.getLiveScores).Methods("GET")
	r.HandleFunc("/matches/{matchId}", router.getMatchDetails).Methods("GET")
	r.HandleFunc("/poll/{lotteryId}", router.getPollResults).Methods("GET")
	r.HandleFunc("/poll/{lotteryId}", router.vote).Methods("POST")
	r.HandleFunc("/login", router.login).Methods("POST")
	r.HandleFunc("/manager/{country}/teams", router.getTeams).Methods("GET")
	r.HandleFunc("/manager/{country}/teams", router.importTeams).Methods("POST")
	r.HandleFunc("/manager/{country}/competitions/{year}", router.getCompetitions).Methods("GET")
	r.HandleFunc("/manager/{country}/competitions/{competitionId}/{year}/matches", router.getMatches).Methods("GET")
	r.HandleFunc("/manager/{country}/competitions", router.importCompetitions).Methods("POST")
	r.HandleFunc("/manager/{country}/competitions/{competitionId}/{year}/matches", router.importMatches).Methods("POST")
	r.HandleFunc("/manager/lotteries", router.createLottery).Methods("POST")
	r.HandleFunc("/manager/odds/{matchId}", router.importOdds).Methods("POST")
	r.Use(router.authenticationMiddleware.Middleware)
	log.Fatal(http.ListenAndServe(addr, r))
}

func (router *Router) login(w http.ResponseWriter, r *http.Request) {
	defer handleErrors(w, r)

	token := r.Header.Get("Token")

	extendedToken, err := router.apiService.Login(token)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid token")
		return
	}
	defer r.Body.Close()

	respondWithJSON(w, http.StatusOK, extendedToken)
}

func (router *Router) createLottery(w http.ResponseWriter, r *http.Request) {
	defer handleErrors(w, r)

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
	defer handleErrors(w, r)

	vars := mux.Vars(r)
	number, _ := strconv.Atoi(vars["number"])

	lottery, err := router.apiService.GetLottery(number)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, lottery)
}

func (router *Router) getLiveScores(w http.ResponseWriter, r *http.Request) {
	defer handleErrors(w, r)

	vars := mux.Vars(r)
	lotteryId, _ := strconv.Atoi(vars["lotteryId"])

	scores, err := router.apiService.GetLiveScores(lotteryId)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, scores)
}

func (router *Router) getMatchDetails(w http.ResponseWriter, r *http.Request) {
	defer handleErrors(w, r)

	vars := mux.Vars(r)
	matchId, _ := strconv.Atoi(vars["matchId"])

	matchDetails, err := router.apiService.GetMatchDetails(matchId)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, matchDetails)
}

func (router *Router) vote(w http.ResponseWriter, r *http.Request) {
	defer handleErrors(w, r)

	var poll domain.Poll
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&poll); err != nil {
		log.Printf("Error [vote]: %v", err)
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	user := r.Context().Value("user").(domain.User)
	err := router.apiService.Vote(poll, user)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "cannot vote twice")
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func (router *Router) getPollResults(w http.ResponseWriter, r *http.Request) {
	defer handleErrors(w, r)

	vars := mux.Vars(r)
	lotteryId, _ := strconv.Atoi(vars["lotteryId"])

	results, err := router.apiService.GetPollResults(lotteryId)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, results)
}

func (router *Router) getCurrentLottery(w http.ResponseWriter, r *http.Request) {
	defer handleErrors(w, r)

	lottery, err := router.apiService.GetCurrentLottery()

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, lottery)
}

func (router *Router) getTeams(w http.ResponseWriter, r *http.Request) {
	defer handleErrors(w, r)

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
	defer handleErrors(w, r)

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
	defer handleErrors(w, r)

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
	defer handleErrors(w, r)

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
	defer handleErrors(w, r)

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
	defer handleErrors(w, r)

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
	defer handleErrors(w, r)

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

func handleErrors(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		log.Println("[Router] panic occurred:", err)
		http.NotFound(w, r)
	}
}
