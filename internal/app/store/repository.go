package store

import (
	"modality/internal/app/model"
)

// ModalityRepository ...
type ModalityRepository interface {
	GetTypes(*model.Types) error
	GetLangs(*model.Languages) error
	AddText(*model.ObjectText) error
	GetPageTexts(*model.ObjectTexts) error
	GetCurText(*model.ObjectText) error
	DeleteCurText(int) error
	UpdateCurText(*model.ObjectText, *model.ObjectText) error
	AddModality(*model.Modality) error
	GetCurModality(*model.Modality) error
	UpdateCurModality(*model.Modality, *model.Modality) error
}
