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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Building is an object representing the database table.
type Building struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Address   string    `boil:"address" json:"address" toml:"address" yaml:"address"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *buildingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L buildingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BuildingColumns = struct {
	ID        string
	Name      string
	Address   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Address:   "address",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var BuildingTableColumns = struct {
	ID        string
	Name      string
	Address   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "building.id",
	Name:      "building.name",
	Address:   "building.address",
	CreatedAt: "building.created_at",
	UpdatedAt: "building.updated_at",
}

// Generated where

var BuildingWhere = struct {
	ID        whereHelperint
	Name      whereHelperstring
	Address   whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"building\".\"id\""},
	Name:      whereHelperstring{field: "\"building\".\"name\""},
	Address:   whereHelperstring{field: "\"building\".\"address\""},
	CreatedAt: whereHelpertime_Time{field: "\"building\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"building\".\"updated_at\""},
}

// BuildingRels is where relationship names are stored.
var BuildingRels = struct {
	Apartments string
}{
	Apartments: "Apartments",
}

// buildingR is where relationships are stored.
type buildingR struct {
	Apartments ApartmentSlice `boil:"Apartments" json:"Apartments" toml:"Apartments" yaml:"Apartments"`
}

// NewStruct creates a new relationship struct
func (*buildingR) NewStruct() *buildingR {
	return &buildingR{}
}

func (r *buildingR) GetApartments() ApartmentSlice {
	if r == nil {
		return nil
	}
	return r.Apartments
}

// buildingL is where Load methods for each relationship are stored.
type buildingL struct{}

var (
	buildingAllColumns            = []string{"id", "name", "address", "created_at", "updated_at"}
	buildingColumnsWithoutDefault = []string{"name", "address", "created_at", "updated_at"}
	buildingColumnsWithDefault    = []string{"id"}
	buildingPrimaryKeyColumns     = []string{"id"}
	buildingGeneratedColumns      = []string{}
)

type (
	// BuildingSlice is an alias for a slice of pointers to Building.
	// This should almost always be used instead of []Building.
	BuildingSlice []*Building
	// BuildingHook is the signature for custom Building hook methods
	BuildingHook func(context.Context, boil.ContextExecutor, *Building) error

	buildingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	buildingType                 = reflect.TypeOf(&Building{})
	buildingMapping              = queries.MakeStructMapping(buildingType)
	buildingPrimaryKeyMapping, _ = queries.BindMapping(buildingType, buildingMapping, buildingPrimaryKeyColumns)
	buildingInsertCacheMut       sync.RWMutex
	buildingInsertCache          = make(map[string]insertCache)
	buildingUpdateCacheMut       sync.RWMutex
	buildingUpdateCache          = make(map[string]updateCache)
	buildingUpsertCacheMut       sync.RWMutex
	buildingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var buildingAfterSelectMu sync.Mutex
var buildingAfterSelectHooks []BuildingHook

var buildingBeforeInsertMu sync.Mutex
var buildingBeforeInsertHooks []BuildingHook
var buildingAfterInsertMu sync.Mutex
var buildingAfterInsertHooks []BuildingHook

var buildingBeforeUpdateMu sync.Mutex
var buildingBeforeUpdateHooks []BuildingHook
var buildingAfterUpdateMu sync.Mutex
var buildingAfterUpdateHooks []BuildingHook

var buildingBeforeDeleteMu sync.Mutex
var buildingBeforeDeleteHooks []BuildingHook
var buildingAfterDeleteMu sync.Mutex
var buildingAfterDeleteHooks []BuildingHook

var buildingBeforeUpsertMu sync.Mutex
var buildingBeforeUpsertHooks []BuildingHook
var buildingAfterUpsertMu sync.Mutex
var buildingAfterUpsertHooks []BuildingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Building) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Building) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Building) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Building) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Building) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Building) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Building) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Building) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Building) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBuildingHook registers your hook function for all future operations.
func AddBuildingHook(hookPoint boil.HookPoint, buildingHook BuildingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		buildingAfterSelectMu.Lock()
		buildingAfterSelectHooks = append(buildingAfterSelectHooks, buildingHook)
		buildingAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		buildingBeforeInsertMu.Lock()
		buildingBeforeInsertHooks = append(buildingBeforeInsertHooks, buildingHook)
		buildingBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		buildingAfterInsertMu.Lock()
		buildingAfterInsertHooks = append(buildingAfterInsertHooks, buildingHook)
		buildingAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		buildingBeforeUpdateMu.Lock()
		buildingBeforeUpdateHooks = append(buildingBeforeUpdateHooks, buildingHook)
		buildingBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		buildingAfterUpdateMu.Lock()
		buildingAfterUpdateHooks = append(buildingAfterUpdateHooks, buildingHook)
		buildingAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		buildingBeforeDeleteMu.Lock()
		buildingBeforeDeleteHooks = append(buildingBeforeDeleteHooks, buildingHook)
		buildingBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		buildingAfterDeleteMu.Lock()
		buildingAfterDeleteHooks = append(buildingAfterDeleteHooks, buildingHook)
		buildingAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		buildingBeforeUpsertMu.Lock()
		buildingBeforeUpsertHooks = append(buildingBeforeUpsertHooks, buildingHook)
		buildingBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		buildingAfterUpsertMu.Lock()
		buildingAfterUpsertHooks = append(buildingAfterUpsertHooks, buildingHook)
		buildingAfterUpsertMu.Unlock()
	}
}

// OneG returns a single building record from the query using the global executor.
func (q buildingQuery) OneG(ctx context.Context) (*Building, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single building record from the query.
func (q buildingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Building, error) {
	o := &Building{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for building")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Building records from the query using the global executor.
func (q buildingQuery) AllG(ctx context.Context) (BuildingSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Building records from the query.
func (q buildingQuery) All(ctx context.Context, exec boil.ContextExecutor) (BuildingSlice, error) {
	var o []*Building

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Building slice")
	}

	if len(buildingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Building records in the query using the global executor
func (q buildingQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Building records in the query.
func (q buildingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count building rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q buildingQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q buildingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if building exists")
	}

	return count > 0, nil
}

// Apartments retrieves all the apartment's Apartments with an executor.
func (o *Building) Apartments(mods ...qm.QueryMod) apartmentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"apartment\".\"building_id\"=?", o.ID),
	)

	return Apartments(queryMods...)
}

// LoadApartments allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (buildingL) LoadApartments(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBuilding interface{}, mods queries.Applicator) error {
	var slice []*Building
	var object *Building

	if singular {
		var ok bool
		object, ok = maybeBuilding.(*Building)
		if !ok {
			object = new(Building)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBuilding)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBuilding))
			}
		}
	} else {
		s, ok := maybeBuilding.(*[]*Building)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBuilding)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBuilding))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &buildingR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &buildingR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`apartment`),
		qm.WhereIn(`apartment.building_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load apartment")
	}

	var resultSlice []*Apartment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice apartment")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on apartment")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for apartment")
	}

	if len(apartmentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Apartments = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &apartmentR{}
			}
			foreign.R.Building = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.BuildingID {
				local.R.Apartments = append(local.R.Apartments, foreign)
				if foreign.R == nil {
					foreign.R = &apartmentR{}
				}
				foreign.R.Building = local
				break
			}
		}
	}

	return nil
}

// AddApartmentsG adds the given related objects to the existing relationships
// of the building, optionally inserting them as new records.
// Appends related to o.R.Apartments.
// Sets related.R.Building appropriately.
// Uses the global database handle.
func (o *Building) AddApartmentsG(ctx context.Context, insert bool, related ...*Apartment) error {
	return o.AddApartments(ctx, boil.GetContextDB(), insert, related...)
}

// AddApartments adds the given related objects to the existing relationships
// of the building, optionally inserting them as new records.
// Appends related to o.R.Apartments.
// Sets related.R.Building appropriately.
func (o *Building) AddApartments(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Apartment) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.BuildingID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"apartment\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"building_id"}),
				strmangle.WhereClause("\"", "\"", 2, apartmentPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.BuildingID = o.ID
		}
	}

	if o.R == nil {
		o.R = &buildingR{
			Apartments: related,
		}
	} else {
		o.R.Apartments = append(o.R.Apartments, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &apartmentR{
				Building: o,
			}
		} else {
			rel.R.Building = o
		}
	}
	return nil
}

// Buildings retrieves all the records using an executor.
func Buildings(mods ...qm.QueryMod) buildingQuery {
	mods = append(mods, qm.From("\"building\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"building\".*"})
	}

	return buildingQuery{q}
}

// FindBuildingG retrieves a single record by ID.
func FindBuildingG(ctx context.Context, iD int, selectCols ...string) (*Building, error) {
	return FindBuilding(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindBuilding retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBuilding(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Building, error) {
	buildingObj := &Building{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"building\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, buildingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from building")
	}

	if err = buildingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return buildingObj, err
	}

	return buildingObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Building) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Building) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no building provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(buildingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	buildingInsertCacheMut.RLock()
	cache, cached := buildingInsertCache[key]
	buildingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			buildingAllColumns,
			buildingColumnsWithDefault,
			buildingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(buildingType, buildingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(buildingType, buildingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"building\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"building\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into building")
	}

	if !cached {
		buildingInsertCacheMut.Lock()
		buildingInsertCache[key] = cache
		buildingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Building record using the global executor.
// See Update for more documentation.
func (o *Building) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Building.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Building) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	buildingUpdateCacheMut.RLock()
	cache, cached := buildingUpdateCache[key]
	buildingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			buildingAllColumns,
			buildingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update building, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"building\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, buildingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(buildingType, buildingMapping, append(wl, buildingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update building row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for building")
	}

	if !cached {
		buildingUpdateCacheMut.Lock()
		buildingUpdateCache[key] = cache
		buildingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q buildingQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q buildingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for building")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for building")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BuildingSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BuildingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"building\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, buildingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in building slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all building")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Building) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Building) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no building provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(buildingColumnsWithDefault, o)

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

	buildingUpsertCacheMut.RLock()
	cache, cached := buildingUpsertCache[key]
	buildingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			buildingAllColumns,
			buildingColumnsWithDefault,
			buildingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			buildingAllColumns,
			buildingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert building, could not build update column list")
		}

		ret := strmangle.SetComplement(buildingAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(buildingPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert building, could not build conflict column list")
			}

			conflict = make([]string, len(buildingPrimaryKeyColumns))
			copy(conflict, buildingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"building\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(buildingType, buildingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(buildingType, buildingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert building")
	}

	if !cached {
		buildingUpsertCacheMut.Lock()
		buildingUpsertCache[key] = cache
		buildingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Building record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Building) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Building record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Building) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Building provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), buildingPrimaryKeyMapping)
	sql := "DELETE FROM \"building\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from building")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for building")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q buildingQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q buildingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no buildingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from building")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for building")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BuildingSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BuildingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(buildingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"building\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from building slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for building")
	}

	if len(buildingAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Building) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Building provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Building) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBuilding(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BuildingSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty BuildingSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BuildingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BuildingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"building\".* FROM \"building\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BuildingSlice")
	}

	*o = slice

	return nil
}

// BuildingExistsG checks if the Building row exists.
func BuildingExistsG(ctx context.Context, iD int) (bool, error) {
	return BuildingExists(ctx, boil.GetContextDB(), iD)
}

// BuildingExists checks if the Building row exists.
func BuildingExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"building\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if building exists")
	}

	return exists, nil
}

// Exists checks if the Building row exists.
func (o *Building) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BuildingExists(ctx, exec, o.ID)
}
