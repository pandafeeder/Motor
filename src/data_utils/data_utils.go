/*
utils dealing with generic slice/mapping operation
*/

package data_utils

// actual value pointed by pointer is used to be compared 
func UniqSliceOfPtrByVal [V comparable] (slice []*V) ([]*V) {
        if len(slice) == 0 { return slice }
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

func UniqSliceOfNonPtrByVal [V comparable] (slice []V) ([]V) {
        if len(slice) == 0 { return slice }
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


func IndexOf [V comparable] (slice []V, ele V) int {
        index := -1
        for id, s := range slice {
                if ele == s {
                        return id
                }
        }
        return index
}


