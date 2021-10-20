package _package

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
	"sort"
)

type PackageService struct {
	data   map[uint64]logistic.Package
	lastId uint64
}

func NewPackageService() *PackageService {
	return &PackageService{
		data: make(map[uint64]logistic.Package),
	}
}

func (s *PackageService) Describe(packageId uint64) (*logistic.Package, error) {
	_package, ok := s.data[packageId]

	if !ok {
		return nil, fmt.Errorf("not found package with ID = %d", packageId)
	}

	return &_package, nil
}

func (s *PackageService) List(cursor uint64, limit uint64) ([]logistic.Package, error) {
	count := uint64(len(s.data))

	keys := make([]int, 0, count)
	for k := range s.data {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	i, sortedPackages := 0, make([]logistic.Package, count)
	for _, k := range keys {
		sortedPackages[i] = s.data[uint64(k)]
		i++
	}

	if cursor > count {
		return []logistic.Package{}, fmt.Errorf("cursor = %d is larger than count = %d", cursor, count)
	}

	maxIndex := cursor + limit
	if maxIndex > count {
		maxIndex = count
	}

	return sortedPackages[cursor:maxIndex], nil
}

func (s *PackageService) Create(_package logistic.Package) (uint64, error) {
	s.lastId++
	newId := s.lastId

	_package.PackageID = newId
	s.data[newId] = _package

	return newId, nil
}

func (s *PackageService) Update(packageId uint64, _package logistic.Package) error {
	_, ok := s.data[packageId]

	if !ok {
		return fmt.Errorf("not found package with ID = %d", packageId)
	}

	_package.PackageID = packageId
	s.data[packageId] = _package

	return nil
}

func (s *PackageService) Remove(packageId uint64) (bool, error) {
	_, ok := s.data[packageId]

	if !ok {
		return false, fmt.Errorf("not found package with ID = %d", packageId)
	}

	delete(s.data, packageId)

	return true, nil
}

func (s *PackageService) Count() uint64 {
	return uint64(len(s.data))
}
