package main

import (
	"bufio"
	"clientgrpc/internal/config"
	"clientgrpc/internal/grpcclient"
	"clientgrpc/internal/storage"
	"log"
	"os"
)

type PathInfo struct {
	BasePath string `json:"base_path"`
	NewPath  string `json:"new_path"`
	JunkPath string `json:"junk_path"`
}

var pathsMap = make(map[string][]string)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	clientService := grpcclient.NewGRPCClientService(cfg)
	clientService.RunGRPCClient()

	err = loadPathsFromFile(cfg.Mailbox_paths_list, cfg)
	if err != nil {
		log.Fatalf("Error loading paths from file: %v", err)
	}

	storage.SavePathsAsJSON(pathsMap)
}

func addPaths(basePath string, cfg *config.Config) {
	newPath := basePath + cfg.New_mail_path
	junkPath := basePath + cfg.Junk_path
	pathsMap[basePath] = []string{newPath, junkPath}
}

func loadPathsFromFile(filePath string, cfg *config.Config) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addPaths(scanner.Text(), cfg)
	}

	return scanner.Err()
}
