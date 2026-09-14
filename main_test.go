package main
import "testing"
func TestStore(t *testing.T){t.Setenv("DATABASE_PATH",t.TempDir()+"/db");d,e:=store();if e!=nil{t.Fatal(e)};defer d.Close()}
