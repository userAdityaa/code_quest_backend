package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/userAdityaa/code_quest_backend/config"
)

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	clientID := config.GetGithubClientID()
	redirectURI := "http://localhost:8080/login/github/callback/"
	encodedRedirectURI := url.QueryEscape(redirectURI)

	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		clientID,
		encodedRedirectURI,
	)

	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func GithubCallBackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	accessToken := getGithubAccessToken(code)
	githubData, err := getGithubData(accessToken)
	if err != nil {
		http.Error(w, "Failed to get GitHub data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.Redirect(w, r, fmt.Sprintf("http://localhost:3000/questions?username=%s&avatar_url=%s", githubData["login"], githubData["avatar_url"]), http.StatusFound)
}

func getGithubAccessToken(code string) string {
	clientID := config.GetGithubClientID()
	clientSecret := config.GetGithubClientSecret()
	requestBody := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJson, _ := json.Marshal(requestBody)
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJson),
	)
	if err != nil {
		log.Panic("Failed to create request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Request to GitHub failed")
	}
	defer resp.Body.Close()

	var response struct {
		AccessToken string `json:"access_token"`
	}
	json.NewDecoder(resp.Body).Decode(&response)

	return response.AccessToken
}

func getGithubData(accessToken string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Panic("Failed to create request")
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("GitHub data request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s", body)
	}

	var userData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		return nil, err
	}

	return userData, nil
}
