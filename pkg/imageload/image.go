// build js,wasm

package imageload

import (
	"errors"

	"github.com/GodsBoss/gggg/dom"
)

type Loader interface {
	Load(imageURL string) (*dom.Image, error)
}

func NewLoader(doc *dom.Document) Loader {
	return &loader{
		doc: doc,
	}
}

type loader struct {
	doc *dom.Document
}

func (l *loader) Load(imageURL string) (*dom.Image, error) {
	img, err := l.doc.CreateImage(imageURL)
	if err != nil {
		return nil, err
	}
	okChan := make(chan struct{})
	errChan := make(chan interface{})

	img.On(
		func() {
			close(okChan)
		},
		func(err interface{}) {
			errChan <- err
		},
	)

	select {
	case <-okChan:
		close(errChan)
		return img, nil
	case <-errChan:
		close(okChan)
		return nil, errors.New("could not load image")
	}
}
