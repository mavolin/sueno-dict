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

		CompoundWords []compoundWord
	}

	compoundWord struct {
		CompoundWordID repository.WordID `gorm:"primary_key"`
		WordID         repository.WordID `gorm:"primary_key"`
		Word           word              `gorm:"foreignKey:WordID"`

		Index int
	}
)

func (ww *word) toRepoType() repository.Word {
	w := ww.Word
	if len(ww.CompoundWords) == 0 {
		return w
	}

	w.CompoundWords = make([]repository.Word, len(ww.CompoundWords))

	for i, cw := range ww.CompoundWords {
		w.CompoundWords[i] = cw.Word.toRepoType()
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
		Preload("CompoundWords.Word").
		Order("(select type from definitions where definitions.word_id = words.id limit 1)").
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
		Preload("CompoundWords.Word").
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
		Preload("CompoundWords.Word").
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

func (r *Repository) SearchTranslation(ctx context.Context, query string) ([]repository.Word, error) {
	var wws []word

	res := r.DB.
		WithContext(ctx).
		Where("id in (?)", r.DB.
			Select("word_id").
			Where("translation like ?", "%"+query+"%").
			Model(&repository.Definition{}),
		).
		Preload("Definitions").
		Preload("CompoundWords", func(db *gorm.DB) *gorm.DB {
			return db.Order("index")
		}).
		Preload("CompoundWords.Word").
		First(&wws)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(res.Error, "postgres: Word")
	}

	ws := make([]repository.Word, len(wws))
	for i, ww := range wws {
		ws[i] = ww.toRepoType()
	}

	return ws, errors.Wrap(res.Error, "postgres: Word")
}

func (r *Repository) migrate() error {
	err := r.DB.AutoMigrate(&word{}, &repository.Definition{}, compoundWord{})
	return errors.Wrap(err, "postgres: could not migrate database")
}
