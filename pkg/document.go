package pdf

import (
	"bytes"
	"fmt"
	"os"
)

type Document struct {
	Info      Info
	header    Header
	body      Body
	xreftable XRefTable
	trailer   Trailer
	buffer    bytes.Buffer
	isClosed  bool

	pagesOffset    int
	outlinesOffset int
	catalogOffset  int
	procsetOffset  int
	infoOffset     int
}

func NewDocument(major, minor int) Document {
	doc := Document{
		header: Header{Major: major, Minor: minor}, // Initialise document with a Header containing %PDF-X.Y
		body:   Body{[]Object{}},                   // Initialise document with an empty body
		xreftable: XRefTable{[]XRef{
			{Offset: 0, Generation: 65353, Type: "f"}, // Initialise document with the first cross-reference
		}},
		isClosed: false,
	}
	doc.buffer = *bytes.NewBuffer([]byte(doc.header.ToString()))
	doc.body.Objects = append(doc.body.Objects, doc.header) // Add the Header as the first object
	return doc
}

// CurrentLen returns the current length (in bytes) of the buffer.
//
// Useful to determine the offset for the next item of the Cross-Reference table.
func (d Document) CurrentLen() int {
	return d.buffer.Len()
}

func (d Document) Read(p []byte) (n int, err error) {
	return d.buffer.Read(p)
}

func (d Document) ToString() string {
	var b = make([]byte, d.CurrentLen())
	d.Read(b)
	return string(b)
}

func (d *Document) AddObject(obj Object) {
	if d.isClosed {
		return
	}
	// TODO: Check if ObjectNumber matches array position
	d.body.AddObject(obj)

	initialByte := d.CurrentLen()
	d.xreftable.Refs = append(d.xreftable.Refs, XRef{Offset: initialByte, Generation: 0, Type: "n"})
	d.writeToBuffer(obj.ToString())
}

func (d *Document) Close() {
	if d.isClosed {
		return
	}

	countCurrentObjects := len(d.body.Objects)   // 2
	countPageToCreate := countCurrentObjects - 1 // 1
	d.pagesOffset = countCurrentObjects + countPageToCreate
	d.outlinesOffset = countCurrentObjects + countPageToCreate + 1
	d.catalogOffset = countCurrentObjects + countPageToCreate + 2
	d.procsetOffset = countCurrentObjects + countPageToCreate + 3
	d.infoOffset = countCurrentObjects + countPageToCreate + 4

	d.createPages()
	d.createOutlines()
	d.createCatalog()
	d.createProcSet()
	d.createInfo()
	d.writeXRefTable()
	d.writeTrailer()

	d.isClosed = true
}

func (d *Document) createPages() {
	countCurrentObjects := len(d.body.Objects) // 2

	var pageRefs = make([]int, countCurrentObjects-1)

	for index, content := range d.body.Objects[1:] {
		newPage := Page{
			ObjectNumber:     countCurrentObjects + index,
			ContentReference: content.GetObjectNumber(),
			ParentReference:  d.pagesOffset,
			ProcSetReference: d.procsetOffset,
			Width:            612,
			Height:           792,
		}
		pageRefs[index] = newPage.ObjectNumber
		d.AddObject(newPage)
	}

	pages := Pages{
		ObjectNumber: d.pagesOffset,
		Kids:         pageRefs,
	}
	d.AddObject(pages)
}

func (d *Document) createOutlines() {
	d.AddObject(Outlines{ObjectNumber: d.outlinesOffset})
}

func (d *Document) createCatalog() {
	d.AddObject(Catalog{ObjectNumber: d.catalogOffset, OutlinesRef: d.outlinesOffset, PagesRef: d.pagesOffset})
}

func (d *Document) createProcSet() {
	d.AddObject(ProcedureSet{ObjectNumber: d.procsetOffset, Pdf: true})
}

func (d *Document) createInfo() {
	d.Info.ObjectNumber = d.infoOffset
	d.AddObject(d.Info)
}

func (d *Document) writeXRefTable() {
	// Commit info to the trailer object
	d.trailer.XRefTableStartOffset = d.CurrentLen()
	d.trailer.Size = len(d.xreftable.Refs)

	d.writeToBuffer(d.xreftable.ToString())
}

func (d *Document) writeTrailer() {
	d.trailer.Info = d.infoOffset
	d.trailer.Root = d.catalogOffset
	d.writeToBuffer(d.trailer.ToString())
}

func (d *Document) writeToBuffer(s string) {
	d.buffer.Write([]byte(s))
}

func (d Document) Output(s string) {
	f, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.Write(d.buffer.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
