package aes

func SubBytes(state []byte) {
  for i, b := range state {
    state[i] = sbox[b]
  }
}
