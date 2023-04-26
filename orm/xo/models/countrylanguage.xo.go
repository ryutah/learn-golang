package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Countrylanguage represents a row from 'public.countrylanguage'.
type Countrylanguage struct {
	Countrycode string  `json:"countrycode" db:"countrycode"` // countrycode
	Language    string  `json:"language" db:"language"`       // language
	Isofficial  bool    `json:"isofficial" db:"isofficial"`   // isofficial
	Percentage  float32 `json:"percentage" db:"percentage"`   // percentage
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Countrylanguage exists in the database.
func (c *Countrylanguage) Exists() bool {
	return c._exists
}

// Deleted returns true when the Countrylanguage has been marked for deletion from
// the database.
func (c *Countrylanguage) Deleted() bool {
	return c._deleted
}

// Insert inserts the Countrylanguage to the database.
func (c *Countrylanguage) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.countrylanguage (` +
		`countrycode, language, isofficial, percentage` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`
	// run
	logf(sqlstr, c.Countrycode, c.Language, c.Isofficial, c.Percentage)
	if _, err := db.ExecContext(ctx, sqlstr, c.Countrycode, c.Language, c.Isofficial, c.Percentage); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Update updates a Countrylanguage in the database.
func (c *Countrylanguage) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.countrylanguage SET ` +
		`isofficial = $1, percentage = $2 ` +
		`WHERE countrycode = $3 AND language = $4`
	// run
	logf(sqlstr, c.Isofficial, c.Percentage, c.Countrycode, c.Language)
	if _, err := db.ExecContext(ctx, sqlstr, c.Isofficial, c.Percentage, c.Countrycode, c.Language); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Countrylanguage to the database.
func (c *Countrylanguage) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Upsert performs an upsert for Countrylanguage.
func (c *Countrylanguage) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.countrylanguage (` +
		`countrycode, language, isofficial, percentage` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (countrycode, language) DO ` +
		`UPDATE SET ` +
		`isofficial = EXCLUDED.isofficial, percentage = EXCLUDED.percentage `
	// run
	logf(sqlstr, c.Countrycode, c.Language, c.Isofficial, c.Percentage)
	if _, err := db.ExecContext(ctx, sqlstr, c.Countrycode, c.Language, c.Isofficial, c.Percentage); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Delete deletes the Countrylanguage from the database.
func (c *Countrylanguage) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM public.countrylanguage ` +
		`WHERE countrycode = $1 AND language = $2`
	// run
	logf(sqlstr, c.Countrycode, c.Language)
	if _, err := db.ExecContext(ctx, sqlstr, c.Countrycode, c.Language); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CountrylanguageByCountrycodeLanguage retrieves a row from 'public.countrylanguage' as a Countrylanguage.
//
// Generated from index 'countrylanguage_pkey'.
func CountrylanguageByCountrycodeLanguage(ctx context.Context, db DB, countrycode, language string) (*Countrylanguage, error) {
	// query
	const sqlstr = `SELECT ` +
		`countrycode, language, isofficial, percentage ` +
		`FROM public.countrylanguage ` +
		`WHERE countrycode = $1 AND language = $2`
	// run
	logf(sqlstr, countrycode, language)
	c := Countrylanguage{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, countrycode, language).Scan(&c.Countrycode, &c.Language, &c.Isofficial, &c.Percentage); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}

// Country returns the Country associated with the Countrylanguage's (Countrycode).
//
// Generated from foreign key 'countrylanguage_countrycode_fkey'.
func (c *Countrylanguage) Country(ctx context.Context, db DB) (*Country, error) {
	return CountryByCode(ctx, db, c.Countrycode)
}
