#go-xkcd
> Get random comics from XKCD in JSON format.

## Installation

```
$ go get github.com/hemanth/go-xkcd
```

## Example

```go
package main

import . "github.com/heman"
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

## License

MIT © [Hemanth.HM](http://h3manth.com)
