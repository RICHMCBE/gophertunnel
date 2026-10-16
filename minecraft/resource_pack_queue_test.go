package minecraft

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/resource"
)

func TestResourcePackQueueRequestsVersionsWithSameUUID(t *testing.T) {
	packUUID := uuid.MustParse("46e26af2-e87b-4227-a297-79e8d9799fd2")
	packs := []*resource.Pack{
		resourcePackFixture(t, packUUID, "3.10.8"),
		resourcePackFixture(t, packUUID, "3.10.9"),
	}
	queue := &resourcePackQueue{packs: packs, chunkSize: 1}
	requested := []string{
		packUUID.String() + "_3.10.8",
		packUUID.String() + "_3.10.9",
	}
	if err := queue.Request(requested); err != nil {
		t.Fatalf("Request(%v): %v", requested, err)
	}
	if got, want := len(queue.packsToDownload), len(requested); got != want {
		t.Fatalf("queued pack count = %d, want %d", got, want)
	}

	versions := make(map[string]bool, len(requested))
	for range len(requested) {
		info, ok := queue.NextPack()
		if !ok {
			t.Fatal("NextPack() returned no pack before the queue was empty")
		}
		versions[info.UUID] = true
	}
	for _, id := range requested {
		if !versions[id] {
			t.Errorf("NextPack() did not return requested pack %q", id)
		}
	}
	if _, ok := queue.NextPack(); ok {
		t.Fatal("NextPack() returned a pack after all requested versions were consumed")
	}
}

func resourcePackFixture(t *testing.T, packUUID uuid.UUID, version string) *resource.Pack {
	t.Helper()
	path := filepath.Join(t.TempDir(), "pack.mcpack")
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	archive := zip.NewWriter(file)
	manifest, err := archive.Create("manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	content := `{"format_version":2,"header":{"name":"fixture","description":"test","uuid":"` + packUUID.String() + `","version":[` + versionComponents(version) + `],"min_engine_version":[1,20,0]},"modules":[{"type":"resources","uuid":"` + uuid.NewString() + `","version":[` + versionComponents(version) + `]}]}`
	if _, err := manifest.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
	pack, err := resource.ReadPath(path)
	if err != nil {
		t.Fatalf("ReadPath(%q): %v", path, err)
	}
	return pack
}

func versionComponents(version string) string {
	if version == "3.10.8" {
		return "3,10,8"
	}
	if version == "3.10.9" {
		return "3,10,9"
	}
	panic("unexpected fixture version: " + version)
}
