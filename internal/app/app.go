package app

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type App struct {
	client  *http.Client
	baseUrl string
	year    string
}

func InitializeApp() *App {
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar:       jar,
		Transport: &http.Transport{},
	}

	return &App{
		client:  client,
		baseUrl: "https://adventofcode.com/",
		year:    "2022",
	}
}

func (a *App) urlToRemoteInput(day int) string {
	return a.baseUrl + filepath.Join(a.year, "day", strconv.Itoa(day), "input")
}

func (a *App) pathToCachedInput(day int) string {
	homedir, _ := os.UserHomeDir()
	return filepath.Join(homedir, ".cache", "advent-of-code-2023", "inputs", strconv.Itoa(day))
}

func (a *App) pathToCookie() string {
	homedir, _ := os.UserHomeDir()
	return filepath.Join(homedir, ".config", "advent-of-code-2023", "cookie.json")
}

func (a *App) urlToAnswer(day int) string {
	return a.baseUrl + filepath.Join(a.year, "day", strconv.Itoa(day), "answer")
}

func (a *App) LoadCookie() {
	file, _ := os.Open(a.pathToCookie())
	contents, _ := io.ReadAll(file)
	domain, _ := url.Parse(a.baseUrl)

	cookie := &http.Cookie{
		Name:  "session",
		Value: strings.TrimSpace(string(contents)),
	}

	a.client.Jar.SetCookies(domain, []*http.Cookie{cookie})
}

func downloadInputForDay(a *App, day int) {
	req, _ := http.NewRequest("GET", a.urlToRemoteInput(day), nil)
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()

	path := a.pathToCachedInput(day)
	out, err := os.Create(path)
	if err != nil {
		err = os.MkdirAll(filepath.Dir(path), 0750)
		out, _ = os.Create(path)
	}
	defer out.Close()

	io.Copy(out, resp.Body)
}

func (a *App) GetInputForDay(day int) string {
	path := a.pathToCachedInput(day)
	file, err := os.Open(path)
	if err != nil {
		downloadInputForDay(a, day)
		file, _ = os.Open(path)
	}
	bytes, _ := io.ReadAll(file)
	return string(bytes)
}

func (a *App) SubmitAnswer(day int, answer string) {
	url := a.urlToAnswer(day)
	payload := strings.NewReader("level=" + strconv.Itoa(day) + "&answer=" + answer)
	req, _ := http.NewRequest(
		"POST",
		url,
		payload,
	)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, _ := a.client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
