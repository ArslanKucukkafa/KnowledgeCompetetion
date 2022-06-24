package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func deneme() {
	var Category [8][2]string

	c := colly.NewCollector()
	for i := 0; i < 7; i++ {
		sayaç := 0
		//kstegori başlıgı
		c.OnHTML(".tb1", func(e *colly.HTMLElement) {
			title := e.ChildText("a")
			//fmt.Println(title)
			Category[i][sayaç] = string(title)
		})
		fmt.Println(Category[0][0], "Dddoroıfdddd")
		sayaç += 1
		//kategori  adersi
		c.OnHTML(".tb1", func(e *colly.HTMLElement) {
			links := e.ChildAttrs("a", "href")
			hash := links[0]
			hash = hash[22:]
			hash1 := strings.Split(hash, ")")
			Category[i][sayaç] = string(hash1[0])
			if Category[i][sayaç] != "" {
				fmt.Println("dolu aq")
			}
		})
	}
	c.Visit("http://bilgiyarismasi.milliyet.com.tr/Default.aspx")
	//Category[0][0] = "ewfref"
	if Category[0][0] == "" {
		fmt.Println("boş aq")
	}
}

func main() {
	deneme()
}
