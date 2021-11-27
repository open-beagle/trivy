// +build integration

package integration

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/afero"

	"github.com/aquasecurity/trivy-db/pkg/db"
)

var update = flag.Bool("update", false, "update golden files")

func gunzipDB() (string, error) {
	gz, err := os.Open("testdata/trivy.db.gz")
	if err != nil {
		return "", err
	}
	zr, err := gzip.NewReader(gz)
	if err != nil {
		return "", err
	}

	tmpDir, err := ioutil.TempDir("", "integration")
	if err != nil {
		return "", err
	}

	dbPath := db.Path(tmpDir)
	dbDir := filepath.Dir(dbPath)
	err = os.MkdirAll(dbDir, 0700)
	if err != nil {
		return "", err
	}

	file, err := os.Create(dbPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err = io.Copy(file, zr); err != nil {
		return "", err
	}

	fs := afero.NewOsFs()
	metadataFile := filepath.Join(dbDir, "metadata.json")
	b, err := json.Marshal(db.Metadata{
		Version:    1,
		Type:       1,
		NextUpdate: time.Time{},
		UpdatedAt:  time.Time{},
	})
	if err != nil {
		return "", err
	}
	err = afero.WriteFile(fs, metadataFile, b, 0600)
	if err != nil {
		return "", err
	}

	return tmpDir, nil
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func waitPort(ctx context.Context, addr string) error {
	for {
		conn, err := net.Dial("tcp", addr)
		if err == nil && conn != nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return err
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
