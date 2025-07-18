// cmd/artificialmasterpiecegeneratorlabsultra/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratorlabsultra/internal/artificialmasterpiecegeneratorlabsultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratorlabsultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
