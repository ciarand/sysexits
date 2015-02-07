package sysexits_test

import (
	"testing"

	"github.com/ciarand/sysexits"
)

// Source: http://www.opensource.apple.com/source/Libc/Libc-320/include/sysexits.h
var codesToTest = []struct {
	expected int
	actual   sysexits.ExitCode
}{
	// #define EX_USAGE	64	/* command line usage error */
	{64, sysexits.ExitUsage},

	// #define EX_DATAERR	65	/* data format error */
	{65, sysexits.ExitDataErr},

	// #define EX_NOINPUT	66	/* cannot open input */
	{66, sysexits.ExitNoInput},

	// #define EX_NOUSER	67	/* addressee unknown */
	{67, sysexits.ExitNoUser},

	// #define EX_NOHOST	68	/* host name unknown */
	{68, sysexits.ExitNoHost},

	// #define EX_UNAVAILABLE	69	/* service unavailable */
	{69, sysexits.ExitUnavailable},

	// #define EX_SOFTWARE	70	/* internal software error */
	{70, sysexits.ExitSoftware},

	// #define EX_OSERR	71	/* system error (e.g., can't fork) */
	{71, sysexits.ExitOsErr},

	// #define EX_OSFILE	72	/* critical OS file missing */
	{72, sysexits.ExitOsFile},

	// #define EX_CANTCREAT	73	/* can't create (user) output file */
	{73, sysexits.ExitCantCreate},

	// #define EX_IOERR	74	/* input/output error */
	{74, sysexits.ExitIoErr},

	// #define EX_TEMPFAIL	75	/* temp failure; user is invited to retry */
	{75, sysexits.ExitTempFail},

	// #define EX_PROTOCOL	76	/* remote error in protocol */
	{76, sysexits.ExitProtocol},

	// #define EX_NOPERM	77	/* permission denied */
	{77, sysexits.ExitNoPerm},

	// #define EX_CONFIG	78	/* configuration error */
	{78, sysexits.ExitConfig},
}

func TestCodes(t *testing.T) {
	for _, c := range codesToTest {
		if c.expected != int(c.actual) {
			t.Errorf("%s was not equal to %d", c.actual.String(), c.expected)
		}
	}
}
