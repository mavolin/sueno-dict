package root

// Code generated by github.com/mavolin/corgi v0.5.6. DO NOT EDIT.

import (
	_bytes "bytes"
	"fmt"
	_io "io"
	"net/url"

	_writeutil "github.com/mavolin/corgi/pkg/writeutil"
	"github.com/mavolin/sueno-dict/pkl/sueno"
	"github.com/mavolin/sueno-dict/repository"
)

func RenderTranslationSearch(_w _io.Writer, query string, results []repository.TranslatedWord) (err error) {
	var (
		_buf    _bytes.Buffer
		_closed bool
		// in case we never use them
		_ = _buf
		_ = _closed
	)

	_closed = false
	err = _writeutil.Write(_w, "<!doctype html><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta content=\"ie=edge\" http-equiv=\"x-ua-compatible\"><link rel=\"stylesheet\" href=\"/static/css/main.css\"><title")
	if err != nil {
		return err
	}
	if !_closed {
		_closed = true
		err = _writeutil.Write(_w, ">")
		if err != nil {
			return err
		}
	}
	err = _writeutil.Write(_w, "Search for ")
	if err != nil {
		return err
	}
	err = _writeutil.WriteHTML(_w, query)
	if err != nil {
		return err
	}
	_closed = false
	err = _writeutil.Write(_w, "</title></head><body")
	if err != nil {
		return err
	}
	if !_closed {
		_closed = true
		err = _writeutil.Write(_w, ">")
		if err != nil {
			return err
		}
	}
	_closed = false
	err = _writeutil.Write(_w, "<div class=\"l-container\"><div class=\"l-page-header\"><a href=\"/\" class=\"c-home-link\">The Sueno Dictionary</a><form action=\"/\" method=\"get\" class=\"c-page-search\"><input type=\"text\" name=\"q\" required aria-label=\"word\" placeholder=\"skribi\"")
	if err != nil {
		return err
	}
	err = _writeutil.WriteAttr(_w, "value", query, false)
	if err != nil {
		return err
	}
	_closed = true
	err = _writeutil.Write(_w, " class=\"c-page-search__input\"><button type=\"submit\" value=\"Search\" class=\"c-page-search__button\"><svg fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" class=\"c-page-search__button-icon\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z\"></path></svg></button></form></div><ol class=\"l-search-results\">")
	if err != nil {
		return err
	}
	for _, result := range results {
		_closed = false
		err = _writeutil.Write(_w, "<li class=\"c-search-result\"><a")
		if err != nil {
			return err
		}
		err = _writeutil.WriteAttr(_w, "href", fmt.Sprintf("/?q=%s&type=word", url.QueryEscape(result.Word.Word)), false)
		if err != nil {
			return err
		}
		_closed = true
		err = _writeutil.Write(_w, " class=\"c-search-result__link\">")
		if err != nil {
			return err
		}
		err = _writeutil.WriteHTML(_w, result.Word.Word)
		if err != nil {
			return err
		}
		_closed = true
		err = _writeutil.Write(_w, "</a><span class=\"c-search-result__type\"> (")
		if err != nil {
			return err
		}
		switch result.MatchedDefinition.Type {
		case sueno.Noun:
			err = _writeutil.Write(_w, "noun")
			if err != nil {
				return err
			}
		case sueno.Verb:
			err = _writeutil.Write(_w, "verb")
			if err != nil {
				return err
			}
		case sueno.Adjective:
			err = _writeutil.Write(_w, "adjective")
			if err != nil {
				return err
			}
		case sueno.Adverb:
			err = _writeutil.Write(_w, "adverb")
			if err != nil {
				return err
			}
		case sueno.Preposition:
			err = _writeutil.Write(_w, "preposition")
			if err != nil {
				return err
			}
		case sueno.Conjunction:
			err = _writeutil.Write(_w, "conjunction")
			if err != nil {
				return err
			}
		case sueno.Interjection:
			err = _writeutil.Write(_w, "interjection")
			if err != nil {
				return err
			}
		case sueno.Pronoun:
			err = _writeutil.Write(_w, "pronoun")
			if err != nil {
				return err
			}
		case sueno.Article:
			err = _writeutil.Write(_w, "article")
			if err != nil {
				return err
			}
		case sueno.Cardinal:
			err = _writeutil.Write(_w, "number")
			if err != nil {
				return err
			}
		case sueno.Abbreviation:
			err = _writeutil.Write(_w, "abbreviation")
			if err != nil {
				return err
			}
		case sueno.Other:
			err = _writeutil.Write(_w, "other")
			if err != nil {
				return err
			}
		}
		_closed = true
		err = _writeutil.Write(_w, ")</span><span class=\"c-search-result__translation\">&mdash;")
		if err != nil {
			return err
		}
		err = _writeutil.WriteHTML(_w, result.MatchedDefinition.Translation)
		if err != nil {
			return err
		}
		err = _writeutil.Write(_w, "</span></li>")
		if err != nil {
			return err
		}
	}
	err = _writeutil.Write(_w, "</ol></div></body>")
	if err != nil {
		return err
	}
	return nil
}
