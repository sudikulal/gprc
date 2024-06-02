package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var MONGO_URI string
var DATABASE string
var JWT_SECRET_KEY string
var ENCRYPTION_KEY string

var REGISTER_TYPE = map[string]int {
    "EMAIL":1,
    "USER_NAME":2,
}

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    MONGO_URI = os.Getenv("MONGO_URI")
    DATABASE = os.Getenv("DATABASE")
    JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
    ENCRYPTION_KEY = os.Getenv("ENCRYPTION_KEY")

    if MONGO_URI == "" || DATABASE == "" || JWT_SECRET_KEY == ""  || ENCRYPTION_KEY == "" {
        log.Fatalf("MONGO_URI / DATABASE / JWT_SECRET_KEY  / ENCRYPTION_KEY environment variable not set")
    }
}

