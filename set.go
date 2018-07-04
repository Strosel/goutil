package goutil

import (
  "strings"
  "fmt"
)

var input string

func Set(_input interface{}) []interface{} {
  input := fmt.Sprintf("%v", _input)
  arr := []interface{}
  for _, v := range strings.Split(input, ""){
    if !In(v, arr){
      arr = append(arr, v)
    }
  }
  return arr
}
