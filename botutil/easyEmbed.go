package botutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

// Everything here is mainly stolen from https://gist.github.com/Necroforger/8b0b70b1a69fa7828b8ad6387ebb3835
// Then changed to actually function properly.
// Reason: I was writing this, searched for constants and discovered someone already invented the wheel.

// Embed structure
type Embed struct {
	*dg.MessageEmbed
}

// Constants for message embed character limits
const (
	EmbedLimitTitle       = 256
	EmbedLimitDescription = 2048
	EmbedLimitFieldValue  = 1024
	EmbedLimitFieldName   = 256
	EmbedLimitField       = 25
	EmbedLimitFooter      = 2048
	EmbedLimit            = 4000
)

// NewEmbed returns a new embed object
func NewEmbed() *Embed {
	return &Embed{&dg.MessageEmbed{}}
}

// SetTitle takes <Title string>
// Sets embed with title set <Title>
func (e *Embed) SetTitle(Title string) {
	if len(Title) > EmbedLimitTitle {
		Title = Title[:EmbedLimitTitle]
	}
	e.Title = Title
}

// SetDescription takes <Description string>
// Sets embed with description set <Description>
func (e *Embed) SetDescription(Description string) {
	if len(Description) > EmbedLimitDescription {
		Description = Description[:EmbedLimitDescription]
	}
	e.Description = Description
}

// AddField takes <Name string> <Value string> <Inline bool>
// Returns an embed with embed field set <Name> <Value> <Inline>
func (e *Embed) AddField(Name string, Value string, Inline bool) {
	if len(Value) > EmbedLimitFieldValue {
		Value = Value[:EmbedLimitFieldValue]
	}

	if len(Name) > EmbedLimitFieldName {
		Name = Name[:EmbedLimitFieldName]
	}

	EmbedField := dg.MessageEmbedField{
		Name:   Name,
		Value:  Value,
		Inline: Inline,
	}

	e.Fields = append(e.Fields, &EmbedField)
}

// SetFooter takes <IconURL string> <Text string>
// Sets embed with embed footer set <Iconurl> <Text>
func (e *Embed) SetFooter(IconURL string, Text string) {
	if len(Text) > EmbedLimitFooter {
		Text = Text[:EmbedLimitFooter]
	}

	e.Footer = &dg.MessageEmbedFooter{
		IconURL: IconURL,
		Text:    Text,
	}
}

// SetImage takes <URL string>
// Sets embed image <URL>
func (e *Embed) SetImage(URL string) {
	e.Image = &dg.MessageEmbedImage{
		URL: URL,
	}
}

// SetThumbnail takes <URL string>
// Sets embed thumbnail to <URL>
func (e *Embed) SetThumbnail(URL string) {
	e.Thumbnail = &dg.MessageEmbedThumbnail{
		URL: URL,
	}
}

// SetAuthor takes <Name string> <IconURL string> <URL string>
// Returns embed with author set <Name> <IconURL> <URL>
func (e *Embed) SetAuthor(Name string, IconURL string, URL string) {
	e.Author = &dg.MessageEmbedAuthor{
		Name:    Name,
		IconURL: IconURL,
		URL:     URL,
	}
}

// SetURL takes <URL string>
// Sets embed url <URL>
func (e *Embed) SetURL(URL string) {
	e.URL = URL
}

// SetColor takes <Color string>
// Sets color of embed to <Color>
// Returns error
func (e *Embed) SetColor(Color string) error {
	Color = strings.Replace(Color, "0x", "", -1)
	Color = strings.Replace(Color, "0X", "", -1)
	Color = strings.Replace(Color, "#", "", -1)
	ColorInt, err := strconv.Atoi(Color)
	if err != nil {
		return err
	}
	e.Color = ColorInt
	return nil
}

// InlineAllFields sets all fields in the embed to be inline
func (e *Embed) InlineAllFields() *Embed {
	for _, v := range e.Fields {
		v.Inline = true
	}
	return e
}

// Truncate truncates the number of embed fields over the character limit.
// Rest of truncation is done on function call
func (e *Embed) Truncate() *Embed {
	e.TruncateFields()
	return e
}

// TruncateFields truncates fields that are too long
func (e *Embed) TruncateFields() *Embed {
	if len(e.Fields) > 25 {
		e.Fields = e.Fields[:EmbedLimitField]
	}
	return e
}

// SendToWebhook takes <Webhook string>
// Sends embed to webhook
// Returns error if invalid embed or error posting to webhook
func (e *Embed) SendToWebhook(Webhook string) error {
	embedArray := append(make([]*dg.MessageEmbed, 0), e.MessageEmbed)
	params := dg.WebhookParams{
		Embeds: embedArray,
	}

	embedJson, err := json.Marshal(params)
	if err != nil {
		return err
	}

	_, err = http.Post(Webhook, "application/json", bytes.NewBuffer(embedJson))
	if err != nil {
		return err
	}

	return nil
}
