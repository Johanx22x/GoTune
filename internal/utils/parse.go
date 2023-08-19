package utils

import (
    "strconv"
)

// StringToInt converts a string to an int 
func StringToInt(s string) (int, error) {
    i, err := strconv.Atoi(s)
    if err != nil {
        return 0, err
    }

    return i, nil
}

// StringToInt64 converts a string to an int64 
func StringToInt64(s string) (int64, error) {
    i, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        return 0, err
    }

    return i, nil
}
