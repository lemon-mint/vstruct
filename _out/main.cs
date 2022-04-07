namespace vstruct.main
{
	using UUID = System.String;
	public enum Speices : byte
	{
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
		Goblin,
	}

	public enum ItemType : byte
	{
		Weapon,
		Armor,
		Potion,
	}

	class Coordinate
	{
		public Coordinate(int size)
		{
			this.value = new byte[size];
		}

		public static Coordinate FromBytes(byte[] bytes)
		{
			var s = new Coordinate(bytes.Length);
			Array.Copy(bytes, 0, s.value, 0, bytes.Length);
			return s;
		}

		public static Coordinate Serialize(Coordinate dst, long X, long Y)
		{
			var __tmp_0 = BitConverter.GetBytes(X);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_0);
			}
			Array.Copy(__tmp_0, 0, dst.value, 0, 8);
			var __tmp_1 = BitConverter.GetBytes(Y);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_1);
			}
			Array.Copy(__tmp_1, 0, dst.value, 8, 8);

			return dst;
		}

		public static Coordinate New(long X, long Y)
		{
			var __vstruct__size = (ulong)16 ;
			var __vstruct__buf = Coordinate.FromBytes(new byte[__vstruct__size]);
			__vstruct__buf = Coordinate.Serialize(__vstruct__buf, X, Y);
			return __vstruct__buf;
		}

		public byte[] value = new byte[0];
		public long GetX()
		{
			var buf = new byte[8];
			Array.Copy(this.value, 0, buf, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf);
			}
			return BitConverter.ToInt64(buf, 0);
		}

		public long GetY()
		{
			var buf = new byte[8];
			Array.Copy(this.value, 8, buf, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf);
			}
			return BitConverter.ToInt64(buf, 0);
		}

	}

	class Item
	{
		public Item(int size)
		{
			this.value = new byte[size];
		}

		public static Item FromBytes(byte[] bytes)
		{
			var s = new Item(bytes.Length);
			Array.Copy(bytes, 0, s.value, 0, bytes.Length);
			return s;
		}

		public static Item Serialize(Item dst, ItemType Type, long Damage, long Armor, string Name)
		{
			dst.value[0] = (byte)Type;
			var __tmp_1 = BitConverter.GetBytes(Damage);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_1);
			}
			Array.Copy(__tmp_1, 0, dst.value, 1, 8);
			var __tmp_2 = BitConverter.GetBytes(Armor);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_2);
			}
			Array.Copy(__tmp_2, 0, dst.value, 9, 8);

			var __index = (ulong)25;
			var __tmp_3 = (ulong)System.Text.Encoding.UTF8.GetBytes(Name).Length +__index;
			var __tmp_3_len = BitConverter.GetBytes(__tmp_3);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_3_len);
			}
			Array.Copy(__tmp_3_len, 0, dst.value, 17, 8);
			var __buf_3 = System.Text.Encoding.UTF8.GetBytes(Name);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__buf_3);
			}
			Array.Copy(__buf_3, 0, dst.value, (long)__index, (long)__tmp_3-(long)__index);
			return dst;
		}

		public static Item New(ItemType Type, long Damage, long Armor, string Name)
		{
			var __vstruct__size = (ulong)25 + (ulong)System.Text.Encoding.UTF8.GetBytes(Name).Length ;
			var __vstruct__buf = Item.FromBytes(new byte[__vstruct__size]);
			__vstruct__buf = Item.Serialize(__vstruct__buf, Type, Damage, Armor, Name);
			return __vstruct__buf;
		}

		public byte[] value = new byte[0];
		public ItemType GetType()
		{
			return (ItemType)(this.value[0]);
		}

		public long GetDamage()
		{
			var buf = new byte[8];
			Array.Copy(this.value, 1, buf, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf);
			}
			return BitConverter.ToInt64(buf, 0);
		}

		public long GetArmor()
		{
			var buf = new byte[8];
			Array.Copy(this.value, 9, buf, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf);
			}
			return BitConverter.ToInt64(buf, 0);
		}

		public string GetName()
		{
			ulong __off0 = 25;
			var buf_0_2 = new byte[8];
			Array.Copy(this.value, 17, buf_0_2, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf_0_2);
			}
			ulong __off1 = ((ulong)BitConverter.ToUInt64(buf_0_2, 0));
			var buf_0_3 = new Byte[__off1-__off0];
			Array.Copy(this.value, (long)__off0, buf_0_3, 0, (long)__off1-(long)__off0);
			return ((string)System.Text.Encoding.UTF8.GetString(buf_0_3));
		}

	}

	class Inventory
	{
		public Inventory(int size)
		{
			this.value = new byte[size];
		}

		public static Inventory FromBytes(byte[] bytes)
		{
			var s = new Inventory(bytes.Length);
			Array.Copy(bytes, 0, s.value, 0, bytes.Length);
			return s;
		}

		public static Inventory Serialize(Inventory dst, Item RightHand, Item LeftHand)
		{

			var __index = (ulong)16;
			var __tmp_0 = (ulong)RightHand.value.Length +__index;
			var __tmp_0_len = BitConverter.GetBytes(__tmp_0);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_0_len);
			}
			Array.Copy(__tmp_0_len, 0, dst.value, 0, 8);
			var __buf_0 = RightHand;
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__buf_0.value);
			}
			Array.Copy(__buf_0.value, 0, dst.value, (long)__index, (long)__tmp_0-(long)__index);
			__index += (ulong)RightHand.value.Length;
			var __tmp_1 = (ulong)LeftHand.value.Length +__index;
			var __tmp_1_len = BitConverter.GetBytes(__tmp_1);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_1_len);
			}
			Array.Copy(__tmp_1_len, 0, dst.value, 8, 8);
			var __buf_1 = LeftHand;
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__buf_1.value);
			}
			Array.Copy(__buf_1.value, 0, dst.value, (long)__index, (long)__tmp_1-(long)__index);
			return dst;
		}

		public static Inventory New(Item RightHand, Item LeftHand)
		{
			var __vstruct__size = (ulong)16 + (ulong)RightHand.value.Length + (ulong)LeftHand.value.Length ;
			var __vstruct__buf = Inventory.FromBytes(new byte[__vstruct__size]);
			__vstruct__buf = Inventory.Serialize(__vstruct__buf, RightHand, LeftHand);
			return __vstruct__buf;
		}

		public byte[] value = new byte[0];
		public Item GetRightHand()
		{
			ulong __off0 = 16;
			var buf_0_2 = new byte[8];
			Array.Copy(this.value, 0, buf_0_2, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf_0_2);
			}
			ulong __off1 = ((ulong)BitConverter.ToUInt64(buf_0_2, 0));
			var buf_0_3 = new Byte[__off1-__off0];
			Array.Copy(this.value, (long)__off0, buf_0_3, 0, (long)__off1-(long)__off0);
			return Item.FromBytes(buf_0_3);
		}

		public Item GetLeftHand()
		{
			var buf_1_1 = new byte[8];
			Array.Copy(this.value, 0, buf_1_1, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf_1_1);
			}
			ulong __off0 = ((ulong)BitConverter.ToUInt64(buf_1_1, 0));
			var buf_1_2 = new byte[8];
			Array.Copy(this.value, 8, buf_1_2, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf_1_2);
			}
			ulong __off1 = ((ulong)BitConverter.ToUInt64(buf_1_2, 0));
			var buf_1_3 = new Byte[__off1-__off0];
			Array.Copy(this.value, (long)__off0, buf_1_3, 0, (long)__off1-(long)__off0);
			return Item.FromBytes(buf_1_3);
		}

	}

	class Entity
	{
		public Entity(int size)
		{
			this.value = new byte[size];
		}

		public static Entity FromBytes(byte[] bytes)
		{
			var s = new Entity(bytes.Length);
			Array.Copy(bytes, 0, s.value, 0, bytes.Length);
			return s;
		}

		public static Entity Serialize(Entity dst, Speices Type, Coordinate Position, long Hp, UUID Id, Inventory Inventory)
		{
			dst.value[0] = (byte)Type;
			Array.Copy(Position.value, 0, dst.value, 1, 16);
			var __tmp_2 = BitConverter.GetBytes(Hp);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_2);
			}
			Array.Copy(__tmp_2, 0, dst.value, 17, 8);

			var __index = (ulong)41;
			var __tmp_3 = (ulong)System.Text.Encoding.UTF8.GetBytes(Id).Length +__index;
			var __tmp_3_len = BitConverter.GetBytes(__tmp_3);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_3_len);
			}
			Array.Copy(__tmp_3_len, 0, dst.value, 25, 8);
			var __buf_3 = System.Text.Encoding.UTF8.GetBytes(Id);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__buf_3);
			}
			Array.Copy(__buf_3, 0, dst.value, (long)__index, (long)__tmp_3-(long)__index);
			__index += (ulong)(ulong)System.Text.Encoding.UTF8.GetBytes(Id).Length;
			var __tmp_4 = (ulong)Inventory.value.Length +__index;
			var __tmp_4_len = BitConverter.GetBytes(__tmp_4);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__tmp_4_len);
			}
			Array.Copy(__tmp_4_len, 0, dst.value, 33, 8);
			var __buf_4 = Inventory;
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(__buf_4.value);
			}
			Array.Copy(__buf_4.value, 0, dst.value, (long)__index, (long)__tmp_4-(long)__index);
			return dst;
		}

		public static Entity New(Speices Type, Coordinate Position, long Hp, UUID Id, Inventory Inventory)
		{
			var __vstruct__size = (ulong)41 + (ulong)System.Text.Encoding.UTF8.GetBytes(Id).Length + (ulong)Inventory.value.Length ;
			var __vstruct__buf = Entity.FromBytes(new byte[__vstruct__size]);
			__vstruct__buf = Entity.Serialize(__vstruct__buf, Type, Position, Hp, Id, Inventory);
			return __vstruct__buf;
		}

		public byte[] value = new byte[0];
		public Speices GetType()
		{
			return (Speices)(this.value[0]);
		}

		public Coordinate GetPosition()
		{
			var buf = new Byte[16];
			Array.Copy(this.value, 1, buf, 0, 16);
			return Coordinate.FromBytes(buf);
		}

		public long GetHp()
		{
			var buf = new byte[8];
			Array.Copy(this.value, 17, buf, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf);
			}
			return BitConverter.ToInt64(buf, 0);
		}

		public UUID GetId()
		{
			ulong __off0 = 41;
			var buf_0_2 = new byte[8];
			Array.Copy(this.value, 25, buf_0_2, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf_0_2);
			}
			ulong __off1 = ((ulong)BitConverter.ToUInt64(buf_0_2, 0));
			var buf_0_3 = new Byte[__off1-__off0];
			Array.Copy(this.value, (long)__off0, buf_0_3, 0, (long)__off1-(long)__off0);
			return ((UUID)System.Text.Encoding.UTF8.GetString(buf_0_3));
		}

		public Inventory GetInventory()
		{
			var buf_1_1 = new byte[8];
			Array.Copy(this.value, 25, buf_1_1, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf_1_1);
			}
			ulong __off0 = ((ulong)BitConverter.ToUInt64(buf_1_1, 0));
			var buf_1_2 = new byte[8];
			Array.Copy(this.value, 33, buf_1_2, 0, 8);
			if (!BitConverter.IsLittleEndian)
			{
				Array.Reverse(buf_1_2);
			}
			ulong __off1 = ((ulong)BitConverter.ToUInt64(buf_1_2, 0));
			var buf_1_3 = new Byte[__off1-__off0];
			Array.Copy(this.value, (long)__off0, buf_1_3, 0, (long)__off1-(long)__off0);
			return Inventory.FromBytes(buf_1_3);
		}

	}


}
