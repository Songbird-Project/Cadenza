package main

import (
	"fmt"

	"github.com/adrg/xdg"
)

func SetTheme() {}

func (a *App) GetTheme() (string, error) {
	configFilePath, err := xdg.ConfigFile(fmt.Sprintf("cadenza/themes/%s/main.css", a.cfg.ThemeName))
	if err != nil {
		return "", fmt.Errorf("Failed to resolve theme path: %w", err)
	}

	return configFilePath, nil
}
