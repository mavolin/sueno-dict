package root

// Code generated by github.com/mavolin/corgi v0.5.6. DO NOT EDIT.

import (
	_bytes "bytes"
	"fmt"
	_io "io"
	"net/url"

	_writeutil "github.com/mavolin/corgi/pkg/writeutil"
	"github.com/mavolin/sueno-dict/pkg/sueno"
	"github.com/mavolin/sueno-dict/repository"
)

func RenderEntry(_w _io.Writer, word repository.Word, otherRootWords []repository.Word) (err error) {
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
	err = _writeutil.WriteHTML(_w, word.Word)
	if err != nil {
		return err
	}
	_closed = false
	err = _writeutil.Write(_w, "&mdash;Definition</title></head><body")
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
	err = _writeutil.Write(_w, "<form style=\"display: none\" method=\"post\"")
	if err != nil {
		return err
	}
	err = _writeutil.WriteAttr(_w, "action", fmt.Sprintf("/%d/delete", word.ID), false)
	if err != nil {
		return err
	}
	_closed = true
	{
		_closed = false
		err = _writeutil.Write(_w, " class=\"js-delete-form\"></form><div class=\"l-container\"><div class=\"l-page-header\"><a href=\"/\" class=\"c-home-link\"><h2>The Sueno Dictionary</h2></a><form action=\"/\" method=\"get\" class=\"c-page-search\"><input type=\"text\" name=\"q\" required aria-label=\"word\" placeholder=\"skribi\" class=\"c-page-search__input\"><button type=\"submit\" value=\"Search\" class=\"c-page-search__button\"><svg fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" class=\"c-page-search__button-icon\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z\"></path></svg></button></form></div><div class=\"c-entry-section\"><div")
		if err != nil {
			return err
		}
		var i int
		var prevType sueno.WordType
		_closed = true
		err = _writeutil.Write(_w, " class=\"l-entries\">")
		if err != nil {
			return err
		}
		for i < len(word.Definitions) {
			{
				_closed = false
				err = _writeutil.Write(_w, "<div")
				if err != nil {
					return err
				}
				firstDefinition := word.Definitions[i]
				_closed = true
				err = _writeutil.Write(_w, " class=\"c-entry\"><div class=\"c-entry__header\"><h2 class=\"c-entry__title\">")
				if err != nil {
					return err
				}
				err = _writeutil.WriteHTML(_w, word.Word)
				if err != nil {
					return err
				}
				_closed = false
				err = _writeutil.Write(_w, "</h2><div")
				if err != nil {
					return err
				}
				switch firstDefinition.Type {
				case sueno.Noun:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "noun")
					if err != nil {
						return err
					}
				case sueno.Verb:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "verb")
					if err != nil {
						return err
					}
				case sueno.Adjective:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "adjective")
					if err != nil {
						return err
					}
				case sueno.Adverb:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "adverb")
					if err != nil {
						return err
					}
				case sueno.Preposition:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "preposition")
					if err != nil {
						return err
					}
				case sueno.Conjunction:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "conjunction")
					if err != nil {
						return err
					}
				case sueno.Interjection:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "interjection")
					if err != nil {
						return err
					}
				case sueno.Pronoun:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "pronoun")
					if err != nil {
						return err
					}
				case sueno.Article:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "article")
					if err != nil {
						return err
					}
				case sueno.Cardinal:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "number")
					if err != nil {
						return err
					}
				case sueno.Abbreviation:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "abbreviation")
					if err != nil {
						return err
					}
				case sueno.Other:
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-entry__type\">")
						if err != nil {
							return err
						}
					}
					err = _writeutil.Write(_w, "other")
					if err != nil {
						return err
					}
				}
				if !_closed {
					_closed = true
					err = _writeutil.Write(_w, " class=\"c-entry__type\">")
					if err != nil {
						return err
					}
				}
				err = _writeutil.Write(_w, "</div>")
				if err != nil {
					return err
				}
				if len(word.CompoundWords) > 0 {
					_closed = true
					err = _writeutil.Write(_w, "<div class=\"c-entry__compound-words\">")
					if err != nil {
						return err
					}
					for i, cw := range word.CompoundWords {
						_closed = false
						err = _writeutil.Write(_w, "<div class=\"c-entry__compound-word\"><a")
						if err != nil {
							return err
						}
						err = _writeutil.WriteAttr(_w, "href", "/?q="+url.QueryEscape(cw.Word)+"&type=word", false)
						if err != nil {
							return err
						}
						_closed = true
						err = _writeutil.Write(_w, ">")
						if err != nil {
							return err
						}
						err = _writeutil.WriteHTML(_w, cw.Word)
						if err != nil {
							return err
						}
						err = _writeutil.Write(_w, "</a></div>")
						if err != nil {
							return err
						}
						if i > 0 {
							err = _writeutil.Write(_w, " + ")
							if err != nil {
								return err
							}
						}
					}
					err = _writeutil.Write(_w, "</div>")
					if err != nil {
						return err
					}
				}
				err = _writeutil.Write(_w, "</div>")
				if err != nil {
					return err
				}
				switch firstDefinition.Type {
				case sueno.Noun:
					_closed = true
					err = _writeutil.Write(_w, "<div class=\"c-entry-section\"><div class=\"c-entry-section__body\"><table class=\"c-entry-fields-table\"><tr class=\"c-entry-field\"><td class=\"c-entry-field__label\">Plural:</td><td class=\"c-entry-field__value\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPluralNoun(word.Word))
					if err != nil {
						return err
					}
					err = _writeutil.Write(_w, "</td></tr></table></div></div>")
					if err != nil {
						return err
					}
				case sueno.Adjective:
					_closed = true
					err = _writeutil.Write(_w, "<div class=\"c-entry-section\"><div class=\"c-entry-section__body\"><table class=\"c-entry-fields-table\"><tr class=\"c-entry-field\"><td class=\"c-entry-field__label\">Comparative:</td><td class=\"c-entry-field__value\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToComparativeAdjective(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr class=\"c-entry-field\"><td class=\"c-entry-field__label\">Superlative:</td><td class=\"c-entry-field__value\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToSuperlativeAdjective(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr></table></div></div><div class=\"c-entry-section\"><h3 class=\"c-entry-section__title\">Adverb</h3><div class=\"c-entry-section__body\"><table class=\"c-entry-fields-table\"><tr class=\"c-entry-field\"><td class=\"c-entry-field__label\">Base:</td><td class=\"c-entry-field__value\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToBaseAdverb(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr class=\"c-entry-field\"><td class=\"c-entry-field__label\">Comparative:</td><td class=\"c-entry-field__value\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToComparativeAdverb(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr class=\"c-entry-field\"><td class=\"c-entry-field__label\">Superlative:</td><td class=\"c-entry-field__value\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToSuperlativeAdverb(word.Word))
					if err != nil {
						return err
					}
					err = _writeutil.Write(_w, "</td></tr></table></div></div>")
					if err != nil {
						return err
					}
				case sueno.Cardinal:
					_closed = true
					err = _writeutil.Write(_w, "<div class=\"c-entry-section\"><div class=\"c-entry-section__body\"><table class=\"c-entry-fields-table\"><tr class=\"c-entry-field\"><td class=\"c-entry-field__label\">Ordinal:</td><td class=\"c-entry-field__value\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToOrdinal(word.Word))
					if err != nil {
						return err
					}
					err = _writeutil.Write(_w, "</td></tr></table></div></div>")
					if err != nil {
						return err
					}
				}
				_closed = true
				err = _writeutil.Write(_w, "<div class=\"c-entry-section\"><h3 class=\"c-entry-section__title\">Definition of <em>")
				if err != nil {
					return err
				}
				err = _writeutil.WriteHTML(_w, word.Word)
				if err != nil {
					return err
				}
				_closed = true
				{
					_closed = false
					err = _writeutil.Write(_w, "</em></h3><div class=\"c-entry-section__body\"><ol")
					if err != nil {
						return err
					}
					prevType = firstDefinition.Type
					if !_closed {
						_closed = true
						err = _writeutil.Write(_w, " class=\"c-defintions\">")
						if err != nil {
							return err
						}
					}
					for i < len(word.Definitions) {
						def := word.Definitions[i]
						if prevType != def.Type {
							break
						}
						_closed = true
						err = _writeutil.Write(_w, "<li class=\"c-definition\"><div><span class=\"c-definition__translation\">")
						if err != nil {
							return err
						}
						err = _writeutil.WriteHTML(_w, def.Translation)
						if err != nil {
							return err
						}
						err = _writeutil.Write(_w, "</span>")
						if err != nil {
							return err
						}
						if def.Definition != "" {
							_closed = true
							err = _writeutil.Write(_w, "<span class=\"c-definition__definition\">&mdash;")
							if err != nil {
								return err
							}
							err = _writeutil.WriteHTML(_w, def.Definition)
							if err != nil {
								return err
							}
							err = _writeutil.Write(_w, "</span>")
							if err != nil {
								return err
							}
						}
						err = _writeutil.Write(_w, "</div>")
						if err != nil {
							return err
						}
						if def.Example != "" {
							_closed = true
							err = _writeutil.Write(_w, "<div class=\"c-definition__example-line\"><span class=\"c-definition__example\">")
							if err != nil {
								return err
							}
							err = _writeutil.WriteHTML(_w, def.Example)
							if err != nil {
								return err
							}
							_closed = true
							err = _writeutil.Write(_w, "</span><span class=\"c-definition__example-translation\">&mdash;")
							if err != nil {
								return err
							}
							err = _writeutil.WriteHTML(_w, def.ExampleTranslation)
							if err != nil {
								return err
							}
							err = _writeutil.Write(_w, "</span></div>")
							if err != nil {
								return err
							}
						}
						err = _writeutil.Write(_w, "</li>")
						if err != nil {
							return err
						}
						i++
					}
				}
				err = _writeutil.Write(_w, "</ol></div></div>")
				if err != nil {
					return err
				}
				if firstDefinition.Type == sueno.Verb {
					_closed = true
					err = _writeutil.Write(_w, "<div class=\"c-entry-section\"><h3 class=\"c-entry-section__title\">Conjugation of <em>")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, word.Word)
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</em></h3><div class=\"c-entry-section__body\"><div class=\"l-conjugation-grid\"><div class=\"c-conjugation-group\"><h4 class=\"c-conjugation-group__title\">Indicative</h4><table class=\"c-conjugation-table c-conjucation-group__table\"><tr><td class=\"c-conjugation-table__tense\">Past:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPastIndicative(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr><td class=\"c-conjugation-table__tense\">Present:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPresentIndicative(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr><td class=\"c-conjugation-table__tense\">Future:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToFutureIndicative(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr></table></div><div class=\"c-conjugation-group\"><h4 class=\"c-conjugation-group__title\">Subjunctive</h4><table class=\"c-conjugation-table c-conjucation-group__table\"><tr><td class=\"c-conjugation-table__tense\">Past:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPastSubjunctive(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr><td class=\"c-conjugation-table__tense\">Present:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPresentSubjunctive(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr><td class=\"c-conjugation-table__tense\">Future:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToFutureSubjunctive(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr></table></div><div class=\"c-conjugation-group\"><h4 class=\"c-conjugation-group__title\">Conditional</h4><table class=\"c-conjugation-table c-conjucation-group__table\"><tr><td class=\"c-conjugation-table__tense\">Past:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPastConditional(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr><td class=\"c-conjugation-table__tense\">Present:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPresentConditional(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr><td class=\"c-conjugation-table__tense\">Future:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToFutureConditional(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr></table></div><div class=\"c-conjugation-group\"><h4 class=\"c-conjugation-group__title\">Participles</h4><table class=\"c-conjugation-table c-conjucation-group__table\"><tr><td class=\"c-conjugation-table__tense\">Past:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPastParticiple(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr><td class=\"c-conjugation-table__tense\">Present:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToPresentParticiple(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr><tr><td class=\"c-conjugation-table__tense\">Future:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToFutureParticiple(word.Word))
					if err != nil {
						return err
					}
					_closed = true
					err = _writeutil.Write(_w, "</td></tr></table></div><div class=\"c-conjugation-group\"><h4 class=\"c-conjugation-group__title\">Imperative</h4><table class=\"c-conjugation-table c-conjucation-group__table\"><td class=\"c-conjugation-table__tense\">Present:</td><td class=\"c-conjugation-table__conjugation\">")
					if err != nil {
						return err
					}
					err = _writeutil.WriteHTML(_w, sueno.ToImperative(word.Word))
					if err != nil {
						return err
					}
					err = _writeutil.Write(_w, "</td></table></div></div></div></div>")
					if err != nil {
						return err
					}
				}
			}
			err = _writeutil.Write(_w, "</div>")
			if err != nil {
				return err
			}
		}
	}
	err = _writeutil.Write(_w, "</div></div>")
	if err != nil {
		return err
	}
	if len(otherRootWords) > 0 {
		_closed = true
		err = _writeutil.Write(_w, "<div class=\"c-entry-section\"><h3 class=\"c-entry-section__title\">Words with the same root <em>")
		if err != nil {
			return err
		}
		err = _writeutil.WriteHTML(_w, word.Root)
		if err != nil {
			return err
		}
		_closed = true
		err = _writeutil.Write(_w, "</em></h3><div class=\"c-entry-section__body\"><ol class=\"c-related-words\">")
		if err != nil {
			return err
		}
		for _, w := range otherRootWords {
			_closed = false
			err = _writeutil.Write(_w, "<li class=\"c-related-word\"><a")
			if err != nil {
				return err
			}
			err = _writeutil.WriteAttr(_w, "href", fmt.Sprintf("/?q=%s&type=word", url.QueryEscape(w.Word)), false)
			if err != nil {
				return err
			}
			_closed = true
			err = _writeutil.Write(_w, " class=\"c-related-word__link\">")
			if err != nil {
				return err
			}
			err = _writeutil.WriteHTML(_w, w.Word)
			if err != nil {
				return err
			}
			_closed = true
			err = _writeutil.Write(_w, "</a><span class=\"c-related-word__type\"> (")
			if err != nil {
				return err
			}
			switch w.Definitions[0].Type {
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
			err = _writeutil.Write(_w, ")</span><span class=\"c-related-word__translation\">&mdash;")
			if err != nil {
				return err
			}
			err = _writeutil.WriteHTML(_w, w.Definitions[0].Translation)
			if err != nil {
				return err
			}
			err = _writeutil.Write(_w, "</span></li>")
			if err != nil {
				return err
			}
		}
		err = _writeutil.Write(_w, "</ol></div></div>")
		if err != nil {
			return err
		}
	}
	err = _writeutil.Write(_w, "</div></body>")
	if err != nil {
		return err
	}
	return nil
}
