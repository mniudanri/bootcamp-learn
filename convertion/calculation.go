package main
func GetSumFloat32Typed(num1, num2 uint8) (float32) {
  // converting uint8 to float32
  var (
  //   num1 uint8 = 20
  //   num2 uint8 = 30
    sum float32
  )

  sum = float32(num1 + num2)

  return sum
}
