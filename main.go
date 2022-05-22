package main
import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/go-rod/rod"
)
func main() {
	fmt.Println("HugoPoster, the server.")
	//cmd := exec.Command("echo", "content/post/piday.md") // For Debugging
	cmd := exec.Command("git", "diff", "--name-only", "HEAD", "origin/main")
	cmd.Dir = "/home/jonte/website/blog"
	out, _ := cmd.Output()
	parts := strings.Split(string(out), "/")
	if parts[0] == "content" {
		page := rod.New().MustConnect().MustPage("https://blog.jontes.page/post/" + strings.ToLower(strings.TrimSuffix(parts[2], "\n")))
		page.MustWaitLoad().MustScreenshot("twitter.png")
		title := page.MustElement("body > div > div > div > h1 > a").MustText()
		twittercmd := exec.Command("python3", "twitter.py", "\"" + title + "\"", "\"https://blog.jontes.page/post/" + strings.ToLower(strings.TrimSuffix(parts[2], "\n")) + "\"")
		twittercmd.Run()
		fmt.Println("Posted to Twitter.")
	}
}