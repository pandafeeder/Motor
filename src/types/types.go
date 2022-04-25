/*
utils dealing with generic slice/mapping operation
*/

package types


func UniqSlice [V comparable] (slice []V) ([]V) {
        uniq_slice := make([]V, 0, 0)
        for _, ele_i := range slice {
                if len(uniq_slice) == 0 {
                        uniq_slice = append(uniq_slice, ele_i)
                        continue
                }
                has_match := 0
                for _, ele_j := range uniq_slice {
                        if ele_i == ele_j {
                                has_match = 1
                                break
                        }
                }
                if has_match == 0 {
                        uniq_slice = append(uniq_slice, ele_i)
                }
        }
        return uniq_slice
}

// pointer is used to be compared NOT the value pointer pointing to!!!
func UniqSliceOfPtr [V comparable] (slice []*V) ([]*V) {
        uniq_slice := make([]*V, 0, 0)
        for _, ele_i := range slice {
                if len(uniq_slice) == 0 {
                        uniq_slice = append(uniq_slice, ele_i)
                        continue
                }
                has_match := 0
                for _, ele_j := range uniq_slice {
                        if ele_i == ele_j {
                                has_match = 1
                                break
                        }
                }
                if has_match == 0 {
                        uniq_slice = append(uniq_slice, ele_i)
                }
        }
        return uniq_slice
}

// actual value pointed by pointer is used to be compared 
func UniqSliceOfVal [V comparable] (slice []*V) ([]*V) {
        uniq_slice := make([]*V, 0, 0)
        for _, ele_i := range slice {
                if len(uniq_slice) == 0 {
                        uniq_slice = append(uniq_slice, ele_i)
                        continue
                }
                has_match := 0
                for _, ele_j := range uniq_slice {
                        if *ele_i == *ele_j {
                                has_match = 1
                                break
                        }
                }
                if has_match == 0 {
                        uniq_slice = append(uniq_slice, ele_i)
                }
        }
        return uniq_slice
}

