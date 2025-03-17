Example :
Contoh :
/*package config
import (
    "log"
    "context"
    "os"
    "github.com/jackc/pgx/v5"  
    "github.com/joho/godotenv"
)

var DB *pgx.Conn 

func LoadEnv () {
    
    err := godotenv.Load()
    
    if err != nil {
        log.Println("Peringatan: Gagal memuat .env file, pakai default atau env sistem.")
    }
    
}

func ConnectCBMI () {
    
    LoadEnv()
    
    dbURL := os.Getenv("DATABASE_URL")
    
    if dbURL == "" {
        dbURL = "(USERNAME:PASSWORD:HOST:PORT:DBNAM)"
    }
    
    conn, err := pgx.Connect(context.Background(), dbURL)
    
    if err != nil {
        log.Fatal("Gagal konek ke database", err)
    }
    
    DB = conn
    
}

func CloseCBMI () {
    
    if DB != nil {
        DB.Close(context.Background())
    }
    
}
*/