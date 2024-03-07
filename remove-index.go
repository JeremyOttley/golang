// If you want to keep your array ordered, you have to shift all of the elements at the right of the deleting index by one to the left. Hopefully, this can be done easily in Golang:
func remove(slice []int, i int) []int {
  copy(slice[i:], slice[i+1:])
  return slice[:len(slice)-1]
}
