package gomesh

type Mesh struct {
	fileLocation            string
	VertexCount, facetCount int
	vertices                []Vertex
	facets                  []Facet
}
