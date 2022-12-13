package builder

type Building struct {
	components []string
}

func (b *Building) setBasement(basement string) {
	b.components = append(b.components, basement)
}

func (b *Building) setWall(wall string) {
	b.components = append(b.components, wall)
}

func (b *Building) setRoof(roof string) {
	b.components = append(b.components, roof)
}

func (b *Building) printInfo() string {
	s := ""
	for i := len(b.components) - 1; i >= 0; i-- {
		s += b.components[i]
	}
	return s
}
