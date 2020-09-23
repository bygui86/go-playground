package main

import (
	"fmt"
	"math"
	"reflect"
)

const (
	sample = `{"integer":1597839200,"floating":1000000.000000001,"value":"TICKER_UPDATED","payload":"eyJzdHJlYW0iOiJidGN1c2R0QHRpY2tlciIsImRhdGEiOnsiZSI6IjI0aHJUaWNrZXIiLCJFIjoxNTk3ODMxMjU3Mjg1LCJzIjoiQlRDVVNEVCIsInAiOiItNDI2LjI2MDAwMDAwIiwiUCI6Ii0zLjQ4OCIsInciOiIxMTk1MC40OTY4Mzk1NCIsIngiOiIxMjIyMS45MTAwMDAwMCIsImMiOiIxMTc5NS42NTAwMDAwMCIsIlEiOiIwLjA0NjQwNDAwIiwiYiI6IjExNzk1LjY0MDAwMDAwIiwiQiI6IjAuMDIxNTQ4MDAiLCJhIjoiMTE3OTUuNjUwMDAwMDAiLCJBIjoiNC41MTg1NjgwMCIsIm8iOiIxMjIyMS45MTAwMDAwMCIsImgiOiIxMjI5Ni42OTAwMDAwMCIsImwiOiIxMTYxMi43MTAwMDAwMCIsInYiOiI4NjQ3NC42MzkwMzIwMCIsInEiOiIxMDMzNDE0OTAwLjQ1MjI0MDYyIiwiTyI6MTU5Nzc0NDg1NzIxNCwiQyI6MTU5NzgzMTI1NzIxNCwiRiI6Mzg3OTI5NTM3LCJMIjozODkyODUxMjQsIm4iOjEzNTU1ODh9fQ=="}`

	ticker = `{"type":"TICKER_UPDATED","generic_type":"ticker","source":"BINANCE","exchange":"EXCHANGE","ingestion_time":1597838775,"base_asset_symbol":"BTC","base_asset_class":"CRYPTO_ASSET","quote_asset_symbol":"USDT","quote_asset_class":"CRYPTO_ASSET","payload":"eyJzdHJlYW0iOiJidGN1c2R0QHRpY2tlciIsImRhdGEiOnsiZSI6IjI0aHJUaWNrZXIiLCJFIjoxNTk3ODMxMjU3Mjg1LCJzIjoiQlRDVVNEVCIsInAiOiItNDI2LjI2MDAwMDAwIiwiUCI6Ii0zLjQ4OCIsInciOiIxMTk1MC40OTY4Mzk1NCIsIngiOiIxMjIyMS45MTAwMDAwMCIsImMiOiIxMTc5NS42NTAwMDAwMCIsIlEiOiIwLjA0NjQwNDAwIiwiYiI6IjExNzk1LjY0MDAwMDAwIiwiQiI6IjAuMDIxNTQ4MDAiLCJhIjoiMTE3OTUuNjUwMDAwMDAiLCJBIjoiNC41MTg1NjgwMCIsIm8iOiIxMjIyMS45MTAwMDAwMCIsImgiOiIxMjI5Ni42OTAwMDAwMCIsImwiOiIxMTYxMi43MTAwMDAwMCIsInYiOiI4NjQ3NC42MzkwMzIwMCIsInEiOiIxMDMzNDE0OTAwLjQ1MjI0MDYyIiwiTyI6MTU5Nzc0NDg1NzIxNCwiQyI6MTU5NzgzMTI1NzIxNCwiRiI6Mzg3OTI5NTM3LCJMIjozODkyODUxMjQsIm4iOjEzNTU1ODh9fQ==","ohlcv_interval":60}`
)

func main() {
	fmt.Printf("sample json-string size: %d\n", deepSize(sample))        // 721
	fmt.Printf("sample json-bytes size: %d\n", deepSize([]byte(sample))) // 729

	fmt.Printf("ticker json-string - deep size: %d\n", deepSize(ticker))        // 906
	fmt.Printf("ticker json-bytes - deep size: %d\n", deepSize([]byte(ticker))) // 914
}

// deepSize returns the real size (in bytes) of the object
func deepSize(object interface{}) int64 {
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
