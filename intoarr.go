package goutil

import (
  "reflect"
  "errors"
)

func intoarr(in interface{}) []interface{} {
  slice, success := takeArg(in, reflect.Slice)
  if !success {
    panic(errors.New("Array read error"))
  }
  c := slice.Len()
  arr := make([]interface{}, c)
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
