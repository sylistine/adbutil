package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const adb = "adb"
const devicesArg = "devices"
const helpSwitch = "--help"

func main() {
	fmt.Println("Android Debug Bridge Utility v0")

	// Check that adb exists.
	_, err := adbCmd(helpSwitch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("adb command successfully located.")

	fmt.Println("Searching for devices...")

	var activeDevice Device
	devices := UpdateDevices()

	if len(devices) < 1 {
		fmt.Println("No devices found. Waiting for devices.")
	} else {
		activeDevice = devices[0]
		fmt.Printf("Device found. Setting active device to %s.\n", activeDevice.ID)
	}

	for {
		devices := UpdateDevices()
		if activeDevice == (Device{}) {
			// Search for devices.
			if len(devices) > 0 {
				fmt.Printf("Device found. Setting active device to %s.\n", devices[0])
				activeDevice = devices[0]
			}
			continue
		} else {
			// Try to get a new default device if the one we've been using has disconnected.
			// Otherwise, loop again.
			activeDevice, err = FindDeviceByID(activeDevice.ID) // Updates device status.
			if err != nil || !DeviceConnected(activeDevice) {
				if len(devices) > 0 {
					fmt.Printf("Active device %s disconnected. Setting active device to %s.\n", activeDevice.ID, devices[0].ID)
					activeDevice = devices[0]
				} else {
					fmt.Printf("Active device %s disconnected. No other devices available.\n", activeDevice.ID)
					fmt.Println("Waiting for device.")
					activeDevice = Device{}
					time.Sleep(1 * time.Second)
					continue
				}
			}
		}

		if activeDevice.ID != "" {
			fmt.Printf("Cmd (%s): ", activeDevice.ID)
			cmd := getUserCmd()

			// Double check that the device still exists.
			UpdateDevices()
			if !DeviceConnected(activeDevice) {
				fmt.Printf("Device %s is no longer available.\n", activeDevice)
				continue
			}

			switch cmd[0] {
			case "exit":
				os.Exit(0)
			case "dev":
				PrintDevices()
			case "device":
				PrintDevices()
			case "devices":
				PrintDevices()
			case "cd":
				// Change device.
				requestedDevice, err := FindDeviceByID(cmd[1])
				if err != nil {
					fmt.Println(err)
					break
				}
				if requestedDevice == activeDevice {
					fmt.Println("Device already selected.")
					break
				}
				fmt.Printf("Setting active device to %s", cmd[1])
				activeDevice = requestedDevice
			case "ls":
				switch cmd[1] {
				case "apk":
					// List apks that CAN BE installed.
					fmt.Println("ls apk not implemented.")
				case "pkg":
					// List packages that ARE installed.
					fmt.Println("ls pkg not implemented.")
				case "activity":
					// List activities within packages that ARE installed.
					fmt.Println("ls activity not implemented.")
				default:
					fmt.Println("ls option unrecognized.")
					// Print list of apk, pkg, or activity.
				}
			case "install":
				break
			default:
				fmt.Println("Command not found.")
			}
		}
	}
}

func getUserCmd() []string {
	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmds := strings.Fields(cmd)

	return cmds
}

func stringInStrings(str string, strs []string) bool {
	found := false
	for _, val := range strs {
		if val == str {
			found = true
			break
		}
	}
	return found
}

func adbCmd(args ...string) (string, error) {
	cmd := exec.Command(adb, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	oBytes, _ := ioutil.ReadAll(stdout)
	oString := string(oBytes)

	return oString, cmd.Wait()
}
