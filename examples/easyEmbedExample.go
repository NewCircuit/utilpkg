package utilpkg

import (
	"github.com/NewCircuit/utilpkg/botutil"
)

// main shows an example of making an embed
//  Using the embed utils in botutil/easyEmbed.go
func main() {
	embed := botutil.NewEmbed()
	embed.SetTitle("Title")
	embed.SetDescription("Description")
	embed.SetURL("https://golang.org/")
	embed.AddField("Name", "Value", true)
	embed.SetThumbnail("https://miro.medium.com/max/607/1*vIZzAnSg-8IwmHzZOPppHg.png")
	embed.SetImage("https://www.hardwinsoftware.com/blog/wp-content/uploads/2018/02/golang-gopher.png")
	embed.SetColor(0xF1B379)
	embed.SetAuthor("Name", "https://ohlinger.co/assets/img/golang.jpg",
		"https://www.wikiwand.com/en/Go_(programming_language)")
	embed.SetFooter("https://hackernoon.com/drafts/0fnv29qd.png", "Text")
	embed.Truncate()
}
