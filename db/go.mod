module example/basic_api/db

go 1.22.5

require (
	example/basic_api/models v0.0.0-00010101000000-000000000000
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.11
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.16.0 // indirect
)

replace example/basic_api/mocks => ../mocks

replace example/basic_api/models => ../models
