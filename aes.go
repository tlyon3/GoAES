package aes

type AES struct {
  key []byte
}

func (a *AES) Encrypt(input []byte) []byte {
  state := make([]byte, 16)
  copy(state, input)
  
  //will need to change the number of rounds to adapt for 192 and 256
  var numRounds = 9

  AddRoundKey(state)
  for i := 0; i < numRounds; i++{
    SubBytes(state)
    ShiftRows(state)
    MixCols(state)
    AddRoundKey(state)
  }

  SubBytes(state)
  ShiftRows(state)
  AddRoundKey(state)

  return state
}
