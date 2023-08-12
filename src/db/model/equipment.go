//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Equipment struct {
	ID                       int32 `sql:"primary_key"`
	Name                     string
	Family                   *string
	Category                 *string
	Subcategory              *string
	Cost                     *string
	DmgS                     *string
	ArmorShieldBonus         *string
	MaximumDexBonus          *string
	DmgM                     *string
	Weight                   *string
	Critical                 *string
	ArmorCheckPenalty        *string
	ArcaneSpellFailureChance *string
	RangeIncrement           *string
	Speed30                  *string
	Type                     *string
	Speed20                  *string
	FullText                 *string
	Reference                *string
}