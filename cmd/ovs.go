package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync/atomic"
	"time"

	"github.com/micro/go-micro/util/log"
)

const (
	// On Windows we need an increased timeout on OVS commands, because
	// adding internal ports on a non Hyper-V enabled host will call
	ovsCommandTimeout = 15
	ovsVsctlCommand   = "/bin/ovs-vsctl"
	ovsOfctlCommand   = "/bin/ovs-ofctl"
	ovnNbctlCommand   = "/bin/ovn-nbctl"
	ovnSbctlCommand   = "/bin/ovn-sbctl"
)

var runCounter uint64

func run(cmdPath string, args ...string) (*bytes.Buffer, *bytes.Buffer, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := exec.Command(cmdPath, args...)

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	counter := atomic.AddUint64(&runCounter, 1)
	logCmd := fmt.Sprintf("%s %s", cmdPath, strings.Join(args, " "))
	log.Log("exec(%d): %s", counter, logCmd)

	err := cmd.Run()
	log.Log("exec(%d): stdout: %q", counter, logCmd)
	log.Log("exec(%d): stderr: %q", counter, logCmd)

	if err != nil {
		log.Log("exec(%d): err: %v", counter, err)
	}
	return stdout, stderr, err
}

// RunOVSOfctl runs a command via ovs-ofctl.
func RunOVSOfctl(args ...string) (string, string, error) {
	stdout, stderr, err := run(ovsOfctlCommand, args...)
	return strings.Trim(stdout.String(), "\" \n"), stderr.String(), err
}

// RunOVSVsctl runs a command via ovs-vsctl.
func RunOVSVsctl(args ...string) (string, string, error) {
	cmdArgs := []string{fmt.Sprintf("--timeout=%d", ovsCommandTimeout)}
	cmdArgs = append(cmdArgs, args...)
	stdout, stderr, err := run(ovsVsctlCommand, cmdArgs...)
	return strings.Trim(strings.TrimSpace(stdout.String()), "\""), stderr.String(), err
}

// Run the ovn-ctl command and retry if "Connection refused"
// poll waitng for service to become available
func runOVNretry(cmdPath string, args ...string) (*bytes.Buffer, *bytes.Buffer, error) {

	retriesLeft := 200
	for {
		stdout, stderr, err := run(cmdPath, args...)
		if err == nil {
			return stdout, stderr, err
		}

		// Connection refused
		// Master may not be up so keep trying
		if strings.Contains(stderr.String(), "Connection refused") {
			if retriesLeft == 0 {
				return stdout, stderr, err
			}
			retriesLeft--
			time.Sleep(2 * time.Second)
		} else {
			// Some other problem for caller to handle
			return stdout, stderr, fmt.Errorf("OVN command '%s %s' failed: %s", cmdPath, strings.Join(args, " "), err)
		}
	}
}

// RunOVNNbctlWithTimeout runs command via ovn-nbctl with a specific timeout
func RunOVNNbctlWithTimeout(timeout int, nburl string, args ...string) (string, string,
	error) {
	var cmdArgs []string

	cmdArgs = []string{
		fmt.Sprintf("--db=%s", nburl),
	}

	cmdArgs = append(cmdArgs, fmt.Sprintf("--timeout=%d", timeout))
	cmdArgs = append(cmdArgs, args...)
	stdout, stderr, err := runOVNretry(ovnNbctlCommand, cmdArgs...)
	return strings.Trim(strings.TrimSpace(stdout.String()), "\""), stderr.String(), err
}

// RunOVNNbctl runs a command via ovn-nbctl.
func RunOVNNbctl(nburl string, args ...string) (string, string, error) {
	return RunOVNNbctlWithTimeout(ovsCommandTimeout, nburl, args...)
}

// RunOVNSbctlWithTimeout runs command via ovn-sbctl with a specific timeout
func RunOVNSbctlWithTimeout(timeout int, sburl string, args ...string) (string, string,
	error) {
	var cmdArgs []string

	cmdArgs = []string{
		fmt.Sprintf("--db=%s", sburl),
	}

	cmdArgs = append(cmdArgs, fmt.Sprintf("--timeout=%d", timeout))
	cmdArgs = append(cmdArgs, args...)
	stdout, stderr, err := runOVNretry(ovnSbctlCommand, cmdArgs...)
	return strings.Trim(strings.TrimSpace(stdout.String()), "\""), stderr.String(), err
}

// RunOVNSbctl runs a command via ovn-sbctl.
func RunOVNSbctl(sburl string, args ...string) (string, string, error) {
	return RunOVNSbctlWithTimeout(ovsCommandTimeout, sburl, args...)
}
