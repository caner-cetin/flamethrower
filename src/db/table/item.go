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

var Item = newItemTable("", "item", "")

type itemTable struct {
	sqlite.Table

	// Columns
	ID              sqlite.ColumnInteger
	Name            sqlite.ColumnString
	Category        sqlite.ColumnString
	Subcategory     sqlite.ColumnString
	SpecialAbility  sqlite.ColumnString
	Aura            sqlite.ColumnString
	CasterLevel     sqlite.ColumnString
	Price           sqlite.ColumnString
	ManifesterLevel sqlite.ColumnString
	Prereq          sqlite.ColumnString
	Cost            sqlite.ColumnString
	Weight          sqlite.ColumnString
	FullText        sqlite.ColumnString
	Reference       sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type ItemTable struct {
	itemTable

	EXCLUDED itemTable
}

// AS creates new ItemTable with assigned alias
func (a ItemTable) AS(alias string) *ItemTable {
	return newItemTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ItemTable with assigned schema name
func (a ItemTable) FromSchema(schemaName string) *ItemTable {
	return newItemTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ItemTable with assigned table prefix
func (a ItemTable) WithPrefix(prefix string) *ItemTable {
	return newItemTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ItemTable with assigned table suffix
func (a ItemTable) WithSuffix(suffix string) *ItemTable {
	return newItemTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newItemTable(schemaName, tableName, alias string) *ItemTable {
	return &ItemTable{
		itemTable: newItemTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newItemTableImpl("", "excluded", ""),
	}
}

func newItemTableImpl(schemaName, tableName, alias string) itemTable {
	var (
		IDColumn              = sqlite.IntegerColumn("id")
		NameColumn            = sqlite.StringColumn("name")
		CategoryColumn        = sqlite.StringColumn("category")
		SubcategoryColumn     = sqlite.StringColumn("subcategory")
		SpecialAbilityColumn  = sqlite.StringColumn("special_ability")
		AuraColumn            = sqlite.StringColumn("aura")
		CasterLevelColumn     = sqlite.StringColumn("caster_level")
		PriceColumn           = sqlite.StringColumn("price")
		ManifesterLevelColumn = sqlite.StringColumn("manifester_level")
		PrereqColumn          = sqlite.StringColumn("prereq")
		CostColumn            = sqlite.StringColumn("cost")
		WeightColumn          = sqlite.StringColumn("weight")
		FullTextColumn        = sqlite.StringColumn("full_text")
		ReferenceColumn       = sqlite.StringColumn("reference")
		allColumns            = sqlite.ColumnList{IDColumn, NameColumn, CategoryColumn, SubcategoryColumn, SpecialAbilityColumn, AuraColumn, CasterLevelColumn, PriceColumn, ManifesterLevelColumn, PrereqColumn, CostColumn, WeightColumn, FullTextColumn, ReferenceColumn}
		mutableColumns        = sqlite.ColumnList{NameColumn, CategoryColumn, SubcategoryColumn, SpecialAbilityColumn, AuraColumn, CasterLevelColumn, PriceColumn, ManifesterLevelColumn, PrereqColumn, CostColumn, WeightColumn, FullTextColumn, ReferenceColumn}
	)

	return itemTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		Name:            NameColumn,
		Category:        CategoryColumn,
		Subcategory:     SubcategoryColumn,
		SpecialAbility:  SpecialAbilityColumn,
		Aura:            AuraColumn,
		CasterLevel:     CasterLevelColumn,
		Price:           PriceColumn,
		ManifesterLevel: ManifesterLevelColumn,
		Prereq:          PrereqColumn,
		Cost:            CostColumn,
		Weight:          WeightColumn,
		FullText:        FullTextColumn,
		Reference:       ReferenceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
