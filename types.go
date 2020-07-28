
package main

type Formula struct {
   Col_outcome  string `yaml:"col_outcome"`
   Cols_features string `yaml:"cols_features"`
}

type ModelSpec struct {
   Colsets  map[string][]string `yaml:"colsets"`
   Themes   map[string][]string `yaml:"themes"`
   Formulas map[string]Formula `yaml:"formulas"`
}

