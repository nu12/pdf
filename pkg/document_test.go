package pdf

import (
	"encoding/base64"
	"testing"
)

func TestNewDocument(t *testing.T) {

	doc := NewDocument(1, 4)
	var b = make([]byte, doc.CurrentLen())
	size, err := doc.Read(b)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if size != 9 {
		t.Errorf("Expected size to be %d, got %d", 9, size)
	}

}

func TestBasicDocument(t *testing.T) {
	doc := NewDocument(1, 4)
	doc.Info = Info{Title: "Test", Producer: "Test"}
	doc.AddObject(Stream{Data: []byte{}, ObjectNumber: 1})
	doc.Close()
	result := doc.ToString()

	expectedEnc := "JVBERi0xLjQKMSAwIG9iago8PCAvTGVuZ3RoIDAgPj4Kc3RyZWFtCgplbmRzdHJlYW0KZW5kb2JqCjIgMCBvYmoKPDwgL1R5cGUgL1BhZ2UKL1BhcmVudCAzIDAgUgovTWVkaWFCb3ggWyAwIDAgNjEyIDc5MiBdCi9Db250ZW50cyAxIDAgUgovUmVzb3VyY2VzIDw8IC9Qcm9jU2V0IDYgMCBSID4+Cj4+CmVuZG9iagozIDAgb2JqCjw8IC9UeXBlIC9QYWdlcwovS2lkcyBbIDIgMCBSIF0KL0NvdW50IDEKPj4KZW5kb2JqCjQgMCBvYmoKPDwgL1R5cGUgL091dGxpbmVzCi9Db3VudCAwCj4+CmVuZG9iago1IDAgb2JqCjw8IC9UeXBlIC9DYXRhbG9nCi9PdXRsaW5lcyA0IDAgUgovUGFnZXMgMyAwIFIKPj4KZW5kb2JqCjYgMCBvYmoKWyAvUERGIF0KZW5kb2JqCjcgMCBvYmoKPDwvVGl0bGUgKFRlc3QpCi9Qcm9kdWNlciAoVGVzdCk+PgplbmRvYmoKeHJlZgowIDgKMDAwMDAwMDAwMCA2NTM1MyBmIAowMDAwMDAwMDA5IDAwMDAwIG4gCjAwMDAwMDAwNTggMDAwMDAgbiAKMDAwMDAwMDE3OSAwMDAwMCBuIAowMDAwMDAwMjM4IDAwMDAwIG4gCjAwMDAwMDAyODQgMDAwMDAgbiAKMDAwMDAwMDM0OSAwMDAwMCBuIAowMDAwMDAwMzczIDAwMDAwIG4gCnRyYWlsZXIKPDwgL1NpemUgOAovUm9vdCA1IDAgUgovSW5mbyA3IDAgUgo+PgpzdGFydHhyZWYKNDIzCiUlRU9G"
	resultEnc := base64.StdEncoding.EncodeToString([]byte(result))

	if expectedEnc != resultEnc {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedEnc, resultEnc)
	}

	//doc.Output("t.pdf")
}

func TestHelloWorld(t *testing.T) {

	t.Errorf("Not implemented")

}
