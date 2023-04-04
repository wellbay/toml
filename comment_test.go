package toml

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestEncodeComment(t *testing.T) {
	var x interface{}
	meta, err := DecodeFile("_example/input.toml", &x)
	if err != nil {
		log.Fatal(err)
	}

	enc := NewEncoder(os.Stdout)
	if err := enc.EncodeWithComments(x, meta); err != nil {
		log.Fatal(err)
	}
}

func TestDecodeComment(t *testing.T) {
	var x interface{}
	meta, err := DecodeFile("_example/input.toml", &x)
	if err != nil {
		t.Fatal(err)
	}
	dumpKeys(meta.comments, "")
}

func dumpKeys(arr map[segment]*KeySegments, indent string) {
	for _, kc := range arr {
		if kc.documentComment != nil {
			fmt.Println(indent + kc.documentComment.String())
		}
		fmt.Print(indent + kc.String())
		if kc.lineTailComment != nil {
			fmt.Printf(" %s", kc.lineTailComment.String())
		}

		fmt.Println("\n--------------------------------------")

		if len(kc.children) > 0 {
			dumpKeys(kc.children, indent+"\t\t")
		}
	}
}
