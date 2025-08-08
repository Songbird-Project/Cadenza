package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

const (
	LogNone = iota - 2
	LogDebug
	LogAll
	LogInfo
	LogWarn
	LogErr
	LogFatal
)

// App struct
type App struct {
	ctx       context.Context
	cfg       Config
	themePath string
	// cfgPath string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	store, err := NewConfigStore()
	if err != nil {
		ConsoleLog(fmt.Sprintf("Failed to initialise config store: %v\n", err), LogErr)
		return
	}

	ConsoleLog(fmt.Sprintf("Config path: %v", store.configPath), LogInfo)
	// a.cfgPath = store.configPath

	cfg, err := store.Config()
	if err != nil {
		ConsoleLog(fmt.Sprintf("Failed to retrieve the config file: %v\n", err), LogErr)
		return
	}

	ConsoleLog(fmt.Sprintf("Config:\n%v", cfg), LogInfo)
	ConsoleLog("Starting Cadenza v0.0.1", LogInfo)
}

func (a *App) Theme() string {
	theme, err := a.GetTheme()
	if err != nil {
		ConsoleLog(fmt.Sprintf("Failed to find theme: %v\n", err), LogErr)
	}

	return theme
}

func ConsoleLog(err string, level int) {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "Cadenza",
	})

	logLevel := 0

	if logLevelStr := os.Getenv("LOG_LEVEL"); logLevelStr != "" {
		if logLevelInt, err := strconv.Atoi(logLevelStr); err == nil {
			logLevel = logLevelInt
		}
	}

	if level < logLevel || logLevel == -2 {
		return
	}

	switch level {
	case -1:
		logger.Debug(fmt.Sprintf("%s", err))
	case 0:
		logger.Print(fmt.Sprintf("%s", err))
	case 1:
		logger.Info(fmt.Sprintf("%s", err))
	case 2:
		logger.Warn(fmt.Sprintf("%s", err))
	case 3:
		logger.Error(fmt.Sprintf("%s", err))
	case 4:
		logger.Fatal(fmt.Sprintf("%s", err))
	}
}
