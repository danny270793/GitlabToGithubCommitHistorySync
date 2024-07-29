package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Commit struct {
	Commit struct {
		Author struct {
			Date time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

func GetCommitsByDate(user string, repo string, accessToken string, since string, until string) (map[string]int, int, error) {
	commits, err := GetAllCommits(user, repo, accessToken, since, until)
	if err != nil {
		return nil, 0, err
	}

	allGithubCommits := 0
	githubEventsByDate := make(map[string]int)
	for _, commit := range commits {
		date := commit.Commit.Author.Date.Format("2006-01-02")
		value, exists := githubEventsByDate[date]
		if !exists {
			githubEventsByDate[date] = 1
		} else {
			githubEventsByDate[date] = value + 1
		}
		allGithubCommits += 1
	}
	return githubEventsByDate, allGithubCommits, nil
}

func GetAllCommits(user string, repo string, accessToken string, since string, until string) ([]Commit, error) {
	allCommits := make([]Commit, 0)
	currentPage := 1

	hasNext := true
	for hasNext {
		commits, err := GetCommits(user, repo, accessToken, currentPage, since, until)
		if err != nil {
			return nil, err
		}
		if len(commits) == 0 {
			hasNext = false
		}
		allCommits = append(allCommits, commits...)
		currentPage += 1
	}
	return allCommits, nil
}

func GetCommits(user string, repo string, accessToken string, page int, since string, until string) ([]Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?page=%d&since=%s&until=%s", user, repo, page, since, until)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var commits []Commit
	if err := json.Unmarshal(body, &commits); err != nil {
		return nil, err
	}

	return commits, nil
}
