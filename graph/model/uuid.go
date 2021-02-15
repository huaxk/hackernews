package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(writer io.Writer) {
		str := u.String()
		io.WriteString(writer, strconv.Quote(str))

	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	str, ok := v.(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("UUID must be string")
	}

	return uuid.FromString(str)
}
