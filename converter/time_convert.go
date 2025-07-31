package converter

import (
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/samber/lo"

	"github.com/go-fox/utils/timeutil"
)

// NewTimestampUnixTimestampPbConverterPair returns a converter pair for converting between int64 and timestamppb.Timestamp
func NewTimestampUnixTimestampPbConverterPair() []copier.TypeConverter {
	srcType := lo.ToPtr(int64(0))
	dstType := &timestamppb.Timestamp{}

	fromFn := timeutil.TimestampUnixToTimestampPb
	toFn := timeutil.TimestampPbToTimestampUnix

	return NewGenericTypeConverterPair(srcType, dstType, fromFn, toFn)
}
