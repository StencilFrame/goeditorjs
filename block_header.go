package goeditorjs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// HeaderHandler is the default HeaderHandler for EditorJS HTML generation
type HeaderHandler struct{}

func (*HeaderHandler) parse(editorJSBlock EditorJSBlock) (*header, error) {
	header := &header{}
	return header, json.Unmarshal(editorJSBlock.Data, header)
}

// Type "header"
func (*HeaderHandler) Type() string {
	return "header"
}

// GenerateHTML generates html for HeaderBlocks
func (h *HeaderHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	header, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<h%d>%s</h%d>", header.Level, header.Text, header.Level), nil
}

// GenerateMarkdown generates markdown for HeaderBlocks
func (h *HeaderHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	header, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", strings.Repeat("#", header.Level), header.Text), nil
}
