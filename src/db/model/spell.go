//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Spell struct {
	ID                       int32 `sql:"primary_key"`
	Name                     string
	Altname                  *string
	School                   *string
	Subschool                *string
	Descriptor               *string
	SpellcraftDc             *string
	Level                    *string
	Components               *string
	CastingTime              *string
	Range                    *string
	Target                   *string
	Area                     *string
	Effect                   *string
	Duration                 *string
	SavingThrow              *string
	SpellResistance          *string
	ShortDescription         *string
	ToDevelop                *string
	MaterialComponents       *string
	ArcaneMaterialComponents *string
	Focus                    *string
	Description              *string
	XpCost                   *string
	ArcaneFocus              *string
	WizardFocus              *string
	VerbalComponents         *string
	SorcererFocus            *string
	BardFocus                *string
	ClericFocus              *string
	DruidFocus               *string
	FullText                 *string
	Reference                *string
}
