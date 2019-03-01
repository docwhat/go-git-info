package cli

import (
	"fmt"
	"log"

	"docwhat.org/go-git-info/internal/pkg/gitinfo"
	git "gopkg.in/src-d/go-git.v4"
)

func Run(version string) {
	fmt.Printf("version: %v\n", version)
	gitinfo.RepoState()
	plainOpenOptions := &git.PlainOpenOptions{DetectDotGit: true}
	if err := plainOpenOptions.Validate(); err != nil {
		log.Panicln("Boom!")
	}
	repo, err := git.PlainOpenWithOptions(".", plainOpenOptions)
	if err != nil {
		log.Panicf("Unable to open repository: %v", err)
	}
	fmt.Printf("narf: %v\n", repo)

	worktree, err := repo.Worktree()
	if err != nil {
		log.Panicf("Unable to open worktree: %v", err)
	}

	status, err := worktree.Status()
	if err != nil {
		log.Panicf("Unable to get worktree status: %v", err)
	}

	fmt.Printf("narf2: %v\n", status)

}
