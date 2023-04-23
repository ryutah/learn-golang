// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Countrylanguage is an object representing the database table.
type Countrylanguage struct {
	Countrycode string  `boil:"countrycode" json:"countrycode" toml:"countrycode" yaml:"countrycode"`
	Language    string  `boil:"language" json:"language" toml:"language" yaml:"language"`
	Isofficial  bool    `boil:"isofficial" json:"isofficial" toml:"isofficial" yaml:"isofficial"`
	Percentage  float32 `boil:"percentage" json:"percentage" toml:"percentage" yaml:"percentage"`

	R *countrylanguageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L countrylanguageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CountrylanguageColumns = struct {
	Countrycode string
	Language    string
	Isofficial  string
	Percentage  string
}{
	Countrycode: "countrycode",
	Language:    "language",
	Isofficial:  "isofficial",
	Percentage:  "percentage",
}

var CountrylanguageTableColumns = struct {
	Countrycode string
	Language    string
	Isofficial  string
	Percentage  string
}{
	Countrycode: "countrylanguage.countrycode",
	Language:    "countrylanguage.language",
	Isofficial:  "countrylanguage.isofficial",
	Percentage:  "countrylanguage.percentage",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var CountrylanguageWhere = struct {
	Countrycode whereHelperstring
	Language    whereHelperstring
	Isofficial  whereHelperbool
	Percentage  whereHelperfloat32
}{
	Countrycode: whereHelperstring{field: "\"countrylanguage\".\"countrycode\""},
	Language:    whereHelperstring{field: "\"countrylanguage\".\"language\""},
	Isofficial:  whereHelperbool{field: "\"countrylanguage\".\"isofficial\""},
	Percentage:  whereHelperfloat32{field: "\"countrylanguage\".\"percentage\""},
}

// CountrylanguageRels is where relationship names are stored.
var CountrylanguageRels = struct {
	CountrycodeCountry string
}{
	CountrycodeCountry: "CountrycodeCountry",
}

// countrylanguageR is where relationships are stored.
type countrylanguageR struct {
	CountrycodeCountry *Country `boil:"CountrycodeCountry" json:"CountrycodeCountry" toml:"CountrycodeCountry" yaml:"CountrycodeCountry"`
}

// NewStruct creates a new relationship struct
func (*countrylanguageR) NewStruct() *countrylanguageR {
	return &countrylanguageR{}
}

func (r *countrylanguageR) GetCountrycodeCountry() *Country {
	if r == nil {
		return nil
	}
	return r.CountrycodeCountry
}

// countrylanguageL is where Load methods for each relationship are stored.
type countrylanguageL struct{}

var (
	countrylanguageAllColumns            = []string{"countrycode", "language", "isofficial", "percentage"}
	countrylanguageColumnsWithoutDefault = []string{"countrycode", "language", "isofficial", "percentage"}
	countrylanguageColumnsWithDefault    = []string{}
	countrylanguagePrimaryKeyColumns     = []string{"countrycode", "language"}
	countrylanguageGeneratedColumns      = []string{}
)

type (
	// CountrylanguageSlice is an alias for a slice of pointers to Countrylanguage.
	// This should almost always be used instead of []Countrylanguage.
	CountrylanguageSlice []*Countrylanguage
	// CountrylanguageHook is the signature for custom Countrylanguage hook methods
	CountrylanguageHook func(context.Context, boil.ContextExecutor, *Countrylanguage) error

	countrylanguageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	countrylanguageType                 = reflect.TypeOf(&Countrylanguage{})
	countrylanguageMapping              = queries.MakeStructMapping(countrylanguageType)
	countrylanguagePrimaryKeyMapping, _ = queries.BindMapping(countrylanguageType, countrylanguageMapping, countrylanguagePrimaryKeyColumns)
	countrylanguageInsertCacheMut       sync.RWMutex
	countrylanguageInsertCache          = make(map[string]insertCache)
	countrylanguageUpdateCacheMut       sync.RWMutex
	countrylanguageUpdateCache          = make(map[string]updateCache)
	countrylanguageUpsertCacheMut       sync.RWMutex
	countrylanguageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var countrylanguageAfterSelectHooks []CountrylanguageHook

var countrylanguageBeforeInsertHooks []CountrylanguageHook
var countrylanguageAfterInsertHooks []CountrylanguageHook

var countrylanguageBeforeUpdateHooks []CountrylanguageHook
var countrylanguageAfterUpdateHooks []CountrylanguageHook

var countrylanguageBeforeDeleteHooks []CountrylanguageHook
var countrylanguageAfterDeleteHooks []CountrylanguageHook

var countrylanguageBeforeUpsertHooks []CountrylanguageHook
var countrylanguageAfterUpsertHooks []CountrylanguageHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Countrylanguage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Countrylanguage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Countrylanguage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Countrylanguage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Countrylanguage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Countrylanguage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Countrylanguage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Countrylanguage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Countrylanguage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range countrylanguageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCountrylanguageHook registers your hook function for all future operations.
func AddCountrylanguageHook(hookPoint boil.HookPoint, countrylanguageHook CountrylanguageHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		countrylanguageAfterSelectHooks = append(countrylanguageAfterSelectHooks, countrylanguageHook)
	case boil.BeforeInsertHook:
		countrylanguageBeforeInsertHooks = append(countrylanguageBeforeInsertHooks, countrylanguageHook)
	case boil.AfterInsertHook:
		countrylanguageAfterInsertHooks = append(countrylanguageAfterInsertHooks, countrylanguageHook)
	case boil.BeforeUpdateHook:
		countrylanguageBeforeUpdateHooks = append(countrylanguageBeforeUpdateHooks, countrylanguageHook)
	case boil.AfterUpdateHook:
		countrylanguageAfterUpdateHooks = append(countrylanguageAfterUpdateHooks, countrylanguageHook)
	case boil.BeforeDeleteHook:
		countrylanguageBeforeDeleteHooks = append(countrylanguageBeforeDeleteHooks, countrylanguageHook)
	case boil.AfterDeleteHook:
		countrylanguageAfterDeleteHooks = append(countrylanguageAfterDeleteHooks, countrylanguageHook)
	case boil.BeforeUpsertHook:
		countrylanguageBeforeUpsertHooks = append(countrylanguageBeforeUpsertHooks, countrylanguageHook)
	case boil.AfterUpsertHook:
		countrylanguageAfterUpsertHooks = append(countrylanguageAfterUpsertHooks, countrylanguageHook)
	}
}

// One returns a single countrylanguage record from the query.
func (q countrylanguageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Countrylanguage, error) {
	o := &Countrylanguage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for countrylanguage")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Countrylanguage records from the query.
func (q countrylanguageQuery) All(ctx context.Context, exec boil.ContextExecutor) (CountrylanguageSlice, error) {
	var o []*Countrylanguage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Countrylanguage slice")
	}

	if len(countrylanguageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Countrylanguage records in the query.
func (q countrylanguageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count countrylanguage rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q countrylanguageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if countrylanguage exists")
	}

	return count > 0, nil
}

// CountrycodeCountry pointed to by the foreign key.
func (o *Countrylanguage) CountrycodeCountry(mods ...qm.QueryMod) countryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"code\" = ?", o.Countrycode),
	}

	queryMods = append(queryMods, mods...)

	return Countries(queryMods...)
}

// LoadCountrycodeCountry allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (countrylanguageL) LoadCountrycodeCountry(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCountrylanguage interface{}, mods queries.Applicator) error {
	var slice []*Countrylanguage
	var object *Countrylanguage

	if singular {
		var ok bool
		object, ok = maybeCountrylanguage.(*Countrylanguage)
		if !ok {
			object = new(Countrylanguage)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCountrylanguage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCountrylanguage))
			}
		}
	} else {
		s, ok := maybeCountrylanguage.(*[]*Countrylanguage)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCountrylanguage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCountrylanguage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &countrylanguageR{}
		}
		args = append(args, object.Countrycode)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &countrylanguageR{}
			}

			for _, a := range args {
				if a == obj.Countrycode {
					continue Outer
				}
			}

			args = append(args, obj.Countrycode)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`country`),
		qm.WhereIn(`country.code in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Country")
	}

	var resultSlice []*Country
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Country")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for country")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for country")
	}

	if len(countryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CountrycodeCountry = foreign
		if foreign.R == nil {
			foreign.R = &countryR{}
		}
		foreign.R.CountrycodeCountrylanguages = append(foreign.R.CountrycodeCountrylanguages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Countrycode == foreign.Code {
				local.R.CountrycodeCountry = foreign
				if foreign.R == nil {
					foreign.R = &countryR{}
				}
				foreign.R.CountrycodeCountrylanguages = append(foreign.R.CountrycodeCountrylanguages, local)
				break
			}
		}
	}

	return nil
}

// SetCountrycodeCountry of the countrylanguage to the related item.
// Sets o.R.CountrycodeCountry to related.
// Adds o to related.R.CountrycodeCountrylanguages.
func (o *Countrylanguage) SetCountrycodeCountry(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Country) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"countrylanguage\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"countrycode"}),
		strmangle.WhereClause("\"", "\"", 2, countrylanguagePrimaryKeyColumns),
	)
	values := []interface{}{related.Code, o.Countrycode, o.Language}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Countrycode = related.Code
	if o.R == nil {
		o.R = &countrylanguageR{
			CountrycodeCountry: related,
		}
	} else {
		o.R.CountrycodeCountry = related
	}

	if related.R == nil {
		related.R = &countryR{
			CountrycodeCountrylanguages: CountrylanguageSlice{o},
		}
	} else {
		related.R.CountrycodeCountrylanguages = append(related.R.CountrycodeCountrylanguages, o)
	}

	return nil
}

// Countrylanguages retrieves all the records using an executor.
func Countrylanguages(mods ...qm.QueryMod) countrylanguageQuery {
	mods = append(mods, qm.From("\"countrylanguage\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"countrylanguage\".*"})
	}

	return countrylanguageQuery{q}
}

// FindCountrylanguage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCountrylanguage(ctx context.Context, exec boil.ContextExecutor, countrycode string, language string, selectCols ...string) (*Countrylanguage, error) {
	countrylanguageObj := &Countrylanguage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"countrylanguage\" where \"countrycode\"=$1 AND \"language\"=$2", sel,
	)

	q := queries.Raw(query, countrycode, language)

	err := q.Bind(ctx, exec, countrylanguageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from countrylanguage")
	}

	if err = countrylanguageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return countrylanguageObj, err
	}

	return countrylanguageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Countrylanguage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no countrylanguage provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(countrylanguageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	countrylanguageInsertCacheMut.RLock()
	cache, cached := countrylanguageInsertCache[key]
	countrylanguageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			countrylanguageAllColumns,
			countrylanguageColumnsWithDefault,
			countrylanguageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(countrylanguageType, countrylanguageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(countrylanguageType, countrylanguageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"countrylanguage\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"countrylanguage\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into countrylanguage")
	}

	if !cached {
		countrylanguageInsertCacheMut.Lock()
		countrylanguageInsertCache[key] = cache
		countrylanguageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Countrylanguage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Countrylanguage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	countrylanguageUpdateCacheMut.RLock()
	cache, cached := countrylanguageUpdateCache[key]
	countrylanguageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			countrylanguageAllColumns,
			countrylanguagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update countrylanguage, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"countrylanguage\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, countrylanguagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(countrylanguageType, countrylanguageMapping, append(wl, countrylanguagePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update countrylanguage row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for countrylanguage")
	}

	if !cached {
		countrylanguageUpdateCacheMut.Lock()
		countrylanguageUpdateCache[key] = cache
		countrylanguageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q countrylanguageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for countrylanguage")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for countrylanguage")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CountrylanguageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countrylanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"countrylanguage\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, countrylanguagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in countrylanguage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all countrylanguage")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Countrylanguage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no countrylanguage provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(countrylanguageColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	countrylanguageUpsertCacheMut.RLock()
	cache, cached := countrylanguageUpsertCache[key]
	countrylanguageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			countrylanguageAllColumns,
			countrylanguageColumnsWithDefault,
			countrylanguageColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			countrylanguageAllColumns,
			countrylanguagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("model: unable to upsert countrylanguage, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(countrylanguagePrimaryKeyColumns))
			copy(conflict, countrylanguagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"countrylanguage\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(countrylanguageType, countrylanguageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(countrylanguageType, countrylanguageMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "model: unable to upsert countrylanguage")
	}

	if !cached {
		countrylanguageUpsertCacheMut.Lock()
		countrylanguageUpsertCache[key] = cache
		countrylanguageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Countrylanguage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Countrylanguage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no Countrylanguage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), countrylanguagePrimaryKeyMapping)
	sql := "DELETE FROM \"countrylanguage\" WHERE \"countrycode\"=$1 AND \"language\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from countrylanguage")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for countrylanguage")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q countrylanguageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no countrylanguageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from countrylanguage")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for countrylanguage")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CountrylanguageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(countrylanguageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countrylanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"countrylanguage\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, countrylanguagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from countrylanguage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for countrylanguage")
	}

	if len(countrylanguageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Countrylanguage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCountrylanguage(ctx, exec, o.Countrycode, o.Language)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CountrylanguageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CountrylanguageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countrylanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"countrylanguage\".* FROM \"countrylanguage\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, countrylanguagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in CountrylanguageSlice")
	}

	*o = slice

	return nil
}

// CountrylanguageExists checks if the Countrylanguage row exists.
func CountrylanguageExists(ctx context.Context, exec boil.ContextExecutor, countrycode string, language string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"countrylanguage\" where \"countrycode\"=$1 AND \"language\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, countrycode, language)
	}
	row := exec.QueryRowContext(ctx, sql, countrycode, language)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if countrylanguage exists")
	}

	return exists, nil
}

// Exists checks if the Countrylanguage row exists.
func (o *Countrylanguage) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CountrylanguageExists(ctx, exec, o.Countrycode, o.Language)
}
