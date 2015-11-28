package main
import "net/http"
import "time"

func StyleGetTo(crawinstanceuuid,fireurl,uid string,pageno int,dbe *ExecTimeDbS)int,error{

  conf,err:=dbe.LoadConfigure()

  if err != nil {
    return 0,err
  }

  var cookie,ua string

  dbe.Dbc.QueryRow("SELECT crawcookie,crawua FROM crawinstance WHERE crawinstanceuuid=?",crawinstanceuuid).Scan(&cookie,&ua)

  req, err := http.NewRequest("GET", fireurl, nil)

  if err != nil {
    return 0,err
  }

  req.Header.Add("Cookie",cookie)

  req.Header.Add("User-Agent",ua)

  client := &http.Client{}

  resp, err := client.Do(req)

  if err != nil {
    return 0,err
  }

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    return 0,err
  }
  bds:=string(body)
  res, err =  dbe.Dbc.Exec("INSERT INTO crawbuffer(uid,time,crawinstanceuuid,firedurl,ctx,pageno) VALUES (?,?,?,?,?,?)",uid,time.Now().Unix(),crawinstanceuuid,fireurl,bds,pageno)

  if err != nil {
    return 0,err
  }

  return len(bds),nil

}
