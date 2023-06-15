package ethereum

import (
	"fmt"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	channel := "pubsub://mint"
	fmt.Printf("`%s`", strings.SplitAfterN(channel, "://", 2)[1])

}
