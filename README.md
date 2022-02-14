
[![GitHub](https://img.shields.io/github/license/lemon-mint/vstruct?style=for-the-badge)](https://github.com/lemon-mint/vstruct/blob/main/LICENSE)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/lemon-mint/vstruct?label=latest&style=for-the-badge)](https://github.com/lemon-mint/vstruct/releases/latest)
[![npm](https://img.shields.io/npm/v/vstruct?color=cb0303&style=for-the-badge)](https://www.npmjs.com/package/vstruct)

# vstruct

Code Generation Based High Speed Data Serialization Tool

# Installation

## 1. From NPM (recommended)

```bash
npm install -g vstruct
```

## 2. From Source

```
git clone https://github.com/lemon-mint/vstruct.git
cd vstruct
go build -o vstruct ./cli/vsc
```

## 3. Pre-compiled binaries

[https://github.com/lemon-mint/vstruct/releases/latest](https://github.com/lemon-mint/vstruct/releases/latest)

# Vstruct Syntax

## 0. Primitive Types

### 0.1. Boolean
---
```
bool: bool # true or false
```

### 0.2. Signed Integers
---
```
int8: int8 # signed 8-bit integer
int16: int16 # signed 16-bit integer
int32: int32 # signed 32-bit integer
int64: int64 # signed 64-bit integer
```

### 0.3. Unsigned Integers
---
```
uint8: uint8 # unsigned 8-bit integer
uint16: uint16 # unsigned 16-bit integer
uint32: uint32 # unsigned 32-bit integer
uint64: uint64 # unsigned 64-bit integer
```

### 0.4. Floating Point
---
```
float32: float32 # 32-bit floating point (IEEE 754)
float64: float64 # 64-bit floating point (IEEE 754)
```

### 0.5. Bytes
---
```
bytes: bytes # variable length bytes
```

### 0.6. String
---
```
string: string # variable length string
```

## 1. Enum

```vstruct
enum MyEnum {
    one,
    two,
    three
}
```

## 2. Struct

```vstruct
struct MyStruct {
    uint8  a;
    uint16 b;
    uint32 c;
    uint64 d;
    string e;
    MyEnum f;
}
```

## 3. Alias

```vstruct
alias UUID = string;
```

# Vstruct CLI Usage

```
vstruct [options] <lang> <package name> <input file>
```

## Options

```
-o <output> Output file name (default: <inputfile>.<.go|.py|.dart|.rs>)
-s          Prints the generated code to stdout
-v          Print version and exit
-h          Print help and exit
-l          Print license and exit
```

## Languages

```
go: Go
python: Python
dart: Dart
rust: Rust (Experimental)
```
