package aes
/*
  0 4 8  12
  1 5 9  13
  2 6 10 14
  3 7 11 15
 */

func MixCols(state []byte) []byte {
  //has to be a better way than creating this everytime...
  // m[row][col]
  m := [][]byte{}
  row1 := []byte{2, 3, 1, 1}
  row2 := []byte{1, 2, 3, 1}
  row3 := []byte{1, 1, 2, 3}
  row4 := []byte{3, 1, 1, 2}
  m = append(m, row1)
  m = append(m, row2)
  m = append(m, row3)
  m = append(m, row4)

  var col1 = [4]byte{state[0], state[1], state[2], state[3]}
  var col2 = [4]byte{state[4], state[5], state[6], state[7]}
  var col3 = [4]byte{state[8], state[9], state[10], state[11]}
  var col4 = [4]byte{state[12], state[13], state[14], state[15]}
  //matrix[col][row]
  var matrix = [4][4]byte{col1, col2, col3, col4}

  for i := 0; i < 4; i++ {
    var col = matrix[i]
    var newcol = matrix[i]
    for j := 0; j < 4; j++ {
      var sum byte = 0
      for k := 0; k < 4; k++ {
        var temp = FiniteFieldMult(col[k], m[j][k])
        sum += temp
      }
      newcol[j] = sum
    }
    matrix[i][0] = newcol[0]
    matrix[i][1] = newcol[1]
    matrix[i][2] = newcol[2]
    matrix[i][3] = newcol[3]
  }
  return state
}

func FiniteFieldMult(a byte, b byte) byte {
  var temp = a
  var sum byte = 0
  for b > 0 {
    if b & 0x01 == 1 {
      sum = sum ^ temp
    }
    temp = xtime(temp)
    b = b >> 1
  }
  return sum
}

func xtime(a byte) byte{
  if a & 0x80 == 0x80 {
    a = a << 1
    a = a ^ 0x1b
  } else {
    a = a << 1
  }
  return a
}
