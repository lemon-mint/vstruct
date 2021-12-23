import 'dart:typed_data';

// Code generated by vstruct. DO NOT EDIT.
// Package Name: main

enum Speices {
    Human,
    Elf,
    Orc,
    Dwarf,
    Gnome,
    Halfling,
    HalfElf,
    HalfOrc,
    Dragonborn,
    Tiefling,
    Gnoll,
    Goblin
}

impl Speices {
    pub fn from_u8(value: u8) -> Speices {
        match value {
            0 => Speices::Human,
            1 => Speices::Elf,
            2 => Speices::Orc,
            3 => Speices::Dwarf,
            4 => Speices::Gnome,
            5 => Speices::Halfling,
            6 => Speices::HalfElf,
            7 => Speices::HalfOrc,
            8 => Speices::Dragonborn,
            9 => Speices::Tiefling,
            10 => Speices::Gnoll,
            11 => Speices::Goblin,
            _ => panic!("invalid value for Speices: {}", value),
        }
    }
}

enum ItemType {
    Weapon,
    Armor,
    Potion
}

impl ItemType {
    pub fn from_u8(value: u8) -> ItemType {
        match value {
            0 => ItemType::Weapon,
            1 => ItemType::Armor,
            2 => ItemType::Potion,
            _ => panic!("invalid value for ItemType: {}", value),
        }
    }
}

class Coordinate {
    buffer: Uint8List;
}

impl Coordinate {
    pub fn new(X: I64, Y: I64) -> Coordinate {
        let mut buffer = Vec::new();

        let __unsigned_X = X as u64;
        buffer.push(__unsigned_X as u8);
        buffer.push((__unsigned_X >> 8) as u8);
        buffer.push((__unsigned_X >> 16) as u8);
        buffer.push((__unsigned_X >> 24) as u8);
        buffer.push((__unsigned_X >> 32) as u8);
        buffer.push((__unsigned_X >> 40) as u8);
        buffer.push((__unsigned_X >> 48) as u8);
        buffer.push((__unsigned_X >> 56) as u8);

        let __unsigned_Y = Y as u64;
        buffer.push(__unsigned_Y as u8);
        buffer.push((__unsigned_Y >> 8) as u8);
        buffer.push((__unsigned_Y >> 16) as u8);
        buffer.push((__unsigned_Y >> 24) as u8);
        buffer.push((__unsigned_Y >> 32) as u8);
        buffer.push((__unsigned_Y >> 40) as u8);
        buffer.push((__unsigned_Y >> 48) as u8);
        buffer.push((__unsigned_Y >> 56) as u8);

        Coordinate { buffer: buffer }
    }

    pub fn as_bytes(&self) -> &[u8] {
        &self.buffer[..]
    }

    pub fn as_bytes_mut(&mut self) -> &mut [u8] {
        &mut self.buffer[..]
    }

    pub fn len(&self) -> usize {
        self.buffer.len()
    }

    pub fn from_bytes(bytes: &[u8]) -> Coordinate {
        let mut buffer = Vec::new();
        buffer.extend_from_slice(bytes);
        Coordinate { buffer: buffer }
    }

    pub fn X(&self) -> i64 {
        let mut __result = self.buffer[0] as i64 
            | (self.buffer[1] as i64) << 8 
            | (self.buffer[2] as i64) << 16 
            | (self.buffer[3] as i64) << 24 
            | (self.buffer[4] as i64) << 32 
            | (self.buffer[5] as i64) << 40 
            | (self.buffer[6] as i64) << 48 
            | (self.buffer[7] as i64) << 56;
        __result
    }

    pub fn Y(&self) -> i64 {
        let mut __result = self.buffer[8] as i64 
            | (self.buffer[9] as i64) << 8 
            | (self.buffer[10] as i64) << 16 
            | (self.buffer[11] as i64) << 24 
            | (self.buffer[12] as i64) << 32 
            | (self.buffer[13] as i64) << 40 
            | (self.buffer[14] as i64) << 48 
            | (self.buffer[15] as i64) << 56;
        __result
    }

}

class Item {
    buffer: Uint8List;
}

impl Item {
    pub fn new(Type: ItemType, Damage: I64, Armor: I64, Name: String) -> Item {
        let mut buffer = Vec::new();

        buffer.push(Type as u8);

        let __unsigned_Damage = Damage as u64;
        buffer.push(__unsigned_Damage as u8);
        buffer.push((__unsigned_Damage >> 8) as u8);
        buffer.push((__unsigned_Damage >> 16) as u8);
        buffer.push((__unsigned_Damage >> 24) as u8);
        buffer.push((__unsigned_Damage >> 32) as u8);
        buffer.push((__unsigned_Damage >> 40) as u8);
        buffer.push((__unsigned_Damage >> 48) as u8);
        buffer.push((__unsigned_Damage >> 56) as u8);

        let __unsigned_Armor = Armor as u64;
        buffer.push(__unsigned_Armor as u8);
        buffer.push((__unsigned_Armor >> 8) as u8);
        buffer.push((__unsigned_Armor >> 16) as u8);
        buffer.push((__unsigned_Armor >> 24) as u8);
        buffer.push((__unsigned_Armor >> 32) as u8);
        buffer.push((__unsigned_Armor >> 40) as u8);
        buffer.push((__unsigned_Armor >> 48) as u8);
        buffer.push((__unsigned_Armor >> 56) as u8);

        let mut __dyn_index = 25;
        __dyn_index += Name.len();
        buffer.push(__dyn_index as u8);
        buffer.push((__dyn_index >> 8) as u8);
        buffer.push((__dyn_index >> 16) as u8);
        buffer.push((__dyn_index >> 24) as u8);
        buffer.push((__dyn_index >> 32) as u8);
        buffer.push((__dyn_index >> 40) as u8);
        buffer.push((__dyn_index >> 48) as u8);
        buffer.push((__dyn_index >> 56) as u8);
        let __string_bytes_Name = Name.as_bytes();
        buffer.extend_from_slice(__string_bytes_Name);

        Item { buffer: buffer }
    }

    pub fn as_bytes(&self) -> &[u8] {
        &self.buffer[..]
    }

    pub fn as_bytes_mut(&mut self) -> &mut [u8] {
        &mut self.buffer[..]
    }

    pub fn len(&self) -> usize {
        self.buffer.len()
    }

    pub fn from_bytes(bytes: &[u8]) -> Item {
        let mut buffer = Vec::new();
        buffer.extend_from_slice(bytes);
        Item { buffer: buffer }
    }

    pub fn Type(&self) -> ItemType {
        ItemType::from_u8(self.buffer[0])
    }

    pub fn Damage(&self) -> i64 {
        let mut __result = self.buffer[1] as i64 
            | (self.buffer[2] as i64) << 8 
            | (self.buffer[3] as i64) << 16 
            | (self.buffer[4] as i64) << 24 
            | (self.buffer[5] as i64) << 32 
            | (self.buffer[6] as i64) << 40 
            | (self.buffer[7] as i64) << 48 
            | (self.buffer[8] as i64) << 56;
        __result
    }

    pub fn Armor(&self) -> i64 {
        let mut __result = self.buffer[9] as i64 
            | (self.buffer[10] as i64) << 8 
            | (self.buffer[11] as i64) << 16 
            | (self.buffer[12] as i64) << 24 
            | (self.buffer[13] as i64) << 32 
            | (self.buffer[14] as i64) << 40 
            | (self.buffer[15] as i64) << 48 
            | (self.buffer[16] as i64) << 56;
        __result
    }

    pub fn Name(&self) -> String {
        let __off0: u64 = 25;
        let __off1: u64 = self.buffer[17] as u64 
            | (self.buffer[18] as u64) << 8 
            | (self.buffer[19] as u64) << 16 
            | (self.buffer[20] as u64) << 24 
            | (self.buffer[21] as u64) << 32 
            | (self.buffer[22] as u64) << 40 
            | (self.buffer[23] as u64) << 48 
            | (self.buffer[24] as u64) << 56;
        // TODO: Think about rust's borrowing rules

        String::from_utf8(self.buffer[__off0 as usize..__off1 as usize].to_vec()).unwrap()
    }

}

class Inventory {
    buffer: Uint8List;
}

impl Inventory {
    pub fn new(RightHand: Item, LeftHand: Item) -> Inventory {
        let mut buffer = Vec::new();

        let mut __dyn_index = 16;
        __dyn_index += RightHand.len();
        buffer.push(__dyn_index as u8);
        buffer.push((__dyn_index >> 8) as u8);
        buffer.push((__dyn_index >> 16) as u8);
        buffer.push((__dyn_index >> 24) as u8);
        buffer.push((__dyn_index >> 32) as u8);
        buffer.push((__dyn_index >> 40) as u8);
        buffer.push((__dyn_index >> 48) as u8);
        buffer.push((__dyn_index >> 56) as u8);
        let __struct_bytes_RightHand = RightHand.as_bytes();
        buffer.extend_from_slice(__struct_bytes_RightHand);

        __dyn_index += LeftHand.len();
        buffer.push(__dyn_index as u8);
        buffer.push((__dyn_index >> 8) as u8);
        buffer.push((__dyn_index >> 16) as u8);
        buffer.push((__dyn_index >> 24) as u8);
        buffer.push((__dyn_index >> 32) as u8);
        buffer.push((__dyn_index >> 40) as u8);
        buffer.push((__dyn_index >> 48) as u8);
        buffer.push((__dyn_index >> 56) as u8);
        let __struct_bytes_LeftHand = LeftHand.as_bytes();
        buffer.extend_from_slice(__struct_bytes_LeftHand);

        Inventory { buffer: buffer }
    }

    pub fn as_bytes(&self) -> &[u8] {
        &self.buffer[..]
    }

    pub fn as_bytes_mut(&mut self) -> &mut [u8] {
        &mut self.buffer[..]
    }

    pub fn len(&self) -> usize {
        self.buffer.len()
    }

    pub fn from_bytes(bytes: &[u8]) -> Inventory {
        let mut buffer = Vec::new();
        buffer.extend_from_slice(bytes);
        Inventory { buffer: buffer }
    }

    pub fn RightHand(&self) -> Item {
        let __off0: u64 = 16;
        let __off1: u64 = self.buffer[0] as u64 
            | (self.buffer[1] as u64) << 8 
            | (self.buffer[2] as u64) << 16 
            | (self.buffer[3] as u64) << 24 
            | (self.buffer[4] as u64) << 32 
            | (self.buffer[5] as u64) << 40 
            | (self.buffer[6] as u64) << 48 
            | (self.buffer[7] as u64) << 56;
        // TODO: Think about rust's borrowing rules

        Item::from_bytes(&self.buffer[__off0 as usize..__off1 as usize])
    }

    pub fn LeftHand(&self) -> Item {
        let __off0: u64 = self.buffer[0] as u64 
            | (self.buffer[1] as u64) << 8 
            | (self.buffer[2] as u64) << 16 
            | (self.buffer[3] as u64) << 24 
            | (self.buffer[4] as u64) << 32 
            | (self.buffer[5] as u64) << 40 
            | (self.buffer[6] as u64) << 48 
            | (self.buffer[7] as u64) << 56;
        let __off1: u64 = self.buffer[8] as u64 
            | (self.buffer[9] as u64) << 8 
            | (self.buffer[10] as u64) << 16 
            | (self.buffer[11] as u64) << 24 
            | (self.buffer[12] as u64) << 32 
            | (self.buffer[13] as u64) << 40 
            | (self.buffer[14] as u64) << 48 
            | (self.buffer[15] as u64) << 56;
        // TODO: Think about rust's borrowing rules

        Item::from_bytes(&self.buffer[__off0 as usize..__off1 as usize])
    }

}

class Entity {
    buffer: Uint8List;
}

impl Entity {
    pub fn new(Type: Speices, Position: Coordinate, Hp: I64, Id: UUID, Inventory: Inventory) -> Entity {
        let mut buffer = Vec::new();

        buffer.push(Type as u8);

        let __struct_bytes_Position = Position.as_bytes();
        buffer.extend_from_slice(__struct_bytes_Position);

        let __unsigned_Hp = Hp as u64;
        buffer.push(__unsigned_Hp as u8);
        buffer.push((__unsigned_Hp >> 8) as u8);
        buffer.push((__unsigned_Hp >> 16) as u8);
        buffer.push((__unsigned_Hp >> 24) as u8);
        buffer.push((__unsigned_Hp >> 32) as u8);
        buffer.push((__unsigned_Hp >> 40) as u8);
        buffer.push((__unsigned_Hp >> 48) as u8);
        buffer.push((__unsigned_Hp >> 56) as u8);

        let mut __dyn_index = 41;
        __dyn_index += Id.len();
        buffer.push(__dyn_index as u8);
        buffer.push((__dyn_index >> 8) as u8);
        buffer.push((__dyn_index >> 16) as u8);
        buffer.push((__dyn_index >> 24) as u8);
        buffer.push((__dyn_index >> 32) as u8);
        buffer.push((__dyn_index >> 40) as u8);
        buffer.push((__dyn_index >> 48) as u8);
        buffer.push((__dyn_index >> 56) as u8);
        let __string_bytes_Id = Id.as_bytes();
        buffer.extend_from_slice(__string_bytes_Id);

        __dyn_index += Inventory.len();
        buffer.push(__dyn_index as u8);
        buffer.push((__dyn_index >> 8) as u8);
        buffer.push((__dyn_index >> 16) as u8);
        buffer.push((__dyn_index >> 24) as u8);
        buffer.push((__dyn_index >> 32) as u8);
        buffer.push((__dyn_index >> 40) as u8);
        buffer.push((__dyn_index >> 48) as u8);
        buffer.push((__dyn_index >> 56) as u8);
        let __struct_bytes_Inventory = Inventory.as_bytes();
        buffer.extend_from_slice(__struct_bytes_Inventory);

        Entity { buffer: buffer }
    }

    pub fn as_bytes(&self) -> &[u8] {
        &self.buffer[..]
    }

    pub fn as_bytes_mut(&mut self) -> &mut [u8] {
        &mut self.buffer[..]
    }

    pub fn len(&self) -> usize {
        self.buffer.len()
    }

    pub fn from_bytes(bytes: &[u8]) -> Entity {
        let mut buffer = Vec::new();
        buffer.extend_from_slice(bytes);
        Entity { buffer: buffer }
    }

    pub fn Type(&self) -> Speices {
        Speices::from_u8(self.buffer[0])
    }

    pub fn Position(&self) -> Coordinate {
        // TODO: Think about rust's borrowing rules
        Coordinate::from_bytes(&self.buffer[1..17])
    }

    pub fn Hp(&self) -> i64 {
        let mut __result = self.buffer[17] as i64 
            | (self.buffer[18] as i64) << 8 
            | (self.buffer[19] as i64) << 16 
            | (self.buffer[20] as i64) << 24 
            | (self.buffer[21] as i64) << 32 
            | (self.buffer[22] as i64) << 40 
            | (self.buffer[23] as i64) << 48 
            | (self.buffer[24] as i64) << 56;
        __result
    }

    pub fn Id(&self) -> UUID {
        let __off0: u64 = 41;
        let __off1: u64 = self.buffer[25] as u64 
            | (self.buffer[26] as u64) << 8 
            | (self.buffer[27] as u64) << 16 
            | (self.buffer[28] as u64) << 24 
            | (self.buffer[29] as u64) << 32 
            | (self.buffer[30] as u64) << 40 
            | (self.buffer[31] as u64) << 48 
            | (self.buffer[32] as u64) << 56;
        // TODO: Think about rust's borrowing rules

        String::from_utf8(self.buffer[__off0 as usize..__off1 as usize].to_vec()).unwrap()
    }

    pub fn Inventory(&self) -> Inventory {
        let __off0: u64 = self.buffer[25] as u64 
            | (self.buffer[26] as u64) << 8 
            | (self.buffer[27] as u64) << 16 
            | (self.buffer[28] as u64) << 24 
            | (self.buffer[29] as u64) << 32 
            | (self.buffer[30] as u64) << 40 
            | (self.buffer[31] as u64) << 48 
            | (self.buffer[32] as u64) << 56;
        let __off1: u64 = self.buffer[33] as u64 
            | (self.buffer[34] as u64) << 8 
            | (self.buffer[35] as u64) << 16 
            | (self.buffer[36] as u64) << 24 
            | (self.buffer[37] as u64) << 32 
            | (self.buffer[38] as u64) << 40 
            | (self.buffer[39] as u64) << 48 
            | (self.buffer[40] as u64) << 56;
        // TODO: Think about rust's borrowing rules

        Inventory::from_bytes(&self.buffer[__off0 as usize..__off1 as usize])
    }

}

type UUID = String;

