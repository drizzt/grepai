package search

import (
	"sort"

	"github.com/yoanbernabeu/grepai/store"
)

// ReciprocalRankFusion merges multiple result lists using RRF.
// k is the RRF constant (typically 60).
// Results are deduplicated by chunk ID and sorted by combined RRF score.
func ReciprocalRankFusion(k float32, limit int, lists ...[]store.SearchResult) []store.SearchResult {
	scores := make(map[string]float32)       // chunkID -> RRF score
	chunkMap := make(map[string]store.Chunk) // chunkID -> chunk

	for _, list := range lists {
		for rank, result := range list {
			id := result.Chunk.ID
			scores[id] += 1.0 / (k + float32(rank) + 1)
			chunkMap[id] = result.Chunk
		}
	}

	results := make([]store.SearchResult, 0, len(scores))
	for id, score := range scores {
		results = append(results, store.SearchResult{
			Chunk: chunkMap[id],
			Score: score,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if limit > 0 && len(results) > limit {
		results = results[:limit]
	}

	return results
}
