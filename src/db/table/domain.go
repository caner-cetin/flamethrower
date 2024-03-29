//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Domain = newDomainTable("", "domain", "")

type domainTable struct {
	sqlite.Table

	// Columns
	ID            sqlite.ColumnInteger
	Name          sqlite.ColumnString
	GrantedPowers sqlite.ColumnString
	Spell1        sqlite.ColumnString
	Spell2        sqlite.ColumnString
	Spell3        sqlite.ColumnString
	Spell4        sqlite.ColumnString
	Spell5        sqlite.ColumnString
	Spell6        sqlite.ColumnString
	Spell7        sqlite.ColumnString
	Spell8        sqlite.ColumnString
	Spell9        sqlite.ColumnString
	FullText      sqlite.ColumnString
	Reference     sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type DomainTable struct {
	domainTable

	EXCLUDED domainTable
}

// AS creates new DomainTable with assigned alias
func (a DomainTable) AS(alias string) *DomainTable {
	return newDomainTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DomainTable with assigned schema name
func (a DomainTable) FromSchema(schemaName string) *DomainTable {
	return newDomainTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DomainTable with assigned table prefix
func (a DomainTable) WithPrefix(prefix string) *DomainTable {
	return newDomainTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DomainTable with assigned table suffix
func (a DomainTable) WithSuffix(suffix string) *DomainTable {
	return newDomainTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDomainTable(schemaName, tableName, alias string) *DomainTable {
	return &DomainTable{
		domainTable: newDomainTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newDomainTableImpl("", "excluded", ""),
	}
}

func newDomainTableImpl(schemaName, tableName, alias string) domainTable {
	var (
		IDColumn            = sqlite.IntegerColumn("id")
		NameColumn          = sqlite.StringColumn("name")
		GrantedPowersColumn = sqlite.StringColumn("granted_powers")
		Spell1Column        = sqlite.StringColumn("spell_1")
		Spell2Column        = sqlite.StringColumn("spell_2")
		Spell3Column        = sqlite.StringColumn("spell_3")
		Spell4Column        = sqlite.StringColumn("spell_4")
		Spell5Column        = sqlite.StringColumn("spell_5")
		Spell6Column        = sqlite.StringColumn("spell_6")
		Spell7Column        = sqlite.StringColumn("spell_7")
		Spell8Column        = sqlite.StringColumn("spell_8")
		Spell9Column        = sqlite.StringColumn("spell_9")
		FullTextColumn      = sqlite.StringColumn("full_text")
		ReferenceColumn     = sqlite.StringColumn("reference")
		allColumns          = sqlite.ColumnList{IDColumn, NameColumn, GrantedPowersColumn, Spell1Column, Spell2Column, Spell3Column, Spell4Column, Spell5Column, Spell6Column, Spell7Column, Spell8Column, Spell9Column, FullTextColumn, ReferenceColumn}
		mutableColumns      = sqlite.ColumnList{NameColumn, GrantedPowersColumn, Spell1Column, Spell2Column, Spell3Column, Spell4Column, Spell5Column, Spell6Column, Spell7Column, Spell8Column, Spell9Column, FullTextColumn, ReferenceColumn}
	)

	return domainTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		Name:          NameColumn,
		GrantedPowers: GrantedPowersColumn,
		Spell1:        Spell1Column,
		Spell2:        Spell2Column,
		Spell3:        Spell3Column,
		Spell4:        Spell4Column,
		Spell5:        Spell5Column,
		Spell6:        Spell6Column,
		Spell7:        Spell7Column,
		Spell8:        Spell8Column,
		Spell9:        Spell9Column,
		FullText:      FullTextColumn,
		Reference:     ReferenceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
