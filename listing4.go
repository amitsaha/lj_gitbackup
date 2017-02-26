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
		1:  Repository{GitURL: "ssh://github.com/amitsaha/gitbackup", Name: "gitbackup"},
		2:  Repository{GitURL: "ssh://github.com/amitsaha/lj_gitbackup", Name: "lj_gitbackup"},
		3:  Repository{GitURL: "ssh://github.com/amitsaha/gitbackup", Name: "gitbackup"},
		4:  Repository{GitURL: "ssh://github.com/amitsaha/lj_gitbackup", Name: "lj_gitbackup"},
		5:  Repository{GitURL: "ssh://github.com/amitsaha/gitbackup", Name: "gitbackup"},
		6:  Repository{GitURL: "ssh://github.com/amitsaha/lj_gitbackup", Name: "lj_gitbackup"},
		7:  Repository{GitURL: "ssh://github.com/amitsaha/gitbackup", Name: "gitbackup"},
		8:  Repository{GitURL: "ssh://github.com/amitsaha/lj_gitbackup", Name: "lj_gitbackup"},
		9:  Repository{GitURL: "ssh://github.com/amitsaha/gitbackup", Name: "gitbackup"},
		10: Repository{GitURL: "ssh://github.com/amitsaha/lj_gitbackup", Name: "lj_gitbackup"},
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
	repositories = append(repositories, getRepo(4))
	repositories = append(repositories, getRepo(5))
	repositories = append(repositories, getRepo(6))
	repositories = append(repositories, getRepo(7))
	repositories = append(repositories, getRepo(8))
	repositories = append(repositories, getRepo(9))
	repositories = append(repositories, getRepo(10))

	// Create a channel of capacity 5
	tokens := make(chan bool, 5)

	for _, r := range repositories {
		if (Repository{}) != r {
			wg.Add(1)
			// Get a token
			tokens <- true
			go func(r Repository) {
				backUp(&r, &wg)
				// release the token
				<-tokens
			}(r)
		}
	}
}
