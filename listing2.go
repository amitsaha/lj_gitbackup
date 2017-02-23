package main

import (
	"log"
)

type Repository struct {
	GitURL string
	Name   string
}

func getRepo(id int) Repository {
	repos := map[int]Repository{
		1: Repository{GitURL: "git+ssh://github.com/amitsaha/gitbackup", Name: "gitbackup"},
		2: Repository{GitURL: "git+ssh://github.com/amitsaha/lj_gitbackup", Name: "lj_gitbackup"},
	}

	return repos[id]
}

func backUp(r *Repository) {
	log.Printf("Backing up %s\n", r.Name)
}

func main() {
	var repositories []Repository
	repositories = append(repositories, getRepo(1))
	repositories = append(repositories, getRepo(2))
	repositories = append(repositories, getRepo(3))

	for _, r := range repositories {
		if (Repository{}) != r {
			backUp(&r)
		}
	}
}
