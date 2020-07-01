package utilpkg

// main shows an example of making an embed
//  Using the embed utils in utils/easyEmbed.go
func main() {
	//Create a new embed object
	embed := NewEmbed()
	embed.SetTitle("Title")
	embed.SetDescription("Description")
	embed.SetURL("URL")
	embed.AddField("Name", "Value", true)
	embed.SetThumbnail("URL")
	embed.SetImage("URL")
	embed.SetColor("#F1B379")
	embed.SetAuthor("Name", "URL", "IconURL")
	embed.SetFooter("Text", "IconURL")
	embed.Truncate()
	embed.SendToWebhook("Webhook URL")
}
