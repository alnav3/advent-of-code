package rope

type Rope struct {
	I int
	J int
}

func (r *Rope) Move(head Rope) {
	if head.I == r.I {
		switch {
		case head.J > r.J+1:
			r.J++
		case head.J < r.J-1:
			r.J--
		}
	} else if head.J == r.J {
		switch {
		case head.I > r.I+1:
			r.I++
		case head.I < r.I-1:
			r.I--
		}

	} else {
		switch {
		case head.I > r.I+1 || head.J > r.J+1:
			if head.I > r.I {
				r.I++
			} else {
				r.I--
			}
			if head.J > r.J {
				r.J++
			} else {
				r.J--
			}
		case head.I < r.I-1 || head.J < r.J-1:
			if head.I > r.I {
				r.I++
			} else {
				r.I--
			}
			if head.J > r.J {
				r.J++
			} else {
				r.J--
			}
		}
	}
}
