package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var MONGO_URI string
var DATABASE string

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
    if MONGO_URI == "" || DATABASE == "" {
        log.Fatalf("MONGO_URI / DATABASE environment variable not set")
    }
}
