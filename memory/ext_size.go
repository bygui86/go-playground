package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"time"
	"unsafe"
)

const (
	payload = `{"stream":"btcusdt@ticker","data":{"e":"24hrTicker","E":1597831257285,"s":"BTCUSDT","p":"-426.26000000","P":"-3.488","w":"11950.49683954","x":"12221.91000000","c":"11795.65000000","Q":"0.04640400","b":"11795.64000000","B":"0.02154800","a":"11795.65000000","A":"4.51856800","o":"12221.91000000","h":"12296.69000000","l":"11612.71000000","v":"86474.63903200","q":"1033414900.45224062","O":1597744857214,"C":1597831257214,"F":387929537,"L":389285124,"n":1355588}}`
)

type SimpleEvent struct {
	Integer  int64   `json:"integer"`
	Floating float64 `json:"floating"`
	Value    string  `json:"value"`
	Payload  []byte  `json:"payload"`
}

type Event struct {
	Type             string `json:"type,omitempty"`         // shared.*PublishType
	GenericType      string `json:"generic_type,omitempty"` // shared.*GenericEventType
	Source           string `json:"source,omitempty"`
	Exchange         string `json:"exchange,omitempty"`
	IngestionTime    int64  `json:"ingestion_time,omitempty"` // t.UnixNano()
	BaseAssetSymbol  string `json:"base_asset_symbol,omitempty"`
	BaseAssetClass   string `json:"base_asset_class,omitempty"`
	QuoteAssetSymbol string `json:"quote_asset_symbol,omitempty"`
	QuoteAssetClass  string `json:"quote_asset_class,omitempty"`
	Payload          []byte `json:"payload,omitempty"`
	OhlcvInterval    int    `json:"ohlcv_interval,omitempty"`
	PublishRetries   int    `json:"-"`
}

func main() {
	// simpleEvent := newSimpleEvent()
	// fmt.Printf("simpleEvent: %+v\n", simpleEvent)
	// fmt.Printf("simpleEvent - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(simpleEvent))        // 8
	// fmt.Printf("simpleEvent - reflect size:\t\t%d\n", reflect.TypeOf(simpleEvent).Size()) // 8
	// fmt.Printf("simpleEvent - deep size:\t\t%d\n", deepSize(simpleEvent))                 // 594

	jsonSimpleEvent, cErr := json.Marshal(newSimpleEvent())
	if cErr != nil {
		panic(cErr)
	}
	fmt.Printf("jsonSimpleEvent: %s\n", string(jsonSimpleEvent))
	fmt.Printf("jsonSimpleEvent length: %d\n", len(jsonSimpleEvent))
	fmt.Printf("jsonSimpleEvent - unsafe sizeof:\t\t%d\n", unsafe.Sizeof(jsonSimpleEvent))        // 24
	fmt.Printf("jsonSimpleEvent - reflect size:\t\t%d\n", reflect.TypeOf(jsonSimpleEvent).Size()) // 24
	fmt.Printf("jsonSimpleEvent - deep size:\t\t%d\n", deepSize(jsonSimpleEvent))                 // 2241

	// event := newEvent()
	// fmt.Printf("event: %+v\n", event)
	// fmt.Printf("event - unsafe sizeof: \t %d\n", unsafe.Sizeof(event))        // 8
	// fmt.Printf("event - reflect size: \t %d\n", reflect.TypeOf(event).Size()) // 8
	// fmt.Printf("event - deep size: \t %d\n", deepSize(event))                 // 886

	jsonEvent, err := json.Marshal(newEvent())
	if err != nil {
		panic(err)
	}
	fmt.Printf("jsonEvent: %s\n", string(jsonEvent))
	fmt.Printf("jsonEvent length: %d\n", len(jsonEvent))
	fmt.Printf("jsonEvent - unsafe sizeof (bytes): \t %d\n", unsafe.Sizeof(jsonEvent))        // 24
	fmt.Printf("jsonEvent - reflect size (bytes): \t %d\n", reflect.TypeOf(jsonEvent).Size()) // 24
	fmt.Printf("jsonEvent - deep size (bytes): \t %d\n", deepSize(jsonEvent))                 // 1058
}

func newSimpleEvent() *SimpleEvent {
	return &SimpleEvent{
		Integer:  time.Now().UTC().Unix(),
		Floating: 1_000_000.000000001,
		Value:    "TICKER_UPDATED",
		Payload:  []byte(payload),
	}
}

func newEvent() *Event {
	return &Event{
		Type:             "TICKER_UPDATED",
		GenericType:      "ticker",
		Source:           "BINANCE",
		Exchange:         "EXCHANGE",
		IngestionTime:    time.Now().UTC().Unix(),
		BaseAssetSymbol:  "BTC",
		BaseAssetClass:   "CRYPTO_ASSET",
		QuoteAssetSymbol: "USDT",
		QuoteAssetClass:  "CRYPTO_ASSET",
		Payload:          []byte(payload),
		OhlcvInterval:    60,
		PublishRetries:   10,
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
