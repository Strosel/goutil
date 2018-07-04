package goutil

import (
  "fmt"
  "strings"
  "errors"
)

func In(in interface{}, _arr interface{}) bool {
  if !strings.Contains(fmt.Sprintf("%T", _arr), "[]"){
    panic(errors.New("In(in, arr) arr must be array/slice"))
  }else if strings.Contains(fmt.Sprintf("%T", in), "[]"){
    panic(errors.New("In(in, arr) in can't be array/slice"))
  }else if fmt.Sprintf("[]%T", in) != fmt.Sprintf("%T", _arr){
    panic(errors.New("In(in, arr) in and arr must share base type"))
  }
  arr := intoarr(_arr)
  b := false
  for _, v := range arr{
    if v == in{
      b = true
      break
    }
  }
  return b
}
