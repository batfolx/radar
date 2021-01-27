package radar

import (
	"github.com/tarm/serial"
)

const MAX_BYTES = 1024

func GetUSBDevice(usb string) (*serial.Port, error) {
	config := &serial.Config{
		Name: usb,
		Baud: 9600,
	}

	return serial.OpenPort(config)
}

func RecvDevice(usb *serial.Port, buffer []byte, until byte) (int, error) {
	// keeps reading until last character is `until` character
	offset := 0
	for {
		n, err := usb.Read(buffer[offset:])
		offset += n
		if offset > MAX_BYTES {
			return 0, nil
		}

		if err != nil {
			return 0, nil
		} else if buffer[offset-1] == until {
			return offset, nil
		}

	}
}
