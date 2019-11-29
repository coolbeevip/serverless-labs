package openfaasfn

import (
	"fmt"
)

func Handler(req []byte) string {
	return fmt.Sprintf("Hi,I'm OpenFaaS. I have received your message '%s'", string(req))
}
