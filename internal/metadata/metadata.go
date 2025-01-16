package metadata

import (
	"database/sql"

	"github.com/quantinium03/conduit/internal/utils"
)

type MetadataDb struct {
    db *sql.DB
    lock utils.Lock[string, *MediaInfo]
}

func StartService() (*Metadata, error) {
    return "", nil
}
