// 代码生成时间: 2025-08-05 16:07:50
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Permission represents the permissions a user can have.
type Permission struct {
    gorm.Model
    Name  string
    Users []User `gorm:"many2many:permission_users;"`
}

// User represents a user with permissions.
type User struct {
    gorm.Model
    Name     string
    Email    string `gorm:"type:varchar(100);uniqueIndex"`
    Password string `gorm:"not null"`
    Permissions []Permission `gorm:"many2many:permission_users;"`
}

// UserPermissionJWN is a JWT payload for a user.
type UserPermissionJWT struct {
    ID       uint   `json:"id"`
    Name     string `json:"name"`
    Email    string `json: