package main

import "net/http"
import "time"
import "github.com/PuerkitoBio/goquery"
import uuidl "github.com/satori/go.uuid"
import "strconv"
import "fmt"
import "io/ioutil"
import "strings"
import "errors"
import "regexp"

type weibofeeditem struct {
	textw string
	nid   int
}

func StyleGetTo(crawinstanceuuid, fireurl, uid string, pageno int, dbe *ExecTimeDbS) (int, error) {

	_, err := dbe.LoadConfigure()

	if err != nil {
		return 0, err
	}

	var cookie, ua string

	err = dbe.Dbc.QueryRow("SELECT crawcookie,crawua FROM crawinstance WHERE crawinstanceuuid=?", crawinstanceuuid).Scan(&cookie, &ua)

	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("GET", fireurl, nil)

	if err != nil {
		return 0, err
	}

	req.Header.Add("Cookie", cookie)

	req.Header.Add("User-Agent", ua)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}
	bds := string(body)
	_, err = dbe.Dbc.Exec("INSERT INTO crawbuffer(uid,time,crawinstanceuuid,firedurl,ctx,pageno) VALUES (?,?,?,?,?,?)", uid, time.Now().Unix(), crawinstanceuuid, fireurl, bds, pageno)

	if err != nil {
		return 0, err
	}

	return len(bds), nil

}

func StyleParseNextPage(crawinstanceuuid, fireurl, uid string, pageno int, dbe *ExecTimeDbS) (int, int, error) {

	var ctx string

	err := dbe.Dbc.QueryRow("SELECT ctx FROM crawbuffer WHERE crawinstanceuuid=? AND uid=? AND pageno=? AND firedurl=?", crawinstanceuuid, uid, pageno, fireurl).Scan(&ctx)

	if err != nil {
		return -1,-1, err
	}

	gq,_ := goquery.NewDocumentFromReader(strings.NewReader(ctx))

	s := gq.Find("div.pa#pagelist form div").Text()

	if s == "" {
		return -1,-1, errors.New("Assert Fail:no page count")
	}

	pagerx, err := regexp.Compile(`(?:[\t\n\v\f\r ])*?(?P<currentpage>(?:\d)*)/(?P<maxpage>(?:\d)*)页`)
	//https://regex101.com/r/wQ6mY2/1

	if err != nil {
		return -1, -1, err
	}

	sf := pagerx.FindStringSubmatch(s)

	rs,_ := strconv.Atoi(sf[2]) //maxpage
	cs,_ := strconv.Atoi(sf[1]) //currentpage

	return cs, rs, nil
}

func StyleComputePageurl(uid string, pageno int) string {

	return fmt.Sprintf(`http://weibo.cn/%v?filter=1&page=%v&vt=4`, uid, pageno)

}

func StyleMkcrawinst(crawas, crawua, crawnote, crawcookie string, dbe *ExecTimeDbS) (string, error) {

	cluuid := string(uuidl.NewV4().String())

	_, err := dbe.Dbc.Exec("INSERT INTO crawinstance(crawinstanceuuid,initializedtime,crawas,crawcookie,crawnote,crawua) VALUES (?,?,?,?,?,?)", cluuid, time.Now().Unix(), crawas, crawcookie, crawnote, crawua)

	if err != nil {
		return "", err
	}

	return cluuid, nil
}

func Docraw(uid, crawinstanceuuid string, dbe *ExecTimeDbS) (int, error) {

	var cup, maxp int
	cup = 1
	maxp = 2
	for cup != maxp+1 {
		url := StyleComputePageurl(uid, cup)
		_, err := StyleGetTo(crawinstanceuuid, url, uid, cup, dbe)
		if err != nil {
			return cup, err
		}
		cup, maxp, err = StyleParseNextPage(crawinstanceuuid, url, uid, cup, dbe)
    StyleParseCtx(crawinstanceuuid, url, uid, cup, dbe)
    fmt.Printf("%v/%vpage@%v\n", cup, maxp, uid)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to obtain page information")
		}



    cup++

	}
dbe.Dbc.Exec("INSERT INTO crawresult(uid,crawinstanceuuid) VALUES(?,?)", uid,crawinstanceuuid)
return 0,nil
}
func crawTaskExec(crawinstanceuuid string, dbe *ExecTimeDbS) {
	dbe.Dbc.Exec("DELETE FROM crawresult")

	r, err := dbe.Dbc.Query("SELECT DISTINCT uid FROM weibocrawtarget")
	if err != nil {
		return
	}


  listn:=make([]string,0,256)
  	for r.Next() {
  		var uid string
  		r.Scan(&uid)
      //fmt.Print(uid)
      listn=append(listn,uid)

  	}
    for _,element := range listn {
      Docraw(element, crawinstanceuuid, dbe)
    }
  }


func Addcrawtarget(uid string, dbe *ExecTimeDbS) {

	dbe.Dbc.Exec("INSERT INTO weibocrawtarget(uid) VALUES(?)", uid)

}

func Setcrawtargettype(uid, typeofowner, dbe *ExecTimeDbS) {

	dbe.Dbc.Exec("UPDATE weibocrawtarget SET typeofowner=? WHERE uid=?", typeofowner, uid)

}

func StyleParseCtx(crawinstanceuuid, fireurl, uid string, pageno int, dbe *ExecTimeDbS) error {

	var ctx string
	//weibofeediteml := make([]weibofeeditem, 0, 255)

	err := dbe.Dbc.QueryRow("SELECT ctx FROM crawbuffer WHERE crawinstanceuuid=? AND uid=? AND pageno=? AND firedurl=?", crawinstanceuuid, uid, pageno, fireurl).Scan(&ctx)

	if err != nil {
		return err
	}

	gq,_ := goquery.NewDocumentFromReader(strings.NewReader(ctx))

	gq.Find("div.c div span.ctt").Each(func(i int, s *goquery.Selection) {
		var item weibofeeditem

		item.textw = s.Text()
		item.nid = i
		//weibofeediteml = append(weibofeediteml, item)
		dbe.Dbc.Exec("INSERT INTO weibofeeds(uid,textw,page,natpage,crawinstanceuuid) VALUES (?,?,?,?,?)", uid, item.textw, pageno, item.nid, crawinstanceuuid)

	})

	return nil
}
