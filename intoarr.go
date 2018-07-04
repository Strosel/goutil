package goutil

import "reflect"

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

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
    val = reflect.ValueOf(arg)
    if val.Kind() == kind {
        ok = true
    }
    return
}
