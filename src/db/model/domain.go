//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Domain struct {
	ID            int32 `sql:"primary_key"`
	Name          string
	GrantedPowers *string
	Spell1        *string
	Spell2        *string
	Spell3        *string
	Spell4        *string
	Spell5        *string
	Spell6        *string
	Spell7        *string
	Spell8        *string
	Spell9        *string
	FullText      *string
	Reference     *string
}
