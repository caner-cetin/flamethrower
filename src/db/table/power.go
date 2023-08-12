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

var Power = newPowerTable("", "power", "")

type powerTable struct {
	sqlite.Table

	// Columns
	ID               sqlite.ColumnInteger
	Name             sqlite.ColumnString
	Discipline       sqlite.ColumnString
	Subdiscipline    sqlite.ColumnString
	Descriptor       sqlite.ColumnString
	Level            sqlite.ColumnString
	Display          sqlite.ColumnString
	ManifestingTime  sqlite.ColumnString
	Range            sqlite.ColumnString
	Target           sqlite.ColumnString
	Area             sqlite.ColumnString
	Effect           sqlite.ColumnString
	Duration         sqlite.ColumnString
	SavingThrow      sqlite.ColumnString
	PowerPoints      sqlite.ColumnString
	PowerResistance  sqlite.ColumnString
	ShortDescription sqlite.ColumnString
	XpCost           sqlite.ColumnString
	Description      sqlite.ColumnString
	Augment          sqlite.ColumnString
	FullText         sqlite.ColumnString
	Reference        sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type PowerTable struct {
	powerTable

	EXCLUDED powerTable
}

// AS creates new PowerTable with assigned alias
func (a PowerTable) AS(alias string) *PowerTable {
	return newPowerTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PowerTable with assigned schema name
func (a PowerTable) FromSchema(schemaName string) *PowerTable {
	return newPowerTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PowerTable with assigned table prefix
func (a PowerTable) WithPrefix(prefix string) *PowerTable {
	return newPowerTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PowerTable with assigned table suffix
func (a PowerTable) WithSuffix(suffix string) *PowerTable {
	return newPowerTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPowerTable(schemaName, tableName, alias string) *PowerTable {
	return &PowerTable{
		powerTable: newPowerTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newPowerTableImpl("", "excluded", ""),
	}
}

func newPowerTableImpl(schemaName, tableName, alias string) powerTable {
	var (
		IDColumn               = sqlite.IntegerColumn("id")
		NameColumn             = sqlite.StringColumn("name")
		DisciplineColumn       = sqlite.StringColumn("discipline")
		SubdisciplineColumn    = sqlite.StringColumn("subdiscipline")
		DescriptorColumn       = sqlite.StringColumn("descriptor")
		LevelColumn            = sqlite.StringColumn("level")
		DisplayColumn          = sqlite.StringColumn("display")
		ManifestingTimeColumn  = sqlite.StringColumn("manifesting_time")
		RangeColumn            = sqlite.StringColumn("range")
		TargetColumn           = sqlite.StringColumn("target")
		AreaColumn             = sqlite.StringColumn("area")
		EffectColumn           = sqlite.StringColumn("effect")
		DurationColumn         = sqlite.StringColumn("duration")
		SavingThrowColumn      = sqlite.StringColumn("saving_throw")
		PowerPointsColumn      = sqlite.StringColumn("power_points")
		PowerResistanceColumn  = sqlite.StringColumn("power_resistance")
		ShortDescriptionColumn = sqlite.StringColumn("short_description")
		XpCostColumn           = sqlite.StringColumn("xp_cost")
		DescriptionColumn      = sqlite.StringColumn("description")
		AugmentColumn          = sqlite.StringColumn("augment")
		FullTextColumn         = sqlite.StringColumn("full_text")
		ReferenceColumn        = sqlite.StringColumn("reference")
		allColumns             = sqlite.ColumnList{IDColumn, NameColumn, DisciplineColumn, SubdisciplineColumn, DescriptorColumn, LevelColumn, DisplayColumn, ManifestingTimeColumn, RangeColumn, TargetColumn, AreaColumn, EffectColumn, DurationColumn, SavingThrowColumn, PowerPointsColumn, PowerResistanceColumn, ShortDescriptionColumn, XpCostColumn, DescriptionColumn, AugmentColumn, FullTextColumn, ReferenceColumn}
		mutableColumns         = sqlite.ColumnList{NameColumn, DisciplineColumn, SubdisciplineColumn, DescriptorColumn, LevelColumn, DisplayColumn, ManifestingTimeColumn, RangeColumn, TargetColumn, AreaColumn, EffectColumn, DurationColumn, SavingThrowColumn, PowerPointsColumn, PowerResistanceColumn, ShortDescriptionColumn, XpCostColumn, DescriptionColumn, AugmentColumn, FullTextColumn, ReferenceColumn}
	)

	return powerTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		Name:             NameColumn,
		Discipline:       DisciplineColumn,
		Subdiscipline:    SubdisciplineColumn,
		Descriptor:       DescriptorColumn,
		Level:            LevelColumn,
		Display:          DisplayColumn,
		ManifestingTime:  ManifestingTimeColumn,
		Range:            RangeColumn,
		Target:           TargetColumn,
		Area:             AreaColumn,
		Effect:           EffectColumn,
		Duration:         DurationColumn,
		SavingThrow:      SavingThrowColumn,
		PowerPoints:      PowerPointsColumn,
		PowerResistance:  PowerResistanceColumn,
		ShortDescription: ShortDescriptionColumn,
		XpCost:           XpCostColumn,
		Description:      DescriptionColumn,
		Augment:          AugmentColumn,
		FullText:         FullTextColumn,
		Reference:        ReferenceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
