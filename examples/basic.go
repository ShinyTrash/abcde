package main

import (
	"github.com/Shinyhero36/go-toast"
)

func main() {
	notif := toast.Notification{
		AppID: "Example App",
		Title: "A Toast in Go",
	}
	_ = notif.Push()
}
