package goeditorjs

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// CodeBoxHandler is the default CodeBoxHandler for EditorJS HTML generation
type CodeBoxHandler struct{}

func (*CodeBoxHandler) parse(editorJSBlock EditorJSBlock) (*codeBox, error) {
	codeBox := &codeBox{}
	return codeBox, json.Unmarshal(editorJSBlock.Data, codeBox)
}

// Type "codeBox"
func (*CodeBoxHandler) Type() string {
	return "codeBox"
}

// GenerateHTML generates html for CodeBoxBlocks
func (h *CodeBoxHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	codeBox, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`<pre><code class="%s">%s</code></pre>`, codeBox.Language, codeBox.Code), nil
}

// GenerateMarkdown generates markdown for CodeBoxBlocks
func (h *CodeBoxHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	codeBox, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	codeBox.Code = strings.ReplaceAll(codeBox.Code, "<div>", "\n")
	codeBox.Code = removeHTMLTags(codeBox.Code)

	return fmt.Sprintf("```%s\n%s\n```", codeBox.Language, codeBox.Code), nil
}

func removeHTMLTags(in string) string {
	// regex to match html tag
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)
	// should replace long string first
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			in = strings.ReplaceAll(in, group, "")
		}
	}
	return in
}
