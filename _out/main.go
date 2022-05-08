package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

type _ = strings.Builder
type _ = unsafe.Pointer

var _ = math.Float32frombits
var _ = math.Float64frombits
var _ = strconv.FormatInt
var _ = strconv.FormatUint
var _ = strconv.FormatFloat
var _ = fmt.Sprint

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

func (e Speices) MatchS(s struct {
	onHuman      func()
	onElf        func()
	onOrc        func()
	onDwarf      func()
	onGnome      func()
	onHalfling   func()
	onHalfElf    func()
	onHalfOrc    func()
	onDragonborn func()
	onTiefling   func()
	onGnoll      func()
	onGoblin     func()
}) {
	switch e {
	case Speices_Human:
		s.onHuman()
	case Speices_Elf:
		s.onElf()
	case Speices_Orc:
		s.onOrc()
	case Speices_Dwarf:
		s.onDwarf()
	case Speices_Gnome:
		s.onGnome()
	case Speices_Halfling:
		s.onHalfling()
	case Speices_HalfElf:
		s.onHalfElf()
	case Speices_HalfOrc:
		s.onHalfOrc()
	case Speices_Dragonborn:
		s.onDragonborn()
	case Speices_Tiefling:
		s.onTiefling()
	case Speices_Gnoll:
		s.onGnoll()
	case Speices_Goblin:
		s.onGoblin()
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

func (e ItemType) MatchS(s struct {
	onWeapon func()
	onArmor  func()
	onPotion func()
}) {
	switch e {
	case ItemType_Weapon:
		s.onWeapon()
	case ItemType_Armor:
		s.onArmor()
	case ItemType_Potion:
		s.onPotion()
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

func (s Coordinate) Vstruct_Validate() bool {
	return len(s) >= 16
}

func (s Coordinate) String() string {
	if !s.Vstruct_Validate() {
		return "Coordinate (invalid)"
	}
	var __b strings.Builder
	__b.WriteString("Coordinate {")
	__b.WriteString("X: ")
	__b.WriteString(strconv.FormatInt(int64(s.X()), 10))
	__b.WriteString(", ")
	__b.WriteString("Y: ")
	__b.WriteString(strconv.FormatInt(int64(s.Y()), 10))
	__b.WriteString("}")
	return __b.String()
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
	_ = s[24]
	var __off0 uint64 = 25
	var __off1 uint64 = uint64(s[17]) |
		uint64(s[18])<<8 |
		uint64(s[19])<<16 |
		uint64(s[20])<<24 |
		uint64(s[21])<<32 |
		uint64(s[22])<<40 |
		uint64(s[23])<<48 |
		uint64(s[24])<<56
	var __v = s[__off0:__off1]

	return *(*string)(unsafe.Pointer(&__v))
}

func (s Item) Vstruct_Validate() bool {
	if len(s) < 25 {
		return false
	}

	_ = s[24]

	var __off0 uint64 = 25
	var __off1 uint64 = uint64(s[17]) |
		uint64(s[18])<<8 |
		uint64(s[19])<<16 |
		uint64(s[20])<<24 |
		uint64(s[21])<<32 |
		uint64(s[22])<<40 |
		uint64(s[23])<<48 |
		uint64(s[24])<<56
	var __off2 uint64 = uint64(len(s))
	return __off0 <= __off1 && __off1 <= __off2
}

func (s Item) String() string {
	if !s.Vstruct_Validate() {
		return "Item (invalid)"
	}
	var __b strings.Builder
	__b.WriteString("Item {")
	__b.WriteString("Type: ")
	__b.WriteString(s.Type().String())
	__b.WriteString(", ")
	__b.WriteString("Damage: ")
	__b.WriteString(strconv.FormatInt(int64(s.Damage()), 10))
	__b.WriteString(", ")
	__b.WriteString("Armor: ")
	__b.WriteString(strconv.FormatInt(int64(s.Armor()), 10))
	__b.WriteString(", ")
	__b.WriteString("Name: ")
	__b.WriteString(strconv.Quote(s.Name()))
	__b.WriteString("}")
	return __b.String()
}

type Inventory []byte

func (s Inventory) RightHand() Item {
	_ = s[7]
	var __off0 uint64 = 16
	var __off1 uint64 = uint64(s[0]) |
		uint64(s[1])<<8 |
		uint64(s[2])<<16 |
		uint64(s[3])<<24 |
		uint64(s[4])<<32 |
		uint64(s[5])<<40 |
		uint64(s[6])<<48 |
		uint64(s[7])<<56
	return Item(s[__off0:__off1])
}

func (s Inventory) LeftHand() Item {
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

func (s Inventory) Vstruct_Validate() bool {
	if len(s) < 16 {
		return false
	}

	_ = s[15]

	var __off0 uint64 = 16
	var __off1 uint64 = uint64(s[0]) |
		uint64(s[1])<<8 |
		uint64(s[2])<<16 |
		uint64(s[3])<<24 |
		uint64(s[4])<<32 |
		uint64(s[5])<<40 |
		uint64(s[6])<<48 |
		uint64(s[7])<<56
	var __off2 uint64 = uint64(s[8]) |
		uint64(s[9])<<8 |
		uint64(s[10])<<16 |
		uint64(s[11])<<24 |
		uint64(s[12])<<32 |
		uint64(s[13])<<40 |
		uint64(s[14])<<48 |
		uint64(s[15])<<56
	var __off3 uint64 = uint64(len(s))
	if __off0 <= __off1 && __off1 <= __off2 && __off2 <= __off3 {
		return s.RightHand().Vstruct_Validate() && s.LeftHand().Vstruct_Validate()
	}

	return false
}

func (s Inventory) String() string {
	if !s.Vstruct_Validate() {
		return "Inventory (invalid)"
	}
	var __b strings.Builder
	__b.WriteString("Inventory {")
	__b.WriteString("RightHand: ")
	__b.WriteString(s.RightHand().String())
	__b.WriteString(", ")
	__b.WriteString("LeftHand: ")
	__b.WriteString(s.LeftHand().String())
	__b.WriteString("}")
	return __b.String()
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
	_ = s[32]
	var __off0 uint64 = 41
	var __off1 uint64 = uint64(s[25]) |
		uint64(s[26])<<8 |
		uint64(s[27])<<16 |
		uint64(s[28])<<24 |
		uint64(s[29])<<32 |
		uint64(s[30])<<40 |
		uint64(s[31])<<48 |
		uint64(s[32])<<56
	var __v = s[__off0:__off1]

	return *(*UUID)(unsafe.Pointer(&__v))
}

func (s Entity) Inventory() Inventory {
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
	return Inventory(s[__off0:__off1])
}

func (s Entity) Vstruct_Validate() bool {
	if len(s) < 41 {
		return false
	}

	_ = s[40]

	var __off0 uint64 = 41
	var __off1 uint64 = uint64(s[25]) |
		uint64(s[26])<<8 |
		uint64(s[27])<<16 |
		uint64(s[28])<<24 |
		uint64(s[29])<<32 |
		uint64(s[30])<<40 |
		uint64(s[31])<<48 |
		uint64(s[32])<<56
	var __off2 uint64 = uint64(s[33]) |
		uint64(s[34])<<8 |
		uint64(s[35])<<16 |
		uint64(s[36])<<24 |
		uint64(s[37])<<32 |
		uint64(s[38])<<40 |
		uint64(s[39])<<48 |
		uint64(s[40])<<56
	var __off3 uint64 = uint64(len(s))
	if __off0 <= __off1 && __off1 <= __off2 && __off2 <= __off3 {
		return s.Inventory().Vstruct_Validate()
	}

	return false
}

func (s Entity) String() string {
	if !s.Vstruct_Validate() {
		return "Entity (invalid)"
	}
	var __b strings.Builder
	__b.WriteString("Entity {")
	__b.WriteString("Type: ")
	__b.WriteString(s.Type().String())
	__b.WriteString(", ")
	__b.WriteString("Position: ")
	__b.WriteString(s.Position().String())
	__b.WriteString(", ")
	__b.WriteString("Hp: ")
	__b.WriteString(strconv.FormatInt(int64(s.Hp()), 10))
	__b.WriteString(", ")
	__b.WriteString("Id: ")
	__b.WriteString(strconv.Quote(string(s.Id())))
	__b.WriteString(", ")
	__b.WriteString("Inventory: ")
	__b.WriteString(s.Inventory().String())
	__b.WriteString("}")
	return __b.String()
}

func Serialize_Coordinate(dst Coordinate, X int64, Y int64) Coordinate {
	_ = dst[15]
	var __tmp_0 = uint64(X)
	dst[0] = byte(__tmp_0)
	dst[1] = byte(__tmp_0 >> 8)
	dst[2] = byte(__tmp_0 >> 16)
	dst[3] = byte(__tmp_0 >> 24)
	dst[4] = byte(__tmp_0 >> 32)
	dst[5] = byte(__tmp_0 >> 40)
	dst[6] = byte(__tmp_0 >> 48)
	dst[7] = byte(__tmp_0 >> 56)
	var __tmp_1 = uint64(Y)
	dst[8] = byte(__tmp_1)
	dst[9] = byte(__tmp_1 >> 8)
	dst[10] = byte(__tmp_1 >> 16)
	dst[11] = byte(__tmp_1 >> 24)
	dst[12] = byte(__tmp_1 >> 32)
	dst[13] = byte(__tmp_1 >> 40)
	dst[14] = byte(__tmp_1 >> 48)
	dst[15] = byte(__tmp_1 >> 56)

	return dst
}

func New_Coordinate(X int64, Y int64) Coordinate {
	var __vstruct__size = 16
	var __vstruct__buf = make(Coordinate, __vstruct__size)
	__vstruct__buf = Serialize_Coordinate(__vstruct__buf, X, Y)
	return __vstruct__buf
}

func Serialize_Item(dst Item, Type ItemType, Damage int64, Armor int64, Name string) Item {
	_ = dst[24]
	dst[0] = byte(Type)
	var __tmp_1 = uint64(Damage)
	dst[1] = byte(__tmp_1)
	dst[2] = byte(__tmp_1 >> 8)
	dst[3] = byte(__tmp_1 >> 16)
	dst[4] = byte(__tmp_1 >> 24)
	dst[5] = byte(__tmp_1 >> 32)
	dst[6] = byte(__tmp_1 >> 40)
	dst[7] = byte(__tmp_1 >> 48)
	dst[8] = byte(__tmp_1 >> 56)
	var __tmp_2 = uint64(Armor)
	dst[9] = byte(__tmp_2)
	dst[10] = byte(__tmp_2 >> 8)
	dst[11] = byte(__tmp_2 >> 16)
	dst[12] = byte(__tmp_2 >> 24)
	dst[13] = byte(__tmp_2 >> 32)
	dst[14] = byte(__tmp_2 >> 40)
	dst[15] = byte(__tmp_2 >> 48)
	dst[16] = byte(__tmp_2 >> 56)

	var __index = uint64(25)
	__tmp_3 := uint64(len(Name)) + __index
	dst[17] = byte(__tmp_3)
	dst[18] = byte(__tmp_3 >> 8)
	dst[19] = byte(__tmp_3 >> 16)
	dst[20] = byte(__tmp_3 >> 24)
	dst[21] = byte(__tmp_3 >> 32)
	dst[22] = byte(__tmp_3 >> 40)
	dst[23] = byte(__tmp_3 >> 48)
	dst[24] = byte(__tmp_3 >> 56)
	copy(dst[__index:__tmp_3], Name)
	return dst
}

func New_Item(Type ItemType, Damage int64, Armor int64, Name string) Item {
	var __vstruct__size = 25 + len(Name)
	var __vstruct__buf = make(Item, __vstruct__size)
	__vstruct__buf = Serialize_Item(__vstruct__buf, Type, Damage, Armor, Name)
	return __vstruct__buf
}

func Serialize_Inventory(dst Inventory, RightHand Item, LeftHand Item) Inventory {
	_ = dst[15]

	var __index = uint64(16)
	__tmp_0 := uint64(len(RightHand)) + __index
	dst[0] = byte(__tmp_0)
	dst[1] = byte(__tmp_0 >> 8)
	dst[2] = byte(__tmp_0 >> 16)
	dst[3] = byte(__tmp_0 >> 24)
	dst[4] = byte(__tmp_0 >> 32)
	dst[5] = byte(__tmp_0 >> 40)
	dst[6] = byte(__tmp_0 >> 48)
	dst[7] = byte(__tmp_0 >> 56)
	copy(dst[__index:__tmp_0], RightHand)
	__index += uint64(len(RightHand))
	__tmp_1 := uint64(len(LeftHand)) + __index
	dst[8] = byte(__tmp_1)
	dst[9] = byte(__tmp_1 >> 8)
	dst[10] = byte(__tmp_1 >> 16)
	dst[11] = byte(__tmp_1 >> 24)
	dst[12] = byte(__tmp_1 >> 32)
	dst[13] = byte(__tmp_1 >> 40)
	dst[14] = byte(__tmp_1 >> 48)
	dst[15] = byte(__tmp_1 >> 56)
	copy(dst[__index:__tmp_1], LeftHand)
	return dst
}

func New_Inventory(RightHand Item, LeftHand Item) Inventory {
	var __vstruct__size = 16 + len(RightHand) + len(LeftHand)
	var __vstruct__buf = make(Inventory, __vstruct__size)
	__vstruct__buf = Serialize_Inventory(__vstruct__buf, RightHand, LeftHand)
	return __vstruct__buf
}

func Serialize_Entity(dst Entity, Type Speices, Position Coordinate, Hp int64, Id UUID, Inventory Inventory) Entity {
	_ = dst[40]
	dst[0] = byte(Type)
	copy(dst[1:17], Position)
	var __tmp_2 = uint64(Hp)
	dst[17] = byte(__tmp_2)
	dst[18] = byte(__tmp_2 >> 8)
	dst[19] = byte(__tmp_2 >> 16)
	dst[20] = byte(__tmp_2 >> 24)
	dst[21] = byte(__tmp_2 >> 32)
	dst[22] = byte(__tmp_2 >> 40)
	dst[23] = byte(__tmp_2 >> 48)
	dst[24] = byte(__tmp_2 >> 56)

	var __index = uint64(41)
	__tmp_3 := uint64(len(Id)) + __index
	dst[25] = byte(__tmp_3)
	dst[26] = byte(__tmp_3 >> 8)
	dst[27] = byte(__tmp_3 >> 16)
	dst[28] = byte(__tmp_3 >> 24)
	dst[29] = byte(__tmp_3 >> 32)
	dst[30] = byte(__tmp_3 >> 40)
	dst[31] = byte(__tmp_3 >> 48)
	dst[32] = byte(__tmp_3 >> 56)
	copy(dst[__index:__tmp_3], Id)
	__index += uint64(len(Id))
	__tmp_4 := uint64(len(Inventory)) + __index
	dst[33] = byte(__tmp_4)
	dst[34] = byte(__tmp_4 >> 8)
	dst[35] = byte(__tmp_4 >> 16)
	dst[36] = byte(__tmp_4 >> 24)
	dst[37] = byte(__tmp_4 >> 32)
	dst[38] = byte(__tmp_4 >> 40)
	dst[39] = byte(__tmp_4 >> 48)
	dst[40] = byte(__tmp_4 >> 56)
	copy(dst[__index:__tmp_4], Inventory)
	return dst
}

func New_Entity(Type Speices, Position Coordinate, Hp int64, Id UUID, Inventory Inventory) Entity {
	var __vstruct__size = 41 + len(Id) + len(Inventory)
	var __vstruct__buf = make(Entity, __vstruct__size)
	__vstruct__buf = Serialize_Entity(__vstruct__buf, Type, Position, Hp, Id, Inventory)
	return __vstruct__buf
}

type UUID = string
