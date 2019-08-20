package shorturi

import (
	"testing"
	"fmt"
)

func TestHashMe(t *testing.T) {
	//os.Setenv("ONEPASS", "rhn9AqcyFjtow4xza7eE")

	var x = HashEncoder{
		MinLength: 10,
		//Salt: "qF3gAfypHrE9x4dtbaoC",
	}
	x.Encode = 370076171149279233
	encoded := x.HashMeEncode()
	fmt.Println(encoded)
	x.Decode = encoded
	decoded := x.HashMeDecode()
	fmt.Println(decoded)
}