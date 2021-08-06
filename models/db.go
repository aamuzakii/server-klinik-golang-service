package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq" //PostgreSQL Driver
    "fmt"
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", "postgres", "user=postgres password=postgres dbname=klinik host=localhost sslmode=disable")
    orm.RegisterModel(new(Users))
    ormObject = orm.NewOrm()
    name := "default"

    // Drop table and re-create.
    force := true

    // Print log.
    verbose := true

    // Error.
    err := orm.RunSyncdb(name, force, verbose)
    if err != nil {
        fmt.Println(err)
    }
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
func GetOrmObject() orm.Ormer {
    return ormObject
}
