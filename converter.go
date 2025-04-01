package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Mappings struct {
	Pairs []Pair `json:"mappings"`
}

type Pair struct {
	Windows string   `json:"windows"`
	Linux   []string `json:"linux"`
}

func main() {
	fmt.Println("Starting converter...")
	pwd, _ := os.Getwd()
	windowsDir := "Unzipped"
	fullPathToWindowsDir := filepath.Join(pwd, windowsDir)
	fmt.Println("Looking in:", fullPathToWindowsDir)

	file, err := os.ReadFile("map.json")
	if err != nil {
		fmt.Println("Error reading map.json:", err)
		return
	}
	var mapped Mappings
	if err := json.Unmarshal(file, &mapped); err != nil {
		fmt.Println("Error unmarshaling map.json:", err)
		return
	}
	fmt.Println("Loaded map.json with", len(mapped.Pairs), "mappings")

	convertDir := "Converted"
	fullPathToConvertedDir := filepath.Join(pwd, convertDir)
	if err := os.MkdirAll(fullPathToConvertedDir, 0700); err != nil {
		fmt.Println("Error creating Converted dir:", err)
		return
	}

	files, err := os.ReadDir(fullPathToWindowsDir)
	if err != nil {
		fmt.Println("Error reading Unzipped dir:", err)
		return
	}
	fmt.Println("Found", len(files), "items in Unzipped")

	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			fmt.Println("Skipping directory:", fileName)
			continue
		}
		if strings.HasPrefix(fileName, ".") {
			fmt.Println("Skipping hidden file:", fileName)
			continue
		}
		if filepath.Ext(fileName) != ".ani" {
			fmt.Println("Skipping non-.ani file:", fileName)
			continue
		}
		inputPath := filepath.Join(fullPathToWindowsDir, fileName)
		baseName := strings.TrimSuffix(fileName, ".ani")
		outputPath := filepath.Join(fullPathToConvertedDir, baseName+".xcur")
		tempOutput := filepath.Join(fullPathToConvertedDir, baseName)

		fmt.Println("Converting:", inputPath, "to:", outputPath)
		cmd := exec.Command("win2xcur", inputPath, "-o", fullPathToConvertedDir)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error converting", fileName, ":", err)
			fmt.Println("win2xcur output:", string(output))
			continue
		}

		if err := os.Rename(tempOutput, outputPath); err != nil {
			fmt.Println("Error renaming", tempOutput, "to", outputPath, ":", err)
			continue
		}
		fmt.Println("Successfully converted", fileName, "to", outputPath)
	}

	fmt.Println("Done!")
}
