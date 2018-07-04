package goutil

import (
  "fmt"
  "strings"
  "errors"
)

func IndexOf(i interface{}, _arr interface{}) int {
  if !strings.Contains(fmt.Sprintf("%T", _arr), "[]"){
    panic(errors.New("IndexOf(i, arr) arr must be array/slice"))
  }else if strings.Contains(fmt.Sprintf("%T", i), "[]"){
    panic(errors.New("IndexOf(i, arr) in can't be array/slice"))
  }else if fmt.Sprintf("[]%T", i) != fmt.Sprintf("%T", _arr){
    panic(errors.New("IndexOf(i, arr) in and arr must share base type"))
  }
  arr := intoarr(_arr)
  index := 0
  for in, v := range arr{
    if v == i{
      index = in
      break
    }
  }
  return index
}
