package main

import (
	"math"
)

var _ = math.Float32frombits
var _ = math.Float64frombits

type Speices uint8

const (
	Speices_Human      Speices = 0
	Speices_Elf        Speices = 1
	Speices_Orc        Speices = 2
	Speices_Dwarf      Speices = 3
	Speices_Gnome      Speices = 4
	Speices_Halfling   Speices = 5
	Speices_HalfElf    Speices = 6
	Speices_HalfOrc    Speices = 7
	Speices_Dragonborn Speices = 8
	Speices_Tiefling   Speices = 9
	Speices_Gnoll      Speices = 10
	Speices_Goblin     Speices = 11
)

func (e Speices) String() string {
	switch e {
	case Speices_Human:
		return "Human"
	case Speices_Elf:
		return "Elf"
	case Speices_Orc:
		return "Orc"
	case Speices_Dwarf:
		return "Dwarf"
	case Speices_Gnome:
		return "Gnome"
	case Speices_Halfling:
		return "Halfling"
	case Speices_HalfElf:
		return "HalfElf"
	case Speices_HalfOrc:
		return "HalfOrc"
	case Speices_Dragonborn:
		return "Dragonborn"
	case Speices_Tiefling:
		return "Tiefling"
	case Speices_Gnoll:
		return "Gnoll"
	case Speices_Goblin:
		return "Goblin"
	}
	return ""
}

func (e Speices) Match(
	onHuman func(),
	onElf func(),
	onOrc func(),
	onDwarf func(),
	onGnome func(),
	onHalfling func(),
	onHalfElf func(),
	onHalfOrc func(),
	onDragonborn func(),
	onTiefling func(),
	onGnoll func(),
	onGoblin func(),
) {
	switch e {
	case Speices_Human:
		onHuman()
	case Speices_Elf:
		onElf()
	case Speices_Orc:
		onOrc()
	case Speices_Dwarf:
		onDwarf()
	case Speices_Gnome:
		onGnome()
	case Speices_Halfling:
		onHalfling()
	case Speices_HalfElf:
		onHalfElf()
	case Speices_HalfOrc:
		onHalfOrc()
	case Speices_Dragonborn:
		onDragonborn()
	case Speices_Tiefling:
		onTiefling()
	case Speices_Gnoll:
		onGnoll()
	case Speices_Goblin:
		onGoblin()
	}
}

type ItemType uint8

const (
	ItemType_Weapon ItemType = 0
	ItemType_Armor  ItemType = 1
	ItemType_Potion ItemType = 2
)

func (e ItemType) String() string {
	switch e {
	case ItemType_Weapon:
		return "Weapon"
	case ItemType_Armor:
		return "Armor"
	case ItemType_Potion:
		return "Potion"
	}
	return ""
}

func (e ItemType) Match(
	onWeapon func(),
	onArmor func(),
	onPotion func(),
) {
	switch e {
	case ItemType_Weapon:
		onWeapon()
	case ItemType_Armor:
		onArmor()
	case ItemType_Potion:
		onPotion()
	}
}

type Coordinate []byte

func (s Coordinate) X() int64 {
	_ = s[7]
	var __v uint64 = uint64(s[0]) |
		uint64(s[1])<<8 |
		uint64(s[2])<<16 |
		uint64(s[3])<<24 |
		uint64(s[4])<<32 |
		uint64(s[5])<<40 |
		uint64(s[6])<<48 |
		uint64(s[7])<<56
	return int64(__v)
}

func (s Coordinate) Y() int64 {
	_ = s[15]
	var __v uint64 = uint64(s[8]) |
		uint64(s[9])<<8 |
		uint64(s[10])<<16 |
		uint64(s[11])<<24 |
		uint64(s[12])<<32 |
		uint64(s[13])<<40 |
		uint64(s[14])<<48 |
		uint64(s[15])<<56
	return int64(__v)
}

type Item []byte

func (s Item) Type() ItemType {
	return ItemType(s[0])
}

func (s Item) Damage() int64 {
	_ = s[8]
	var __v uint64 = uint64(s[1]) |
		uint64(s[2])<<8 |
		uint64(s[3])<<16 |
		uint64(s[4])<<24 |
		uint64(s[5])<<32 |
		uint64(s[6])<<40 |
		uint64(s[7])<<48 |
		uint64(s[8])<<56
	return int64(__v)
}

func (s Item) Armor() int64 {
	_ = s[16]
	var __v uint64 = uint64(s[9]) |
		uint64(s[10])<<8 |
		uint64(s[11])<<16 |
		uint64(s[12])<<24 |
		uint64(s[13])<<32 |
		uint64(s[14])<<40 |
		uint64(s[15])<<48 |
		uint64(s[16])<<56
	return int64(__v)
}

func (s Item) Name() string {
	_ = s[40]
	var __off0 uint64 = uint64(s[25]) |
		uint64(s[26])<<8 |
		uint64(s[27])<<16 |
		uint64(s[28])<<24 |
		uint64(s[29])<<32 |
		uint64(s[30])<<40 |
		uint64(s[31])<<48 |
		uint64(s[32])<<56
	var __off1 uint64 = uint64(s[33]) |
		uint64(s[34])<<8 |
		uint64(s[35])<<16 |
		uint64(s[36])<<24 |
		uint64(s[37])<<32 |
		uint64(s[38])<<40 |
		uint64(s[39])<<48 |
		uint64(s[40])<<56
	return string(s[__off0:__off1])
}

type Inventory []byte

func (s Inventory) RightHand() Item {
	_ = s[15]
	var __off0 uint64 = uint64(s[0]) |
		uint64(s[1])<<8 |
		uint64(s[2])<<16 |
		uint64(s[3])<<24 |
		uint64(s[4])<<32 |
		uint64(s[5])<<40 |
		uint64(s[6])<<48 |
		uint64(s[7])<<56
	var __off1 uint64 = uint64(s[8]) |
		uint64(s[9])<<8 |
		uint64(s[10])<<16 |
		uint64(s[11])<<24 |
		uint64(s[12])<<32 |
		uint64(s[13])<<40 |
		uint64(s[14])<<48 |
		uint64(s[15])<<56
	return Item(s[__off0:__off1])
}

func (s Inventory) LeftHand() Item {
	_ = s[23]
	var __off0 uint64 = uint64(s[8]) |
		uint64(s[9])<<8 |
		uint64(s[10])<<16 |
		uint64(s[11])<<24 |
		uint64(s[12])<<32 |
		uint64(s[13])<<40 |
		uint64(s[14])<<48 |
		uint64(s[15])<<56
	var __off1 uint64 = uint64(s[16]) |
		uint64(s[17])<<8 |
		uint64(s[18])<<16 |
		uint64(s[19])<<24 |
		uint64(s[20])<<32 |
		uint64(s[21])<<40 |
		uint64(s[22])<<48 |
		uint64(s[23])<<56
	return Item(s[__off0:__off1])
}

type Entity []byte

func (s Entity) Type() Speices {
	return Speices(s[0])
}

func (s Entity) Position() Coordinate {
	return Coordinate(s[1:17])
}

func (s Entity) Hp() int64 {
	_ = s[24]
	var __v uint64 = uint64(s[17]) |
		uint64(s[18])<<8 |
		uint64(s[19])<<16 |
		uint64(s[20])<<24 |
		uint64(s[21])<<32 |
		uint64(s[22])<<40 |
		uint64(s[23])<<48 |
		uint64(s[24])<<56
	return int64(__v)
}

func (s Entity) Id() UUID {
	_ = s[40]
	var __off0 uint64 = uint64(s[25]) |
		uint64(s[26])<<8 |
		uint64(s[27])<<16 |
		uint64(s[28])<<24 |
		uint64(s[29])<<32 |
		uint64(s[30])<<40 |
		uint64(s[31])<<48 |
		uint64(s[32])<<56
	var __off1 uint64 = uint64(s[33]) |
		uint64(s[34])<<8 |
		uint64(s[35])<<16 |
		uint64(s[36])<<24 |
		uint64(s[37])<<32 |
		uint64(s[38])<<40 |
		uint64(s[39])<<48 |
		uint64(s[40])<<56
	return UUID(s[__off0:__off1])
}

func (s Entity) Inventory() Inventory {
	_ = s[48]
	var __off0 uint64 = uint64(s[33]) |
		uint64(s[34])<<8 |
		uint64(s[35])<<16 |
		uint64(s[36])<<24 |
		uint64(s[37])<<32 |
		uint64(s[38])<<40 |
		uint64(s[39])<<48 |
		uint64(s[40])<<56
	var __off1 uint64 = uint64(s[41]) |
		uint64(s[42])<<8 |
		uint64(s[43])<<16 |
		uint64(s[44])<<24 |
		uint64(s[45])<<32 |
		uint64(s[46])<<40 |
		uint64(s[47])<<48 |
		uint64(s[48])<<56
	return Inventory(s[__off0:__off1])
}

type UUID = string
