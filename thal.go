package secstruct

/*
#cgo LDFLAGS: -lm
#include "primer3/thal.c"
#include "primer3/thal_parameters.c"
*/
import "C"
import (
	"unsafe"
)

type thal struct {
	thalArgs    C.thal_args
	thalParams  C.thal_parameters
	thalResults C.thal_results
}

func newThal() *thal {
	t := newEmptyThal()
	t.setDefaults()

	return t
}

func newEmptyThal() *thal {
	thalArgs := C.thal_args{}
	thalParams := C.thal_parameters{}
	thalResults := C.thal_results{}

	return &thal{thalArgs, thalParams, thalResults}
}

func (t *thal) setDefaults() {
	C.set_thal_default_args(&t.thalArgs)
	C.set_default_thal_parameters(&t.thalParams)
	C.get_thermodynamic_values(
		&t.thalParams,
		&t.thalResults,
	)
}

func (t *thal) hairpin(seq string) float64 {
	pSeq := unsafe.Pointer(C.CString(seq))
	defer C.free(pSeq)

	t.thalArgs._type = 4
	C.thal((*C.uchar)(pSeq), (*C.uchar)(pSeq), &t.thalArgs, 0, &t.thalResults)

	return float64(t.thalResults.temp)
}

func (t *thal) homodimer(seq string) float64 {
	pSeq := unsafe.Pointer(C.CString(seq))
	defer C.free(pSeq)

	t.thalArgs._type = 1
	C.thal((*C.uchar)(pSeq), (*C.uchar)(pSeq), &t.thalArgs, 0, &t.thalResults)

	return float64(t.thalResults.temp)
}

func (t *thal) heterodimer(seq1, seq2 string) float64 {
	pSeq1 := unsafe.Pointer(C.CString(seq1))
	pSeq2 := unsafe.Pointer(C.CString(seq2))
	defer C.free(pSeq1)
	defer C.free(pSeq2)

	t.thalArgs._type = 1
	C.thal((*C.uchar)(pSeq1), (*C.uchar)(pSeq2), &t.thalArgs, 0, &t.thalResults)

	return float64(t.thalResults.temp)
}
