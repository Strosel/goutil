package goutil

import (
  "strings"
  "fmt"
)

func Set(_input interface{}) []string {
  input := fmt.Sprintf("%v", _input)
  arr := []string{}
  for _, v := range strings.Split(input, ""){
    if !In(v, arr){
      arr = append(arr, v)
    }
  }
  return arr
}
