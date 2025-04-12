package validation

import (
	"database/sql"
	"errors"
	"github.com/valyala/fasthttp"
	"govtech-hackaton-backend/cmd/processing/db"
)

func ProcessRequest(ctx *fasthttp.RequestCtx, request *BovineDuplicateRequest) error {
	var bovineCount int

	err := db.DB.QueryRow(`SELECT count FROM user_animals WHERE user_id=$1 AND animal_type='bovine'`, request.UserID).Scan(&bovineCount)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return err
	}

	status := "approved"

	if bovineCount == 0 {
		status = "declined"
	} else if len(request.Duplicates) > bovineCount {
		request.Duplicates = request.Duplicates[:bovineCount]
	}

	tx, err := db.DB.Begin()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return err
	}

	defer tx.Rollback()

	var requestID int
	err = tx.QueryRow(`
		INSERT INTO animal_duplicate_requests (
			user_id, event_code, institution_or_name, fiscal_or_personal_code, farm_name,
			address, locality, representative_name, phone, email, full_name, status
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id`,
		request.UserID, request.EventCode, request.InstitutionOrName, request.FiscalOrPersonalCode,
		request.FarmName, request.Address, request.Locality, request.RepresentativeName,
		request.Phone, request.Email, request.FullName, status).Scan(&requestID)

	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return err
	}

	for _, d := range request.Duplicates {
		_, err = tx.Exec(`INSERT INTO animal_duplicates (request_id, tag_number) VALUES ($1,$2)`, requestID, d.TagNumber)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return err
	}

	return nil
}
