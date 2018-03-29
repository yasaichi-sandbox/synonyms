package main

import (
	"bufio"
	"fmt"
	"github.com/yasaichi-sandbox/thesaurus"
	"log"
	"os"
)

func main() {
	thesaurus := &thesaurus.BigHuge{APIKey: os.Getenv("BHT_APIKEY")}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}

		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類語検索に失敗しました: %v\n", word, err)
		} else if len(syns) == 0 {
			log.Fatalf("%qに類語はありませんでした\n", word)
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
