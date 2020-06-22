package bank_example_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBankExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BankExample Suite")
}
