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

var Equipment = newEquipmentTable("", "equipment", "")

type equipmentTable struct {
	sqlite.Table

	// Columns
	ID                       sqlite.ColumnInteger
	Name                     sqlite.ColumnString
	Family                   sqlite.ColumnString
	Category                 sqlite.ColumnString
	Subcategory              sqlite.ColumnString
	Cost                     sqlite.ColumnString
	DmgS                     sqlite.ColumnString
	ArmorShieldBonus         sqlite.ColumnString
	MaximumDexBonus          sqlite.ColumnString
	DmgM                     sqlite.ColumnString
	Weight                   sqlite.ColumnString
	Critical                 sqlite.ColumnString
	ArmorCheckPenalty        sqlite.ColumnString
	ArcaneSpellFailureChance sqlite.ColumnString
	RangeIncrement           sqlite.ColumnString
	Speed30                  sqlite.ColumnString
	Type                     sqlite.ColumnString
	Speed20                  sqlite.ColumnString
	FullText                 sqlite.ColumnString
	Reference                sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type EquipmentTable struct {
	equipmentTable

	EXCLUDED equipmentTable
}

// AS creates new EquipmentTable with assigned alias
func (a EquipmentTable) AS(alias string) *EquipmentTable {
	return newEquipmentTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EquipmentTable with assigned schema name
func (a EquipmentTable) FromSchema(schemaName string) *EquipmentTable {
	return newEquipmentTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EquipmentTable with assigned table prefix
func (a EquipmentTable) WithPrefix(prefix string) *EquipmentTable {
	return newEquipmentTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EquipmentTable with assigned table suffix
func (a EquipmentTable) WithSuffix(suffix string) *EquipmentTable {
	return newEquipmentTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEquipmentTable(schemaName, tableName, alias string) *EquipmentTable {
	return &EquipmentTable{
		equipmentTable: newEquipmentTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newEquipmentTableImpl("", "excluded", ""),
	}
}

func newEquipmentTableImpl(schemaName, tableName, alias string) equipmentTable {
	var (
		IDColumn                       = sqlite.IntegerColumn("id")
		NameColumn                     = sqlite.StringColumn("name")
		FamilyColumn                   = sqlite.StringColumn("family")
		CategoryColumn                 = sqlite.StringColumn("category")
		SubcategoryColumn              = sqlite.StringColumn("subcategory")
		CostColumn                     = sqlite.StringColumn("cost")
		DmgSColumn                     = sqlite.StringColumn("dmg_s")
		ArmorShieldBonusColumn         = sqlite.StringColumn("armor_shield_bonus")
		MaximumDexBonusColumn          = sqlite.StringColumn("maximum_dex_bonus")
		DmgMColumn                     = sqlite.StringColumn("dmg_m")
		WeightColumn                   = sqlite.StringColumn("weight")
		CriticalColumn                 = sqlite.StringColumn("critical")
		ArmorCheckPenaltyColumn        = sqlite.StringColumn("armor_check_penalty")
		ArcaneSpellFailureChanceColumn = sqlite.StringColumn("arcane_spell_failure_chance")
		RangeIncrementColumn           = sqlite.StringColumn("range_increment")
		Speed30Column                  = sqlite.StringColumn("speed_30")
		TypeColumn                     = sqlite.StringColumn("type")
		Speed20Column                  = sqlite.StringColumn("speed_20")
		FullTextColumn                 = sqlite.StringColumn("full_text")
		ReferenceColumn                = sqlite.StringColumn("reference")
		allColumns                     = sqlite.ColumnList{IDColumn, NameColumn, FamilyColumn, CategoryColumn, SubcategoryColumn, CostColumn, DmgSColumn, ArmorShieldBonusColumn, MaximumDexBonusColumn, DmgMColumn, WeightColumn, CriticalColumn, ArmorCheckPenaltyColumn, ArcaneSpellFailureChanceColumn, RangeIncrementColumn, Speed30Column, TypeColumn, Speed20Column, FullTextColumn, ReferenceColumn}
		mutableColumns                 = sqlite.ColumnList{NameColumn, FamilyColumn, CategoryColumn, SubcategoryColumn, CostColumn, DmgSColumn, ArmorShieldBonusColumn, MaximumDexBonusColumn, DmgMColumn, WeightColumn, CriticalColumn, ArmorCheckPenaltyColumn, ArcaneSpellFailureChanceColumn, RangeIncrementColumn, Speed30Column, TypeColumn, Speed20Column, FullTextColumn, ReferenceColumn}
	)

	return equipmentTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                       IDColumn,
		Name:                     NameColumn,
		Family:                   FamilyColumn,
		Category:                 CategoryColumn,
		Subcategory:              SubcategoryColumn,
		Cost:                     CostColumn,
		DmgS:                     DmgSColumn,
		ArmorShieldBonus:         ArmorShieldBonusColumn,
		MaximumDexBonus:          MaximumDexBonusColumn,
		DmgM:                     DmgMColumn,
		Weight:                   WeightColumn,
		Critical:                 CriticalColumn,
		ArmorCheckPenalty:        ArmorCheckPenaltyColumn,
		ArcaneSpellFailureChance: ArcaneSpellFailureChanceColumn,
		RangeIncrement:           RangeIncrementColumn,
		Speed30:                  Speed30Column,
		Type:                     TypeColumn,
		Speed20:                  Speed20Column,
		FullText:                 FullTextColumn,
		Reference:                ReferenceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
