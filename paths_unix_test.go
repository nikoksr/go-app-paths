//go:build !darwin && !windows
// +build !darwin,!windows

package gap

import (
	"fmt"
	"testing"
)

func TestPaths(t *testing.T) {
	tests := []struct {
		scope      *Scope
		dataDirs   []string
		configDirs []string
		cacheDir   string
		configFile string
		dataFile   string
		logFile    string
		dataDir    string
		logDir     string
	}{
		{
			scope:      NewScope(System, "foobar"),
			dataDirs:   []string{"/usr/local/share/foobar", "/usr/share/foobar"},
			configDirs: []string{"/etc/xdg/foobar", "/etc/foobar"},
			cacheDir:   "/var/cache/foobar",
			configFile: "/etc/xdg/foobar/foobar.conf",
			dataFile:   "/usr/local/share/foobar/foobar.data",
			logFile:    "/var/log/foobar/foobar.log",
			dataDir:    "/usr/local/share/foobar",
			logDir:     "/var/log/foobar",
		},
		{
			scope:      NewVendorScope(System, "barcorp", "foobar"),
			dataDirs:   []string{"/usr/local/share/barcorp/foobar", "/usr/share/barcorp/foobar"},
			configDirs: []string{"/etc/xdg/barcorp/foobar", "/etc/barcorp/foobar"},
			cacheDir:   "/var/cache/barcorp/foobar",
			configFile: "/etc/xdg/barcorp/foobar/foobar.conf",
			dataFile:   "/usr/local/share/barcorp/foobar/foobar.data",
			logFile:    "/var/log/barcorp/foobar/foobar.log",
			dataDir:    "/usr/local/share/barcorp/foobar",
			logDir:     "/var/log/barcorp/foobar",
		},
		{
			scope:      NewScope(User, "foobar"),
			dataDirs:   []string{"~/.local/share/foobar", "/usr/local/share/foobar", "/usr/share/foobar"},
			configDirs: []string{"~/.config/foobar", "/etc/xdg/foobar", "/etc/foobar"},
			cacheDir:   "~/.cache/foobar",
			configFile: "~/.config/foobar/foobar.conf",
			dataFile:   "~/.local/share/foobar/foobar.data",
			logFile:    "~/.local/share/foobar/foobar.log",
			dataDir:    "~/.local/share/foobar",
			logDir:     "~/.local/share/foobar",
		},
		{
			scope:      NewCustomHomeScope("/tmp", "", "foobar"),
			dataDirs:   []string{"/tmp/.local/share/foobar"},
			configDirs: []string{"/tmp/.config/foobar"},
			cacheDir:   "/tmp/.cache/foobar",
			configFile: "/tmp/.config/foobar/foobar.conf",
			dataFile:   "/tmp/.local/share/foobar/foobar.data",
			logFile:    "/tmp/.local/share/foobar/foobar.log",
			dataDir:    "/tmp/.local/share/foobar",
			logDir:     "/tmp/.local/share/foobar",
		},
	}

	for _, tt := range tests {
		paths, err := tt.scope.DataDirs()
		if err != nil {
			t.Errorf("Error retrieving data dir: %s", err)
		}

		if paths[0] != expandUser(tt.dataDirs[0]) {
			t.Errorf("Expected data dir: %s - got: %s", tt.dataDirs[0], paths[0])
		}

		paths, err = tt.scope.ConfigDirs()
		if err != nil {
			t.Errorf("Error retrieving config dir: %s", err)
		}

		if paths[0] != expandUser(tt.configDirs[0]) {
			t.Errorf("Expected config dir: %s - got: %s", tt.configDirs[0], paths[0])
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
		if path != expandUser(tt.dataDir) {
			t.Errorf("Expected data dir: %s - got: %s", tt.dataDir, path)
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

func TestConfigLookups(t *testing.T) {
	tests := []struct {
		scope      *Scope
		configFile string
		result     []string
	}{
		{NewScope(System, "ssh"), "sshd_config", []string{"/etc/ssh/sshd_config"}},
		{NewScope(User, "ssh"), "sshd_config", []string{"/etc/ssh/sshd_config"}},
	}

	for _, tt := range tests {
		r, err := tt.scope.LookupConfig(tt.configFile)
		if err != nil {
			t.Errorf("Error looking up config: %s", err)
		}
		if len(r) == 0 {
			fmt.Println(r)
			t.Skipf("Expected config file not found: %s (integration test, skipping)", tt.result[0])
			continue
		}
		if r[0] != tt.result[0] {
			t.Errorf("Expected config file: %s - got: %s", tt.result[0], r[0])
		}
	}
}

func TestDataLookups(t *testing.T) {
	tests := []struct {
		scope    *Scope
		dataFile string
		result   []string
	}{
		{NewVendorScope(System, "terminfo", "x"), "xterm+256color", []string{"/usr/share/terminfo/x/xterm+256color"}},
		{NewVendorScope(User, "terminfo", "x"), "xterm+256color", []string{"/usr/share/terminfo/x/xterm+256color"}},
	}

	for _, tt := range tests {
		r, err := tt.scope.LookupDataFile(tt.dataFile)
		if err != nil {
			t.Errorf("Error looking up data file: %s", err)
		}
		if len(r) == 0 {
			fmt.Println(r)
			t.Skipf("Expected data file not found: %s (integration test, skipping)", tt.result[0])
			continue
		}
		if r[0] != tt.result[0] {
			t.Errorf("Expected data file: %s - got: %s", tt.result[0], r[0])
		}
	}
}
