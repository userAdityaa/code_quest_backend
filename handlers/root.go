package handlers

import (
	"fmt"
	"net/http"

	"github.com/userAdityaa/code_quest_backend/utils"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<a href="/login/github/">LOGIN</a>`)
}

func LoggedinHandler(w http.ResponseWriter, r *http.Request, githubData string) {
	if githubData == "" {
		fmt.Fprintf(w, "UNAUTHORIZED!")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, utils.PrettifyJSON(githubData))
}
