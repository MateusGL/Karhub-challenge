package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"karhub-beer-api/internal/dtos"
	"net/http"
	"net/url"
	"os"
	"time"
)

type spotifyPlaylistResponse struct {
	Name   string `json:"name"`
	Tracks struct {
		Items []struct {
			Track struct {
				Name    string `json:"name"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
}

var (
	spotifyToken    string
	spotifyTokenExp time.Time
)

func GetSpotifyPlaylistForBeer(beerStyle string) (dtos.Playlist, error) {
	token, err := getSpotifyToken()
	if err != nil {
		return dtos.Playlist{}, err
	}

	fmt.Printf("token: %s\n", token)

	playlist, err := searchPlaylist(beerStyle, token)
	if err != nil {
		fmt.Println(err)

		return dtos.Playlist{}, err
	}

	return playlist, nil
}

func getSpotifyToken() (string, error) {

	envToken := os.Getenv("SPOTIFY_ACCESS_TOKEN")
	if envToken != "" {
		spotifyToken = envToken
		spotifyTokenExp = time.Now().Add(1 * time.Hour)
		return spotifyToken, nil
	}

	return "", errors.New("spotify access token not found in environment variables")
}

func searchPlaylist(query, token string) (dtos.Playlist, error) {
	searchURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=playlist&limit=1", url.QueryEscape(query))
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return dtos.Playlist{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return dtos.Playlist{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Spotify search error: %s\n", string(body))
		return dtos.Playlist{}, fmt.Errorf("spotify search error: %s", string(body))
	}

	var searchResult struct {
		Playlists struct {
			Items []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"items"`
		} `json:"playlists"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return dtos.Playlist{}, err
	}

	if len(searchResult.Playlists.Items) == 0 {
		return dtos.Playlist{}, errors.New("no playlist found")
	}

	playlistID := searchResult.Playlists.Items[0].ID
	playlistName := searchResult.Playlists.Items[0].Name

	playlistURL := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", playlistID)
	req, err = http.NewRequest("GET", playlistURL, nil)
	if err != nil {
		return dtos.Playlist{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return dtos.Playlist{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return dtos.Playlist{}, fmt.Errorf("spotify playlist error: %s", string(body))
	}

	var playlistResp spotifyPlaylistResponse
	if err := json.NewDecoder(resp.Body).Decode(&playlistResp); err != nil {
		return dtos.Playlist{}, err
	}

	var tracks []dtos.Track
	for _, item := range playlistResp.Tracks.Items {
		track := item.Track
		if track.Name == "" || len(track.Artists) == 0 {
			continue
		}
		tracks = append(tracks, dtos.Track{
			Name:   track.Name,
			Artist: track.Artists[0].Name,
			Link:   track.ExternalUrls.Spotify,
		})
	}

	return dtos.Playlist{
		Name:   playlistName,
		Tracks: tracks,
	}, nil
}
