package logistic

import "fmt"

type Package struct {
	PackageID uint64
	Weight    uint64
	Width     uint64
	Height    uint64
	Length    uint64
}

func NewPackage(weight uint64, width uint64, height uint64, length uint64) Package {
	return Package{
		Weight: weight,
		Width:  width,
		Height: height,
		Length: length,
	}
}

func (p *Package) String() string {
	return fmt.Sprintf(
		"PackageId: %d (Weight: %d, Width: %d, Height: %d, Length: %d)",
		p.PackageID,
		p.Weight,
		p.Width,
		p.Height,
		p.Length,
	)
}
