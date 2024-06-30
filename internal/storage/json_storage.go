package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PathInfo struct {
	BasePath string `json:"base_path"`
	NewPath  string `json:"new_path"`
	JunkPath string `json:"junk_path"`
}

// SavePathsAsJSON сериализует данные о путях и сохраняет их в JSON файл
func SavePathsAsJSON(pathsMap map[string][]string) {
	var pathInfos []PathInfo
	for basePath, paths := range pathsMap {
		pathInfo := PathInfo{
			BasePath: basePath,
			NewPath:  paths[0],
			JunkPath: paths[1],
		}
		pathInfos = append(pathInfos, pathInfo)
	}

	jsonData, err := json.MarshalIndent(pathInfos, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Создание директории, если она не существует
	dir := "data" // Указываем нужную директорию
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
	}

	// Сохранение JSON в файл
	filename := fmt.Sprintf("%s/mailbox_struct.json", dir)
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		log.Fatalf("Error writing JSON to file: %v", err)
	}
}
