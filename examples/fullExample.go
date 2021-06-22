package main

import (
	"fmt"
	"github.com/Shinyhero36/go-toast"
)

func main() {
	notif := toast.Notification{
		AppID:       "Example App",
		Title:       "A Toast in Go",
		Icon:        "C:/Users/geral/GolandProjects/toast/sonic.jpg",
		Message:     "Hello World",
		Banner:      "C:/Users/geral/GolandProjects/toast/banner.jpg",
		Image:       "C:/Users/geral/GolandProjects/toast/banner.jpg",
		Signature:   "by GG !!",
		Audio:       toast.Mail,
		Loop:        false,
		Duration:    toast.Short,
		ProgressBar: toast.ProgressBar{"Downloading...", 0.75, "a.png", "75/100 files"},
	}
	err := notif.Push()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
