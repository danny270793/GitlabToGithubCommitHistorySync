package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	ID          int64       `json:"id"`
	ProjectID   int         `json:"project_id"`
	ActionName  string      `json:"action_name"`
	TargetID    interface{} `json:"target_id"`
	TargetIid   interface{} `json:"target_iid"`
	TargetType  interface{} `json:"target_type"`
	AuthorID    int         `json:"author_id"`
	TargetTitle interface{} `json:"target_title"`
	CreatedAt   time.Time   `json:"created_at"`
	Author      struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		Locked    bool   `json:"locked"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	Imported     bool   `json:"imported"`
	ImportedFrom string `json:"imported_from"`
	PushData     struct {
		CommitCount int         `json:"commit_count"`
		Action      string      `json:"action"`
		RefType     string      `json:"ref_type"`
		CommitFrom  string      `json:"commit_from"`
		CommitTo    string      `json:"commit_to"`
		Ref         string      `json:"ref"`
		CommitTitle string      `json:"commit_title"`
		RefCount    interface{} `json:"ref_count"`
	} `json:"push_data"`
	AuthorUsername string `json:"author_username"`
}

func GetCommitsByDate(userId string, token string, after string, before string) (map[string]int, int, error) {
	events, err := GetAllEvents(userId, token, after, before)
	if err != nil {
		return nil, 0, err
	}

	allGitlabCommits := 0
	gitlabEventsByDate := make(map[string]int)
	for _, event := range events {
		date := event.CreatedAt.Format("2006-01-02")
		value, exists := gitlabEventsByDate[date]
		if !exists {
			gitlabEventsByDate[date] = event.PushData.CommitCount
		} else {
			gitlabEventsByDate[date] = value + event.PushData.CommitCount
		}
		allGitlabCommits += event.PushData.CommitCount
	}
	return gitlabEventsByDate, allGitlabCommits, nil
}

func GetAllEvents(userId string, token string, after string, before string) ([]Event, error) {
	allEvents := make([]Event, 0)
	currentPage := 1

	hasNext := true
	for hasNext {
		events, nextPage, err := GetEvents(userId, token, currentPage, after, before)
		if err != nil {
			return nil, err
		}
		if nextPage == -1 {
			hasNext = false
		}
		allEvents = append(allEvents, events...)
		currentPage = nextPage
	}
	return allEvents, nil
}

func GetEvents(userId string, token string, page int, after string, before string) ([]Event, int, error) {
	url := fmt.Sprintf("https://gitlab.com/api/v4/users/%s/events?action=pushed&page=%d&after=%s&before=%s", userId, page, after, before)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, -1, err
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}

	var events []Event

	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, -1, err
	}

	nextPage := resp.Header.Get("x-next-page")
	if nextPage == "" {
		return events, -1, nil
	}

	num, err := strconv.Atoi(nextPage)
	if err != nil {
		return nil, -1, err
	}

	return events, num, nil
}
