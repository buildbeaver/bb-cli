package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const kindIDSeparator = ":"

// buildBeaverUUIDNamespace is a namespace for our UUIDs to ensure no collisions with UUIDs
// generated by any other means, as per RFC 4122. Namespace was generated randomly.
var buildBeaverUUIDNamespace = uuid.Must(uuid.Parse("BCE4438A-1DD3-4F29-96ED-7A74E23E19CB"))

type ResourceID struct {
	kind ResourceKind
	id   string
}

func ParseResourceID(str string) (ResourceID, error) {
	parts := strings.Split(str, kindIDSeparator)
	if len(parts) != 2 {
		return ResourceID{}, fmt.Errorf("expected: 2 parts in %q, found: %d", str, len(parts))
	}
	kindPart := ResourceKind(parts[0])
	idPart := parts[1]
	id := &ResourceID{kind: kindPart, id: idPart}
	if !id.Valid() {
		return ResourceID{}, fmt.Errorf("error invalid id format: %s", str)
	}
	return *id, nil
}

func NewResourceID(kind ResourceKind) ResourceID {
	return ResourceID{kind: kind, id: uuid.New().String()}
}

// NewResourceIDFromUniqueData creates a new resource ID by hashing the supplied data (including kind)
// to form a UUID. The UUID will always different for different data, but always the same
// for the same combination of kind and data.
// A UUID namespace is used to ensure these UUIDs never clash with any other UUID.
func NewResourceIDFromUniqueData(kind ResourceKind, data string) ResourceID {
	toHash := []byte(kind.String() + data)
	uuidFromData := uuid.NewSHA1(buildBeaverUUIDNamespace, toHash)
	return ResourceID{kind: kind, id: uuidFromData.String()}
}

func (s ResourceID) Kind() ResourceKind {
	return s.kind
}

func (s ResourceID) Equal(that ResourceID) bool {
	return s.kind == that.kind && s.id == that.id
}

func (s ResourceID) MarshalJSON() ([]byte, error) {
	if !s.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(s.String())
}

func (s *ResourceID) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	if str == "" {
		return nil
	}
	id, err := ParseResourceID(str)
	if err != nil {
		return err
	}
	*s = id
	return nil
}

func (s *ResourceID) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("cannot scan from %T", src)
	}
	id, err := ParseResourceID(str)
	if err != nil {
		return err
	}
	*s = id
	return nil
}

func (s ResourceID) Value() (driver.Value, error) {
	if !s.Valid() {
		return nil, nil
	}
	return s.String(), nil
}

func (s ResourceID) IsZero() bool {
	return s.kind == "" && s.id == ""
}

func (s ResourceID) Valid() bool {
	_, err := uuid.Parse(s.id)
	if err != nil {
		return false
	}
	return s.kind != ""
}

func (s ResourceID) String() string {
	if !s.Valid() {
		return "INVALID: nil resource id"
	}
	return fmt.Sprintf("%s%s%s", s.kind, kindIDSeparator, s.id)
}

// ShortString returns a string version of the ResourceID with only 16 hex digits rather than 32;
// this is still unique enough to use as a unique identifier within a build or a runner.
// See https://en.wikipedia.org/wiki/Birthday_problem in the 'Probability table' section; 6000 items each with
// 16 hex digits (e.g. concurrent docker container names, or job names) will have a clash one in a trillion
// (1,000,000,000,000) times.
func (s ResourceID) ShortString() string {
	if !s.Valid() {
		return "INVALID: nil resource id"
	}
	shortID := s.id[:18] // 16 hex digits and 2 '-' separators; runes in a UUID are 1 byte each
	return fmt.Sprintf("%s%s%s", s.kind, kindIDSeparator, shortID)
}

// SanitizeFilePathID returns a version of the supplied resource ID that is safe for use as a file name.
func SanitizeFilePathID(id ResourceID) string {
	return strings.Replace(id.String(), kindIDSeparator, "_", 1)
}

// SanitizeFilePathShortID returns a shortened version of the supplied resource ID that is safe for use
// as a file name. See ResourceID.ShortString for details of why it's normally OK to use this short version.
func SanitizeFilePathShortID(id ResourceID) string {
	return strings.Replace(id.ShortString(), kindIDSeparator, "_", 1)
}
