package utils

import (
  "github.com/fatih/color"
)
var (danger *color.Color
  info *color.Color)

func init() {
   danger = color.New(color.FgRed)
   info = color.New(color.FgCyan)
}

func PrintStackTrace(funcName string, err error) {
  danger.Printf("\n\n------------ %v error \n %+v \n-----------\n", funcName, err)
}
