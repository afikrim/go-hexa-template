package repository

type (
	BaseRepository interface {
		GetBusinessTypeRepository() BusinessTypeRepository

		Migrate() error
	}
)
