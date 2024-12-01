package goeditorjs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ImageHandler is the default ImageHandler for EditorJS HTML generation
type ImageHandler struct {
	// Options are made available to the GenerateHTML and GenerateMarkdown functions.
	// If not provided, DefaultImageHandlerOptions will be used.
	Options *ImageHandlerOptions
}

// ImageHandlerOptions are the options available to the ImageHandler
type ImageHandlerOptions struct {
	BorderClass     string
	StretchClass    string
	BackgroundClass string
}

// DefaultImageHandlerOptions are the default options available to the ImageHandler
var DefaultImageHandlerOptions = &ImageHandlerOptions{
	StretchClass:    "image-tool--stretched",
	BorderClass:     "image-tool--withBorder",
	BackgroundClass: "image-tool--withBackground"}

func (*ImageHandler) parse(editorJSBlock EditorJSBlock) (*image, error) {
	image := &image{}
	return image, json.Unmarshal(editorJSBlock.Data, image)
}

// Type "image"
func (*ImageHandler) Type() string {
	return "image"
}

// GenerateHTML generates html for ImageBlocks
func (h *ImageHandler) GenerateHTML(editorJSBlock EditorJSBlock) (string, error) {
	image, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	return h.generateHTML(image)
}

// GenerateMarkdown generates markdown for ImageBlocks
func (h *ImageHandler) GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error) {
	image, err := h.parse(editorJSBlock)
	if err != nil {
		return "", err
	}

	if image.Stretched || image.WithBackground || image.WithBorder {
		return h.generateHTML(image)
	}
	return fmt.Sprintf(`![alt text](%s "%s")`, image.File.URL, image.Caption), nil

}

func (h *ImageHandler) generateHTML(image *image) (string, error) {
	if h.Options == nil {
		h.Options = DefaultImageHandlerOptions
	}

	classes := []string{}
	if image.Stretched {
		classes = append(classes, h.Options.StretchClass)
	}

	if image.WithBorder {
		classes = append(classes, h.Options.BorderClass)
	}

	if image.WithBackground {
		classes = append(classes, h.Options.BackgroundClass)
	}

	class := ""
	if len(classes) > 0 {
		class = fmt.Sprintf(`class="%s"`, strings.Join(classes, " "))
	}

	return fmt.Sprintf(`<img src="%s" alt="%s" %s/>`, image.File.URL, image.Caption, class), nil
}
