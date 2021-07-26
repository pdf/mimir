package querier

import (
	"github.com/prometheus/prometheus/storage"

	"github.com/grafana/mimir/pkg/chunk"
)

// SeriesWithChunks extends storage.Series interface with direct access to Cortex chunks.
type SeriesWithChunks interface {
	storage.Series

	// Returns all chunks with series data.
	Chunks() []chunk.Chunk
}
