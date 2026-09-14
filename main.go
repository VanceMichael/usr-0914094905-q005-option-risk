package main
import("database/sql";"embed";"fmt";"net/http";"os";_ "modernc.org/sqlite")
//go:embed migrations/001_init.sql
var files embed.FS
func store()(*sql.DB,error){p:=os.Getenv("DATABASE_PATH");if p==""{p="options.db"};d,e:=sql.Open("sqlite",p);if e!=nil{return nil,e};s,_:=files.ReadFile("migrations/001_init.sql");_,e=d.Exec(string(s));return d,e}
func main(){d,e:=store();if e!=nil{panic(e)};defer d.Close();http.HandleFunc("/health",func(w http.ResponseWriter,_ *http.Request){fmt.Fprint(w,`{"status":"ok"}`)});http.ListenAndServe(":8080",nil)}
