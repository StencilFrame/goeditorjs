package goeditorjs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ListHandler is the default ListHandler for EditorJS HTML generation
type ListHandler struct{}

func (*ListHandler) parse(editorJSBlock EditorJSBlock) (*list, error) {
	list := &list{}
	return list, json.Unmarshal(editorJSBlock.Data, list)
}

// Type "list"
func (*ListHandler) Type() string {
	return "list"
}

// GenerateHTML generates html for ListBlocks
func (h *ListHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	list, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	result := ""
	if list.Style == "ordered" {
		result = "<ol>%s</ol>"
	} else {
		result = "<ul>%s</ul>"
	}

	innerData := ""
	for _, s := range list.Items {
		innerData += fmt.Sprintf("<li>%s</li>", s.Content)
	}

	return fmt.Sprintf(result, innerData), nil
}

// GenerateMarkdown generates markdown for ListBlocks
func (h *ListHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	list, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	listItemPrefix := ""
	if list.Style == "ordered" {
		listItemPrefix = "1. "
	} else {
		listItemPrefix = "- "
	}

	results := []string{}
	for _, s := range list.Items {
		results = append(results, listItemPrefix+s.Content)
	}

	return strings.Join(results, "\n"), nil
}
