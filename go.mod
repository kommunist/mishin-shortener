module mishin-shortener

replace internal/hasher => ./internal/app/hasher

replace internal/storage => ./internal/app/storage

replace internal/handlers => ./internal/app/handlers

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-chi/chi/v5 v5.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	internal/handlers v0.0.0-00010101000000-000000000000 // indirect
	internal/hasher v0.0.0-00010101000000-000000000000 // indirect
	internal/storage v0.0.0-00010101000000-000000000000 // indirect
)

go 1.22.0
