package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ModalityType simple type of modality
type ModalityType struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

// Types simple types of modality
type Types struct {
	Types []ModalityType `json:"types"`
}

// ObjectText text object for analysis
type ObjectText struct {
	ID       int64  `json:"id"`
	Text     string `json:"text,omitempty"`
	Language `json:"lang,omitempty"`
	URL      string `json:"url,omitempty"`
	// AddDateTime Time   `json:"add_date_time,omitempty"`
	// Active      bool   `json:"add_date_time,omitempty"`
}

// Filter ...
type Filter struct {
	TextLike string `json:"text_like,omitempty"`
	URLLike  string `json:"url_like,omitempty"`
	LangID   int    `json:"lang_id,omitempty"`
}

// SortBy ...
type SortBy struct {
	Name      string `json:"name"`
	Ascending bool   `json:"ascending"`
}

// ObjectTexts ...
type ObjectTexts struct {
	ObjectTexts []ObjectText `json:"obect_texts"`
	Page        int          `json:"page,omitempty"`
	Limit       int          `json:"limit,omitempty"`
	Count       int          `json:"count,omitempty"`
	SortBy      []SortBy     `json:"sort_by,omitempty"`
	Filter      `json:"filter,omitempty"`
}

// Language ...
type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

// Languages ...
type Languages struct {
	Languages []Language `json:"languages"`
}

// Modality fucken modality
type Modality struct {
	ID          int64  `json:"id"`
	Text        string `json:"text,omitempty"`
	TypeID      int    `json:"type_id,omitempty"`
	TextID      int    `json:"text_id,omitempty"`
	StartSymbol int    `json:"start_symbol,omitempty"`
}

// Modalities ...
type Modalities struct {
	Modalities []Modality `json:"modalities"`
}

// ValidateURL ...
func (ot *ObjectText) ValidateURL() error {
	return validation.ValidateStruct(
		ot,
		validation.Field(&ot.URL, validation.Required, is.URL),
	)
}
