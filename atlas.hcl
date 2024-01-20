data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/db/loader.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "sqlite://dev?mode=memory"
  url = "sqlite3://bin/data.db"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}