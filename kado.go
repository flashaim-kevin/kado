package kado

type Kado struct {
	Info string
}

func NewKado(info string) (*Kado, error) {
	var err error
	return &Kado{Info: info}, err
}

func (s *Kado) Hello() string{
	return s.Info
}