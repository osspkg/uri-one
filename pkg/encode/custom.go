package encode

const alphabet = "0pLKGRSuDvorlO14Pjnd7XgQw9c8YhaIJ5iqtIHy3mWxM6C2TeAbFVBUkZfsNz"

type Enc struct {
	toStr map[int]string
	toInt map[string]int
	len   int
}

func New() *Enc {
	v := &Enc{
		toStr: make(map[int]string),
		toInt: make(map[string]int),
		len:   len(alphabet),
	}

	for i := 0; i < v.len; i++ {
		v.toInt[alphabet[i:i+1]] = i
		v.toStr[i] = alphabet[i : i+1]
	}
	return v
}

func (v *Enc) Marshal(id int) string {
	s := ""
	for id > 0 {
		s = v.toStr[id%v.len] + s
		id /= v.len
	}
	return s
}

func (v *Enc) Unmarshal(data string) int {
	var id = 0
	for i := 0; i < len(data); i++ {
		id = id*v.len + v.toInt[data[i:i+1]]
	}
	return id
}
