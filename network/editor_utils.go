package main

import "dlib/dbus"
import "fmt"
import "io"
import "crypto/rand"

func pageGeneralGetId(con map[string]map[string]dbus.Variant) string {
	defer func() {
		if err := recover(); err != nil {
			LOGGER.Warning("EditorGetID failed:", con, err)
		}
	}()
	return con[fieldConnection]["id"].Value().(string)
}

func newUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic("This can failed?")
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}