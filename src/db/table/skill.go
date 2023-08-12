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

var Skill = newSkillTable("", "skill", "")

type skillTable struct {
	sqlite.Table

	// Columns
	ID          sqlite.ColumnInteger
	Name        sqlite.ColumnString
	Subtype     sqlite.ColumnString
	KeyAbility  sqlite.ColumnString
	Psionic     sqlite.ColumnString
	Trained     sqlite.ColumnString
	ArmorCheck  sqlite.ColumnString
	Description sqlite.ColumnString
	SkillCheck  sqlite.ColumnString
	Action      sqlite.ColumnString
	TryAgain    sqlite.ColumnString
	Special     sqlite.ColumnString
	Restriction sqlite.ColumnString
	Synergy     sqlite.ColumnString
	EpicUse     sqlite.ColumnString
	Untrained   sqlite.ColumnString
	FullText    sqlite.ColumnString
	Reference   sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type SkillTable struct {
	skillTable

	EXCLUDED skillTable
}

// AS creates new SkillTable with assigned alias
func (a SkillTable) AS(alias string) *SkillTable {
	return newSkillTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SkillTable with assigned schema name
func (a SkillTable) FromSchema(schemaName string) *SkillTable {
	return newSkillTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SkillTable with assigned table prefix
func (a SkillTable) WithPrefix(prefix string) *SkillTable {
	return newSkillTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SkillTable with assigned table suffix
func (a SkillTable) WithSuffix(suffix string) *SkillTable {
	return newSkillTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSkillTable(schemaName, tableName, alias string) *SkillTable {
	return &SkillTable{
		skillTable: newSkillTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newSkillTableImpl("", "excluded", ""),
	}
}

func newSkillTableImpl(schemaName, tableName, alias string) skillTable {
	var (
		IDColumn          = sqlite.IntegerColumn("id")
		NameColumn        = sqlite.StringColumn("name")
		SubtypeColumn     = sqlite.StringColumn("subtype")
		KeyAbilityColumn  = sqlite.StringColumn("key_ability")
		PsionicColumn     = sqlite.StringColumn("psionic")
		TrainedColumn     = sqlite.StringColumn("trained")
		ArmorCheckColumn  = sqlite.StringColumn("armor_check")
		DescriptionColumn = sqlite.StringColumn("description")
		SkillCheckColumn  = sqlite.StringColumn("skill_check")
		ActionColumn      = sqlite.StringColumn("action")
		TryAgainColumn    = sqlite.StringColumn("try_again")
		SpecialColumn     = sqlite.StringColumn("special")
		RestrictionColumn = sqlite.StringColumn("restriction")
		SynergyColumn     = sqlite.StringColumn("synergy")
		EpicUseColumn     = sqlite.StringColumn("epic_use")
		UntrainedColumn   = sqlite.StringColumn("untrained")
		FullTextColumn    = sqlite.StringColumn("full_text")
		ReferenceColumn   = sqlite.StringColumn("reference")
		allColumns        = sqlite.ColumnList{IDColumn, NameColumn, SubtypeColumn, KeyAbilityColumn, PsionicColumn, TrainedColumn, ArmorCheckColumn, DescriptionColumn, SkillCheckColumn, ActionColumn, TryAgainColumn, SpecialColumn, RestrictionColumn, SynergyColumn, EpicUseColumn, UntrainedColumn, FullTextColumn, ReferenceColumn}
		mutableColumns    = sqlite.ColumnList{NameColumn, SubtypeColumn, KeyAbilityColumn, PsionicColumn, TrainedColumn, ArmorCheckColumn, DescriptionColumn, SkillCheckColumn, ActionColumn, TryAgainColumn, SpecialColumn, RestrictionColumn, SynergyColumn, EpicUseColumn, UntrainedColumn, FullTextColumn, ReferenceColumn}
	)

	return skillTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		Subtype:     SubtypeColumn,
		KeyAbility:  KeyAbilityColumn,
		Psionic:     PsionicColumn,
		Trained:     TrainedColumn,
		ArmorCheck:  ArmorCheckColumn,
		Description: DescriptionColumn,
		SkillCheck:  SkillCheckColumn,
		Action:      ActionColumn,
		TryAgain:    TryAgainColumn,
		Special:     SpecialColumn,
		Restriction: RestrictionColumn,
		Synergy:     SynergyColumn,
		EpicUse:     EpicUseColumn,
		Untrained:   UntrainedColumn,
		FullText:    FullTextColumn,
		Reference:   ReferenceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}