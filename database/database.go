package database

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"  // ← Cambiado de mysql a postgres
    "gorm.io/gorm"
)

var Database = func() (db *gorm.DB) {
    // Se valida existencia .env y variables de entorno
    errorVariables := godotenv.Load()
    if errorVariables != nil {
        panic(errorVariables)
    }
    
    // String de conexión para PostgreSQL
    dsn := os.Getenv("DATABASE_URL")
    
    if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
        fmt.Println("Error de conexión a PostgreSQL")
        panic(err)
    } else {
        fmt.Println("Conexión exitosa a PostgreSQL")
        return db
    }
}()