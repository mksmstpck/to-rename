package translate

import (
	"context"
	"errors"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type Translate struct {
	lang language.Tag
	c    *translate.Client
}

func (t *Translate) UaToEng(text string) (string, error) {
	resp, err := t.c.Translate(context.Background(), []string{text}, t.lang, nil)
	if err != nil {
		return "", err
	}
	if len(resp) == 0 {
		return "", errors.New("Something went wrong")
	}

	return resp[0].Text, nil
}
