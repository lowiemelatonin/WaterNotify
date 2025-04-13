package main

import (
	"path/filepath"

	"github.com/go-toast/toast"
)

func main() {
	iconPath, err := filepath.Abs("icon.ico")
	if err != nil {
		panic(err)
	}

	notification := toast.Notification{
		AppID:   "waternotify",
		Title:   "Did you drink some water already?",
		Message: "Your kidneys would be thankfull, and your entire body too!",
		Icon:    iconPath,
	}

	err = notification.Push()
	if err != nil {
		panic(err)
	}
}
