// Package gap provides cross-platform access to application-specific directories.
//
// It retrieves standard locations for storing application data, configuration files,
// cache, and logs across Unix, macOS, and Windows platforms.
//
// On Unix systems, the package is fully compliant with the XDG Base Directory
// Specification (https://specifications.freedesktop.org/basedir-spec/).
//
// # Usage
//
// Create a Scope with the desired scope type and application name:
//
//	scope := gap.NewScope(gap.User, "myapp")
//
// Then query for standard directories:
//
//	dataDir, _ := scope.DataDir()        // ~/.local/share/myapp
//	configDir, _ := scope.ConfigDir()    // ~/.config/myapp
//	cacheDir, _ := scope.CacheDir()      // ~/.cache/myapp
//	logDir, _ := scope.LogDir()          // ~/.local/share/myapp
//
// # Scopes
//
// Three scope types are available:
//   - gap.User: User-specific paths (default for most applications)
//   - gap.System: System-wide paths (for shared data)
//   - gap.CustomHome: Custom home directory path
//
// # Vendor Support
//
// For applications that belong to a vendor/organization, use NewVendorScope:
//
//	scope := gap.NewVendorScope(gap.User, "mycompany", "myapp")
//
// This will prefix all paths with the vendor name.
//
// # Path Lookup
//
// The package also provides lookup methods to find existing files across
// standard directories:
//
//	configs, _ := scope.LookupConfig("myapp.conf")
//	dataFiles, _ := scope.LookupDataFile("data.json")
package gap
