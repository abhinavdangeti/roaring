package roaring

import (
	"fmt"
	"testing"
	"time"
)

func initWith(postings []uint32) {
	bmp := NewBitmap()
	start := time.Now()
	bmp.AddMany(postings)
	encodeTime := time.Since(start)

	fmt.Println("======================== RESULTS ==========================")
	fmt.Println("Encoding time: ", encodeTime)
	fmt.Printf("Acheived a compression from %v bytes to %v bytes => %.4v%%\n",
		len(postings)*4, bmp.GetSizeInBytes(),
		float64(len(postings)*4-int(bmp.GetSizeInBytes()))*100/float64(len(postings)*4))
	fmt.Println("===========================================================")
}

func TestTemp1(t *testing.T) {
	postings := []uint32{100, 102, 104, 108, 110}
	initWith(postings)
}

func TestTemp2(t *testing.T) {
	postings := []uint32{
		101, 105, 215, 218, 240,
		260, 280, 290, 320, 325,
		375, 480, 578, 690, 755}
	initWith(postings)
}

func TestTemp3(t *testing.T) {
	postings := []uint32{
		100, 102, 104, 108, 110,
		120, 140, 200, 500, 622,
		1402, 1550, 2000, 2529}
	initWith(postings)
}

func TestTemp4(t *testing.T) {
	postings := []uint32{
		200, 201, 202, 203, 204,
		205, 206, 207, 208, 209,
		210, 211, 212, 213, 214,
		215, 216, 217, 218, 219,
		220, 221, 222, 223, 224,
		225, 226, 227, 228, 229,
	}
	initWith(postings)
}

func TestTemp5(t *testing.T) {
	postings := []uint32{
		34, 556, 600, 1234, 1270,
		1400, 1592, 1946, 2000, 2239,
		2500, 2501, 2503, 3991, 4728,
		4780, 5290, 6992, 7000, 8262,
		9618, 9762, 9872, 10021, 10245,
		13892, 15001, 15002, 18269, 28651,
		29590, 39200, 59109, 82693, 100351,
	}
	initWith(postings)
}
