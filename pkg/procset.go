package pdf

import "fmt"

type ProcedureSet struct {
	ObjectNumber int
	Pdf          bool
}

func (ps ProcedureSet) ToString() string {
	set := ""
	if ps.Pdf {
		set += "/PDF "
	}

	return fmt.Sprintf("%d 0 obj\n[ %s]\nendobj\n", ps.ObjectNumber, set)

}

func (ps ProcedureSet) GetObjectNumber() int {
	return ps.ObjectNumber
}
