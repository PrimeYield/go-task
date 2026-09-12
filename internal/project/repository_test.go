package project

import (
	"sync"
	"testing"
)

func TestMemoryRepository_Create(t *testing.T) {
	repo := NewMemoryRepository()

	created := repo.Create("GoTask Backend")

	if created.ID != 1 {
		t.Fatalf("expected ID 1, got %d", created.ID)
	}

	if created.Name != "GoTask Backend" {
		t.Fatalf("expected name %q, got %q", "GoTask Backend", created.Name)
	}
}

func TestMemoryRepository_CreateIncrementsID(t *testing.T) {
	repo := NewMemoryRepository()

	first := repo.Create("Project A")
	second := repo.Create("Project B")
	third := repo.Create("Project C")

	if first.ID != 1 {
		t.Fatalf("expected first ID 1, got %d", first.ID)
	}

	if second.ID != 2 {
		t.Fatalf("expected second ID 2, got %d", second.ID)
	}

	if third.ID != 3 {
		t.Fatalf("expected third ID 3, got %d", third.ID)
	}
}

func TestMemoryRepository_List(t *testing.T) {
	repo := NewMemoryRepository()

	repo.Create("Project A")
	repo.Create("Project B")
	repo.Create("Project C")

	projects := repo.List()

	if len(projects) != 3 {
		t.Fatalf("expected 3 projects, got %d", len(projects))
	}
}

func TestMemoryRepository_ConcurrentCreate(t *testing.T) {
	repo := NewMemoryRepository()

	const workers = 100

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			repo.Create("Concurrent Project")
		}()
	}

	wg.Wait()

	projects := repo.List()

	if len(projects) != workers {
		t.Fatalf(
			"expected %d projects, got %d",
			workers,
			len(projects),
		)
	}
}
