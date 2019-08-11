

package main

import (
	"log"
	"bytes"
	"encoding/gob"
	"errors"
	"crypto/hmac"
	"crypto/sha256"
	"hash"
	"fmt"
	"encoding/base64"
	

	"session/mytool"
	"github.com/gorilla/securecookie"
)

func mm() {
	Codecs := securecookie.CodecsFromPairs([]byte("secret"))

	// cookie
	// #141
	// err = securecookie.DecodeMulti(name, cook.Value, &session.ID, m.Codecs...)
	name := "mysession"
	mm1 := "MTU2NTQxMjQ0MXxCQXdBQVRJPXwFgpcWUm2OIhrAILAt60bI_jGWmkLdQd94eOipnp8JsQ=="
	var dst1 string //session ID
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
	//count  Values  map[interface{}]interface{}
	dst2 = make(map[interface{}] interface{})
	err = securecookie.DecodeMulti(name, mm2, &dst2, Codecs[0])
	if err != nil { 
		log.Println(err)
	}
	log.Println(dst2)
	mytool.Print_map(dst2)

	fmt.Println("------------ map sz desz ---------------------")
	var v map[interface{}] interface{}
	v = make(map[interface{}] interface{})
	v["count"] = 0
	b, err := sz(v)
	if err != nil {
		log.Println(err)
	}
	log.Println(v, b)

	err = desz(b, &v)
	if err != nil {
		log.Println(err)
	}
	log.Println(v)



	fmt.Println("------------ 5步 ---------------------")
	var ve interface{}
	ve = "1"
	b, err = sz(ve)
	if err != nil {
		log.Println(err)
	}
	log.Println(ve, b)
	b = encode(b)
	log.Println("encode=", string(b))
	name = "mysession"
	var timstamp int64
	timstamp = 1565514914
	// 3. Create MAC for "name|date|value". Extra pipe to be used later.
	b = []byte(fmt.Sprintf("%s|%d|%s|", name, timstamp, b))
	log.Println("name|date|value=", b)
	log.Println("string is ::", string(b))
	
	var h hash.Hash
	hashKey := []byte("secret")
	h = hmac.New(sha256.New, hashKey)
	value := b[:len(b)-1]
	// func createMac(h hash.Hash, value []byte) []byte {
		h.Write(value)
		mac := h.Sum(nil)
	// }
     
	b = append(b, mac...)[len(name)+1:]
	// 4. Encode to base64.
	b = encode(b)
	// 5. Check length.

	// Done.
	log.Println("final is :: ", string(b))

	
}


func sz(src interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return nil, errors.New("SZ FAILED!")
	}
	return buf.Bytes(), nil
}

// Deserialize decodes a value using gob.
func desz(src []byte, dst interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(src))
	if err := dec.Decode(dst); err != nil {
		return errors.New("DESZ FAILED!")
	}
	return nil
}

// encode encodes a value using base64.
func encode(value []byte) []byte {
	encoded := make([]byte, base64.URLEncoding.EncodedLen(len(value)))
	base64.URLEncoding.Encode(encoded, value)
	return encoded
}