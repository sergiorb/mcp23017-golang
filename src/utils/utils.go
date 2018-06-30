package utils

func SetBit(n int, pos uint) int {
  n |= (1 << pos)
  return n
}

func ClearBit(n int, pos uint) int {
    mask := ^(1 << pos)
    n &= mask
    return n
}

func HasBit(n int, pos uint) bool {
    val := n & (1 << pos)
    return (val > 0)
}

func GetBPort(port uint) uint {
  if port == 8 { return 0 }
  if port == 9 { return 1 }
  if port == 10 { return 2 }
  if port == 11 { return 3 }
  if port == 12 { return 4 }
  if port == 13 { return 5 }
  if port == 14 { return 6 }
  return 7
}
