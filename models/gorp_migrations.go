// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// GorpMigration is an object representing the database table.
type GorpMigration struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	AppliedAt null.Time `boil:"applied_at" json:"applied_at,omitempty" toml:"applied_at" yaml:"applied_at,omitempty"`

	R *gorpMigrationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gorpMigrationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GorpMigrationColumns = struct {
	ID        string
	AppliedAt string
}{
	ID:        "id",
	AppliedAt: "applied_at",
}

var GorpMigrationTableColumns = struct {
	ID        string
	AppliedAt string
}{
	ID:        "gorp_migrations.id",
	AppliedAt: "gorp_migrations.applied_at",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var GorpMigrationWhere = struct {
	ID        whereHelperstring
	AppliedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"gorp_migrations\".\"id\""},
	AppliedAt: whereHelpernull_Time{field: "\"gorp_migrations\".\"applied_at\""},
}

// GorpMigrationRels is where relationship names are stored.
var GorpMigrationRels = struct {
}{}

// gorpMigrationR is where relationships are stored.
type gorpMigrationR struct {
}

// NewStruct creates a new relationship struct
func (*gorpMigrationR) NewStruct() *gorpMigrationR {
	return &gorpMigrationR{}
}

// gorpMigrationL is where Load methods for each relationship are stored.
type gorpMigrationL struct{}

var (
	gorpMigrationAllColumns            = []string{"id", "applied_at"}
	gorpMigrationColumnsWithoutDefault = []string{"id"}
	gorpMigrationColumnsWithDefault    = []string{"applied_at"}
	gorpMigrationPrimaryKeyColumns     = []string{"id"}
	gorpMigrationGeneratedColumns      = []string{}
)

type (
	// GorpMigrationSlice is an alias for a slice of pointers to GorpMigration.
	// This should almost always be used instead of []GorpMigration.
	GorpMigrationSlice []*GorpMigration
	// GorpMigrationHook is the signature for custom GorpMigration hook methods
	GorpMigrationHook func(context.Context, boil.ContextExecutor, *GorpMigration) error

	gorpMigrationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gorpMigrationType                 = reflect.TypeOf(&GorpMigration{})
	gorpMigrationMapping              = queries.MakeStructMapping(gorpMigrationType)
	gorpMigrationPrimaryKeyMapping, _ = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, gorpMigrationPrimaryKeyColumns)
	gorpMigrationInsertCacheMut       sync.RWMutex
	gorpMigrationInsertCache          = make(map[string]insertCache)
	gorpMigrationUpdateCacheMut       sync.RWMutex
	gorpMigrationUpdateCache          = make(map[string]updateCache)
	gorpMigrationUpsertCacheMut       sync.RWMutex
	gorpMigrationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gorpMigrationAfterSelectMu sync.Mutex
var gorpMigrationAfterSelectHooks []GorpMigrationHook

var gorpMigrationBeforeInsertMu sync.Mutex
var gorpMigrationBeforeInsertHooks []GorpMigrationHook
var gorpMigrationAfterInsertMu sync.Mutex
var gorpMigrationAfterInsertHooks []GorpMigrationHook

var gorpMigrationBeforeUpdateMu sync.Mutex
var gorpMigrationBeforeUpdateHooks []GorpMigrationHook
var gorpMigrationAfterUpdateMu sync.Mutex
var gorpMigrationAfterUpdateHooks []GorpMigrationHook

var gorpMigrationBeforeDeleteMu sync.Mutex
var gorpMigrationBeforeDeleteHooks []GorpMigrationHook
var gorpMigrationAfterDeleteMu sync.Mutex
var gorpMigrationAfterDeleteHooks []GorpMigrationHook

var gorpMigrationBeforeUpsertMu sync.Mutex
var gorpMigrationBeforeUpsertHooks []GorpMigrationHook
var gorpMigrationAfterUpsertMu sync.Mutex
var gorpMigrationAfterUpsertHooks []GorpMigrationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GorpMigration) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GorpMigration) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GorpMigration) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GorpMigration) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GorpMigration) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GorpMigration) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GorpMigration) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GorpMigration) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GorpMigration) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gorpMigrationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGorpMigrationHook registers your hook function for all future operations.
func AddGorpMigrationHook(hookPoint boil.HookPoint, gorpMigrationHook GorpMigrationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		gorpMigrationAfterSelectMu.Lock()
		gorpMigrationAfterSelectHooks = append(gorpMigrationAfterSelectHooks, gorpMigrationHook)
		gorpMigrationAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		gorpMigrationBeforeInsertMu.Lock()
		gorpMigrationBeforeInsertHooks = append(gorpMigrationBeforeInsertHooks, gorpMigrationHook)
		gorpMigrationBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		gorpMigrationAfterInsertMu.Lock()
		gorpMigrationAfterInsertHooks = append(gorpMigrationAfterInsertHooks, gorpMigrationHook)
		gorpMigrationAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		gorpMigrationBeforeUpdateMu.Lock()
		gorpMigrationBeforeUpdateHooks = append(gorpMigrationBeforeUpdateHooks, gorpMigrationHook)
		gorpMigrationBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		gorpMigrationAfterUpdateMu.Lock()
		gorpMigrationAfterUpdateHooks = append(gorpMigrationAfterUpdateHooks, gorpMigrationHook)
		gorpMigrationAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		gorpMigrationBeforeDeleteMu.Lock()
		gorpMigrationBeforeDeleteHooks = append(gorpMigrationBeforeDeleteHooks, gorpMigrationHook)
		gorpMigrationBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		gorpMigrationAfterDeleteMu.Lock()
		gorpMigrationAfterDeleteHooks = append(gorpMigrationAfterDeleteHooks, gorpMigrationHook)
		gorpMigrationAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		gorpMigrationBeforeUpsertMu.Lock()
		gorpMigrationBeforeUpsertHooks = append(gorpMigrationBeforeUpsertHooks, gorpMigrationHook)
		gorpMigrationBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		gorpMigrationAfterUpsertMu.Lock()
		gorpMigrationAfterUpsertHooks = append(gorpMigrationAfterUpsertHooks, gorpMigrationHook)
		gorpMigrationAfterUpsertMu.Unlock()
	}
}

// OneG returns a single gorpMigration record from the query using the global executor.
func (q gorpMigrationQuery) OneG(ctx context.Context) (*GorpMigration, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single gorpMigration record from the query.
func (q gorpMigrationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GorpMigration, error) {
	o := &GorpMigration{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for gorp_migrations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all GorpMigration records from the query using the global executor.
func (q gorpMigrationQuery) AllG(ctx context.Context) (GorpMigrationSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all GorpMigration records from the query.
func (q gorpMigrationQuery) All(ctx context.Context, exec boil.ContextExecutor) (GorpMigrationSlice, error) {
	var o []*GorpMigration

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GorpMigration slice")
	}

	if len(gorpMigrationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all GorpMigration records in the query using the global executor
func (q gorpMigrationQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all GorpMigration records in the query.
func (q gorpMigrationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count gorp_migrations rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q gorpMigrationQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q gorpMigrationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if gorp_migrations exists")
	}

	return count > 0, nil
}

// GorpMigrations retrieves all the records using an executor.
func GorpMigrations(mods ...qm.QueryMod) gorpMigrationQuery {
	mods = append(mods, qm.From("\"gorp_migrations\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"gorp_migrations\".*"})
	}

	return gorpMigrationQuery{q}
}

// FindGorpMigrationG retrieves a single record by ID.
func FindGorpMigrationG(ctx context.Context, iD string, selectCols ...string) (*GorpMigration, error) {
	return FindGorpMigration(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindGorpMigration retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGorpMigration(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*GorpMigration, error) {
	gorpMigrationObj := &GorpMigration{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"gorp_migrations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gorpMigrationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from gorp_migrations")
	}

	if err = gorpMigrationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return gorpMigrationObj, err
	}

	return gorpMigrationObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *GorpMigration) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GorpMigration) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no gorp_migrations provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gorpMigrationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gorpMigrationInsertCacheMut.RLock()
	cache, cached := gorpMigrationInsertCache[key]
	gorpMigrationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gorpMigrationAllColumns,
			gorpMigrationColumnsWithDefault,
			gorpMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"gorp_migrations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"gorp_migrations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into gorp_migrations")
	}

	if !cached {
		gorpMigrationInsertCacheMut.Lock()
		gorpMigrationInsertCache[key] = cache
		gorpMigrationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single GorpMigration record using the global executor.
// See Update for more documentation.
func (o *GorpMigration) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the GorpMigration.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GorpMigration) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gorpMigrationUpdateCacheMut.RLock()
	cache, cached := gorpMigrationUpdateCache[key]
	gorpMigrationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gorpMigrationAllColumns,
			gorpMigrationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update gorp_migrations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"gorp_migrations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gorpMigrationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, append(wl, gorpMigrationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update gorp_migrations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for gorp_migrations")
	}

	if !cached {
		gorpMigrationUpdateCacheMut.Lock()
		gorpMigrationUpdateCache[key] = cache
		gorpMigrationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q gorpMigrationQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q gorpMigrationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for gorp_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for gorp_migrations")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GorpMigrationSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GorpMigrationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gorpMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"gorp_migrations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gorpMigrationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in gorpMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all gorpMigration")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *GorpMigration) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GorpMigration) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no gorp_migrations provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gorpMigrationColumnsWithDefault, o)

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

	gorpMigrationUpsertCacheMut.RLock()
	cache, cached := gorpMigrationUpsertCache[key]
	gorpMigrationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			gorpMigrationAllColumns,
			gorpMigrationColumnsWithDefault,
			gorpMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			gorpMigrationAllColumns,
			gorpMigrationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert gorp_migrations, could not build update column list")
		}

		ret := strmangle.SetComplement(gorpMigrationAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(gorpMigrationPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert gorp_migrations, could not build conflict column list")
			}

			conflict = make([]string, len(gorpMigrationPrimaryKeyColumns))
			copy(conflict, gorpMigrationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"gorp_migrations\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gorpMigrationType, gorpMigrationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert gorp_migrations")
	}

	if !cached {
		gorpMigrationUpsertCacheMut.Lock()
		gorpMigrationUpsertCache[key] = cache
		gorpMigrationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single GorpMigration record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *GorpMigration) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single GorpMigration record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GorpMigration) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GorpMigration provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gorpMigrationPrimaryKeyMapping)
	sql := "DELETE FROM \"gorp_migrations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from gorp_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for gorp_migrations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q gorpMigrationQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q gorpMigrationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no gorpMigrationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gorp_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for gorp_migrations")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o GorpMigrationSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GorpMigrationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gorpMigrationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gorpMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"gorp_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gorpMigrationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gorpMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for gorp_migrations")
	}

	if len(gorpMigrationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *GorpMigration) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no GorpMigration provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GorpMigration) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGorpMigration(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GorpMigrationSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty GorpMigrationSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GorpMigrationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GorpMigrationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gorpMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"gorp_migrations\".* FROM \"gorp_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gorpMigrationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GorpMigrationSlice")
	}

	*o = slice

	return nil
}

// GorpMigrationExistsG checks if the GorpMigration row exists.
func GorpMigrationExistsG(ctx context.Context, iD string) (bool, error) {
	return GorpMigrationExists(ctx, boil.GetContextDB(), iD)
}

// GorpMigrationExists checks if the GorpMigration row exists.
func GorpMigrationExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"gorp_migrations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if gorp_migrations exists")
	}

	return exists, nil
}

// Exists checks if the GorpMigration row exists.
func (o *GorpMigration) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GorpMigrationExists(ctx, exec, o.ID)
}
