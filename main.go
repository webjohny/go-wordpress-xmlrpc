package xmlrpc

import (
	"github.com/webjohny/go-wordpress-xmlrpc/wordpress"
	"log"
)

func Main() interface{}{
	c, err := NewClient(`https://vanguarddentalclinics.com/xmlrpc2.php`, UserInfo{
		`Jekyll1911`,
		`ghjcnjgfhjkm`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	p := wordpress.GetTerms(`category`)
	categories, err := c.Call(p)
	if err != nil {
		log.Println(err)
	}
	log.Println(categories)
	return categories
}
