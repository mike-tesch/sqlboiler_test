// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTestTables(t *testing.T) {
	t.Parallel()

	query := TestTables()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTestTablesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTestTablesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TestTables().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTestTablesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TestTableSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTestTablesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TestTableExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TestTable exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TestTableExists to return true, but got false.")
	}
}

func testTestTablesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	testTableFound, err := FindTestTable(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if testTableFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTestTablesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TestTables().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTestTablesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TestTables().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTestTablesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	testTableOne := &TestTable{}
	testTableTwo := &TestTable{}
	if err = randomize.Struct(seed, testTableOne, testTableDBTypes, false, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}
	if err = randomize.Struct(seed, testTableTwo, testTableDBTypes, false, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = testTableOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = testTableTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TestTables().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTestTablesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	testTableOne := &TestTable{}
	testTableTwo := &TestTable{}
	if err = randomize.Struct(seed, testTableOne, testTableDBTypes, false, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}
	if err = randomize.Struct(seed, testTableTwo, testTableDBTypes, false, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = testTableOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = testTableTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTableBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTableAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTableAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTableBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTableAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTableBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTableAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTableBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTableAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TestTable) error {
	*o = TestTable{}
	return nil
}

func testTestTablesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TestTable{}
	o := &TestTable{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, testTableDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TestTable object: %s", err)
	}

	AddTestTableHook(boil.BeforeInsertHook, testTableBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	testTableBeforeInsertHooks = []TestTableHook{}

	AddTestTableHook(boil.AfterInsertHook, testTableAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	testTableAfterInsertHooks = []TestTableHook{}

	AddTestTableHook(boil.AfterSelectHook, testTableAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	testTableAfterSelectHooks = []TestTableHook{}

	AddTestTableHook(boil.BeforeUpdateHook, testTableBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	testTableBeforeUpdateHooks = []TestTableHook{}

	AddTestTableHook(boil.AfterUpdateHook, testTableAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	testTableAfterUpdateHooks = []TestTableHook{}

	AddTestTableHook(boil.BeforeDeleteHook, testTableBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	testTableBeforeDeleteHooks = []TestTableHook{}

	AddTestTableHook(boil.AfterDeleteHook, testTableAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	testTableAfterDeleteHooks = []TestTableHook{}

	AddTestTableHook(boil.BeforeUpsertHook, testTableBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	testTableBeforeUpsertHooks = []TestTableHook{}

	AddTestTableHook(boil.AfterUpsertHook, testTableAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	testTableAfterUpsertHooks = []TestTableHook{}
}

func testTestTablesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTestTablesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(testTableColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTestTablesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTestTablesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TestTableSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTestTablesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TestTables().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	testTableDBTypes = map[string]string{`ID`: `uuid`, `Date`: `datetime`}
	_                = bytes.MinRead
)

func testTestTablesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(testTablePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(testTableAllColumns) == len(testTablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTestTablesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(testTableAllColumns) == len(testTablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TestTable{}
	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, testTableDBTypes, true, testTablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(testTableAllColumns, testTablePrimaryKeyColumns) {
		fields = testTableAllColumns
	} else {
		fields = strmangle.SetComplement(
			testTableAllColumns,
			testTablePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TestTableSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTestTablesUpsert(t *testing.T) {
	t.Parallel()

	if len(testTableAllColumns) == len(testTablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TestTable{}
	if err = randomize.Struct(seed, &o, testTableDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TestTable: %s", err)
	}

	count, err := TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, testTableDBTypes, false, testTablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TestTable struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TestTable: %s", err)
	}

	count, err = TestTables().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
