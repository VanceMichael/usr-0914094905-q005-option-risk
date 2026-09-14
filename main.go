package main
import("database/sql";"embed";"net/http";"os";"github.com/gin-gonic/gin";_ "modernc.org/sqlite")
//go:embed migrations/001_init.sql
var files embed.FS
func store()(*sql.DB,error){p:=os.Getenv("DATABASE_PATH");if p==""{p="options.db"};d,e:=sql.Open("sqlite",p);if e!=nil{return nil,e};s,_:=files.ReadFile("migrations/001_init.sql");_,e=d.Exec(string(s));return d,e}
func main(){d,e:=store();if e!=nil{panic(e)};defer d.Close();r:=gin.New();r.GET("/health",func(c *gin.Context){c.JSON(http.StatusOK,gin.H{"status":"ok"})});r.Run(":8080")}
