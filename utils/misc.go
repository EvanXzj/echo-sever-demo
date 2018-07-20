package utils

import "strconv"

// ConvertID : String to uint parser
func ConvertID(sID string) (uint, error) {
	id, err := strconv.ParseUint(sID, 10, 32)

	if err != nil {
		return uint(0), err
	}

	return uint(id), nil
}
