//go:build aix

// Package sqlite registers the ncruces/go-sqlite3 driver under the name
// "sqlite" on AIX.
//
// On AIX, ncruces/go-sqlite3's default locking backend falls into
// lock_other.go, which doesn't implement locking.  This is because the OFD
// (open-file-description) locks used on Linux, and the flock() calls used on
// BSD and illumos, are both absent on AIX.  AIX does support fcntl locking,
// but ncruces' unix VFS does not use it; OFD and flock are the only two
// unix-family backends.  Build trivy on AIX with the sqlite3_dotlk tag to
// use ncruces' dot-lock VFS, which is portable and works correctly:
//
//	GOOS=aix GOARCH=ppc64 CGO_ENABLED=0 go build -tags sqlite3_dotlk ./cmd/trivy
//
// Without that tag the binary will compile but every SQLite open will
// return a "disk I/O error" at runtime.  The aix_require_dotlk.go file
// turns that into a compile-time failure.
package sqlite

import (
	"database/sql"

	"github.com/ncruces/go-sqlite3/driver"
)

func init() {
	// ncruces/go-sqlite3/driver registers under "sqlite3" by default.
	// Re-register the same driver under the name expected by go-rpmdb and
	// trivy-java-db ("sqlite") so we don't need to patch those upstreams.
	sql.Register("sqlite", &driver.SQLite{})
}
