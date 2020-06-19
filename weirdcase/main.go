package main

import (
    "strings"
)

func toWeirdCase(str string) string {
      retVal := []byte{}

  // Your code here and happy coding!
      letterIndex := 0
      for _, letter := range str {
         if letterIndex % 2 == 0 {
            val := strings.ToUpper(string(letter))
            retVal = append(retVal, val[0])
            letterIndex++
         } else if letterIndex % 2 != 0 {
            val := strings.ToLower(string(letter))
            retVal = append(retVal, val[0])
            letterIndex++
         }
         if string(letter) == " " {
             letterIndex = 0
         } 
      }
  return string(retVal)
}
