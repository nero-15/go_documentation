package main

func main() {
	// 	Package user allows user account lookups by name or id.
	//
	// For most Unix systems, this package has two internal implementations of resolving user and group ids to names. One is written in pure Go and parses /etc/passwd and /etc/group. The other is cgo-based and relies on the standard C library (libc) routines such as getpwuid_r and getgrnam_r.
	//
	// When cgo is available, cgo-based (libc-backed) code is used by default. This can be overridden by using osusergo build tag, which enforces the pure Go implementation.
}
