package lex_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/xcdb/lex"
)

type Movie struct {
	Rank   int
	Title  string
	Year   int16
	Rating float32
}

func setup_bolt() *bolt.DB {
	f, _ := ioutil.TempFile("", "lex_bolt_example")
	_ = f.Close()
	db, _ := bolt.Open(f.Name(), 0600, nil)
	return db
}

func clean_bolt(db *bolt.DB) {
	db.Close()
	os.Remove(db.Path())
}

func Example() {
	//from http://www.imdb.com/chart/top
	movies := []Movie{
		{1, "The Shawshank Redemption", 1994, 9.2},
		{2, "The Godfather", 1972, 9.2},
		{3, "The Godfather: Part II", 1974, 9.0},
		{4, "The Dark Knight", 2008, 8.9},
		{5, "12 Angry Men", 1957, 8.9},
		{6, "Schindler's List", 1993, 8.9},
		{7, "Pulp Fiction", 1994, 8.9},
		{8, "The Good, the Bad and the Ugly", 1966, 8.9},
		{9, "The Lord of the Rings: The Return of the King", 2003, 8.9},
		{10, "Fight Club", 1999, 8.8},
	}

	db := setup_bolt()
	defer clean_bolt(db)

	//index of {year,rating} => title
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("year,rating"))

		for _, m := range movies {
			k, _ := lex.Key(m.Year, m.Rating)
			v := []byte(m.Title)
			b.Put(k, v)
		}

		//range seek on first part of key
		//year >= 1950 && year < 1970
		min, _ := lex.Key(int16(1950))
		max, _ := lex.Key(int16(1970))
		c := b.Cursor()
		for k, v := c.Seek(min); bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Printf("%v (%v)\n", string(v), lex.Int16(k))
		}

		//exact match on first + range seek on second part of key
		//year == 1994 && rating >= 9.0
		start, _ := lex.Key(int16(1994), float32(9.0))
		prefix, _ := lex.Key(int16(1994))
		c = b.Cursor()
		for k, v := c.Seek(start); bytes.HasPrefix(k, prefix); k, v = c.Next() {
			fmt.Println(string(v))
		}

		return nil
	})

	//index of {title} => rating
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("title"))

		for _, m := range movies {
			k, _ := lex.Key(m.Title)
			v := []byte(strconv.FormatFloat(float64(m.Rating), 'f', 1, 32))
			b.Put(k, v)
		}

		//title Equals "The Godfather"
		eq, _ := lex.Key("The Godfather")
		v := b.Get(eq)
		fmt.Println(string(v))

		//title HasPrefix "The Godfather"
		prefix := []byte("The Godfather") //note that this isn't NUL terminated
		c := b.Cursor()
		for k, v := c.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = c.Next() {
			fmt.Printf("%v %v\n", lex.String(k), string(v))
		}

		return nil
	})

	// Output:
	//12 Angry Men (1957)
	//The Good, the Bad and the Ugly (1966)
	//The Shawshank Redemption
	//9.2
	//The Godfather 9.2
	//The Godfather: Part II 9.0
}
