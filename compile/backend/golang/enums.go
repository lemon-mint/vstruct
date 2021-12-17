package golang

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeEnums(w io.Writer, i *ir.IR) {
	for _, e := range i.Enums {
		fmt.Fprintf(w, "type %s int%d\n", NameConv(e.Name), e.Size*8)
		fmt.Fprintf(w, "const (\n")
		for i, o := range e.Options {
			fmt.Fprintf(w, "%s_%s %s = %d\n", NameConv(e.Name), NameConv(o), NameConv(e.Name), i)
		}
		fmt.Fprintf(w, ")\n\n")

		fmt.Fprintf(w, "func (e %s) String() string {\n", NameConv(e.Name))
		fmt.Fprintf(w, "switch e {\n")
		for _, o := range e.Options {
			fmt.Fprintf(w, "case %s_%s:\n", NameConv(e.Name), NameConv(o))
			fmt.Fprintf(w, "return \"%s\"\n", NameConv(o))
		}
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "return \"\"\n")
		fmt.Fprintf(w, "}\n\n")

		fmt.Fprintf(w, "func (e %s) Match(\n", NameConv(e.Name))
		for _, o := range e.Options {
			fmt.Fprintf(w, "on%s func(),\n", NameConv(o))
		}
		fmt.Fprintf(w, ") {\n")
		fmt.Fprintf(w, "switch e {\n")
		for _, o := range e.Options {
			fmt.Fprintf(w, "case %s_%s:\n", NameConv(e.Name), NameConv(o))
			fmt.Fprintf(w, "on%s()\n", NameConv(o))
		}
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "}\n\n")
	}
}

type Speices int8

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

type ItemType int8

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
