package types

import (
        "testing"
)

type Person struct {
        name string
        age  uint16
}

func TestUniqSliceFuncs(t *testing.T) {
        p1 := Person{"Fry", 2158}
        p2 := Person{"Fry", 2158}
        p3 := Person{"Lilla",23}
        persons_val := []Person{p1, p2, p3}
        uniq_slice_of_persons_val := UniqSlice(persons_val)

        if len(uniq_slice_of_persons_val) != 2 {
                t.Errorf("UniqSlice: return wrong length as %d", len(uniq_slice_of_persons_val))
        } else {
                e1 := uniq_slice_of_persons_val[0]
                e2 := uniq_slice_of_persons_val[1]
                if e1 != p1 || e2 != p3 {
                        t.Errorf("UniqSlice: return wrong value of elements as [%v, %v]", e1, e2)
                }
        }

        p1_p := &p1
        p2_p := p1_p
        p3_p := &p3
        p4_p := &p2
        persons_ptr := []*Person{p1_p, p2_p, p3_p, p4_p}

        uniq_persons_by_ptr_address := UniqSliceOfPtr(persons_ptr)
        if len(uniq_persons_by_ptr_address) != 3 {
                t.Errorf("UniqSliceOfPtr return wrong length as %d", len(uniq_persons_by_ptr_address))
        } else {
                ee1 := uniq_persons_by_ptr_address[0]
                ee2 := uniq_persons_by_ptr_address[1]
                ee3 := uniq_persons_by_ptr_address[2]
                if *ee1 != p1 || *ee2 != p3 || *ee3 != p2 {
                        t.Errorf("UniqSliceOfPtr: return wrong value of elements as [%v, %v, %v]", ee1, ee2 ,ee3)
                }
        }

        uniq_persons_by_ptr_value   := UniqSliceOfVal(persons_ptr)
        if len(uniq_persons_by_ptr_value) != 2 {
                t.Errorf("UniqSliceOfVal return wrong length as %d", len(uniq_persons_by_ptr_value))
        } else {
                eee1 := uniq_persons_by_ptr_value[0]
                eee2 := uniq_persons_by_ptr_value[1]
                if *eee1 != p1 || *eee2 != p3 {
                        t.Errorf("UniqSliceOfVal: return wrong value of elements as [%v, %v]", eee1, eee2)
                }
        }
}
