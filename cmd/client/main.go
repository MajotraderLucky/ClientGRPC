package main

import (
	"bufio"
	"clientgrpc/internal/config"
	"clientgrpc/internal/grpcclient"
	"fmt"
	"log"
	"os"
)

// Глобальная карта для хранения путей
var pathsMap = make(map[string][]string)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	clientService := grpcclient.NewGRPCClientService(cfg)
	clientService.RunGRPCClient()

	// Загрузка путей из файла
	err = loadPathsFromFile(cfg.Mailbox_paths_list)
	if err != nil {
		log.Fatalf("Error loading paths from file: %v", err)
	}

	// Пример использования полученных путей
	// Проходим по всем путям, которые были загружены из файла
	for basePath, paths := range pathsMap {
		fmt.Printf("Base Path: %s\n", basePath)
		fmt.Println("New Path:", paths[0])  // Путь к новым сообщениям
		fmt.Println("Junk Path:", paths[1]) // Путь к спаму
	}
}

// Функция для добавления новых путей в карту
func addPaths(basePath string) {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	newPath := basePath + cfg.New_mail_path
	junkPath := basePath + cfg.Junk_path
	pathsMap[basePath] = []string{newPath, junkPath}
}

// Функция для загрузки путей из файла и добавления в карту
func loadPathsFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addPaths(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
