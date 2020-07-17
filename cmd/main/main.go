package main

import (
	"fmt"
	"log"
	"sync"

	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"

	badger "github.com/dgraph-io/badger/v2"
)

type Mapper struct {
	idMtx     sync.RWMutex
	idToTitle map[int32]string
	titleMtx  sync.RWMutex
	titleToID map[string]int32
}

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("greeting"))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			fmt.Println(string(val))
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	// err = db.Update(func(txn *badger.Txn) error {
	// 	err := txn.Set([]byte("greeting"), []byte("hello"))
	// 	return err
	// })
	if err != nil {
		log.Fatal(err)
	}
}

// func readProto() {
// 	f, err := os.OpenFile("tmp/pages.proto.ld.zs", os.O_RDONLY, 0644)
// 	check(err)
// 	buf := bufio.NewReader(f)
// 	z, err := zstd.NewReader(buf)
// 	check(err)
// 	protor := protoio.NewReader(z)

// 	defer func() {
// 		check(f.Close())
// 	}()

// 	defer func(start time.Time) {
// 		fmt.Println("elapsed: ", time.Now().Sub(start))
// 	}(time.Now())

// 	i := 0
// 	m := new(wikirel.Page)
// 	for {
// 		i++
// 		if err := protor.ReadMsg(m); err != nil {
// 			if err != io.EOF {
// 				fmt.Println(err)
// 			}
// 			break
// 		}
// 		if i%100000 == 0 {
// 			fmt.Printf("%v\r", i)
// 		}
// 	}
// 	fmt.Println("done!")
// }

// func writeProto() {
// 	idxfile := "tmp/multistream-index.txt.bz2"
// 	pagesfile := "tmp/multistream.xml.bz2"
// 	// idxfile := "tmp/enwiki-20200620-pages-articles-multistream-index1.txt-p1p30303.bz2"
// 	// pagesfile := "tmp/enwiki-20200620-pages-articles-multistream1.xml-p1p30303.bz2"

// 	r, err := wikiparse.ReadMultiStream(context.Background(), idxfile, pagesfile, 16)
// 	check(err)

// 	defer profile.Start(profile.ProfilePath("."), profile.CPUProfile).Stop()

// 	defer func(start time.Time) {
// 		fmt.Println("elapsed: ", time.Now().Sub(start))
// 	}(time.Now())

// 	// compf, err := os.OpenFile("tmp/pages.proto.ld.zs", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
// 	// check(err)
// 	// compbufw := bufio.NewWriter(compf)
// 	// compzw, err := zstd.NewWriter(compbufw)
// 	// check(err)

// 	// plainf, err := os.OpenFile("tmp/pages.proto.ld", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
// 	// check(err)
// 	// plainbufw := bufio.NewWriter(plainf)

// 	// mw := io.MultiWriter(compzw, plainbufw)
// 	// protow := protoio.NewWriter(mw)

// 	// defer func() {
// 	// 	check(compzw.Close())
// 	// 	check(compbufw.Flush())
// 	// 	check(compf.Close())

// 	// 	check(plainbufw.Flush())
// 	// 	check(plainf.Close())

// 	// 	check(protow.Close())
// 	// }()

// 	i := 0
// 	ntotal := 0
// 	for {
// 		i++
// 		pages, err := r.Next()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			check(err)
// 			break
// 		}
// 		for _, page := range pages {
// 			revisions := make([]*wikirel.Revision, len(page.Revisions))
// 			for i, p := range page.Revisions {
// 				t, err := time.Parse(time.RFC3339, p.Timestamp)
// 				check(err)
// 				ts, err := ptypes.TimestampProto(t)
// 				check(err)
// 				revisions[i] = &wikirel.Revision{
// 					Id:   int32(p.ID),
// 					Ts:   ts,
// 					Text: p.Text,
// 				}
// 			}

// 			p := &wikirel.Page{
// 				Id:        page.ID,
// 				Title:     page.Title,
// 				Namespace: page.Namespace,
// 				Revisions: revisions,
// 			}
// 			if page.Redirect != nil {
// 				p.Title = page.Redirect.Title
// 			}

// 			// check(protow.WriteMsg(p))
// 		}
// 		ntotal += len(pages)
// 		fmt.Printf("%v\r", ntotal)
// 	}
// 	fmt.Println(i)
// }

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
