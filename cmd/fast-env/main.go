package main

import (
	"fastenv/internal/builder"
	"fastenv/internal/hasher"
	"fastenv/internal/linker"
	"fastenv/internal/store"
	"fmt"
	"log"
)

func main() {

	manfiestPath := "requirements.txt"

	fmt.Println("Hashing requirements.txt...")
	hash, err := hasher.HashFile(manfiestPath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Manifest Hash: %s \n", hash)
	s, err := store.NewStore()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	targetPath := s.GetEnvPath(hash)
	exist, err := s.Exist(hash)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if exist {
		fmt.Println("Cache Hit! Resuing existing enviroment.")
	} else {
		fmt.Println("Cache Miss! Building new venv...")
		b, err := builder.NewBuilder()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		if err := b.BuildVenv(targetPath); err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Println("Venv creation complete.")
	}

	l := linker.NewLinker(".venv")
	if err := l.LinkSymlink(targetPath); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Successfully linked .ven -> %s\n", targetPath)
	fmt.Println("My new favorite.")
}
