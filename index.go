package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"danny270793.github.com/commithistorysync/command"
	"danny270793.github.com/commithistorysync/env"
	"danny270793.github.com/commithistorysync/github"
	"danny270793.github.com/commithistorysync/gitlab"
)

func getYesterday() string {
	now := time.Now()

	yesterday := now.Add(-24 * time.Hour)

	return yesterday.Format("2006-01-02")
}

func getToday() string {
	now := time.Now()

	return now.Format("2006-01-02")
}

func main() {
	err := env.LoadFromFile("./.env")
	if err != nil {
		panic(err)
	}

	today := os.Getenv("SYNC_START_DATE")
	yesterday := os.Getenv("SYNC_END_DATE")

	log.Printf("getting commits since %s until %s", yesterday, today)

	gitlabEventsByDate, allGitlabCommits, err := gitlab.GetCommitsByDate(os.Getenv("GITLAB_USERID"), os.Getenv("GITLAB_ACCESS_TOKEN"), yesterday, today)
	if err != nil {
		panic(err)
	}
	log.Printf("found %d gitlab commits", allGitlabCommits)
	for key, value := range gitlabEventsByDate {
		log.Printf("\t%s %d", key, value)
	}

	githubEventsByDate, allGithubCommits, err := github.GetCommitsByDate(os.Getenv("GITHUB_USERNAME"), os.Getenv("GITHUB_REPOSITORY"), os.Getenv("GITHUB_ACCESS_TOKEN"), yesterday+"T00:00:00Z", today+"T23:59:59Z")
	if err != nil {
		panic(err)
	}
	log.Printf("found %d github commits", allGithubCommits)
	for key, value := range githubEventsByDate {
		log.Printf("\t%s %d", key, value)
	}

	var commitsToSyncByDate = make(map[string]int)

	for gitlabDate, gitlabCommit := range gitlabEventsByDate {
		commitsOnGithub := 0
		for githubDate, githubCommit := range githubEventsByDate {
			if gitlabDate == githubDate {
				commitsOnGithub = githubCommit
				log.Printf("on date %v, %d gitlab commits was found vs %d github commits\n", gitlabDate, gitlabCommit, githubCommit)
			}
		}
		differences := gitlabCommit - commitsOnGithub
		if differences > 0 {
			commitsToSyncByDate[gitlabDate] = differences
		}
	}

	totalCommits := 0
	for _, commits := range commitsToSyncByDate {
		totalCommits += commits
	}

	doneCommits := 0
	for date, commits := range commitsToSyncByDate {
		percentage := 100 * doneCommits / totalCommits

		for i := 0; i < commits; i++ {
			log.Printf("%d of %d (%d) will create %d commits at %s\n", doneCommits, totalCommits, percentage, commits, date)
			_, err := command.Run(".", fmt.Sprintf("sh -x index.sh %s", date), false)
			if err != nil {
				panic(err)
			}
			doneCommits += 1
		}
	}
}
