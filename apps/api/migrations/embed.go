package migrations

import "embed"

// Files embeds SQL migrations so schema bootstrap does not depend on the working directory.
//
//go:embed *.sql
var Files embed.FS
