// Copyright 2015-2016, Cyrill @ Schumacher.fm and the CoreStore contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package csdb_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/corestoreio/csfw/storage/csdb"
	"github.com/corestoreio/csfw/storage/dbr"
	"github.com/stretchr/testify/assert"
	"strings"
)

type typeWriter struct {
	Write *csdb.ResurrectStmt
}

func newTypeWriterMocked(db *sql.DB) *typeWriter {
	tw := &typeWriter{
		Write: csdb.NewResurrectStmt(db, "INSERT INTO `xtable` (`path`,`value`) VALUES (?,?)"),
	}
	tw.Write.Idle = time.Millisecond * 50
	return tw
}

func newTypeWriterReal(db *sql.DB) *typeWriter {
	tw := &typeWriter{
		Write: csdb.NewResurrectStmt(db, "REPLACE INTO `core_config_data` (`path`,`value`) VALUES (?,?)"),
	}
	tw.Write.Idle = time.Millisecond * 50
	return tw
}

func (tw *typeWriter) Save(key string, value int) error {
	tw.Write.StartStmtUse()
	defer tw.Write.StopStmtUse()

	stmt, err := tw.Write.Stmt()
	if err != nil {
		return err
	}

	result, err := stmt.Exec(key, value)
	if err != nil {
		return err
	}

	liID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if liID < 1 && ra < 1 {
		return fmt.Errorf("No rows inserted (%d) nor affected (%d)", liID, ra)
	}
	return nil
}

func TestResurrectStmtSqlMockNoTicker(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO `xtable` \\(`path`,`value`\\) VALUES .+").
		ExpectExec().WithArgs("gopher", 3141).WillReturnResult(sqlmock.NewResult(1, 1))

	tw := newTypeWriterMocked(db)

	assert.NoError(t, tw.Save("gopher", 3141))
	assert.False(t, tw.Write.IsIdle())

	assert.NoError(t, tw.Write.StopIdleChecker())

	mock.ExpectPrepare("INSERT INTO `xtable` \\(`path`,`value`\\) VALUES .+").
		ExpectExec().WithArgs("gopher", 3144).WillReturnResult(sqlmock.NewResult(1, 1))

	assert.NoError(t, tw.Save("gopher", 3144))

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestResurrectStmtSqlMockShouldPrepareOnceAndThenBecomeIdle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO `xtable` \\(`path`,`value`\\) VALUES .+").
		ExpectExec().WithArgs("gopher", 3141).WillReturnResult(sqlmock.NewResult(1, 1))

	tw := newTypeWriterMocked(db)
	tw.Write.StartIdleChecker()
	tw.Write.StartIdleChecker() // 2x

	assert.NoError(t, tw.Save("gopher", 3141))
	assert.False(t, tw.Write.IsIdle())

	time.Sleep(time.Millisecond * 60)
	assert.True(t, tw.Write.IsIdle())

	assert.NoError(t, tw.Write.StopIdleChecker())

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestResurrectStmtSqlMockShouldPrepareTwoTimesWithThreeCalls(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO `xtable` \\(`path`,`value`\\) VALUES .+").
		ExpectExec().
		WithArgs("gopher", 3141).
		WillReturnResult(sqlmock.NewResult(1, 0))

	tw := newTypeWriterMocked(db)
	tw.Write.StartIdleChecker()

	assert.NoError(t, tw.Save("gopher", 3141))
	assert.False(t, tw.Write.IsIdle())

	mock.
		ExpectExec("INSERT INTO `xtable` \\(`path`,`value`\\) VALUES .+").
		WithArgs("gopher", 3142).
		WillReturnResult(sqlmock.NewResult(1, 0))

	assert.NoError(t, tw.Save("gopher", 3142))
	assert.False(t, tw.Write.IsIdle())

	time.Sleep(time.Millisecond * 60)
	assert.True(t, tw.Write.IsIdle())
	assert.NoError(t, tw.Write.StopIdleChecker())

	mock.ExpectPrepare("INSERT INTO `xtable` \\(`path`,`value`\\) VALUES .+").
		ExpectExec().
		WithArgs("gopher", 271828).
		WillReturnResult(sqlmock.NewResult(1, 0))

	assert.NoError(t, tw.Save("gopher", 271828))

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestResurrectStmtRealDB(t *testing.T) {
	debugLogBuf.Reset()
	infoLogBuf.Reset()
	defer debugLogBuf.Reset()
	defer infoLogBuf.Reset()

	dbc := csdb.MustConnectTest()
	defer func() { assert.NoError(t, dbc.Close()) }()

	tw := newTypeWriterReal(dbc.DB)
	tw.Write.StartIdleChecker()

	assert.NoError(t, tw.Save("RSgopher1", 1))
	assert.False(t, tw.Write.IsIdle())

	assert.NoError(t, tw.Save("RSgopher2", 2))
	assert.False(t, tw.Write.IsIdle())

	time.Sleep(time.Millisecond * 60) // 1. close
	assert.True(t, tw.Write.IsIdle())
	assert.NoError(t, tw.Write.StopIdleChecker()) // 2.close

	assert.NoError(t, tw.Save("RSgopher3", 3))
	assert.NoError(t, tw.Save("RSgopher4", 4))

	//	println("\n", debugLogBuf.String(), "\n")
	//	println("\n", infoLogBuf.String(), "\n")

	// to be more precise you must check the order of the logged values
	assert.Exactly(t, 2, strings.Count(debugLogBuf.String(), `csdb.ResurrectStmt.stmt.Close SQL: "REPLACE INTO`))
	assert.Exactly(t, 2, strings.Count(debugLogBuf.String(), `csdb.ResurrectStmt.stmt.Prepare SQL: "REPLACE INTO`))

	res, err := dbc.NewSession().DeleteFrom("core_config_data").Where(dbr.ConditionRaw("path like \"RSgopher%\"")).Exec()
	assert.NoError(t, err)
	ar, err := res.RowsAffected()
	assert.NoError(t, err)
	assert.Exactly(t, int64(4), ar)
}
