package sysexits

//go:generate stringer -type=ExitCode

// ExitCode is a valid exit code.
type ExitCode int

// Source: http://www.opensource.apple.com/source/Libc/Libc-320/include/sysexits.h
const (
	// ExitOk indicates a successful termination.
	ExitOk ExitCode = 0

	// ExitUsage indicates a command line usage error.
	ExitUsage ExitCode = 64

	// ExitDataErr indicates a data format error.
	ExitDataErr ExitCode = 65

	// ExitNoInput indicates missing input.
	ExitNoInput ExitCode = 66

	// ExitNoUser indicates the addressee is unknown.
	ExitNoUser ExitCode = 67

	// ExitNoHost indicates the host name is unknown.
	ExitNoHost ExitCode = 68

	// ExitUnavailable indicates the service is unavailable.
	ExitUnavailable ExitCode = 69

	// ExitSoftware indicates there was an internal software error.
	ExitSoftware ExitCode = 70

	// ExitOsErr indicates there was a system error (e.g., couldn't fork).
	ExitOsErr ExitCode = 71

	// ExitOsFile indicates a critical OS file is missing.
	ExitOsFile ExitCode = 72

	// ExitCantCreate indicates there was an error creating a user output file.
	ExitCantCreate ExitCode = 73

	// ExitIoErr indicates there was an error related to IO.
	ExitIoErr ExitCode = 74

	// ExitTempFail indicates there was a temporary error and the user is
	// invited to retry.
	ExitTempFail ExitCode = 75

	// ExitProtocol indicates there was a remote error in the protocol.
	ExitProtocol ExitCode = 76

	// ExitNoPerm indicates permission was denied.
	ExitNoPerm ExitCode = 77

	// ExitConfig indicates there was a configuration error.
	ExitConfig ExitCode = 78
)
