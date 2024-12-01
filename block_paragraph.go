package goeditorjs

import (
	"encoding/json"
	"fmt"
)

// ParagraphHandler is the default ParagraphHandler for EditorJS HTML generation
type ParagraphHandler struct{}

func (*ParagraphHandler) parse(editorJSBlock EditorJSBlock) (*paragraph, error) {
	paragraph := &paragraph{}
	return paragraph, json.Unmarshal(editorJSBlock.Data, paragraph)
}

// Type "paragraph"
func (*ParagraphHandler) Type() string {
	return "paragraph"
}

// GenerateHTML generates html for ParagraphBlocks
func (h *ParagraphHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	paragraph, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`<p>%s</p>`, paragraph.Text), nil
}

// GenerateMarkdown generates markdown for ParagraphBlocks
func (h *ParagraphHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	paragraph, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return paragraph.Text, nil
}
