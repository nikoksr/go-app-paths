# go-app-paths

[![Latest Release](https://img.shields.io/github/release/nikoksr/go-app-paths.svg)](https://github.com/nikoksr/go-app-paths/releases)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/nikoksr/go-app-paths?tab=doc)
[![Go ReportCard](https://goreportcard.com/badge/nikoksr/go-app-paths)](https://goreportcard.com/report/nikoksr/go-app-paths)

Lets you retrieve platform-specific paths (like directories for app-data, cache,
config, and logs). It is fully compliant with the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html)
on Unix, but also provides implementations for macOS and Windows systems.

## Fork of [muesli/go-app-paths](https://github.com/muesli/go-app-paths)

This is a fork of [muesli/go-app-paths](https://github.com/muesli/go-app-paths) which addresses some of the open issues in the
original repository.

Since I'm using this fork for my personal purposes and am not planning on merging it
back to the original repository, I changed the import path to `github.com/nikoksr/go-app-paths`
to avoid confusion.

All credits for the original implementation go to [muesli](https://github.com/muesli).

## Installation

Make sure you have a working Go environment (Go 1.2 or higher is required).
See the [install instructions](https://golang.org/doc/install.html).

To install go-app-paths, simply run:

    go get github.com/nikoksr/go-app-paths

## Usage

```go
import gap "github.com/nikoksr/go-app-paths"
```

### Scopes

You can initialize `gap` with either the `gap.User` or `gap.System` scope to
retrieve user- and/or system-specific base directories and paths:

```go
scope := gap.NewScope(gap.User, "app")
```

Alternatively, you can initialize `gap` with an additional vendor name:

```go
scope := gap.NewVendorScope(gap.User, "vendor", "app")
```

This will cause the `app` directory to be prefixed by a `vendor` directory in
all the following tables.

### Directories

`DataDirs` retrieves a priority-sorted list of data directories:

```go
dirs, err := scope.DataDirs()
```

| Platform | User Scope                                                       | System Scope                               |
| -------- | ---------------------------------------------------------------- | ------------------------------------------ |
| Unix     | ["~/.local/share/app", "/usr/local/share/app", "/usr/share/app"] | ["/usr/local/share/app", "/usr/share/app"] |
| macOS    | ["~/Library/Application Support/app"]                            | ["/Library/Application Support/app"]       |
| Windows  | ["%LOCALAPPDATA%/app"]                                           | ["%PROGRAMDATA%/app"]                      |

---

`ConfigDirs` retrieves a priority-sorted list of config directories:

```go
dirs, err := scope.ConfigDirs()
```

| Platform | User Scope                                    | System Scope                 |
| -------- | --------------------------------------------- | ---------------------------- |
| Unix     | ["~/.config/app", "/etc/xdg/app", "/etc/app"] | ["/etc/xdg/app", "/etc/app"] |
| macOS    | ["~/Library/Application Support/app"]         | ["/Library/Application Support/app"] |
| Windows  | ["%LOCALAPPDATA%/app/Config"]                 | ["%PROGRAMDATA%/app/Config"] |

---

`DataDir` retrieves the app's primary data directory:

```go
dir, err := scope.DataDir()
```

| Platform | User Scope                       | System Scope                  |
| -------- | -------------------------------- | ----------------------------- |
| Unix     | ~/.local/share/app              | /usr/local/share/app         |
| macOS    | ~/Library/Application Support/app | /Library/Application Support/app |
| Windows  | %LOCALAPPDATA%/app             | %PROGRAMDATA%/app            |

---

`LogDir` retrieves the app's log directory:

```go
dir, err := scope.LogDir()
```

| Platform | User Scope                       | System Scope                  |
| -------- | -------------------------------- | ----------------------------- |
| Unix     | ~/.local/share/app              | /var/log/app                   |
| macOS    | ~/Library/Logs/app             | /Library/Logs/app              |
| Windows  | %LOCALAPPDATA%/app/Logs        | %PROGRAMDATA%/app/Logs         |

---

`CacheDir` retrieves the app's cache directory:

```go
dir, err := scope.CacheDir()
```

| Platform | User Scope               | System Scope            |
| -------- | ------------------------ | ----------------------- |
| Unix     | ~/.cache/app             | /var/cache/app          |
| macOS    | ~/Library/Caches/app     | /Library/Caches/app     |
| Windows  | %LOCALAPPDATA%/app/Cache | %PROGRAMDATA%/app/Cache |

### Default File Paths

`DataPath` retrieves the default path for a data file:

```go
path, err := scope.DataPath("filename")
```

| Platform | User Scope                                 | System Scope                              |
| -------- | ------------------------------------------ | ----------------------------------------- |
| Unix     | ~/.local/share/app/filename                | /usr/local/share/app/filename             |
| macOS    | ~/Library/Application Support/app/filename | /Library/Application Support/app/filename |
| Windows  | %LOCALAPPDATA%/app/filename                | %PROGRAMDATA%/app/filename                |

---

`ConfigPath` retrieves the default path for a config file:

```go
path, err := scope.ConfigPath("filename.conf")
```

| Platform | User Scope                              | System Scope                           |
| -------- | --------------------------------------- | -------------------------------------- |
| Unix     | ~/.config/app/filename.conf             | /etc/xdg/app/filename.conf             |
| macOS    | ~/Library/Application Support/app/filename.conf | /Library/Application Support/app/filename.conf |
| Windows  | %LOCALAPPDATA%/app/Config/filename.conf | %PROGRAMDATA%/app/Config/filename.conf |

---

`LogPath` retrieves the default path for a log file:

```go
path, err := scope.LogPath("filename.log")
```

| Platform | User Scope                           | System Scope                        |
| -------- | ------------------------------------ | ----------------------------------- |
| Unix     | ~/.local/share/app/filename.log      | /var/log/app/filename.log           |
| macOS    | ~/Library/Logs/app/filename.log      | /Library/Logs/app/filename.log      |
| Windows  | %LOCALAPPDATA%/app/Logs/filename.log | %PROGRAMDATA%/app/Logs/filename.log |

### Lookup Methods

`LookupData` retrieves a priority-sorted list of paths for existing data files
with the name `filename`:

```go
path, err := scope.LookupData("filename")
```

| Platform | User Scope                                                                                  | System Scope                                                 |
| -------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| Unix     | ["~/.local/share/app/filename", "/usr/local/share/app/filename", "/usr/share/app/filename"] | ["/usr/local/share/app/filename", "/usr/share/app/filename"] |
| macOS    | ["~/Library/Application Support/app/filename"]                                              | ["/Library/Application Support/app/filename"]                |
| Windows  | ["%LOCALAPPDATA%/app/filename"]                                                             | ["%PROGRAMDATA%/app/filename"]                               |

---

`LookupConfig` retrieves a priority-sorted list of paths for existing config
files with the name `filename.conf`:

```go
path, err := scope.LookupConfig("filename.conf")
```

| Platform | User Scope                                                                              | System Scope                                             |
| -------- | --------------------------------------------------------------------------------------- | -------------------------------------------------------- |
| Unix     | ["~/.config/app/filename.conf", "/etc/xdg/app/filename.conf", "/etc/app/filename.conf"] | ["/etc/xdg/app/filename.conf", "/etc/app/filename.conf"] |
| macOS    | ["~/Library/Preferences/app/filename.conf"]                                             | ["/Library/Preferences/app/filename.conf"]               |
| Windows  | ["%LOCALAPPDATA%/app/Config/filename.conf"]                                             | ["%PROGRAMDATA%/app/Config/filename.conf"]               |
