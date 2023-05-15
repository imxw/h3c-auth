package notify

import (
	"log"
	"runtime"

	"git.sr.ht/~jackmordaunt/go-toast"
	gosxnotifier "github.com/deckarep/gosx-notifier"
)

const appIcon = "gopher.png"

type Notification struct {
	Message string
	Title   string
}

func NewNotification(message string) *Notification {
	return &Notification{Message: message}
}

func (n *Notification) Push() error {

	if n.Title == "" {
		n.Title = "H3C认证"
	}
	switch runtime.GOOS {
	case "windows":
		note := toast.Notification{
			AppID: "H3CAUTH",
			Title: n.Title,
			Body:  n.Message,
		}
		return note.Push()
	case "darwin":
		note := gosxnotifier.NewNotification(n.Message)
		note.Title = n.Title
		note.Sound = gosxnotifier.Default
		note.Sender = "com.apple.Safari"
		note.AppIcon = appIcon

		return note.Push()
	default:
		log.Fatalf("暂不支持%s", runtime.GOOS)

	}

	return nil
}
