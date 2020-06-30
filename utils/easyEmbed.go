package main

import (
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

// SetTitle takes <title string>
// Returns an embed with title set <title>
func (e *Embed) SetTitle(title string) *Embed {
	if len(title) > EmbedLimitTitle {
		title = title[:EmbedLimitTitle]
	}
	e.Title = title
	return e
}

// SetDescription takes <description string>
// Returns an embed with description set <description>
func (e *Embed) SetDescription(description string) *Embed {
	if len(description) > EmbedLimitDescription {
		description = description[:EmbedLimitDescription]
	}
	e.Description = description
	return e
}

// AddField takes <name string> <value string>
// Returns an embed with embed field set <name> <value>
func (e *Embed) AddField(name string, value string) *Embed {
	if len(value) > EmbedLimitFieldValue {
		value = value[:EmbedLimitFieldValue]
	}

	if len(name) > EmbedLimitFieldName {
		name = name[:EmbedLimitFieldName]
	}

	e.Fields = append(e.Fields, &dg.MessageEmbedField{
		Name:  name,
		Value: value,
	})
	return e
}

// SetFooter takes <iconurl string> <text string> <proxyurl string>
// Returns embed with embed footer set <iconurl> <text> <proxyurl>
func (e *Embed) SetFooter(iconURL string, text string, proxyURL string) *Embed {
	if len(text) > EmbedLimitFooter {
		text = text[:EmbedLimitFooter]
	}

	e.Footer = &dg.MessageEmbedFooter{
		IconURL:      iconURL,
		Text:         text,
		ProxyIconURL: proxyURL,
	}
	return e
}

// SetImage takes <url string> <proxyurl string>
// Returns embed with embed image set <url> <proxyurl>
func (e *Embed) SetImage(URL string, proxyURL string) *Embed {
	e.Image = &dg.MessageEmbedImage{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

// SetThumbnail takes <url string> <proxyurl string> <height int> <width int>
// Returns embed with embed thumbnail set <url> <proxyurl> <height> <width>
func (e *Embed) SetThumbnail(URL string, proxyURL string, height int, width int) *Embed {
	e.Thumbnail = &dg.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
		Height:   height,
		Width:    width,
	}
	return e
}

// SetAuthor takes <name string> <iconurl string> <url string> <proxyurl string>
// Returns embed with author set <name> <iconurl> <url> <proxyurl>
func (e *Embed) SetAuthor(name string, iconURL string, URL string, proxyURL string) *Embed {
	e.Author = &dg.MessageEmbedAuthor{
		Name:         name,
		IconURL:      iconURL,
		URL:          URL,
		ProxyIconURL: proxyURL,
	}
	return e
}

// SetURL takes <url string>
// Returns embed with url set <url>
func (e *Embed) SetURL(URL string) *Embed {
	e.URL = URL
	return e
}

// SetColor takes <color int>
// Returns embed with color set <color>
func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
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
