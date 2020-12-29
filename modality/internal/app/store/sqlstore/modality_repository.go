package sqlstore

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"modality/internal/app/model"
)

// ModalityRepository ...
type ModalityRepository struct {
	store *Store
}

// AddText add new text
func (r *ModalityRepository) AddText(objText *model.ObjectText) error {

	addDateTime := time.Now()

	if objText.Text == "" {
		return errors.New("Text object is empty")
	}

	if objText.URL != "" {
		if err := objText.ValidateURL(); err != nil {
			return err
		}
	}

	stmt, err := r.store.db.Prepare("INSERT INTO input_texts (object_text, lang_id, add_date_time, url) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(
		objText.Text,
		objText.Language.ID,
		addDateTime.Format("2006-01-02 15:04:05"),
		objText.URL,
	)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("new object didn`t added")
		}
		if err != nil {
			return err
		}
	}

	if objText.ID, err = commandTag.LastInsertId(); err != nil {
		return err
	}

	return nil

}

// GetTypes get all types of modality
func (r *ModalityRepository) GetTypes(types *model.Types) error {

	stmt, err := r.store.db.Prepare("SELECT id, name FROM simple_types WHERE active=TRUE")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {

		var modType model.ModalityType

		err := rows.Scan(&modType.ID, &modType.Name)
		if err != nil {
			return err
		}

		types.Types = append(types.Types, modType)

	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		return err
	}

	return nil

}

// GetLangs get all languages
func (r *ModalityRepository) GetLangs(langs *model.Languages) error {

	stmt, err := r.store.db.Prepare("SELECT id, name FROM languages")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {

		var lang model.Language

		err := rows.Scan(&lang.ID, &lang.Name)
		if err != nil {
			return err
		}

		langs.Languages = append(langs.Languages, lang)

	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		return err
	}

	return nil

}

// GetPageTexts get page text objects
func (r *ModalityRepository) GetPageTexts(texts *model.ObjectTexts) error {

	strSQL := "SELECT COUNT(id) FROM input_texts AS it"
	strSQL += " WHERE it.active=TRUE"
	strSQL += " AND LOWER(it.object_text) LIKE LOWER($1)"
	strSQL += " AND LOWER(it.url) LIKE LOWER($2)"
	if texts.Filter.LangID != 0 {
		strSQL += " AND it.lang_id=" + strconv.Itoa(texts.Filter.LangID)
	}

	stmt, err := r.store.db.Prepare(strSQL)
	if err != nil {
		return err
	}

	if err := stmt.QueryRow(
		"%"+texts.Filter.TextLike+"%",
		"%"+texts.Filter.URLLike+"%",
	).Scan(&texts.Count); err != nil {
		return err
	}
	stmt.Close()

	var orderBy []string
	for _, val := range texts.SortBy {
		if !val.Ascending {
			val.Name += " DESC"
		}
		orderBy = append(orderBy, val.Name)
	}

	strSQL = "SELECT id FROM input_texts AS it"
	strSQL += " WHERE it.active=TRUE AND LOWER(it.object_text) LIKE LOWER($1)"
	strSQL += " AND LOWER(it.url) LIKE LOWER($2)"
	if texts.Filter.LangID != 0 {
		strSQL += " AND it.lang_id=" + strconv.Itoa(texts.Filter.LangID)
	}
	if len(orderBy) > 0 {
		strSQL += " ORDER BY " + strings.Join(orderBy, ",")
	}
	strSQL += " LIMIT $3 OFFSET $4"

	offset := (texts.Page - 1) * texts.Limit
	if offset < 0 {
		offset = 0
	}

	stmt, err = r.store.db.Prepare(strSQL)
	if err != nil {
		return err
	}

	rows, err := stmt.Query(
		"%"+texts.Filter.TextLike+"%",
		"%"+texts.Filter.URLLike+"%",
		texts.Limit,
		offset,
	)
	if err != nil {
		return err
	}

	var in []string
	// Iterate through the result set
	for rows.Next() {
		var id int
		var err = rows.Scan(&id)
		if err != nil {
			return err
		}
		in = append(in, strconv.Itoa(id))
	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		return err
	}

	rows.Close()
	stmt.Close()

	strSQL = "SELECT it.id, SUBSTR(it.object_text, 0, 50), it.url, l.name FROM input_texts AS it"
	strSQL += " LEFT JOIN languages AS l ON l.id=it.lang_id"
	strSQL += " WHERE it.id IN (" + strings.Join(in, ",") + ")"
	if len(orderBy) > 0 {
		strSQL += " ORDER BY " + strings.Join(orderBy, ",")
	}

	stmt, err = r.store.db.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err = stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {

		var objText model.ObjectText
		var id sql.NullInt64
		var text, url, langName sql.NullString
		var err = rows.Scan(&id, &text, &url, &langName)
		if err != nil {
			return err
		}
		objText.ID = r.store.getInt64(id)
		objText.Text = r.store.getString(text)
		objText.URL = r.store.getString(url)
		objText.Language.Name = r.store.getString(langName)

		texts.ObjectTexts = append(texts.ObjectTexts, objText)
	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		return err
	}

	return nil

}

// GetCurText get current text object
func (r *ModalityRepository) GetCurText(text *model.ObjectText) error {

	strSQL := "SELECT it.id, it.object_text, it.url, l.id, l.name FROM input_texts AS it"
	strSQL += " LEFT JOIN languages AS l ON l.id=it.lang_id"
	strSQL += " WHERE it.id=$1 AND it.active=TRUE"

	stmt, err := r.store.db.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id, langID sql.NullInt64
	var objText, url, langName sql.NullString
	if err = stmt.QueryRow(text.ID).Scan(&id, &objText, &url, &langID, &langName); err != nil {
		return err
	}

	text.ID = r.store.getInt64(id)
	text.Text = r.store.getString(objText)
	text.URL = r.store.getString(url)
	text.Language.Name = r.store.getString(langName)
	text.Language.ID = r.store.getInt(langID)

	return nil

}

// DeleteCurText unactive current text object
func (r *ModalityRepository) DeleteCurText(textID int) error {

	var err error

	// start transaction
	tx, err := r.store.db.Begin()
	if err != nil {
		return err
	}

	if err := r.deleteCurText(textID, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := r.deleteAllModalitiesCurTextObj(textID, tx); err != nil {
		tx.Rollback()
		return err
	}

	// commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

// deleteCurText unactive current text object
func (r *ModalityRepository) deleteCurText(textID int, tx *sql.Tx) error {

	strSQL := "UPDATE input_texts SET active=FALSE WHERE id=$1"
	stmt, err := tx.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(textID)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("text id=" + strconv.Itoa(textID) + " didn`t deleted")
		}
		if err != nil {
			return err
		}
	}
	return nil

}

// deleteAllModalitiesCurTextObj unactive all modalities from current text object
func (r *ModalityRepository) deleteAllModalitiesCurTextObj(textID int, tx *sql.Tx) error {

	strSQL := "UPDATE modalities SET active=FALSE WHERE text_id=$1"
	stmt, err := tx.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(textID); err != nil {
		return err
	}
	return nil

}

// UpdateCurText update current text object
func (r *ModalityRepository) UpdateCurText(textNew, textOld *model.ObjectText) error {

	var err error
	var updated bool

	// start transaction
	tx, err := r.store.db.Begin()
	if err != nil {
		return err
	}

	// correct params ...
	// text if changed
	if textNew.Text != textOld.Text && textNew.Text != "" {
		if err := r.updateObjectText(textNew, tx); err != nil {
			tx.Rollback()
			return err
		}
		updated = true
	}

	// lang id if changed
	if textNew.Language.ID != textOld.Language.ID && textNew.Language.ID != 0 {
		if err := r.updateObjectLanguage(textNew, tx); err != nil {
			tx.Rollback()
			return err
		}
		updated = true
	}

	// url if changed
	if textNew.URL != textOld.URL && textNew.URL != "" {
		if err := textNew.ValidateURL(); err != nil {
			tx.Rollback()
			return err
		}
		if err := r.updateObjectURL(textNew, tx); err != nil {
			tx.Rollback()
			return err
		}
		updated = true
	}

	// didn`t change any data
	if !updated {
		tx.Rollback()
		return errors.New("old and new datas is equal")
	}

	// commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// updateObjectText ...
func (r *ModalityRepository) updateObjectText(textNew *model.ObjectText, tx *sql.Tx) error {

	if count, err := r.getCountOfModalities(textNew.ID); count > 0 || err != nil {
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("text object has modalities, delete them before text changes")
		}
	}

	strSQL := "UPDATE input_texts SET object_text=$1 WHERE id=$2"
	stmt, err := tx.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(textNew.Text, textNew.ID)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("input_text.text_object id=" + strconv.FormatInt(textNew.ID, 10) + " didn`t updated")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// updateObjectLanguage ...
func (r *ModalityRepository) updateObjectLanguage(textNew *model.ObjectText, tx *sql.Tx) error {

	strSQL := "UPDATE input_texts SET lang_id=$1 WHERE id=$2"
	stmt, err := tx.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(textNew.Language.ID, textNew.ID)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("input_text.lang_id id=" + strconv.FormatInt(textNew.ID, 10) + " didn`t updated")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// updateObjectURL ...
func (r *ModalityRepository) updateObjectURL(textNew *model.ObjectText, tx *sql.Tx) error {

	strSQL := "UPDATE input_texts SET url=$1 WHERE id=$2"
	stmt, err := tx.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(textNew.URL, textNew.ID)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("input_text.url id=" + strconv.FormatInt(textNew.ID, 10) + " didn`t updated")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// AddModality add new modality
func (r *ModalityRepository) AddModality(modality *model.Modality) error {

	addDateTime := time.Now()

	if modality.Text == "" {
		return errors.New("text object is empty")
	}

	if modality.TypeID == 0 {
		return errors.New("modality type is empty")
	}

	if modality.TextID == 0 {
		return errors.New("modality parent object is empty")
	}

	stmt, err := r.store.db.Prepare("INSERT INTO modalities (modality_text, type_id, text_id, start_symbol, add_date_time) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(
		modality.Text,
		modality.TypeID,
		modality.TextID,
		modality.StartSymbol,
		addDateTime.Format("2006-01-02 15:04:05"),
	)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("new object didn`t added")
		}
		if err != nil {
			return err
		}
	}
	if modality.ID, err = commandTag.LastInsertId(); err != nil {
		return err
	}
	return nil

}

// GetCurModality get current modality object
func (r *ModalityRepository) GetCurModality(modality *model.Modality) error {

	strSQL := "SELECT id, modality_text, type_id, text_id, start_symbol FROM modalities"
	strSQL += " WHERE id=$1 AND active=TRUE"

	stmt, err := r.store.db.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id, typeID, textID, startSymbol sql.NullInt64
	var modalText sql.NullString
	if err = stmt.QueryRow(modality.ID).Scan(&id, &modalText, &typeID, &textID, &startSymbol); err != nil {
		return err
	}

	modality.ID = r.store.getInt64(id)
	modality.Text = r.store.getString(modalText)
	modality.TypeID = r.store.getInt(typeID)
	modality.TextID = r.store.getInt(textID)
	modality.StartSymbol = r.store.getInt(startSymbol)

	return nil

}

// getCountOfModalities get count of modalities from current text object
func (r *ModalityRepository) getCountOfModalities(textID int64) (int, error) {

	strSQL := "SELECT COUNT(id) FROM modalities"
	strSQL += " WHERE text_id=$1 AND active=TRUE"

	stmt, err := r.store.db.Prepare(strSQL)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int
	if err = stmt.QueryRow(textID).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil

}

// UpdateCurModality update current modality object
func (r *ModalityRepository) UpdateCurModality(modalNew, modalOld *model.Modality) error {

	var err error
	var updated bool

	// start transaction
	tx, err := r.store.db.Begin()
	if err != nil {
		return err
	}

	// correct params ...
	// text if changed
	if modalNew.Text != modalOld.Text && modalNew.Text != "" {
		if err := r.updateModalText(modalNew, tx); err != nil {
			tx.Rollback()
			return err
		}
		updated = true
	}

	// type id if changed
	if modalNew.TypeID != modalOld.TypeID && modalNew.TypeID != 0 {
		if err := r.updateModalType(modalNew, tx); err != nil {
			tx.Rollback()
			return err
		}
		updated = true
	}

	// number of start symbol if changed
	if modalNew.StartSymbol != modalOld.StartSymbol && modalNew.StartSymbol != 0 {
		if err := r.updateModalStartSymbol(modalNew, tx); err != nil {
			tx.Rollback()
			return err
		}
		updated = true
	}

	// didn`t change any data
	if !updated {
		tx.Rollback()
		return errors.New("old and new datas is equal")
	}

	// commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// updateModalText ...
func (r *ModalityRepository) updateModalText(modalNew *model.Modality, tx *sql.Tx) error {

	strSQL := "UPDATE modalities SET modality_text=$1 WHERE id=$2"
	stmt, err := tx.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(modalNew.Text, modalNew.ID)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("modality.text id=" + strconv.FormatInt(modalNew.ID, 10) + " didn`t updated")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// updateModalType ...
func (r *ModalityRepository) updateModalType(modalNew *model.Modality, tx *sql.Tx) error {

	strSQL := "UPDATE modalities SET type_id=$1 WHERE id=$2"
	stmt, err := tx.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(modalNew.TypeID, modalNew.ID)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("modality.type_id id=" + strconv.FormatInt(modalNew.ID, 10) + " didn`t updated")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// updateModalStartSymbol ...
func (r *ModalityRepository) updateModalStartSymbol(modalNew *model.Modality, tx *sql.Tx) error {

	strSQL := "UPDATE modalities SET start_symbol=$1 WHERE id=$2"
	stmt, err := tx.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(modalNew.StartSymbol, modalNew.ID)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("modality.start_symbol id=" + strconv.FormatInt(modalNew.ID, 10) + " didn`t updated")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteCurModality unactive current modality object
func (r *ModalityRepository) DeleteCurModality(textID int) error {

	strSQL := "UPDATE modalities SET active=FALSE WHERE id=$1"
	stmt, err := r.store.db.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	commandTag, err := stmt.Exec(textID)

	if err != nil {
		return err
	}
	if countRows, err := commandTag.RowsAffected(); countRows != 1 || err != nil {
		if countRows != 1 {
			return errors.New("modalities id=" + strconv.Itoa(textID) + " didn`t deleted")
		}
		if err != nil {
			return err
		}
	}
	return nil

}

// GetAllModalitiesFromTextObject get all modalities from current text object
func (r *ModalityRepository) GetAllModalitiesFromTextObject(modalities *model.Modalities, textID int64) error {

	strSQL := "SELECT id, modality_text, type_id, text_id, start_symbol FROM modalities"
	strSQL += " WHERE text_id=$1 AND active=TRUE"
	strSQL += " ORDER BY start_symbol"

	stmt, err := r.store.db.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(textID)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {

		var id, typeID, textID, startSymbol sql.NullInt64
		var modalText sql.NullString
		var modality model.Modality

		var err = rows.Scan(&id, &modalText, &typeID, &textID, &startSymbol)
		if err != nil {
			return err
		}

		modality.ID = r.store.getInt64(id)
		modality.Text = r.store.getString(modalText)
		modality.TypeID = r.store.getInt(typeID)
		modality.TextID = r.store.getInt(textID)
		modality.StartSymbol = r.store.getInt(startSymbol)

		modalities.Modalities = append(modalities.Modalities, modality)
	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		return err
	}

	return nil

}

// GetLangsStatistic get statistic modalities from text objects on languages
func (r *ModalityRepository) GetLangsStatistic(statisticLangs *model.StatisticLanguages, typeIDs []string) error {

	strSQL := "SELECT AVG(a.count_modal) FROM"
	strSQL += " (SELECT COUNT(m.id) count_modal FROM input_texts AS it"
	strSQL += " LEFT JOIN modalities as m ON it.id=m.text_id AND m.type_id IN (" + strings.Join(typeIDs, ",") + ") AND m.active=TRUE"
	strSQL += " WHERE it.lang_id=$1 and it.active=TRUE"
	strSQL += " GROUP BY it.id) AS a"

	stmt, err := r.store.db.Prepare(strSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for key, lang := range statisticLangs.SatatisticLanguages {

		var avg sql.NullFloat64

		err = stmt.QueryRow(lang.ID).Scan(&avg)
		if err != nil {
			return err
		}

		statisticLangs.SatatisticLanguages[key].AVGCount = r.store.getFloat(avg, 1000)

	}

	return nil

}
