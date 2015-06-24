package flydb

import (
    "fmt"
)

func convertToInt32(v interface{}) (int32, error) {
    switch typedValue := v.(type) {
    case int:
        return int32(typedValue), nil
    case int32:
        return int32(typedValue), nil
    case int64:
        return int32(typedValue), nil
    case float32:
        return int32(typedValue), nil
    case float64:
        return int32(typedValue), nil
    default:
        return 0, fmt.Errorf("Not a number")
    }
}

func convertToInt64(v interface{}) (int64, error) {
    switch typedValue := v.(type) {
    case int:
        return int64(typedValue), nil
    case int32:
        return int64(typedValue), nil
    case int64:
        return int64(typedValue), nil
    case float32:
        return int64(typedValue), nil
    case float64:
        return int64(typedValue), nil
    default:
        return 0, fmt.Errorf("Not a number")
    }
}

func convertToInt(v interface{}) (int, error) {
    switch typedValue := v.(type) {
    case int:
        return int(typedValue), nil
    case int32:
        return int(typedValue), nil
    case int64:
        return int(typedValue), nil
    case float32:
        return int(typedValue), nil
    case float64:
        return int(typedValue), nil
    default:
        return 0, fmt.Errorf("Not a number")
    }
}

func convertToFloat32(v interface{}) (float32, error) {
    switch typedValue := v.(type) {
    case int:
        return float32(typedValue), nil
    case int32:
        return float32(typedValue), nil
    case int64:
        return float32(typedValue), nil
    case float32:
        return float32(typedValue), nil
    case float64:
        return float32(typedValue), nil
    default:
        return 0, fmt.Errorf("Not a number")
    }
}

func convertToFloat64(v interface{}) (float64, error) {
    switch typedValue := v.(type) {
    case int:
        return float64(typedValue), nil
    case int32:
        return float64(typedValue), nil
    case int64:
        return float64(typedValue), nil
    case float32:
        return float64(typedValue), nil
    case float64:
        return float64(typedValue), nil
    default:
        return 0, fmt.Errorf("Not a number")
    }
}