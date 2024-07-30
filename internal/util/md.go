package util

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"strings"
)

var parser = goldmark.New(goldmark.WithExtensions(extension.GFM), goldmark.WithRendererOptions(html.WithHardWraps(), html.WithXHTML()))

func ToMarkDownHtml(s string) (string, error) {

	htmlBuilder := strings.Builder{}

	err := parser.Convert([]byte(s), &htmlBuilder)
	if err != nil {
		return "", err
	}
	htmlStr := htmlBuilder.String()

	return htmlStr, nil

}
