package main

import (
	"fmt"
	"log"
	"strings"
)

// UpdateDevices calls `adb devices` and parses the output into a list of type Device.
func UpdateDevices() (devices []Device) {
	output, err := adbCmd(devicesArg)
	if err != nil {
		log.Fatal(err)
	}

	outputLines := strings.Split(output, "\n")[1:]

	for i := 0; i < len(outputLines); i++ {
		fields := strings.Fields(outputLines[i])
		if len(fields) == 2 {
			devices = append(devices, Device{fields[0], fields[1]})
		} else if len(fields) != 0 {
			fmt.Println("Unexpected number of items in devices List.")
			log.Panicf("Devices list item:\n%v\n", fields)
		}
	}
	return devices
}

// DeviceNotFoundError notifies function callers that a requested device was unavailable.
type DeviceNotFoundError string

func (e DeviceNotFoundError) Error() string {
	return fmt.Sprintf("Device with ID %s could not be located.\n", e)
}

// DeviceConnected returns true if the specified device is connected.
func DeviceConnected(d Device) bool {
	devices := UpdateDevices()
	for i := 0; i < len(devices); i++ {
		if devices[i].ID == d.ID {
			return true
		}
	}
	return false
}

// FindDeviceByID attempts to locate and return a device with the specified ID.
func FindDeviceByID(id string) (Device, error) {
	devices := UpdateDevices()
	for i := 0; i < len(devices); i++ {
		if devices[i].ID == id {
			return devices[i], nil
		}
	}
	return Device{}, DeviceNotFoundError(id)
}

// PrintDevices writes list of connected devices to standard output.
func PrintDevices() {
	devices := UpdateDevices()
	fmt.Println("Connected devices: ")
	for idx, v := range devices {
		fmt.Printf("Device %v\n\tDevice: %v\n\tStatus: %v\n", idx, v.ID, v.Status)
	}
}
