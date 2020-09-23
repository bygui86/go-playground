package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"time"
	"unsafe"
)

type EventA struct {
	Integer  int64   `json:"integer"`
	Floating float64 `json:"floating"`
}

type EventB struct {
	Integer  int64   `json:"integer"`
	Floating float64 `json:"floating"`
	Value    string  `json:"value"`
}

type EventC struct {
	Integer  int64   `json:"integer"`
	Floating float64 `json:"floating"`
	Value    string  `json:"value"`
	Payload  []byte  `json:"payload"`
}

type EventD struct {
	Integer int64 `json:"integer"`
}

type EventE struct {
	Floating float64 `json:"floating"`
}

type EventF struct {
	Floating float32 `json:"floating"`
}

type EventG struct {
	Integer    int64   `json:"integer"`
	EventC     *EventC `json:"eventC"`
	JsonIgnore int64   `json:"-"`
}

func main() {
	// eventA := newEventA()
	// fmt.Printf("eventA: %+v", eventA)
	// fmt.Printf("eventA - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(eventA))        // 8
	// fmt.Printf("eventA - reflect size:\t\t%d\n", reflect.TypeOf(eventA).Size()) // 8
	// fmt.Printf("eventA - deep size:\t\t%d\n", deepSize(eventA))                 // 40

	// eventB := newEventB()
	// fmt.Printf("eventB: %+v", eventB)
	// fmt.Printf("eventB - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(eventB))        // 8
	// fmt.Printf("eventB - reflect size:\t\t%d\n", reflect.TypeOf(eventB).Size()) // 8
	// fmt.Printf("eventB - deep size:\t\t%d\n", deepSize(eventB))                 // 86

	// eventC := newEventC()
	// fmt.Printf("eventC: %+v\n", eventC)
	// fmt.Printf("eventC - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(eventC))        // 8
	// fmt.Printf("eventC - reflect size:\t\t%d\n", reflect.TypeOf(eventC).Size()) // 8
	// fmt.Printf("eventC - deep size:\t\t%d\n", deepSize(eventC))                 // 594

	jsonEventC, cErr := json.Marshal(newEventC())
	if cErr != nil {
		panic(cErr)
	}
	fmt.Printf("jsonEventC: %s\n", string(jsonEventC))
	fmt.Printf("jsonEventC - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(jsonEventC))        // 24
	fmt.Printf("jsonEventC - reflect size:\t\t%d\n", reflect.TypeOf(jsonEventC).Size()) // 24
	fmt.Printf("jsonEventC - deep size:\t\t%d\n", deepSize(jsonEventC))                 // 2241

	// eventD := newEventD()
	// fmt.Printf("eventD: %+v", eventD)
	// fmt.Printf("eventD - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(eventD))        // 8
	// fmt.Printf("eventD - reflect size:\t\t%d\n", reflect.TypeOf(eventD).Size()) // 8
	// fmt.Printf("eventD - deep size:\t\t%d\n", deepSize(eventD))                 // 24

	// eventE := newEventE()
	// fmt.Printf("eventE: %+v", eventE)
	// fmt.Printf("eventE - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(eventE))        // 8
	// fmt.Printf("eventE - reflect size:\t\t%d\n", reflect.TypeOf(eventE).Size()) // 8
	// fmt.Printf("eventE - deep size:\t\t%d\n", deepSize(eventE))                 // 24

	// eventF := newEventF()
	// fmt.Printf("eventF: %+v", eventF)
	// fmt.Printf("eventF - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(eventF))        // 8
	// fmt.Printf("eventF - reflect size:\t\t%d\n", reflect.TypeOf(eventF).Size()) // 8
	// fmt.Printf("eventF - deep size:\t\t%d\n", deepSize(eventF))                 // 16

	// eventG := newEventG()
	// fmt.Printf("eventG: %+v", eventG)
	// fmt.Printf("eventG - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(eventG))        // 8
	// fmt.Printf("eventG - reflect size:\t\t%d\n", reflect.TypeOf(eventG).Size()) // 8
	// fmt.Printf("eventG - deep size:\t\t%d\n", deepSize(eventG))                 // 642
	// jsonEventG, gErr := json.Marshal(eventG)
	// if gErr != nil {
	// 	panic(gErr)
	// }
	// fmt.Printf("jsonEventG: %s\n", string(jsonEventG))
	// fmt.Printf("jsonEventG - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(jsonEventG))        // 24
	// fmt.Printf("jsonEventG - reflect size:\t\t%d\n", reflect.TypeOf(jsonEventG).Size()) // 24
	// fmt.Printf("jsonEventG - deep size:\t\t%d\n", deepSize(jsonEventG))                 // 1505
}

func newEventA() *EventA {
	return &EventA{
		Integer:  time.Now().UTC().Unix(),
		Floating: 1_000_000.000000001,
	}
}

func newEventB() *EventB {
	return &EventB{
		Integer:  time.Now().UTC().Unix(),
		Floating: 1_000_000.000000001,
		Value:    "TICKER_UPDATED",
	}
}

func newEventC() *EventC {
	return &EventC{
		Integer:  time.Now().UTC().Unix(),
		Floating: 1_000_000.000000001,
		Value:    "TICKER_UPDATED",
		Payload:  []byte(`{"stream":"btcusdt@ticker","data":{"e":"24hrTicker","E":1597831257285,"s":"BTCUSDT","p":"-426.26000000","P":"-3.488","w":"11950.49683954","x":"12221.91000000","c":"11795.65000000","Q":"0.04640400","b":"11795.64000000","B":"0.02154800","a":"11795.65000000","A":"4.51856800","o":"12221.91000000","h":"12296.69000000","l":"11612.71000000","v":"86474.63903200","q":"1033414900.45224062","O":1597744857214,"C":1597831257214,"F":387929537,"L":389285124,"n":1355588}}`),
	}
}

func newEventD() *EventD {
	return &EventD{
		Integer: time.Now().UTC().Unix(),
	}
}

func newEventE() *EventE {
	return &EventE{
		Floating: 1_000_000_000.000000001111,
	}
}

func newEventF() *EventF {
	return &EventF{
		Floating: 1_000_000.000000001,
	}
}

func newEventG() *EventG {
	return &EventG{
		Integer:    time.Now().UTC().Unix(),
		EventC:     newEventC(),
		JsonIgnore: time.Now().UTC().Unix(),
	}
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
