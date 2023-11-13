package util

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
)

type Config struct {
	Cookie string `json:"cookie"`
}

func GetConfig() (*Config, error) {
	var config Config
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%v/.config/aoc23.json", homedir)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &config)
	return &config, nil
}

func GetCookie(config *Config) (*http.Cookie, error) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: config.Cookie,
	}

	return cookie, nil
}

func GetClient() (*http.Client, error) {
	client := &http.Client{
		Jar:       &cookiejar.Jar{},
		Transport: &http.Transport{},
	}
	return client, nil
}

func FetchInput() ([]string, error) {
	data := make([]string, 25)
	for day := 0; day < 25; day++ {
		url := fmt.Sprintf("https://adventofcode.com/2022/day/%v/input", day)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		client, err := GetClient()
		config, err := GetConfig()
		if err != nil {
			fmt.Println(err)
		}
		cookie, err := GetCookie(config)
		req.AddCookie(cookie)
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		data = append(data, string(body))
	}
	return data, nil
}
