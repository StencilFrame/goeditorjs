package goeditorjs

import "encoding/json"

// RawHTMLHandler is the default raw handler for EditorJS HTML generation
type RawHTMLHandler struct{}

// Type "raw"
func (*RawHTMLHandler) Type() string {
	return "raw"
}

// GenerateHTML generates html for rawBlocks
func (h *RawHTMLHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	return h.raw(editorJSBlock)

}

// GenerateMarkdown generates markdown for rawBlocks
func (h *RawHTMLHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	return h.raw(editorJSBlock)
}

func (h *RawHTMLHandler) raw(editorJSBlock EditorJSBlock) (string, error) {
	raw := &raw{}
	err := json.Unmarshal(editorJSBlock.Data, raw)
	if err != nil {
		return "", err
	}

	return raw.HTML, nil
}
