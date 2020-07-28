package main

import (
   "github.com/jinzhu/gorm"
)

type Variable struct {
   gorm.Model
   Name      string `gorm:"primary_key"`
   Sets      []*Set `gorm:"many2many:variable_sets;association_foreignkey:name;foreignkey:name;"`
}

type Set struct {
   gorm.Model
   Name      string `gorm:"primary_key"`
   Variables []*Variable `gorm:"many2many:variable_sets;association_foreignkey:name;foreignkey:name;"`
   Themes    []*Theme`gorm:"many2many:set_themes;association_foreignkey:name;foreignkey:name;"`
}

type Theme struct {
   gorm.Model
   Name      string `gorm:"primary_key"`
   Sets      []*Set `gorm:"many2many:set_themes;association_foreignkey:name;foreignkey:name;"`
}
