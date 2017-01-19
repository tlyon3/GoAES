package aes

/*
  0 4 8  12
  1 5 9  13
  2 6 10 14
  3 7 11 15
 */

func ShiftRows(state []byte) {
  //swap 2nd row
  swap(state, 1, 5)
  swap(state, 5, 9)
  swap(state, 9, 13)

  //swap 3rd row
  swap(state, 6, 10)
  swap(state, 10, 14)
  swap(state, 2, 6)
  swap(state, 6, 10)

  //swap 4th row
  swap(state, 11, 15)
  swap(state, 7, 11)
  swap(state, 3, 7)
}

func swap(state []byte, i int, j int){
  temp = state[i]
  state[i] = state[j]
  state[j] = temp
}
