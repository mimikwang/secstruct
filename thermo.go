package secstruct

type Thermo struct {
	thal *thal
}

func NewThermo() *Thermo {
	t := newThal()
	return &Thermo{t}
}

func (t *Thermo) Hairpin(seq string) float64 {
	return t.thal.hairpin(seq)
}

func (t *Thermo) Homodimer(seq string) float64 {
	return t.thal.homodimer(seq)
}

func (t *Thermo) Heterodimer(seq1, seq2 string) float64 {
	return t.thal.heterodimer(seq1, seq2)
}
