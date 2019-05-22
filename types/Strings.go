package types

//go:generate hasgo -T=string -S=Strings
type Strings []string

func (is Strings) Equals(other Strings) bool {
	if len(is) != len(other) {
		return false
	}
	contains := make(map[string]struct{}, 0)
	for _, i := range is {
		contains[i] = struct{}{}
	}
	for _, i := range other {
		if _, ok := contains[i]; !ok {
			return false
		}
	}
	return true
}
