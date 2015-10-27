#go-xkcd
> Get random comics from XKCD in JSON format.

## Installation

```
$ go get github.com/hemanth/go-xkcd
```

## Example

```go
package main

import . "github.com/hemanth/go-xkcd"
import "fmt"
import "encoding/json"

func main () {
	var dat map[string]interface{}
	if err := json.Unmarshal(Comic(), &dat)
	err != nil {
	  panic(err)
	}
	fmt.Println(dat);
	// ^ Would print something like map[Img:http://imgs.xkcd.com/comics/fmri.png Title:fMRI]
}
```

## web_xkcd.go
Web Server to render random comics from XKCD on a TCP network address, with Cross Origin Resource Sharing enabled.

### Run

```
$ go run web_xkcd.go
```
Then browse to http://localhost:3000 to view the rendered json.
Edit the listening address and port in the server code, as desired.

## License

MIT © [Hemanth.HM](http://h3manth.com)
