package versionproxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/blang/semver"
)

func New() http.Handler {
	return nil
}

func modifyVersion(req *http.Request, w http.ResponseWriter) http.ResponseWriter {
	if isVersionRequest(req) {
		return &responseWriterWrapper{internal: w}
	}
	return w
}

func isVersionRequest(req *http.Request) bool {
	return strings.HasSuffix(req.URL.Path, "/version")
}

type responseWriterWrapper struct {
	internal http.ResponseWriter
}

func (w *responseWriterWrapper) Header() http.Header {
	return w.internal.Header()
}

func (w *responseWriterWrapper) WriteHeader(n int) {
	w.internal.WriteHeader(n)
}

func (w *responseWriterWrapper) Write(data []byte) (int, error) {
	var versionData map[string]interface{}
	err := json.Unmarshal(data, &versionData)
	if err != nil {
		fmt.Printf("error decoding version data: %v\n", err)
		return w.internal.Write(data)
	}
	version := versionData["Version"]
	if version == nil {
		fmt.Printf("version not found in version data: %v\n", versionData)
		return w.internal.Write(data)
	}
	versionStr := fmt.Sprintf("%v", version)
	if _, err := semver.Parse(versionStr); err == nil {
		fmt.Printf("version %s parses ok, passing through\n", versionStr)
		return w.internal.Write(data)
	}
	parts := strings.SplitN(versionStr, ".", 3)
	// Remove leading 0's
	parts[0] = strings.TrimLeft(parts[0], "0")
	parts[1] = strings.TrimLeft(parts[1], "0")
	if len(parts[0]) == 0 {
		parts[0] = "1"
	}
	if len(parts[1]) == 0 {
		parts[1] = "1"
	}

	versionData["Version"] = strings.Join(parts, ".")
	newData, err := json.Marshal(versionData)
	if err != nil {
		fmt.Printf("error encoding version data: %v\n", err)
		return w.internal.Write(data)
	}
	fmt.Printf("Modified version %s to %s\n", versionData["Version"])
	return w.internal.Write(newData)
}
