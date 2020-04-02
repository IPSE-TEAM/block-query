package logger

import (
    "encoding/json"
    "fmt"
    "github.com/zhangkuns/logger"
)

var Log *logger.LocalLogger

type FileLoggerConfig struct {
    Filename   string `json:"filename"`
    Append     bool   `json:"append"`
    MaxLines   int    `json:"maxlines"`
    MaxSize    int    `json:"maxsize"`
    Daily      bool   `json:"daily"`
    MaxDays    int64  `json:"maxdays"`
    Level      string `json:"level"`
    PermitMask string `json:"permit"`
}

type ConsoleLoggerConfig struct {
    Level    string `json:"level"`
    Colorful bool   `json:"color"`
}

func NewDefaultFileLoggerCon(path string, level string) *FileLoggerConfig {
    return &FileLoggerConfig{
        Filename:   path,
        Append:     true,
        MaxLines:   100000,
        MaxSize:    10,
        Daily:      true,
        MaxDays:    -3,
        Level:      level,
        PermitMask: "0644",
    }
}

func InitLogWithAdapterConsole(level string, colorful bool) error {
    if Log == nil {
        Log = logger.NewLogger(2)
    }

    consoleLogger := ConsoleLoggerConfig{level, colorful}
    conf, err := json.Marshal(consoleLogger)
    if err != nil {
        return fmt.Errorf("json marshal console logger error: %s", err)
    }

    if err = Log.SetLogger(logger.AdapterConsole, string(conf)); err != nil {
        return fmt.Errorf("set console logger error: %s", err)
    }
    Log.SetLogPathTrim("token-query")
    Log.SetShowFileLineNo(false)

    return nil
}

func InitLogWithAdapterFile(config *FileLoggerConfig) error {
    if Log == nil {
        Log = logger.NewLogger(3)
    }

    conf, err := json.Marshal(config)
    if err != nil {
        return fmt.Errorf("json marshal file logger error: %s", err)
    }

    if err = Log.SetLogger(logger.AdapterFile, string(conf)); err != nil {
        return fmt.Errorf("set file logger error: %s", err)
    }

    return nil
}
