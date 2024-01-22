package util

import (
  "fmt"
  "testing"
)

func TestBitSet(t *testing.T) {
  b8 := uint8(0)
  b16 := uint16(0)
  b32 := uint32(0)
  b64 := uint64(0)
  b := uint(0)
  BitSet(&b8, 0, true)
  if b8 != 1 {
    t.Errorf("expected %d got %d", 1, b8)
  }
  BitSet(&b16, 0, true)
  if b16 != 1 {
    t.Errorf("expected %d got %d", 1, b16)
  }
  BitSet(&b32, 0, true)
  if b32 != 1 {
    t.Errorf("expected %d got %d", 1, b32)
  }
  BitSet(&b64, 0, true)
  if b64 != 1 {
    t.Errorf("expected %d got %d", 1, b64)
  }
  BitSet(&b, 0, true)
  if b != 1 {
    t.Errorf("expected %d got %d", 1, b)
  }
  
  BitSet(&b8, 7, true)
  if b8 != 129 {
    t.Errorf("expected %d got %d", 129, b8)
  }
  BitSet(&b16, 15, true)
  if b16 != 32769 {
    t.Errorf("expected %d got %d", 32769, b16)
  }
  BitSet(&b32, 31, true)
  if b32 != 2147483649 {
    t.Errorf("expected %d got %d", 2147483649, b32)
  }
  BitSet(&b64, 63, true)
  if b64 != 9223372036854775809 {
    t.Errorf("expected %d got %d", uint64(9223372036854775809), b64)
  }
  BitSet(&b, 1, true)
  if b != 3 {
    t.Errorf("expected %d got %d", 3, b)
  }
  
  BitSet(&b8, 0, false)
  if b8 != 128 {
    t.Errorf("expected %d got %d", 128, b8)
  }
  BitSet(&b16, 0, false)
  if b16 != 32768 {
    t.Errorf("expected %d got %d", 32768, b16)
  }
  BitSet(&b32, 0, false)
  if b32 != 2147483648 {
    t.Errorf("expected %d got %d", 2147483648, b32)
  }
  BitSet(&b64, 0, false)
  if b64 != 9223372036854775808 {
    t.Errorf("expected %d got %d", uint64(9223372036854775809), b64)
  }
  BitSet(&b, 0, false)
  if b != 2 {
    t.Errorf("expected %d got %d", 2, b)
  }
  
  e := "e"
  BitSet(&e, 0, true)
  if e != "e" {
    t.Errorf("expected %s got %s", "e", e)
  }
}

func ExampleBitSet() {
  v := 0
  BitSet(&v, 1, true)
  fmt.Println(v)
  ui := uint8(0)
  BitSet(&ui, 7, true)
  fmt.Println(ui)
  i := int8(0)
  BitSet(&i, 7, true)
  fmt.Println(i)
  
  // Output:
  // 2
  // 128
  // -128
}
