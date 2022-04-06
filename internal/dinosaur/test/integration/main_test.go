package integration

import (
	"flag"
	"os"
	"runtime"
	"testing"

	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	flag.Parse()
	glog.V(10).Infof("Starting integration test using go version %s", runtime.Version())
	os.Exit(m.Run())
}

// skip integration tests that won't pass until the missing logic is implemented
// https://bf2.zulipchat.com/#narrow/stream/315461-factorized-fleet-manager/topic/Integration.20tests
func skipTest(t *testing.T) {
	t.Skip("Implementation missing")
}
