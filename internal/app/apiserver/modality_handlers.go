package apiserver

import (
	"errors"
	"net/http"
	"strconv"

	"modality/internal/app/model"

	"github.com/gin-gonic/gin"
)

// handleTextAdd add new text to DB
func (s *server) handleTextAdd() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		objText := &model.ObjectText{}
		if err := ctx.BindJSON(objText); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Modality().AddText(objText); err != nil {
			s.respondWithError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		ctx.JSON(http.StatusOK, objText)

	}
}

// handleTypesGet get all types
func (s *server) handleTypesGet() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		types := &model.Types{}

		err := s.store.Modality().GetTypes(types)
		if err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, types)
		return

	}
}

// handleLangsGet get all languages
func (s *server) handleLangsGet() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		langs := &model.Languages{}

		if err := s.store.Modality().GetLangs(langs); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, langs)
		return

	}
}

// handlePageTextsGet get page of texts
func (s *server) handlePageTextsGet() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		objTexts := &model.ObjectTexts{}
		if err := ctx.BindJSON(objTexts); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		if objTexts.Page == 0 {
			objTexts.Page = 1
		}
		if objTexts.Limit == 0 {
			objTexts.Limit = 20
		}

		if err := s.store.Modality().GetPageTexts(objTexts); err != nil {
			s.respondWithError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		ctx.JSON(http.StatusOK, objTexts)

	}
}

// handleCurTextGet get current text object
func (s *server) handleCurTextGet() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		keys := ctx.Request.URL.Query()
		var err error

		if len(keys) > 0 {
			if idStr, ok := keys["id"]; ok {

				objText := &model.ObjectText{}

				objText.ID, err = strconv.ParseInt(idStr[0], 10, 64)
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				err = s.store.Modality().GetCurText(objText)
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}
				ctx.JSON(http.StatusOK, objText)
				return

			}
		} else {
			s.respondWithError(ctx, http.StatusBadRequest, errors.New("absent params in request"))
			return
		}

	}
}

// handleCurTextDelete inactive current text object
func (s *server) handleCurTextDelete() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		keys := ctx.Request.URL.Query()

		if len(keys) > 0 {
			if idStr, ok := keys["id"]; ok {

				textID, err := strconv.Atoi(idStr[0])
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				err = s.store.Modality().DeleteCurText(textID)
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}
				ctx.Status(http.StatusOK)
				return

			}
		} else {
			s.respondWithError(ctx, http.StatusBadRequest, errors.New("absent params in request"))
			return
		}

	}
}

// handleCurTextUpdate edit current text object
func (s *server) handleCurTextUpdate() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		textNew := &model.ObjectText{}
		if err := ctx.BindJSON(textNew); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		textOld := &model.ObjectText{
			ID: textNew.ID,
		}

		if err := s.store.Modality().GetCurText(textOld); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Modality().UpdateCurText(textNew, textOld); err != nil {
			s.respondWithError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		ctx.JSON(http.StatusOK, textNew)

	}
}

// handleModalityAdd add new modality at text object
func (s *server) handleModalityAdd() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		modality := &model.Modality{}
		if err := ctx.BindJSON(modality); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Modality().AddModality(modality); err != nil {
			s.respondWithError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		ctx.JSON(http.StatusOK, modality)

	}
}

// handleModalityGet get current modality
func (s *server) handleModalityGet() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		keys := ctx.Request.URL.Query()
		var err error

		if len(keys) > 0 {
			if idStr, ok := keys["id"]; ok {

				objModality := &model.Modality{}

				objModality.ID, err = strconv.ParseInt(idStr[0], 10, 64)
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				err := s.store.Modality().GetCurModality(objModality)
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				ctx.JSON(http.StatusOK, objModality)
				return

			}
		} else {
			s.respondWithError(ctx, http.StatusBadRequest, errors.New("absent params in request"))
			return
		}
	}
}

// handleCurModalityDelete inactive current modality object
func (s *server) handleCurModalityDelete() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		keys := ctx.Request.URL.Query()

		if len(keys) > 0 {
			if idStr, ok := keys["id"]; ok {

				modalID, err := strconv.Atoi(idStr[0])
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				err = s.store.Modality().DeleteCurModality(modalID)
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}
				ctx.Status(http.StatusOK)
				return

			}
		} else {
			s.respondWithError(ctx, http.StatusBadRequest, errors.New("absent params in request"))
			return
		}

	}
}

// handleModalityUpdate edit current text object
func (s *server) handleModalityUpdate() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		modalNew := &model.Modality{}
		if err := ctx.BindJSON(modalNew); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		modalOld := &model.Modality{
			ID: modalNew.ID,
		}

		if err := s.store.Modality().GetCurModality(modalOld); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Modality().UpdateCurModality(modalNew, modalOld); err != nil {
			s.respondWithError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		ctx.JSON(http.StatusOK, modalNew)

	}
}

// handleModalitiesGet get all modalities from current text object
func (s *server) handleModalitiesGet() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		keys := ctx.Request.URL.Query()

		if len(keys) > 0 {
			if idStr, ok := keys["id"]; ok {

				modalities := &model.Modalities{}

				textID, err := strconv.ParseInt(idStr[0], 10, 64)
				if err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				if err = s.store.Modality().GetAllModalitiesFromTextObject(modalities, textID); err != nil {
					s.respondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				ctx.JSON(http.StatusOK, modalities)
				return

			}
		} else {
			s.respondWithError(ctx, http.StatusBadRequest, errors.New("absent params in request"))
			return
		}
	}
}

// handleStatisticLanguagesGet get statistic lnguages
func (s *server) handleStatisticLanguagesGet() gin.HandlerFunc {

	type ModalIDs struct {
		TypeIDs []int64 `json:"type_ids"`
	}

	return func(ctx *gin.Context) {

		var typeIDs ModalIDs
		if err := ctx.BindJSON(&typeIDs); err != nil {
			s.respondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		var typeIDsStr []string
		for _, val := range typeIDs.TypeIDs {

			typeIDsStr = append(typeIDsStr, strconv.FormatInt(val, 10))

		}

		var langs model.Languages
		if err := s.store.Modality().GetLangs(&langs); err != nil {
			s.respondWithError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		var statLangs model.StatisticLanguages
		for _, lang := range langs.Languages {

			var statLang model.StatisticLanguage
			statLang.Language = lang
			statLangs.SatatisticLanguages = append(statLangs.SatatisticLanguages, statLang)

		}

		if err := s.store.Modality().GetLangsStatistic(&statLangs, typeIDsStr); err != nil {
			s.respondWithError(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		ctx.JSON(http.StatusOK, &statLangs)

	}
}
