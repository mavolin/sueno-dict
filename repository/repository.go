package repository

import (
	"context"
	"strconv"

	"github.com/mavolin/sueno-dict/pkl/sueno"
)

type Repository interface {
	// WordRoot returns the words associated with the given root, sorted
	// by [Word.Type] and then by name.
	WordRoot(ctx context.Context, root string) (*WordRoot, error)

	CreateWord(context.Context, Word) (WordID, error)
	Word(context.Context, WordID) (*Word, error)
	SearchWord(context.Context, string) (*Word, error)
	SearchTranslation(context.Context, string) ([]TranslatedWord, error)
}

// ============================================================================
// WordRoot
// ======================================================================================

type WordRoot struct {
	Root  string
	Words []Word
}

// ============================================================================
// Word
// ======================================================================================

type WordID uint64

func ParseWordID(s string) (WordID, error) {
	idInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return WordID(idInt), nil
}

type (
	Word struct {
		ID   WordID `gorm:"primary_key"`
		Word string
		Root string

		Definitions   []Definition `gorm:"foreignKey:WordID"`
		CompoundWords []Word       `gorm:"-"`
	}

	TranslatedWord struct {
		Word
		MatchedDefinition Definition `gorm:"-"`
	}
)

// ============================================================================
// Definition
// ======================================================================================

type DefinitionID uint64

func ParseDefinitionID(s string) (DefinitionID, error) {
	idInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return DefinitionID(idInt), nil
}

type Definition struct {
	ID     DefinitionID `gorm:"primary_key"`
	WordID WordID
	Type   sueno.WordType

	Definition  string
	Translation string

	Example            string
	ExampleTranslation string
}
