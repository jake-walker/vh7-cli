package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/pterm/pterm"
	quark "github.com/quark-links/quark-cli/quark"
)

type HistoryItem struct {
	Data quark.Response `json:"data"`
	Date time.Time      `json:"date"`
}

const maxHistoryLength int = 20

func getHistoryLocation() string {
	baseConfigDir, err := os.UserConfigDir()
	if err != nil {
		pterm.Fatal.Println("There was a problem fetching the config directory!", err)
	}

	appConfigDir := filepath.Join(baseConfigDir, "/quark-cli")
	err = os.MkdirAll(appConfigDir, os.ModeDir)
	if !(err == nil || os.IsExist(err)) {
		pterm.Fatal.Println("There was a problem creating the config directory!", err)
	}
	return filepath.Join(appConfigDir, "/history.json")
}

func storeHistory(items []HistoryItem) {
	filename := getHistoryLocation()

	if len(items) > maxHistoryLength {
		// Select last X number of items from the list, discarding the rest
		items = items[len(items)-maxHistoryLength:]
	}

	// Convert the list to JSON bytes
	data, err := json.Marshal(items)
	if err != nil {
		pterm.Fatal.Println("There was a problem converting history!", err)
	}

	// Save the bytes to the history.json file
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		pterm.Fatal.Println("There was a problem saving history!", err)
	}
}

func loadHistory() []HistoryItem {
	filename := getHistoryLocation()

	if !fileExists(filename) {
		return []HistoryItem{}
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		pterm.Fatal.Println("There was a problem reading history!", err)
	}

	items := []HistoryItem{}
	err = json.Unmarshal(data, &items)
	if err != nil {
		pterm.Fatal.Println("There was a problem converting history!", err)
	}

	return items
}

func saveToHistory(data quark.Response) {
	item := HistoryItem{
		Date: time.Now(),
		Data: data,
	}

	items := loadHistory()
	items = append(items, item)
	storeHistory(items)
}
