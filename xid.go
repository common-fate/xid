// Package xid generates prefixed KSUIDs.
package xid

import ksuid "github.com/segmentio/ksuid"

// New creates a new KSUID with the specified prefix.
func New(prefix string) string {
	return prefix + "_" + ksuid.New().String()
}
