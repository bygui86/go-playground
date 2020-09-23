package main

import (
	"math"
	"reflect"
)

/*
	Improved version of https://github.com/creachadair/misctools/blob/4c938d6077e47ed42f6e7d6c8794b54e236dd23d/sizeof/size.go

	DeepSize reports the size of v in bytes, as reflect.Size, but also including
	all recursive substructures of v via pointers, slices, maps, structs and strings.
	If v contains any cycle, the size of each pointer (re)introducing the cycle is
	also included.

	Only values whose size and structure can be obtained by the reflect package
	are counted. Some values have components that are not visible by reflection,
	so are not counted or may be under-counted. In particular:
	   . The space occupied by code and data, reachable through variables captured in
		 the closure of a function pointer, are not counted. A value of function type
		 is counted only as a pointer.
	   . The unused buckets of a map cannot be inspected by the reflect package.
		 Their size is estimated by assuming unfilled slots contain zeroes of their
		 type.
	   . The unused capacity of the array under a slice is estimated by assuming the
		 unused slots contain zeroes of their type. It is possible they contain non
		 zero values from sharing or re-slicing, but without explicitly re-slicing the
		 reflect package cannot touch them.
*/
func DeepSize(object interface{}) int64 {
	return int64(
		deepSizeRecurse(
			reflect.ValueOf(object),
			make(map[uintptr]bool),
		),
	)
}

// deepSizeRecurse is the recursive part of deepSize
func deepSizeRecurse(object reflect.Value, seen map[uintptr]bool) uintptr {
	baseTypeSize := object.Type().Size()
	// fmt.Printf("[DEBUG] deepsize - type (%v) kind (%v) baseTypeSize (%d)\n", object.Type(), object.Kind(), baseTypeSize)

	totalSize := baseTypeSize

	switch object.Kind() {

	case reflect.Ptr:
		// fmt.Printf("[DEBUG] deepsize - pointer (%v)\n", object)
		p := object.Pointer()
		if !seen[p] && !object.IsNil() {
			seen[p] = true
			elemSize := deepSizeRecurse(object.Elem(), seen)
			// fmt.Printf("[DEBUG] deepsize - pointer (%v) elemSize (%d)\n", object, elemSize)
			totalSize += elemSize
		}

	case reflect.Slice:
		// fmt.Printf("[DEBUG] deepsize - slice (%v)\n", object)
		length := object.Len()
		for i := 0; i < length; i++ {
			elemSize := deepSizeRecurse(object.Index(i), seen)
			// fmt.Printf("[DEBUG] deepsize - slice (%v) index (%d) elemSize (%d)\n", object, i, elemSize)
			totalSize += elemSize
		}

		// Account for the parts of the array not covered by this slice.  Since
		// we can't get the values directly, assume they're zeroes. That may be
		// incorrect, in which case we may underestimate.
		capacity := object.Cap()
		if capacity > length {
			addSliceSize := object.Type().Size() * uintptr(capacity-length)
			// fmt.Printf("[DEBUG] deepsize - slice (%v) addSliceSize (%d)\n", object, addSliceSize)
			totalSize += addSliceSize
		}

	case reflect.Map:
		// fmt.Printf("[DEBUG] deepsize - map (%v)\n", object)

		// A map m has len(m) / 6.5 buckets, rounded up to a power of two, and
		// a minimum of one bucket. Each bucket is 16 bytes + 8*(keysize + valsize).
		bucketsNumber := uintptr(math.Pow(2, math.Ceil(math.Log(float64(object.Len())/6.5)/math.Log(2))))
		if bucketsNumber == 0 {
			bucketsNumber = 1
		}

		// We can't tell which keys are in which bucket by reflection, however,
		// here we count the 16-byte header for each bucket, and then just add
		// in the computed key and value sizes.
		bucketsTotalSize := 16 * bucketsNumber
		// fmt.Printf("[DEBUG] deepsize - map (%v) bucketsTotalSize(%d) \n", object, bucketsTotalSize)
		totalSize += bucketsTotalSize
		for _, key := range object.MapKeys() {
			keySize := deepSizeRecurse(key, seen)
			// fmt.Printf("[DEBUG] deepsize - map (%v) mapKeySize(%d) \n", object, keySize)
			totalSize += keySize

			valueSize := deepSizeRecurse(object.MapIndex(key), seen)
			// fmt.Printf("[DEBUG] deepsize - map (%v) mapValueSize(%d) \n", object, deepSizeRecurse)
			totalSize += valueSize
		}

		// We have 'bucketsNumber' buckets of 8 slots each, and object.Len() slots are filled.
		// The remaining slots we will assume contain zero key/value pairs.
		zeroKey := object.Type().Key().Size()    // a zero key
		zeroValue := object.Type().Elem().Size() // a zero value
		emptyPartSize := (8*bucketsNumber - uintptr(object.Len())) * (zeroKey + zeroValue)
		// fmt.Printf("[DEBUG] deepsize - map (%v) emptyPartSize(%d) \n", object, emptyPartSize)
		totalSize += emptyPartSize

	case reflect.Struct:
		// fmt.Printf("[DEBUG] deepsize - struct (%v)\n", object)

		// Chase fields and add their size.
		for i := 0; i < object.NumField(); i++ {
			f := object.Field(i)
			fieldSize := deepSizeRecurse(f, seen)
			// fmt.Printf("[DEBUG] deepsize - struct (%v) - field (%v) size (%d)\n", object, f, fieldSize)
			totalSize += fieldSize
		}

	case reflect.String:
		strSize := uintptr(object.Len())
		// fmt.Printf("[DEBUG] deepsize - string (%v) size (%d)\n", object, strSize)
		totalSize += strSize
	}

	// fmt.Printf("[INFO] deepsize - type (%v) kind (%v) size (%d)\n", object.Type(), object.Kind(), totalSize)
	return totalSize
}
