package repository

import (
	"context"
	"strconv"

	"github.com/mavolin/sueno-dict/pkg/sueno"
)

type Repository interface {
	// WordRoot returns the words associated with the given root, sorted
	// by [Word.Type] and then by name.
	WordRoot(ctx context.Context, root string) (*WordRoot, error)

	CreateWord(context.Context, Word) (WordID, error)
	Word(context.Context, WordID) (*Word, error)
	SearchWord(context.Context, string) (*Word, error)
	SearchTranslation(context.Context, string) ([]TranslatedWord, error)
	DeleteWord(context.Context, WordID) error
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
		ID   WordID `gorm:"primaryKey;autoIncrement;not null"`
		Word string `gorm:"not null"`
		Root string `gorm:"not null;default:''"`

		// CustomPlural allows overriding the plural form of the word.
		//
		// If this is empty, the plural form is grammatically derived from the word.
		//
		// Used for words such as 'ja' (plural: 'jaos').
		CustomPlural string `gorm:"not null;default:''"`

		Definitions   []Definition `gorm:"foreignKey:WordID;references:ID;constraint:OnDelete:cascade"`
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
	ID     DefinitionID   `gorm:"primaryKey;autoIncrement;not null"`
	WordID WordID         `gorm:"not null"`
	Type   sueno.WordType `gorm:"not null"`

	// all the below except translation are optional, denoted by an empty value

	Definition  string `gorm:"not null;default:''"`
	Translation string `gorm:"not null"`

	Example            string `gorm:"not null;default:''"`
	ExampleTranslation string `gorm:"not null;default:''"`
}
