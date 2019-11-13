package main

import (
  "fmt"
  "os"
)

const (
  InfoColor    = "\033[1;34m%s\033[0m\n"
  NoticeColor  = "\033[1;36m%s\033[0m\n"
  WarningColor = "\033[1;33m%s\033[0m\n"
  ErrorColor   = "\033[1;31m%s\033[0m\n"
  DebugColor   = "\033[0;36m%s\033[0m\n"
)

func main() {
  err := start()
  if err != nil {
    fmt.Fprintf(os.Stderr, "error: %v", err)
    os.Exit(1)
  }
}

func start() error {
  if home := homeDir(); home != "" {
    fmt.Printf(NoticeColor,"Found home dir")
  } else {
    fmt.Printf(ErrorColor, "Homedir not found")
  }
    return nil
}

func homeDir() string {
  if h := os.Getenv("HOME"); h != "" {
    return h
  }
  return os.Getenv("USERPROFILE") // windows
}
