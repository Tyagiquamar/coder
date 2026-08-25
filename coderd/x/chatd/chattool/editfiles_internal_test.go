package chattool

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestEditFileEdit_RejectsDeprecatedKeys pins that the chat ingress
// stays closed to the pre-rename search/replace keys: editFileEdit
// has no UnmarshalJSON and must never gain one via
// workspacesdk.FileEdit (see its doc comment).
func TestEditFileEdit_RejectsDeprecatedKeys(t *testing.T) {
	t.Parallel()

	var e editFileEdit
	require.NoError(t, json.Unmarshal([]byte(`{"search":"x","replace":"y"}`), &e))
	require.Equal(t, "", e.OldText, "deprecated search key must not populate old_text")
	require.Equal(t, "", e.NewText, "deprecated replace key must not populate new_text")
}
