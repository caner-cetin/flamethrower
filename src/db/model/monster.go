//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Monster struct {
	ID               int32 `sql:"primary_key"`
	Family           *string
	Name             string
	Altname          *string
	Size             *string
	Type             *string
	Descriptor       *string
	HitDice          *string
	Initiative       *string
	Speed            *string
	ArmorClass       *string
	BaseAttack       *string
	Grapple          *string
	Attack           *string
	FullAttack       *string
	Space            *string
	Reach            *string
	SpecialAttacks   *string
	SpecialQualities *string
	Saves            *string
	Abilities        *string
	Skills           *string
	BonusFeats       *string
	Feats            *string
	EpicFeats        *string
	Environment      *string
	Organization     *string
	ChallengeRating  *string
	Treasure         *string
	Alignment        *string
	Advancement      *string
	LevelAdjustment  *string
	SpecialAbilities *string
	StatBlock        *string
	FullText         *string
	Reference        *string
}
