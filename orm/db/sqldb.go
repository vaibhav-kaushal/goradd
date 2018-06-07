package db

import (
	"context"
	"database/sql"
	"github.com/spekary/goradd"
	. "github.com/spekary/goradd/orm/query"
	"time"
)

type SqlDbI interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (r sql.Result, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (r *sql.Rows, err error)
	TypeTableSuffix() string
	AssociationTableSuffix() string

	generateSelectSql(QueryBuilderI) (sql string, args []interface{})
	generateDeleteSql(QueryBuilderI) (sql string, args []interface{})
}

// ProfileEntry contains the data collected during sql profiling
type ProfileEntry struct {
	DbKey     string
	BeginTime time.Time
	EndTime   time.Time
	Typ       string
	Sql       string
}

// SqlContext is what is stored in the current context to keep track of queries. You must save a copy of this in the
// current context with the SqlContext key before calling database functions in order to use transactions or
// database profiling, or anything else the context is required for. The framework does this for you, but you will need
// to do this yourself if using the orm without the framework.
type SqlContext struct {
	tx      *sql.Tx
	txCount int // Keeps track of when to close a transaction

	profiles []ProfileEntry
}

type SqlDb struct {
	dbKey string  // key of the database as used in the global database map
	db    *sql.DB // Internal copy of golang database

	// codegen options
	typeTableSuffix        string // Primarily for sql tables
	associationTableSuffix string // Primarily for sql tables
	idSuffix               string // suffix to strip off the ends of names of foreign keys when converting them to internal names

	// These codegen options may be moved higher up in hierarchy some day
	goStructPrefix         string // Helps differentiate objects when different databases have the same name.
	associatedObjectPrefix string // Helps differentiate between objects and local values

	profiling bool
}

func NewSqlDb(dbKey string) SqlDb {
	s := SqlDb{
		dbKey:                  dbKey,
		typeTableSuffix:        "_type",
		associationTableSuffix: "_assn",
		idSuffix:               "_id",
	}
	return s
}

// Begin starts a transaction. You must use the context returned from this function for all subsequent
func (s *SqlDb) Begin(ctx context.Context) {
	var c *SqlContext

	i := ctx.Value(goradd.SqlContext)
	if i == nil {
		panic("Can't use transactions without pre-loading a context")
	} else {
		c = i.(*SqlContext)
	}
	c.txCount++

	if c.txCount == 1 {
		var err error

		c.tx, err = s.db.Begin()
		if err != nil {
			panic(err.Error())
		}
	}
	if i == nil {
		ctx = context.WithValue(ctx, goradd.SqlContext, c)
	}
}

func (s *SqlDb) Commit(ctx context.Context) {
	var c *SqlContext
	i := ctx.Value(goradd.SqlContext)
	if i == nil {
		panic("Can't use transactions without pre-loading a context")
	} else {
		c = i.(*SqlContext)
	}

	c.txCount--
	if c.txCount < 0 {
		panic("Called Commit without a matching Begin")
	}
	if c.txCount == 0 {
		err := c.tx.Commit()
		if err != nil {
			panic(err.Error())
		}
		c.tx = nil
	}
}

func (s *SqlDb) Rollback(ctx context.Context) {
	var c *SqlContext
	i := ctx.Value(goradd.SqlContext)
	if i == nil {
		panic("Can't use transactions without pre-loading a context")
	} else {
		c = i.(*SqlContext)
	}

	if c.tx == nil {
		panic("Called Rollback without matching Begin")
	} else {
		err := c.tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		c.tx = nil
		c.txCount = 0
	}
}

func (s *SqlDb) Exec(ctx context.Context, sql string, args ...interface{}) (r sql.Result, err error) {
	var c *SqlContext
	i := ctx.Value(goradd.SqlContext)
	if i != nil {
		c = i.(*SqlContext)
	}

	var beginTime = time.Now()

	if c != nil && c.tx != nil {
		r, err = c.tx.ExecContext(ctx, sql, args...)
	} else {
		r, err = s.db.ExecContext(ctx, sql, args...)
	}

	var endTime = time.Now()

	if c != nil && s.profiling {
		c.profiles = append(c.profiles, ProfileEntry{DbKey: s.dbKey, BeginTime: beginTime, EndTime: endTime, Typ: "Exec", Sql: sql})
	}

	return
}

func (s *SqlDb) Prepare(ctx context.Context, sql string) (r *sql.Stmt, err error) {
	var c *SqlContext
	i := ctx.Value(goradd.SqlContext)
	if i != nil {
		c = i.(*SqlContext)
	}

	var beginTime = time.Now()
	if c != nil && c.tx != nil {
		r, err = c.tx.Prepare(sql)
	} else {
		r, err = s.db.Prepare(sql)
	}
	var endTime = time.Now()
	if c != nil && s.profiling {
		c.profiles = append(c.profiles, ProfileEntry{DbKey: s.dbKey, BeginTime: beginTime, EndTime: endTime, Typ: "Prepare", Sql: sql})
	}

	return
}

func (s *SqlDb) Query(ctx context.Context, sql string, args ...interface{}) (r *sql.Rows, err error) {
	var c *SqlContext
	i := ctx.Value(goradd.SqlContext)
	if i != nil {
		c = i.(*SqlContext)
	}

	var beginTime = time.Now()
	if c != nil && c.tx != nil {
		r, err = c.tx.QueryContext(ctx, sql, args...)
	} else {
		r, err = s.db.QueryContext(ctx, sql, args...)
	}
	var endTime = time.Now()
	if c != nil && s.profiling {
		c.profiles = append(c.profiles, ProfileEntry{DbKey: s.dbKey, BeginTime: beginTime, EndTime: endTime, Typ: "Query", Sql: sql})
	}

	return
}

func (s *SqlDb) DbKey() string {
	return s.dbKey
}

func (s *SqlDb) SetTypeTableSuffix(suffix string) {
	s.typeTableSuffix = suffix
}

func (s *SqlDb) SetAssociationTableSuffix(suffix string) {
	s.associationTableSuffix = suffix
}

func (s *SqlDb) TypeTableSuffix() string {
	return s.typeTableSuffix
}

func (s *SqlDb) AssociationTableSuffix() string {
	return s.associationTableSuffix
}

func (s *SqlDb) SetGoStructPrefix(prefix string) {
	s.goStructPrefix = prefix
}

func (s *SqlDb) SetAssociatedObjectPrefix(prefix string) {
	s.associatedObjectPrefix = prefix
}

func (s *SqlDb) GoStructPrefix() string {
	return s.goStructPrefix
}

func (s *SqlDb) AssociatedObjectPrefix() string {
	return s.associatedObjectPrefix
}

func (s *SqlDb) IdSuffix() string {
	return s.idSuffix
}

func (s *SqlDb) StartProfiling() {
	s.profiling = true
}

func GetProfiles(ctx context.Context) []ProfileEntry {
	var c *SqlContext
	i := ctx.Value(goradd.SqlContext)
	if i == nil {
		panic("Profiling requires a preloaded context.")
	} else {
		c = i.(*SqlContext)
	}

	if c != nil {
		p := c.profiles
		c.profiles = nil
		return p
	}
	return nil
}