package main

import (
	"log"
	"sync"
)

type Repository struct {
	GitURL string
	Name   string
}

func getRepo(id int) Repository {

	repos := map[int]Repository{
		1: Repository{GitURL: "ssh://github.com/amitsaha/gitbackup", Name: "gitbackup"},
		2: Repository{GitURL: "ssh://github.com/amitsaha/lj_gitbackup", Name: "lj_gitbackup"},
	}

	return repos[id]
}

func backUp(r *Repository, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Backing up %s\n", r.Name)
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	var repositories []Repository
	repositories = append(repositories, getRepo(1))
	repositories = append(repositories, getRepo(2))
	repositories = append(repositories, getRepo(3))

	for _, r := range repositories {
		if (Repository{}) != r {
			wg.Add(1)
			go func(r Repository) {
				backUp(&r, &wg)
			}(r)
		}
	}
}
