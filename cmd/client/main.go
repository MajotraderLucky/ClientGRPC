package main

import (
	"bufio"
	"clientgrpc/internal/config"
	"clientgrpc/internal/grpcclient"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	savePathsAsJSON()
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

func savePathsAsJSON() {
	for basePath, paths := range pathsMap {
		pathInfo := PathInfo{
			BasePath: basePath,
			NewPath:  paths[0],
			JunkPath: paths[1],
		}
		jsonData, err := json.MarshalIndent(pathInfo, "", "    ")
		if err != nil {
			log.Fatalf("Error marshaling JSON: %v", err)
		}

		// Создание директории, если она не существует
		dir := filepath.Dir("data")
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0755) // создает все необходимые родительские директории
		}

		// Сохранение JSON в файл
		filename := fmt.Sprintf("data/%s.json", "mailbox_struct")
		if err := os.WriteFile(filename, jsonData, 0644); err != nil {
			log.Fatalf("Error writing JSON to file: %v", err)
		}
	}
}
