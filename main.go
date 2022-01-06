package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// rm -rf /path/to/directory/*      - удаление всех файлов/директорий/поддиректорий (не удалит скрытые файлы/директории/поддиректории)
// rm -rf /path/to/directory/{*,.*} - удаление всех файлов/директорий/поддиректорий и скрытых файлов/директорий/поддиректорий

const (
	shellToUse    = "zsh"
	defaultConfig = `
{
  "commands": [
	"cd /",
    "sudo find . -name .DS_Store -type f -delete",
    "sudo rm $HOME/.viminfo",
    "sudo rm $HOME/.NERDTreeBookmarks",
    "sudo rm -r $HOME/.cache",
    "sudo rm $HOME/.zsh_history",
    "sudo rm -r $HOME/.zsh_sessions/",
    "sudo rm -rf $HOME/Library/Caches/*",
    "sudo qlmanage -r cache"
  ]
}
`
)

type config struct {
	Commands []string `json:"commands"`
}

func main() {
	path := fmt.Sprintf("%s/.config/cleansys", os.Getenv("HOME"))
	if err := createConfigDirIfNotExists(path); err != nil {
		log.Fatalln(err)
	}

	filePath := fmt.Sprintf("%s/config.json", path)
	file, err := createConfigFileIfNotExists(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if err = setDefaultConfig(file); err != nil {
		log.Fatalln(err)
	}

	configFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	var cfg config
	if err = json.Unmarshal(configFile, &cfg); err != nil {
		log.Fatalln(err)
	}

	execute(cfg.Commands)
}

func createConfigDirIfNotExists(path string) error {
	if err := os.MkdirAll(path, 0777); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}

	return nil
}

func createConfigFileIfNotExists(name string) (*os.File, error) {
	file, err := os.Create(name)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			return nil, err
		}
	}

	return file, nil
}

func setDefaultConfig(file *os.File) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		if _, err = file.Write([]byte(defaultConfig)); err != nil {
			return err
		}
	}

	return nil
}

func execute(commands []string) {
	for _, command := range commands {
		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd := exec.Command(shellToUse, "-c", command)
		cmd.Stdin = os.Stdin
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			log.Printf("stdout: %s, stderr: %s", stdout.String(), stderr.String())
		}
	}
}
