package adb

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/httprunner/httprunner/v4/hrp/pkg/uixt"
)

func format(data map[string]string) string {
	result, _ := json.MarshalIndent(data, "", "\t")
	return string(result)
}

var listAndroidDevicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List all Android devices",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceList, err := uixt.GetAndroidDevices(serial)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		for _, d := range deviceList {
			if isDetail {
				fmt.Println(format(d.DeviceInfo()))
			} else {
				fmt.Println(d.Serial(), d.Usb())
			}
		}
		return nil
	},
}

var (
	serial   string
	isDetail bool
)

func init() {
	listAndroidDevicesCmd.Flags().StringVarP(&serial, "serial", "s", "", "filter by device's serial")
	listAndroidDevicesCmd.Flags().BoolVarP(&isDetail, "detail", "d", false, "print device's detail")
	androidRootCmd.AddCommand(listAndroidDevicesCmd)
}
