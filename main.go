package main

import (
	"path/filepath"
	"time"

	"github.com/go-toast/toast"
)

func sendnotification() {
	iconPath, err := filepath.Abs("knt.ico")
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

func main() {
	notified := map[int]bool{
		9:  false,
		12: false,
		15: false,
		18: false,
		21: false,
	}

	for {
		now := time.Now()
		hour := now.Hour()

		if (hour == 9 || hour == 12 || hour == 15 || hour == 18 || hour == 21) && !notified[hour] {
			sendnotification()
			notified[hour] = true
		}

		if hour == 0 {
			for k := range notified {
				notified[k] = false
			}
		}

		time.Sleep(time.Minute)
	}
}
