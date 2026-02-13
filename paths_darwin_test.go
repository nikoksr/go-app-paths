//go:build darwin
// +build darwin

package gap

import (
	"testing"
)

func TestPaths(t *testing.T) {
	tests := []struct {
		scope      *Scope
		dataDir    string
		cacheDir   string
		configFile string
		dataFile   string
		logFile    string
		dataDir2   string
		logDir     string
	}{
		{
			scope:      NewScope(System, "foobar"),
			dataDir:    "/Library/Application Support/foobar",
			cacheDir:   "/Library/Caches/foobar",
			configFile: "/Library/Application Support/foobar/foobar.conf",
			dataFile:   "/Library/Application Support/foobar/foobar.data",
			logFile:    "/Library/Logs/foobar/foobar.log",
			dataDir2:   "/Library/Application Support/foobar",
			logDir:     "/Library/Logs/foobar",
		},
		{
			scope:      NewVendorScope(System, "barcorp", "foobar"),
			dataDir:    "/Library/Application Support/barcorp/foobar",
			cacheDir:   "/Library/Caches/barcorp/foobar",
			configFile: "/Library/Application Support/barcorp/foobar/foobar.conf",
			dataFile:   "/Library/Application Support/barcorp/foobar/foobar.data",
			logFile:    "/Library/Logs/barcorp/foobar/foobar.log",
			dataDir2:   "/Library/Application Support/barcorp/foobar",
			logDir:     "/Library/Logs/barcorp/foobar",
		},
		{
			scope:      NewScope(User, "foobar"),
			dataDir:    "~/Library/Application Support/foobar",
			cacheDir:   "~/Library/Caches/foobar",
			configFile: "~/Library/Application Support/foobar/foobar.conf",
			dataFile:   "~/Library/Application Support/foobar/foobar.data",
			logFile:    "~/Library/Logs/foobar/foobar.log",
			dataDir2:   "~/Library/Application Support/foobar",
			logDir:     "~/Library/Logs/foobar",
		},
		{
			scope:      NewCustomHomeScope("/tmp", "", "foobar"),
			dataDir:    "/tmp/Library/Application Support/foobar",
			cacheDir:   "/tmp/Library/Caches/foobar",
			configFile: "/tmp/Library/Application Support/foobar/foobar.conf",
			dataFile:   "/tmp/Library/Application Support/foobar/foobar.data",
			logFile:    "/tmp/Library/Logs/foobar/foobar.log",
			dataDir2:   "/tmp/Library/Application Support/foobar",
			logDir:     "/tmp/Library/Logs/foobar",
		},
	}

	for _, tt := range tests {
		paths, err := tt.scope.DataDirs()
		if err != nil {
			t.Errorf("Error retrieving data dir: %s", err)
		}
		if paths[0] != expandUser(tt.dataDir) {
			t.Errorf("Expected data dir: %s - got: %s", tt.dataDir, paths[0])
		}

		path, err := tt.scope.CacheDir()
		if err != nil {
			t.Errorf("Error retrieving cache dir: %s", err)
		}
		if path != expandUser(tt.cacheDir) {
			t.Errorf("Expected cache dir: %s - got: %s", tt.cacheDir, path)
		}

		path, err = tt.scope.ConfigPath(tt.scope.App + ".conf")
		if err != nil {
			t.Errorf("Error retrieving config path: %s", err)
		}
		if path != expandUser(tt.configFile) {
			t.Errorf("Expected config path: %s - got: %s", tt.configFile, path)
		}

		path, err = tt.scope.DataPath(tt.scope.App + ".data")
		if err != nil {
			t.Errorf("Error retrieving data path: %s", err)
		}
		if path != expandUser(tt.dataFile) {
			t.Errorf("Expected data path: %s - got: %s", tt.dataFile, path)
		}

		path, err = tt.scope.LogPath(tt.scope.App + ".log")
		if err != nil {
			t.Errorf("Error retrieving log path: %s", err)
		}
		if path != expandUser(tt.logFile) {
			t.Errorf("Expected log path: %s - got: %s", tt.logFile, path)
		}

		path, err = tt.scope.DataDir()
		if err != nil {
			t.Errorf("Error retrieving data dir: %s", err)
		}
		if path != expandUser(tt.dataDir2) {
			t.Errorf("Expected data dir: %s - got: %s", tt.dataDir2, path)
		}

		path, err = tt.scope.LogDir()
		if err != nil {
			t.Errorf("Error retrieving log dir: %s", err)
		}
		if path != expandUser(tt.logDir) {
			t.Errorf("Expected log dir: %s - got: %s", tt.logDir, path)
		}
	}
}
