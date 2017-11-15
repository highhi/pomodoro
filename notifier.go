package pomodoro

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Content.
type Content struct {
	Title   string
	Message string
}

const (
	cmdTerminalNotifier = "terminal-notifier"
	cmdWhich            = "which"
	cmdSwVers           = "sw_vers"
	optTitle            = "-title"
	optMessage          = "-message"
	optProductVers      = "-productVersion"
)

func checkTerminalNotifier() bool {
	if err := exec.Command(cmdWhich, cmdTerminalNotifier).Run(); err != nil {
		return false
	}

	return true
}

func checkOsxVersion() bool {
	result, _ := exec.Command(cmdSwVers, optProductVers).Output()
	version := strings.Split(strings.TrimSpace(string(result)), ".")
	major, _ := strconv.Atoi(version[0])
	minor, _ := strconv.Atoi(version[1])

	if major == 10 && minor > 9 {
		return true
	}

	return false
}

func push(title string, message string) {
	exec.Command(cmdTerminalNotifier, optTitle, title, optMessage, message).Run()
}

// Notifier.
func Notifier(w Content, b Content) {
	if !checkTerminalNotifier() {
		return
	}

	if !checkOsxVersion() {
		return
	}

	for {
		push(w.Title, w.Message)
		time.Sleep(25 * time.Minute)
		push(b.Title, b.Message)
		time.Sleep(10 * time.Minute)
	}
}

// Example.

// import(
// 	"github.com/highhi/pomodoro"
// )

// func main()  {
// 	w := pomodoro.Content{Title: "作業開始", Message: "集中しよう!"}
// 	b := pomodoro.Content{Title: "休憩", Message: "ちょっと一休み"}
// 	pomodoro.Notifier(w, b)
// }
