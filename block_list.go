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

	start := ""
	end := ""
	if list.Style == ListStyleChecklist {
		start = `<ul class="checklist">`
		end = "</ul>"
	} else if list.Style == ListStyleOrdered {
		startAttr := ""
		typeAttr := ""

		if list.Meta.Start > 1 {
			startAttr = fmt.Sprintf(` start="%d"`, list.Meta.Start)
		}

		if list.Meta.CounterType != "" {
			typeAttr = fmt.Sprintf(` type="%s"`, counterTypeToHTML[list.Meta.CounterType])
		}

		start = fmt.Sprintf("<ol%s%s>", startAttr, typeAttr)
		end = "</ol>"
	} else {
		start = "<ul>"
		end = "</ul>"
	}

	innerData := ""
	for _, s := range list.Items {
		if list.Style == ListStyleChecklist && s.Meta.Checked {
			innerData += fmt.Sprintf(`<li class="checked">%s</li>`, s.Content)
			continue
		}
		innerData += fmt.Sprintf("<li>%s</li>", s.Content)
	}

	return start + innerData + end, nil
}

// GenerateMarkdown generates markdown for ListBlocks
func (h *ListHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	list, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	listItemPrefix := "- "

	results := []string{}
	for i, s := range list.Items {
		if list.Style == ListStyleOrdered {
			listItemPrefix = fmt.Sprintf("%d. ", i+1)
		}
		results = append(results, listItemPrefix+s.Content)
	}

	return strings.Join(results, "\n"), nil
}
