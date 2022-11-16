package repository

import "avito/src/data/entity"

type Repository[T entity.Entity] interface {
	getById(id int) T
	getAll() T
	create() T
	update(id int) T
	delete(id int) T
}
