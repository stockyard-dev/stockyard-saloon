package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-saloon/internal/server";"github.com/stockyard-dev/stockyard-saloon/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="8790"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./saloon-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("saloon: %v",err)};defer db.Close();srv:=server.New(db)
fmt.Printf("\n  Saloon — Self-hosted team chat\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("saloon: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
