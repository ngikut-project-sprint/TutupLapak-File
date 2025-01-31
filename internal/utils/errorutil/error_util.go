package errorutil

import (
	"fmt"
	"runtime"
	"strings"
)

func ErrWithContext(
	err error,
	messages ...string,
) error {
	_, filename, line, ok := runtime.Caller(1)

	var context []string

	if ok {
		context = append(context, fmt.Sprintf("%s:%d", filename, line))
	} else {
		context = append(context, "?")
	}

	context = append(context, messages...)

	return fmt.Errorf("%s\n%w", strings.Join(context, " | "), err)
}
