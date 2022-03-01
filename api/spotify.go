package api

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

//Load ID and Secret from .env file and assign to variables;
var _ = godotenv.Load("spotify.env")
var spotify_client_id = os.Getenv("SPOTIFY_CLIENT_ID")
var spotify_secret = os.Getenv("SPOTIFY_CLIENT_SECRET")

func SpotifyLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" {
		http.Error(w, "Expected GET request", http.StatusNotFound)
	} else {
		authScopes := "user-modify-playback-state user-read-playback-state user-read-currently-playing user-read-recently-played user-read-private user-read-email streaming playlist-modify-private playlist-read-private"
		query := r.URL.Query()
		query.Add("response_type", "code")
		query.Add("client_id", spotify_client_id)
		query.Add("redirect_uri", "http://localhost:3222/auth/callback")
		query.Add("scope", authScopes)
		http.Redirect(w, r, "https://accounts.spotify.com/authorize/?"+query.Encode(), http.StatusSeeOther)
		return
	}
}

func SpotifyCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CALLBACK")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	stringToEncode := spotify_client_id + ":" + spotify_secret
	auth64 := base64.StdEncoding.EncodeToString([]byte(stringToEncode))
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", r.FormValue("code"))
	data.Set("redirect_uri", "http://localhost:3222/auth/callback")
	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))

	req.Header = http.Header{
		"Authorization": []string{"Basic " + auth64},
		"Content-Type":  []string{"application/x-www-form-urlencoded"},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Status:", resp.Status)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
	// w.Write(body)
	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}
