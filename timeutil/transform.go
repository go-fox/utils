package timeutil

import (
	"time"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimestampUnixToTimestampPb converts a unix timestamp to a protobuf timestamp
func TimestampUnixToTimestampPb(unix *int64) *timestamppb.Timestamp {
	timeUnix := time.Unix(lo.FromPtr(unix), 0)
	return timestamppb.New(timeUnix)
}

// TimestampUnixMilliToTimestampPb converts a unix timestamp to a protobuf timestamp
func TimestampUnixMilliToTimestampPb(unixMilli *int64) *timestamppb.Timestamp {
	timeUnixMilli := time.UnixMilli(lo.FromPtr(unixMilli))
	return timestamppb.New(timeUnixMilli)
}

// TimestampPbToTimestampUnix converts a protobuf timestamp to a unix timestamp
func TimestampPbToTimestampUnix(timestamp *timestamppb.Timestamp) *int64 {
	return lo.ToPtr(timestamp.AsTime().Unix())
}

// TimestampPbToTimestampUnixMilli converts a protobuf timestamp to a unix timestamp
func TimestampPbToTimestampUnixMilli(timestamp *timestamppb.Timestamp) *int64 {
	return lo.ToPtr(timestamp.AsTime().UnixMilli())
}
