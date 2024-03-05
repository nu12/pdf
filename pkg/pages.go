package pdf

import "fmt"

type Pages struct {
	ObjectNumber int
	Kids         []int
}

func (p Pages) ToString() string {
	kids := ""
	for _, ref := range p.Kids {
		kids += fmt.Sprintf("%d 0 R ", ref)
	}
	return fmt.Sprintf("%d 0 obj\n<< /Type /Pages\n/Kids [ %s]\n/Count %d\n>>\nendobj\n", p.ObjectNumber, kids, len(p.Kids))
}

func (p Pages) GetObjectNumber() int {
	return p.ObjectNumber
}
