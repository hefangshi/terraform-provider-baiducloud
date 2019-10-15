package baiducloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/terraform/helper/resource"
)

// timeout for common product, bcc e.g.
const DefaultTimeout = 180 * time.Second
const DefaultDebugMsg = "\n*************** %s Response *************** \n%+v\n%s******************************\n\n"

const (
	PAYMENT_TIMING_POSTPAID = "Postpaid"
	PAYMENT_TIMING_PREPAID  = "Prepaid"
)

func debugOn() bool {
	for _, part := range strings.Split(os.Getenv("DEBUG"), ",") {
		if strings.TrimSpace(part) == "terraform" {
			return true
		}
	}
	return false
}

func addDebug(action, content interface{}) {
	if debugOn() {
		trace := "[DEBUG TRACE]:\n"
		for skip := 1; skip < 5; skip++ {
			_, filepath, line, _ := runtime.Caller(skip)
			trace += fmt.Sprintf("%s:%d\n", filepath, line)
		}

		fmt.Printf(DefaultDebugMsg, action, content, trace)
		log.Printf(DefaultDebugMsg, action, content, trace)
	}
}

// write data to file
func writeToFile(filePath string, data interface{}) error {
	if strings.HasPrefix(filePath, "~") {
		usr, errCurrent := user.Current()
		if errCurrent != nil {
			return fmt.Errorf("get current user error: %s", errCurrent.Error())
		}
		if usr.HomeDir != "" {
			filePath = strings.Replace(filePath, "~", usr.HomeDir, 1)
		}
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("stat file error: %s", err.Error())
	}

	if fileInfo != nil {
		if errRemove := os.Remove(filePath); errRemove != nil {
			return fmt.Errorf("delete old file error: %s", errRemove.Error())
		}
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json marshal error: %s", err.Error())
	}

	return ioutil.WriteFile(filePath, []byte(bytes), 0644)
}

func buildClientToken() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

func buildStateConf(pending, target []string, timeout time.Duration, f resource.StateRefreshFunc) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Delay:      10 * time.Second,
		Pending:    pending,
		Refresh:    f,
		Target:     target,
		Timeout:    timeout,
		MinTimeout: 3 * time.Second,
	}
}

func stringInSlice(strs []string, value string) bool {
	for _, str := range strs {
		if value == str {
			return true
		}
	}

	return false
}

// check two strings are equal or not
// if both strings are one of defaultStr value, return true
func stringEqualWithDefault(s1, s2 string, defaultStr []string) bool {
	isDefaultS1 := stringInSlice(defaultStr, s1)
	isDefaultS2 := stringInSlice(defaultStr, s2)

	if isDefaultS1 != isDefaultS2 {
		return false
	}

	if s1 != s2 {
		if !isDefaultS1 {
			return false
		}
	}

	return true
}