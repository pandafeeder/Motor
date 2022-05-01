package data_utils

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
        p4 := Person{"Lilla",23}
        p1_p := &p1
        p2_p := p1_p
        p3_p := &p3
        p4_p := &p2
        persons_ptr := []*Person{p1_p, p2_p, p3_p, p4_p}

        uniq_persons_by_ptr_address := UniqSliceOfPtrByVal(persons_ptr)
        if len(uniq_persons_by_ptr_address) != 2 {
                t.Errorf("UniqSliceOfPtrByVal return wrong length as %d", len(uniq_persons_by_ptr_address))
        } else {
                ee1 := uniq_persons_by_ptr_address[0]
                ee2 := uniq_persons_by_ptr_address[1]
                if *ee1 != p1 || *ee2 != p3 {
                        t.Errorf("UniqSliceOfPtrByVal: return wrong value of elements as [%v, %v]", ee1, ee2)
                }
        }

        persons := []Person{p1, p2, p3, p4}
        uniq_persons_by_ptr_value   := UniqSliceOfNonPtrByVal(persons)
        if len(uniq_persons_by_ptr_value) != 2 {
                t.Errorf("UniqSliceOfNonPtrByVal return wrong length as %d", len(uniq_persons_by_ptr_value))
        } else {
                eee1 := uniq_persons_by_ptr_value[0]
                eee2 := uniq_persons_by_ptr_value[1]
                if eee1 != p1 || eee2 != p3 {
                        t.Errorf("UniqSliceOfNonPtrByVal: return wrong value of elements as [%v, %v]", eee1, eee2)
                }
        }
}

func TestIndexOf(t *testing.T) {
        int_slice := []int{1,2,3,4,5}
        no_existing_ele_id := IndexOf(int_slice, 0)
        if no_existing_ele_id != -1 {
                t.Error("Non existing ele index should be -1")
        }
        two_id := IndexOf(int_slice, 2)
        if two_id != 1 {
                t.Error("index of 2 in int_slice should be 1")
        }
}
