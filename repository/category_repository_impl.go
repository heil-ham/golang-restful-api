package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
	}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	saveSql := "INSERT INTO category (name) VALUES (?)"
	sqlResult, errExec := tx.ExecContext(ctx, saveSql, category.Name)

	helper.PanicIfError(errExec)

	id, errId := sqlResult.LastInsertId()

	helper.PanicIfError(errId)
	
	category.Id = int(id)
	return category
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	updateSql := "UPDATE category SET name = ? WHERE id = ?"
	_, errExec := tx.ExecContext(ctx, updateSql, category.Name, category.Id)

	helper.PanicIfError(errExec)
	
	return category
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	deleteSql := "DELETE FROM category WHERE id = ?"
	_, errExec := tx.ExecContext(ctx, deleteSql, category.Id)

	helper.PanicIfError(errExec)
	
}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int32) (domain.Category, error) {
	FindByIdSql := "SELECT id, name FROM category WHERE id = ?"
	sqlRow, errExec := tx.QueryContext(ctx, FindByIdSql, categoryId)
	
	helper.PanicIfError(errExec)

	defer sqlRow.Close()

	category := domain.Category{
	}

	if sqlRow.Next() {
		errScan := sqlRow.Scan(&category.Id, &category.Name)
		helper.PanicIfError(errScan)
		return category, nil
	} else {
		return category, errors.New("tidak ditemukan")
	}
}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	FindAllSql := "SELECT id, name FROM category"
	sqlRow, errExec := tx.QueryContext(ctx, FindAllSql)
	
	helper.PanicIfError(errExec)

	defer sqlRow.Close()

	categories := []domain.Category{
	}

	category := domain.Category{
	}

	for sqlRow.Next() {
		errScan := sqlRow.Scan(&category.Id, &category.Name)
		helper.PanicIfError(errScan)
		categories = append(categories, category)
	}
	return categories
}