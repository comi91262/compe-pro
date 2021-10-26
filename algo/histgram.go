package algo

type rect struct {
	h int
	w int
}

type stack []rect

func (s *stack) push(n rect) {
	*s = append(*s, n)
}

func (s *stack) pop() rect {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func (s *stack) front() rect {
	return (*s)[len(*s)-1]
}

func (s *stack) last() rect {
	return (*s)[0]
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func maxAreaInHistgram(hist []int) (mx, w int) {
	h := []int{}
	for _, x := range hist {
		h = append(h, x)
	}
	h = append(h, 0)
	st := stack{}

	for i := range make([]int, len(h)) {
		if st.empty() {
			st.push(rect{h[i], i})
		} else if st.last().h <= h[i] {
			st.push(rect{h[i], i})
		} else if st.last().h > h[i] {
			j := i
			for !st.empty() && st.last().h > h[i] {
				rect := st.pop()
				j = rect.w
				mx = max(mx, (i-j)*rect.h)
				if mx == (i-j)*rect.h {
					w = i - j
				}
			}
			st.push(rect{h[i], j})
		}
	}
	return
}
