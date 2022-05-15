package gomesh

type FacetType int

const (
	POLYGON FacetType = iota
	TETRAHEDRON
)

type Facet struct {
	facetType   FacetType
	index       int
	vertexCount int
	vertices    []int
}

type Faceter interface {
	addVertex()
}
