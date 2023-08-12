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

var Feat = newFeatTable("", "feat", "")

type featTable struct {
	sqlite.Table

	// Columns
	ID           sqlite.ColumnInteger
	Name         sqlite.ColumnString
	Type         sqlite.ColumnString
	Multiple     sqlite.ColumnString
	Stack        sqlite.ColumnString
	Choice       sqlite.ColumnString
	Prerequisite sqlite.ColumnString
	Benefit      sqlite.ColumnString
	Normal       sqlite.ColumnString
	Special      sqlite.ColumnString
	FullText     sqlite.ColumnString
	Reference    sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type FeatTable struct {
	featTable

	EXCLUDED featTable
}

// AS creates new FeatTable with assigned alias
func (a FeatTable) AS(alias string) *FeatTable {
	return newFeatTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FeatTable with assigned schema name
func (a FeatTable) FromSchema(schemaName string) *FeatTable {
	return newFeatTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FeatTable with assigned table prefix
func (a FeatTable) WithPrefix(prefix string) *FeatTable {
	return newFeatTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FeatTable with assigned table suffix
func (a FeatTable) WithSuffix(suffix string) *FeatTable {
	return newFeatTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFeatTable(schemaName, tableName, alias string) *FeatTable {
	return &FeatTable{
		featTable: newFeatTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newFeatTableImpl("", "excluded", ""),
	}
}

func newFeatTableImpl(schemaName, tableName, alias string) featTable {
	var (
		IDColumn           = sqlite.IntegerColumn("id")
		NameColumn         = sqlite.StringColumn("name")
		TypeColumn         = sqlite.StringColumn("type")
		MultipleColumn     = sqlite.StringColumn("multiple")
		StackColumn        = sqlite.StringColumn("stack")
		ChoiceColumn       = sqlite.StringColumn("choice")
		PrerequisiteColumn = sqlite.StringColumn("prerequisite")
		BenefitColumn      = sqlite.StringColumn("benefit")
		NormalColumn       = sqlite.StringColumn("normal")
		SpecialColumn      = sqlite.StringColumn("special")
		FullTextColumn     = sqlite.StringColumn("full_text")
		ReferenceColumn    = sqlite.StringColumn("reference")
		allColumns         = sqlite.ColumnList{IDColumn, NameColumn, TypeColumn, MultipleColumn, StackColumn, ChoiceColumn, PrerequisiteColumn, BenefitColumn, NormalColumn, SpecialColumn, FullTextColumn, ReferenceColumn}
		mutableColumns     = sqlite.ColumnList{NameColumn, TypeColumn, MultipleColumn, StackColumn, ChoiceColumn, PrerequisiteColumn, BenefitColumn, NormalColumn, SpecialColumn, FullTextColumn, ReferenceColumn}
	)

	return featTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		Name:         NameColumn,
		Type:         TypeColumn,
		Multiple:     MultipleColumn,
		Stack:        StackColumn,
		Choice:       ChoiceColumn,
		Prerequisite: PrerequisiteColumn,
		Benefit:      BenefitColumn,
		Normal:       NormalColumn,
		Special:      SpecialColumn,
		FullText:     FullTextColumn,
		Reference:    ReferenceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
