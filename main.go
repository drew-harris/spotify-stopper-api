package main

import (
	"context"
	"fmt"
	handler "github.com/drew-harris/spotify-stopper-api/api"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

var ctx = context.Background()

var client *spotify.Client

func main() {
	// Set up redis
	godotenv.Load()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	auth := spotifyauth.New(spotifyauth.WithRedirectURL(os.Getenv("CALLBACK")),
		spotifyauth.WithScopes("user-read-currently-playing", "user-read-playback-state", "user-modify-playback-state"))

	url := auth.AuthURL("")

	tokenString := rdb.Get(ctx, "access_token").Val()
	refreshString := rdb.Get(ctx, "refresh_token").Val()

	// Create oauth 2 token
	token := &oauth2.Token{
		AccessToken:  tokenString,
		RefreshToken: refreshString,
	}

	client = spotify.New(auth.Client(ctx, token))

	spotifyHandleCallback := func(w http.ResponseWriter, r *http.Request) {
		tok, err := auth.Token(ctx, "", r)

		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		fmt.Fprintf(w, "Access token: %s", tok.AccessToken)

		rdb.Set(ctx, "access_token", tok.AccessToken, 0)
		rdb.Set(ctx, "refresh_token", tok.RefreshToken, 0)

		client = spotify.New(auth.Client(ctx, tok))
	}

	pause := func(w http.ResponseWriter, r *http.Request) {
		err := client.Pause(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		// Set status 200
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Paused")
	}

	fmt.Println(url)

	http.HandleFunc("/", handler.Test)
	http.HandleFunc("/callback", spotifyHandleCallback)
	http.HandleFunc("/pause", pause)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
