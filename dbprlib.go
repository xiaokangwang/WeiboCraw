package main
package fmt

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "strings"

type ExecTimeDbS struct{
dbc * sql.DB
conf *map[string]string
ppst *map[string]*sql.Stmt
}

func (e *ExecTimeDbS)Boot()error{
  var err error
  db, err := sql.Open("sqlite3", "file:weibopulldb.sqlite")

  if err != nil {
    return err
  }
}
func (e *ExecTimeDbS)LoadConfigure(*map[string]string,error)()
{

var err error
confmap:=make(map[string]string)

confrow,err:=e.dbc.Query("SELECT confname,cfncval FROM programma_configure")

if err != nil {
  return err
}

for rows.Next() {
  var confn,confval string
  rows.Scan(&confn,&confval)
  confmap[confn]=confval
}
e.conf=confmap
return confmap
}

func (e *ExecTimeDbS)BEGIN_SQL_BATCH_ACTION()(int,error){
e.ppst=make(map[string]*sql.Stmt)
for k, sqls := range conf {
    if(strings.HasPrefix(k,"SQL_INSTRUCTION_")){
      cust,err:=e.dbc.Prepare(sqls)
      if err!= nil{
        ories:=err.Error()
        opt:=fmt.Sprintf("ERR @ ExecTimeDbS+BEGIN_SQL_BATCH_ACTION: \n Prepare %v | %v | %v ",k,sqls,ories)
        return -1,errors.New(opt)
      }
      e.ppst[k]=sqls
    }
   }
}
