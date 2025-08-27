package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	lines    bool
	words    bool
	chars    bool
	bytes    bool
	filename string
}

type FileStats struct {
	lines int
	words int
	chars int
}

func main() {
	config := parseFlags()

	if config.filename == "" {
		fmt.Println("Ошибка: укажите имя файла")
		fmt.Println("Использование: go run main.go [флаги] filename.txt")
		os.Exit(1)
	}

	if !config.lines && !config.words && !config.chars && !config.bytes {
		config.lines = true
		config.words = true
		config.chars = true
	}

	err := processFile(config)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}
}

func parseFlags() Config {
	var config Config

	flag.BoolVar(&config.lines, "l", false, "Посчитать строки")
	flag.BoolVar(&config.words, "w", false, "Посчитать слова")
	flag.BoolVar(&config.chars, "m", false, "Посчитать символы")
	flag.BoolVar(&config.bytes, "c", false, "Посчитать байты")

	flag.Parse()

	if flag.NArg() > 0 {
		config.filename = flag.Arg(0)
	}

	return config
}

func processFile(config Config) error {
	file, err := os.Open(config.filename)
	if err != nil {
		return fmt.Errorf("не могу открыть файл: %w", err)
	}
	defer file.Close()

	if config.lines || config.words || config.chars {
		stats, err := countStats(file)
		if err != nil {
			return err
		}

		if config.lines {
			fmt.Printf("Строк: %d\n", stats.lines)
		}
		if config.words {
			fmt.Printf("Слов: %d\n", stats.words)
		}
		if config.chars {
			fmt.Printf("Символов: %d\n", stats.chars)
		}

		file.Seek(0, 0)
	}

	if config.bytes {
		bytes, err := countBytes(file)
		if err != nil {
			return err
		}
		fmt.Printf("Байт: %d\n", bytes)
	}

	return nil
}

func countStats(file *os.File) (FileStats, error) {
	var stats FileStats
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return stats, fmt.Errorf("ошибка чтения: %w", err)
		}

		if line != "" {
			stats.lines++
			stats.words += len(strings.Fields(line))
			stats.chars += len(line)
		}

		if err == io.EOF {
			break
		}
	}

	return stats, nil
}

func countBytes(file *os.File) (int64, error) {
	info, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("не могу получить информацию о файле: %w", err)
	}
	return info.Size(), nil
}
