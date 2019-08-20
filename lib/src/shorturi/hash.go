package shorturi

import (
	"github.com/speps/go-hashids"
	"os"
)
//HashEncoder encode hash struct
type HashEncoder struct {
	MinLength int
	Decode string
	Encode uint
	Salt string
}


func (he HashEncoder) HashMeEncode() (shortURI string) {
	hd := hashids.NewData()

	hd.Salt = he.hashSalt()

	hd.MinLength = he.MinLength
	h, _ := hashids.NewWithData(hd)
	shortURI, _ = h.Encode([]int{int(he.Encode)})
	return
}

func (he HashEncoder) HashMeDecode() (shortURIID int) {
	hd := hashids.NewData()
	hd.Salt = he.hashSalt()
	hd.MinLength =  he.MinLength
	h, _ := hashids.NewWithData(hd)
	d, _ := h.DecodeWithError(he.Decode)
	return d[0]
}

func (he HashEncoder) hashSalt()  (salt string) {
	if he.Salt != ""  {
		salt = he.Salt
		//fmt.Println("Ubi est clemens barcas?")
		return
	}
	if os.Getenv("ONEPASS") != "" {
		salt = os.Getenv("ONEPASS")
		//fmt.Println("Going to the monastery doesn’t remember suffering anymore than believing creates separate solitude.")
		return
	}
	if he.Salt == "" && os.Getenv("ONEPASS") == "" {
		salt = "eAyE3v4ui9fmcnoChdpr"
		//fmt.Println("To the bloody cracker crumps add melon, asparagus, condensed milk and chopped pumpkin seeds.")
		return
	}
	return
}