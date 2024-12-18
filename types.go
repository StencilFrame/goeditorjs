package goeditorjs

import (
	"encoding/json"
	"errors"
)

// editorJS rpresents the Editor JS data
type editorJS struct {
	Blocks []EditorJSBlock `json:"blocks"`
}

// EditorJSBlock type
type EditorJSBlock struct {
	Type string `json:"type"`
	// Data is the Data for an editorJS block in the form of RawMessage ([]byte). It is left up to the Handler to parse the Data field
	Data json.RawMessage `json:"data"`
}

var (
	//ErrBlockHandlerNotFound is returned from GenerateHTML when the HTML engine doesn't have a registered handler
	//for that type and the HTMLEngine is set to return on errors.
	ErrBlockHandlerNotFound = errors.New("Handler not found for block type")
)

// header represents header data from EditorJS
type header struct {
	Text  string `json:"text"`
	Level int    `json:"level"`
}

// paragraph represents paragraph data from EditorJS
type paragraph struct {
	Text      string `json:"text"`
	Alignment string `json:"alignment"`
}

// ListStyle represents the style of list
type ListStyle string

// listStyle constants
const (
	ListStyleOrdered   ListStyle = "ordered"
	ListStyleUnordered ListStyle = "unordered"
	ListStyleChecklist ListStyle = "checklist"
)

// list represents list data from EditorJS
type list struct {
	Style ListStyle  `json:"style"`
	Meta  listMeta   `json:"meta"`
	Items []listItem `json:"items"`
}

// CounterType represents the type of counter for ordered lists
type CounterType string

// CounterType constants
const (
	CounterTypeNumeric    CounterType = "numeric"
	CounterTypeLowerRoman CounterType = "lower-roman"
	CounterTypeUpperRoman CounterType = "upper-roman"
	CounterTypeLowerAlpha CounterType = "lower-alpha"
	CounterTypeUpperAlpha CounterType = "upper-alpha"
	CounterTypeDefault                = CounterTypeNumeric
)

// counterTypeToHTML maps CounterType to HTML
var counterTypeToHTML = map[CounterType]string{
	CounterTypeNumeric:    "1",
	CounterTypeLowerRoman: "i",
	CounterTypeUpperRoman: "I",
	CounterTypeLowerAlpha: "a",
	CounterTypeUpperAlpha: "A",
}

// listMeta represents list meta data from EditorJS
type listMeta struct {
	Checked     bool        `json:"checked,omitempty"`     // for Checklist
	Start       int         `json:"start,omitempty"`       // for Ordered list
	CounterType CounterType `json:"counterType,omitempty"` // for Ordered list
}

// listItem represents list item data from EditorJS
type listItem struct {
	Content string     `json:"content"`
	Meta    listMeta   `json:"meta"`
	Items   []listItem `json:"items"`
}

// codeBox represents code box data from EditorJS
type codeBox struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

// raw represents raw html data from EditorJS
type raw struct {
	HTML string `json:"html"`
}

// image represents image data from EditorJS
type image struct {
	File           file   `json:"file"`
	Caption        string `json:"caption"`
	WithBorder     bool   `json:"withBorder"`
	WithBackground bool   `json:"withBackground"`
	Stretched      bool   `json:"stretched"`
}

type file struct {
	URL string `json:"url"`
}

type table struct {
	Content [][]string `json:"content"`
}
