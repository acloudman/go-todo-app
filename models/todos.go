package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-todo-app/configs"
)

func GetAllTodos(todo *[]Todo) (err error) {
	if err = configs.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err = configs.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *Todo, id string) (err error) {
	if err := configs.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	configs.DB.Save(todo)
	return nil
}

func DeleteATodo(todo *Todo, id string) (err error) {
	configs.DB.Where("id = ?", id).Delete(todo)
	return nil
}