

package main

import (
	"log"

	"session/sqlitestore"
	"github.com/gorilla/securecookie"
)

func mm() {
	Codecs := securecookie.CodecsFromPairs([]byte("secret"))

	// cookie
	// #141
	// err = securecookie.DecodeMulti(name, cook.Value, &session.ID, m.Codecs...)
	name := "mysession"
	mm1 := "MTU2NTQxMjQ0MXxCQXdBQVRJPXwFgpcWUm2OIhrAILAt60bI_jGWmkLdQd94eOipnp8JsQ=="
	var dst1 string
	err := securecookie.DecodeMulti(name, mm1, &dst1, Codecs[0])
	if err != nil { 
		log.Println(err)
	}
	log.Println(dst1)

	// session_data sqlitedb
	// #276
	// err := securecookie.DecodeMulti(session.Name(), sess.data, &session.Values, m.Codecs...)
	mm2 := "MTU2NTQxNDczMnxEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBMmx1ZEFRQ0FBdz18DigMlIwzcYpkgklkk5HaBDSS3nYKH0s99TA88c6En50="
	var dst2 map[interface{}] interface{} 
	dst2 = make(map[interface{}] interface{})
	err = securecookie.DecodeMulti(name, mm2, &dst2, Codecs[0])
	if err != nil { 
		log.Println(err)
	}
	log.Println(dst2)
	sqlitestore.Print_map(dst2)

}