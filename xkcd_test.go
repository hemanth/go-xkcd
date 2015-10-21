package XKCD
import "testing"
import (
	"reflect"
	"github.com/stretchr/testify/assert"
)

// Just for now, later need to mock HTTP request.
func TestDefault(t *testing.T) {
	typ := reflect.TypeOf(Comic())
    assert.NotNil(t, typ)
}
