package login

// define behavior-oriented tests using Ginkgo and Gomega.
import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// in terminal,
// go install github.com/onsi/ginkgo/v2/ginkgo@latest
// go get github.com/onsi/gomega
// In the end, using command `ginkgo` to run the tests
func TestLoginFeature(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Login Feature Suite")
}

var _ = Describe("Login Feature", func() {
	Context("With valid credentials", func() {
		It("should allow access when username and password are correct", func() {
			Expect(ValidateCredentials("testuser", "correctpassword")).To(BeTrue())
		})
	})

	Context("With invalid credentials", func() {
		It("should deny access when the password is incorrect", func() {
			Expect(ValidateCredentials("testuser", "wrongpassword")).To(BeFalse())
		})

		It("should deny access when the username does not exist", func() {
			Expect(ValidateCredentials("nonexistentuser", "password")).To(BeFalse())
		})
	})
})
