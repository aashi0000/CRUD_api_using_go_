module example/basic_api/services

go 1.22.5

replace example/basic_api/models => ../models

replace example/basic_api/db => ../db

require (
	example/basic_api/models v0.0.0-00010101000000-000000000000
	gorm.io/gorm v1.25.11
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.16.0 // indirect
)
