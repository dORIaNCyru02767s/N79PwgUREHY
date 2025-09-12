// 代码生成时间: 2025-09-13 01:54:00
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User represents the user model with permissions
type User struct {
    gorm.Model
    Username string
    Permissions []Permission `gorm:"many2many:user_permissions;"`
# FIXME: 处理边界情况
}

// Permission represents a permission model
# TODO: 优化性能
type Permission struct {
# 扩展功能模块
    gorm.Model
    Name string
# FIXME: 处理边界情况
    Users []User `gorm:"many2many:user_permissions;"`
}

// UserPermission is a join table for User and Permission
type UserPermission struct {
    gorm.Model
    UserID uint
    PermissionID uint
}

// NewUser creates a new user with permissions
func NewUser(db *gorm.DB, username string, permissions []string) (*User, error) {
    user := User{Username: username}
    if err := db.Create(&user).Error; err != nil {
        return nil, err
    }
    for _, name := range permissions {
        var permission Permission
        if result := db.Where(Permission{Name: name}).First(&permission); result.Error != nil {
            if result.Error == gorm.ErrRecordNotFound {
                db.Create(&Permission{Name: name})
            } else {
                return nil, result.Error
            }
        }
# 改进用户体验
        db.Model(&user).Association("Permissions