package postgres

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/mavolin/sueno-dict/repository"
)

type (
	word struct {
		repository.Word

		CompoundWords []compoundWord `gorm:"foreignKey:WordID;references:ID;constraint:OnDelete:cascade"`
	}

	compoundWord struct {
		// CompoundWordID is the [repository.WordID] of the compound word.
		CompoundWordID repository.WordID `gorm:"not null"`
		CompoundWord   word
		// WordID is the [repository.WordID] of the word the [CompoundWordID]
		// is part of.
		WordID repository.WordID `gorm:"primaryKey;not null"`

		Index int `gorm:"primaryKey;not null"`
	}
)

func (ww *word) toRepoType() repository.Word {
	w := ww.Word
	if len(ww.CompoundWords) == 0 {
		return w
	}

	w.CompoundWords = make([]repository.Word, len(ww.CompoundWords))

	for i, cw := range ww.CompoundWords {
		w.CompoundWords[i] = cw.CompoundWord.toRepoType()
	}

	return w
}

func wordFromRepoType(w repository.Word) *word {
	ww := &word{Word: w}

	if len(w.CompoundWords) == 0 {
		return ww
	}

	ww.CompoundWords = make([]compoundWord, len(w.CompoundWords))

	for i, cw := range w.CompoundWords {
		ww.CompoundWords[i] = compoundWord{
			CompoundWordID: cw.ID,
			WordID:         ww.ID,
			Index:          i,
		}
	}

	return ww
}

func (r *Repository) WordRoot(ctx context.Context, root string) (*repository.WordRoot, error) {
	var wws []word

	res := r.DB.
		WithContext(ctx).
		Where(&repository.Word{Root: root}).
		Preload("Definitions").
		Preload("CompoundWords", func(db *gorm.DB) *gorm.DB {
			return db.Order("index")
		}).
		Preload("CompoundWords.CompoundWord").
		Order("(select type from definitions where definitions.word_id = words.id order by definitions.id limit 1)").
		Order("word").
		Find(&wws)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "postgres: WordRoot")
	}

	if len(wws) == 0 {
		return nil, nil
	}

	ws := make([]repository.Word, len(wws))
	for i, ww := range wws {
		ws[i] = ww.toRepoType()
	}

	return &repository.WordRoot{Root: root, Words: ws}, nil
}

func (r *Repository) CreateWord(ctx context.Context, w repository.Word) (repository.WordID, error) {
	w.ID = 0

	ww := wordFromRepoType(w)

	res := r.DB.
		WithContext(ctx).
		Create(&ww)
	if res.Error != nil {
		return 0, errors.Wrap(res.Error, "postgres: CreateWord")
	}

	return ww.ID, nil
}

func (r *Repository) Word(ctx context.Context, id repository.WordID) (*repository.Word, error) {
	var ww word

	res := r.DB.
		WithContext(ctx).
		Where(&repository.Word{ID: id}).
		Preload("Definitions").
		Preload("CompoundWords", func(db *gorm.DB) *gorm.DB {
			return db.Order("index")
		}).
		Preload("CompoundWords.CompoundWord").
		First(&ww)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(res.Error, "postgres: Word")
	}

	w := ww.toRepoType()
	return &w, errors.Wrap(res.Error, "postgres: Word")
}

func (r *Repository) SearchWord(ctx context.Context, query string) (*repository.Word, error) {
	var ww word

	res := r.DB.
		WithContext(ctx).
		Where(&repository.Word{Word: query}).
		Preload("Definitions", func(db *gorm.DB) *gorm.DB {
			return db.Order("id")
		}).
		Preload("CompoundWords", func(db *gorm.DB) *gorm.DB {
			return db.Order("index")
		}).
		Preload("CompoundWords.CompoundWord").
		First(&ww)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(res.Error, "postgres: Word")
	}

	w := ww.toRepoType()
	return &w, errors.Wrap(res.Error, "postgres: Word")
}

func (r *Repository) SearchTranslation(ctx context.Context, query string) ([]repository.TranslatedWord, error) {
	var tws []repository.TranslatedWord

	err := r.DB.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var ds []repository.Definition

			res := tx.
				Where("translation like ?", "%"+query+"%").
				Order("(select type from definitions as ds " +
					"where ds.word_id = definitions.word_id " +
					"order by ds.id limit 1)").
				Order("(select word from words where words.id = definitions.word_id)").
				Limit(75).
				Find(&ds)
			if res.Error != nil {
				return errors.Wrap(res.Error, "could not load matching definitions")
			}

			ids := make([]repository.WordID, len(ds))
			for i, d := range ds {
				ids[i] = d.WordID
			}

			var wws []word

			res = r.DB.
				WithContext(ctx).
				Where("words.id in (?)", ids).
				Preload("Definitions").
				Preload("CompoundWords", func(db *gorm.DB) *gorm.DB {
					return db.Order("index")
				}).
				Preload("CompoundWords.CompoundWord").
				Order("(select type from definitions where definitions.word_id = words.id order by definitions.id limit 1)").
				Order("word").
				Find(&wws)
			if res.Error != nil {
				return errors.Wrap(res.Error, "could not load matching words")
			}

			tws = make([]repository.TranslatedWord, len(ds))

			var defIndex int

			for _, ww := range wws {
				for defIndex < len(ds) {
					def := ds[defIndex]
					if def.WordID != ww.ID {
						break
					}

					tws[defIndex] = repository.TranslatedWord{
						Word:              ww.toRepoType(),
						MatchedDefinition: def,
					}

					defIndex++
				}
			}

			return nil
		})

	return tws, errors.WithMessage(err, "postgres: SearchTranslation")
}

func (r *Repository) CompoundWordsContaining(ctx context.Context, id repository.WordID) ([]repository.Word, error) {
	var wws []word

	res := r.DB.
		WithContext(ctx).
		Where("id in (select word_id from compound_words where compound_word_id = ?)", id).
		Preload("Definitions").
		Preload("CompoundWords", func(db *gorm.DB) *gorm.DB {
			return db.Order("index")
		}).
		Preload("CompoundWords.CompoundWord").
		Find(&wws)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "postgres: CompoundWordsContaining")
	}

	ws := make([]repository.Word, len(wws))
	for i, ww := range wws {
		ws[i] = ww.toRepoType()
	}

	return ws, nil
}

func (r *Repository) DeleteWord(ctx context.Context, id repository.WordID) error {
	res := r.DB.
		WithContext(ctx).
		Delete(&repository.Word{}, id)
	if res.Error != nil {
		return errors.Wrap(res.Error, "postgres: DeleteWord")
	}

	return nil
}

func (r *Repository) migrate() error {
	err := r.DB.AutoMigrate(&repository.Word{}, &repository.Definition{}, &word{}, &compoundWord{})
	return errors.Wrap(err, "postgres: could not migrate database")
}
