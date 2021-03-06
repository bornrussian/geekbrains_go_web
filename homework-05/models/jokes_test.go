// Code generated by SQLBoiler 4.0.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testJokes(t *testing.T) {
	t.Parallel()

	query := Jokes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testJokesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
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

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJokesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Jokes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJokesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JokeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJokesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := JokeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Joke exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JokeExists to return true, but got false.")
	}
}

func testJokesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	jokeFound, err := FindJoke(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if jokeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testJokesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Jokes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testJokesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Jokes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJokesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jokeOne := &Joke{}
	jokeTwo := &Joke{}
	if err = randomize.Struct(seed, jokeOne, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}
	if err = randomize.Struct(seed, jokeTwo, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jokeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jokeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Jokes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJokesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jokeOne := &Joke{}
	jokeTwo := &Joke{}
	if err = randomize.Struct(seed, jokeOne, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}
	if err = randomize.Struct(seed, jokeTwo, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jokeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jokeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func jokeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func testJokesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Joke{}
	o := &Joke{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jokeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Joke object: %s", err)
	}

	AddJokeHook(boil.BeforeInsertHook, jokeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jokeBeforeInsertHooks = []JokeHook{}

	AddJokeHook(boil.AfterInsertHook, jokeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jokeAfterInsertHooks = []JokeHook{}

	AddJokeHook(boil.AfterSelectHook, jokeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jokeAfterSelectHooks = []JokeHook{}

	AddJokeHook(boil.BeforeUpdateHook, jokeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jokeBeforeUpdateHooks = []JokeHook{}

	AddJokeHook(boil.AfterUpdateHook, jokeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jokeAfterUpdateHooks = []JokeHook{}

	AddJokeHook(boil.BeforeDeleteHook, jokeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jokeBeforeDeleteHooks = []JokeHook{}

	AddJokeHook(boil.AfterDeleteHook, jokeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jokeAfterDeleteHooks = []JokeHook{}

	AddJokeHook(boil.BeforeUpsertHook, jokeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jokeBeforeUpsertHooks = []JokeHook{}

	AddJokeHook(boil.AfterUpsertHook, jokeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jokeAfterUpsertHooks = []JokeHook{}
}

func testJokesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJokesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(jokeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJokesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
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

func testJokesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JokeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testJokesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Jokes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jokeDBTypes = map[string]string{`ID`: `bigint`, `Autor`: `varchar`, `Date`: `varchar`, `Header`: `varchar`, `Content`: `text`}
	_           = bytes.MinRead
)

func testJokesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(jokePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(jokeAllColumns) == len(jokePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testJokesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jokeAllColumns) == len(jokePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jokeAllColumns, jokePrimaryKeyColumns) {
		fields = jokeAllColumns
	} else {
		fields = strmangle.SetComplement(
			jokeAllColumns,
			jokePrimaryKeyColumns,
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

	slice := JokeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testJokesUpsert(t *testing.T) {
	t.Parallel()

	if len(jokeAllColumns) == len(jokePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLJokeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Joke{}
	if err = randomize.Struct(seed, &o, jokeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Joke: %s", err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, jokeDBTypes, false, jokePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Joke: %s", err)
	}

	count, err = Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
