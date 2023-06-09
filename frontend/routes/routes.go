package routes

import (
	"github.com/gorilla/mux"
	"../handlers"
	"../templates/disciplinas"
	"../templates/professores"
	"../templates/emprestimos"
	"../templates/home"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", home.Home).Methods("GET")
	router.HandleFunc("/disciplinas", handlers.DisciplinasHandler).Methods("GET")
	router.HandleFunc("/disciplinas", disciplinas.Insert).Methods("POST")
	router.HandleFunc("/visualizar-disciplinas", disciplinas.Read).Methods("GET")

	router.HandleFunc("/professores", handlers.ProfessoresHandler).Methods("GET")
	router.HandleFunc("/professores", professores.Insert).Methods("POST")
	router.HandleFunc("/visualizar-professores", professores.Read).Methods("GET")
	
	router.HandleFunc("/emprestimos", handlers.EmprestimosHandler).Methods("GET")
	router.HandleFunc("/emprestimos", emprestimos.Insert).Methods("POST")
	router.HandleFunc("/visualizar-emprestimos", emprestimos.Read).Methods("GET")
}
