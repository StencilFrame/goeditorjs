package goeditorjs

import (
	"encoding/json"
	"strings"
)

// Table
type TableHandler struct{}

func (*TableHandler) parse(editorJSBlock EditorJSBlock) (*table, error) {
	table := &table{}
	return table, json.Unmarshal(editorJSBlock.Data, table)
}

// Type "table"
func (*TableHandler) Type() string {
	return "table"
}

// GenerateHTML generates html for TableBlocks
func (h *TableHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	table, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	sb.WriteString("<table>\n")
	for _, row := range table.Content {
		sb.WriteString("<tr>\n")
		for _, cell := range row {
			sb.WriteString("<td>" + cell + "</td>\n")
		}
		sb.WriteString("</tr>\n")
	}
	sb.WriteString("</table>")
	return sb.String(), nil
}

// GenerateMarkdown generates markdown for TableBlocks
func (h *TableHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	table, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	for i, row := range table.Content {
		sb.WriteString("| " + strings.Join(row, " | ") + " |\n")
		if i == 0 {
			separator := "|"
			for range row {
				separator += " --- |"
			}
			sb.WriteString(separator + "\n")
		}
	}
	return sb.String(), nil
}
