package osutil

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
)

// GetHostName
// @Description:获取主机名
// @receiver h
// @return string
func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return hostname
}

// IsWindows check if current os is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux check if current os is linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac check if current os is macos
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOsEnv gets the value of the environment variable named by the key.
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv sets the value of the environment variable named by the key.
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// RemoveOsEnv remove a single environment variable.
func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

// CompareOsEnv gets env named by the key and compare it with comparedEnv
func CompareOsEnv(key, comparedEnv string) bool {
	env := GetOsEnv(key)
	if env == "" {
		return false
	}
	return env == comparedEnv
}

// ExecCommand use shell /bin/bash -c to execute command
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errout bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("cmd")
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err = cmd.Run()

	if err != nil {
		stderr = errout.String()
	}
	stdout = out.String()

	return
}

// GetOsBits get this system bits 32bit or 64bit
// return bit int (32/64)
func GetOsBits() int {
	return 32 << (^uint(0) >> 63)
}
