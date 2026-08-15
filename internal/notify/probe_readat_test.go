package notify

import (
	"testing"
	"time"
)

func TestMarkRead_ReturnedReadAtMustBeIndependent(t *testing.T) {
	s := New()
	base := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	s.Create(CreateInput{ID: "R1", Recipient: "u", Content: "c"}, base)
	s.MarkSent("R1", base.Add(time.Hour))

	readTime := base.Add(2 * time.Hour)
	returned, err := s.MarkRead("R1", readTime)
	if err != nil {
		t.Fatal(err)
	}
	if returned.ReadAt == nil {
		t.Fatal("ReadAt should not be nil after MarkRead")
	}

	*returned.ReadAt = base.Add(999 * time.Hour)

	got, _ := s.Get("R1")
	if !got.ReadAt.Equal(readTime) {
		t.Errorf("store corrupted: got ReadAt=%v, want %v", got.ReadAt, readTime)
	}
}
