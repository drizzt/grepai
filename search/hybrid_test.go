package search

import (
	"testing"

	"github.com/yoanbernabeu/grepai/store"
)

func TestReciprocalRankFusion(t *testing.T) {
	list1 := []store.SearchResult{
		{Chunk: store.Chunk{ID: "a"}, Score: 0.9},
		{Chunk: store.Chunk{ID: "b"}, Score: 0.8},
		{Chunk: store.Chunk{ID: "c"}, Score: 0.7},
	}

	list2 := []store.SearchResult{
		{Chunk: store.Chunk{ID: "b"}, Score: 0.95},
		{Chunk: store.Chunk{ID: "d"}, Score: 0.85},
		{Chunk: store.Chunk{ID: "a"}, Score: 0.75},
	}

	results := ReciprocalRankFusion(60, 10, list1, list2)

	// "a" is rank 0 in list1 (1/61) and rank 2 in list2 (1/63) = ~0.0322
	// "b" is rank 1 in list1 (1/62) and rank 0 in list2 (1/61) = ~0.0324
	// So "b" should be first, then "a"

	if len(results) != 4 {
		t.Errorf("expected 4 results, got %d", len(results))
		return
	}

	// "b" should have highest combined score
	if results[0].Chunk.ID != "b" {
		t.Errorf("expected first result to be 'b', got '%s'", results[0].Chunk.ID)
	}

	// "a" should be second
	if results[1].Chunk.ID != "a" {
		t.Errorf("expected second result to be 'a', got '%s'", results[1].Chunk.ID)
	}
}

func TestReciprocalRankFusion_Limit(t *testing.T) {
	list1 := []store.SearchResult{
		{Chunk: store.Chunk{ID: "a"}, Score: 0.9},
		{Chunk: store.Chunk{ID: "b"}, Score: 0.8},
		{Chunk: store.Chunk{ID: "c"}, Score: 0.7},
	}

	results := ReciprocalRankFusion(60, 2, list1)

	if len(results) != 2 {
		t.Errorf("expected 2 results with limit, got %d", len(results))
	}
}

func TestReciprocalRankFusion_EmptyLists(t *testing.T) {
	results := ReciprocalRankFusion(60, 10)
	if len(results) != 0 {
		t.Errorf("expected 0 results for empty lists, got %d", len(results))
	}
}
