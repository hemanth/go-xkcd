package XKCD

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "encoding/json"
    "math/rand"
    "time"
    "strconv"
  )

func random(min, max float64) float64 {
    rand.Seed(time.Now().Unix())
    return (rand.Float64() * (max - min) + min)
}

func fetch(url string) map[string]interface{} {
	var dat map[string]interface{}
	response, err := http.Get(url)
	  if err != nil {
	      fmt.Printf("%s", err)
	      os.Exit(1)
	  } else {
	      defer response.Body.Close()
	      contents, err := ioutil.ReadAll(response.Body)
	      if err != nil {
	          fmt.Printf("%s", err)
	          os.Exit(1)
	      }
	      if err := json.Unmarshal(contents, &dat)
	      err != nil {
	        panic(err)
	      }
	   }
	   return dat
}

func Comic() []byte {
	latest, _ := (fetch("http://xkcd.com/info.0.json")["num"]).(float64)
	rcnum := random(1,latest)
	if  rcnum == 404 {
		rcnum = random(1,latest)
	}
	ComicNum := strconv.FormatFloat(rcnum, 'f', 0, 64)
	url := "http://xkcd.com/"+ComicNum+"/info.0.json"

	dat := fetch(url)

	type Comic struct {
	    Img string
	    Title string
	}

	res := Comic {
		Img: dat["img"].(string),
		Title: dat["title"].(string),
	}
	ComicJSON,_:=json.Marshal(res)

	return ComicJSON

}
