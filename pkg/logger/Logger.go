package logger

import (
    "github.com/fatih/color"
    "io"
    "log"
    "os"
    "time"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
    TrafficLogger *log.Logger
)

func getDate() string {
    dt := time.Now()
    return dt.Format("2006-01-02")
}

func InitLogger() {
    if _, err := os.Stat("logs/"); os.IsNotExist(err) {
        os.Mkdir("logs/", 0762)
    }
    file, err := os.OpenFile("logs/logs-"+getDate()+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    multi := io.MultiWriter(file, color.Output)

    WarningLogger = log.New(multi, color.YellowString("WARNING: "), log.Ldate|log.Ltime|log.Lshortfile)
    InfoLogger = log.New(multi, color.BlueString("INFO: "), log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(multi, color.RedString("ERROR: "), log.Ldate|log.Ltime|log.Lshortfile)
    TrafficLogger = log.New(multi, color.GreenString("TRAFFIC: "), log.Ldate|log.Ltime|log.Lshortfile)
}