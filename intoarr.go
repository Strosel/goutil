package goutil

func intoarr(in interface{}) []interface{} {
  slice, success := takeArg(in, reflect.Slice)
  if !success {
    panic()
  }
  arr := make([]interface{}, slice.Len())
  for i := 0; i < c; i++ {
    arr[i] = slice.Index(i).Interface()
  }
  return arr
}
