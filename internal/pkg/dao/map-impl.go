package dao

import (
	"errors"
	"fmt"
)

type MapDao struct {
	unique map[string]string
}

func NewMapDao(collection map[string]string) DAO {
	return &MapDao{
		unique: collection,
	}
}

func (dao *MapDao) Save(hash string, path string) error {

	if _, ok := dao.unique[hash]; ok {
		return errors.New(`duplicate`)
	}

	return nil
}

func (dao *MapDao) Delete(hash string) error {

	if _, ok := dao.unique[hash]; ok {
		delete(dao.unique, hash)
	} else {
		return errors.New(fmt.Sprintf("cannot delete hash %s not found", hash))
	}

	return nil
}
