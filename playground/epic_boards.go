package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

func main() {
	baseURL := os.Getenv("GITLAB_BASE_URL")
	fmt.Printf("GitLab Base URL Name: %s\n", baseURL)
	fmt.Println("-------------------")
	token := os.Getenv("GITLAB_ACCESS_TOKEN")
	group := 24946 // 545
	epic_id := 13

	// Create a new GitLab client with your access token
	gitlabClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	if err != nil {
		log.Fatal(err)
	}

	// Create new epic board
	my_board, _, err := gitlabClient.GroupEpicBoards.GetGroupEpicBoard(group, epic_id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Retrieved epic board: %s\n", my_board.Name)
	fmt.Println("-------------------")

	boards, _, err := gitlabClient.GroupEpicBoards.ListGroupEpicBoards(group, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print project names and their respective URLs
	for _, board := range boards {
		fmt.Printf("Board Name: %s\n", board.Name)
		fmt.Println("-------------------")
	}
}
