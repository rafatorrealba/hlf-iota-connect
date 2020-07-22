package integration_test

import (
	. "github.com/iotaledger/iota.go/api"
	. "github.com/iotaledger/iota.go/api/integration/samples"
	. "github.com/iotaledger/iota.go/consts"
	"strings"

	. "github.com/iotaledger/iota.go/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("GetInclusionStates()", func() {

	api, err := ComposeAPI(HTTPClientSettings{}, nil)
	if err != nil {
		panic(err)
	}

	Context("call", func() {
		It("resolves to correct response", func() {
			states, err := api.GetInclusionStates(DefaultHashes(), strings.Repeat("M", 81))
			Expect(err).ToNot(HaveOccurred())
			Expect(states[0]).To(BeTrue())
			Expect(states[1]).To(BeFalse())
		})
	})

	Context("invalid input", func() {
		It("returns an error for invalid transaction hashes", func() {
			_, err := api.GetInclusionStates(Hashes{"balalaika"}, strings.Repeat("M", 81))
			Expect(errors.Cause(err)).To(Equal(ErrInvalidTransactionHash))
		})

		It("returns an error for invalid tip hashes", func() {
			_, err := api.GetInclusionStates(DefaultHashes(), "balalaika")
			Expect(errors.Cause(err)).To(Equal(ErrInvalidTransactionHash))
		})
	})

})
