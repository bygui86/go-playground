package main

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var capacity = 750_000

var dynamicCapacity = 1_000

type Update struct {
	MarkPrice            float32              `protobuf:"fixed32,1,opt,name=mark_price,json=markPrice,proto3" json:"mark_price,omitempty"`
	FundingRate          float32              `protobuf:"fixed32,2,opt,name=funding_rate,json=fundingRate,proto3" json:"funding_rate,omitempty"`
	EventTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func BenchmarkAppendNoCapacity(b *testing.B) {
	// b.SkipNow()
	for x := 0; x < b.N; x++ {
		a := make([]Update, 0)
		// version 1 - simple
		// protoTime, _ := ptypes.TimestampProto(time.Now())
		for i := 0; i < capacity; i++ {
			// version 2 - more compliant to our use case
			protoTime, _ := ptypes.TimestampProto(time.Now())
			a = append(
				a,
				Update{
					MarkPrice:   11.1,
					FundingRate: 11.1,
					EventTime:   protoTime,
				},
			)
		}
	}
}

func BenchmarkAppendWithCapacity(b *testing.B) {
	// b.SkipNow()
	for x := 0; x < b.N; x++ {
		a := make([]Update, 0, capacity)
		// version 1 - simple
		// protoTime, _ := ptypes.TimestampProto(time.Now())
		for i := 0; i < capacity; i++ {
			// version 2 - more compliant to our use case
			protoTime, _ := ptypes.TimestampProto(time.Now())
			a = append(
				a,
				Update{
					MarkPrice:   22.2,
					FundingRate: 22.2,
					EventTime:   protoTime,
				},
			)
		}
	}
}

func BenchmarkArrayAccessWithSizeAndCapacity(b *testing.B) {
	// b.SkipNow()
	for x := 0; x < b.N; x++ {
		a := make([]Update, dynamicCapacity, capacity)
		// version 1 - simple
		// protoTime, _ := ptypes.TimestampProto(time.Now())
		for i := 0; i < capacity; i++ {
			// version 2 - more compliant to our use case
			protoTime, _ := ptypes.TimestampProto(time.Now())
			if i != 0 {
				if i%dynamicCapacity == 0 {
					a = a[:len(a)+dynamicCapacity]
				}
			}

			a[i] = Update{
				MarkPrice:   33.3,
				FundingRate: 33.3,
				EventTime:   protoTime,
			}
		}
	}
}
