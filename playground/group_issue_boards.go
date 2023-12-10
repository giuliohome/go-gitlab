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
	group := 25247
	board_id := 673

	// Create a new GitLab client with your access token
	gitlabClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	if err != nil {
		log.Fatal(err)
	}

	// Read group issue board
	my_board, _, err := gitlabClient.GroupIssueBoards.GetGroupIssueBoard(group, board_id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Retrieved group issue board: %s\n", my_board.Name)
	fmt.Println("-------------------")

	for _, label := range my_board.Labels {
		fmt.Printf("Scoped Label %s color %s", label.Name, label.Color)
		fmt.Println("-------------------")
	}

}