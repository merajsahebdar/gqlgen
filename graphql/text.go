package graphql

import (
	"context"
	"encoding"
	"fmt"
	"io"
)

func TextMarshaler(ctx context.Context, v encoding.TextMarshaler) Marshaler {
	return textMarshalerAdapter{ctx, v}
}

type textMarshalerAdapter struct {
	Context context.Context
	encoding.TextMarshaler
}

func (a textMarshalerAdapter) MarshalGQL(w io.Writer) {
	b, err := a.MarshalText()
	if err != nil {
		AddError(a.Context, err)
		Null.MarshalGQL(w)
	} else {
		writeQuotedString(w, string(b))
	}
}

func UnmarshalText(res encoding.TextUnmarshaler, v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}

	return res.UnmarshalText([]byte(s))
}
