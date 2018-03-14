package joi_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/softbrewery/gojoi/pkg/joi"

	. "github.com/softbrewery/gojoi/pkg/joi"
)

var _ = Describe("Any", func() {

	Describe("Any", func() {

		It("Should create a new schema", func() {
			s := Any()
			Expect(s).ToNot(BeNil())
		})
	})

	Describe("Kind", func() {

		It("Should return interface", func() {
			s := Any()
			Expect(s.Kind()).To(Equal("interface"))
		})
	})

	Describe("Root", func() {

		It("Should equal itselves", func() {
			s := Any()
			Expect(s.Root()).To(Equal(s))
		})
	})

	Describe("Description", func() {

		// TODO
		It("Should be stored", func() {
			s := Any().Description("my description")
			Expect(s).NotTo(BeNil())
		})
	})

	Describe("Required", func() {

		It("Error should be nil if value is required and set", func() {
			s := Any().Required()
			Expect(s.Validate(10)).To(BeNil())
		})

		It("Error should be not nil if value is required and not set", func() {
			s := Any().Required()
			Expect(s.Validate(nil)).To(Equal(ErrRequired))
		})
	})

	Describe("Forbidden", func() {

		It("Error should be nil if value is forbidden and not set", func() {
			s := Any().Forbidden()
			Expect(s.Validate(nil)).To(BeNil())
		})

		It("Error should be not nil if value is forbidden and set", func() {
			s := Any().Forbidden()
			Expect(s.Validate(10)).To(Equal(ErrForbidden))
		})
	})

	Describe("Allow", func() {

		It("Error should be nil if value is in allow list (int)", func() {
			s := Any().Allow(0, 10, 20)
			Expect(s.Validate(0)).To(BeNil())
			Expect(s.Validate(10)).To(BeNil())
			Expect(s.Validate(20)).To(BeNil())

			Expect(s.Validate(100)).To(Equal(ErrAllow))
		})

		It("Error should be nil if value is in allow list (string)", func() {
			s := Any().Allow("id", "name", "isbn")
			Expect(s.Validate("id")).To(BeNil())
			Expect(s.Validate("name")).To(BeNil())
			Expect(s.Validate("isbn")).To(BeNil())

			Expect(s.Validate("author")).To(Equal(ErrAllow))
		})

		It("Error should be nil if value is in allow list (bool)", func() {
			s := Any().Allow(true)
			Expect(s.Validate(true)).To(BeNil())

			Expect(s.Validate(false)).To(Equal(ErrAllow))
		})
	})

	Describe("Disallow", func() {

		It("Error should be not nil if value is in disallow list (int)", func() {
			s := Any().Disallow(0, 10, 20)
			Expect(s.Validate(0)).To(Equal(ErrDisallow))
			Expect(s.Validate(10)).To(Equal(ErrDisallow))
			Expect(s.Validate(20)).To(Equal(ErrDisallow))

			Expect(s.Validate(100)).To(BeNil())
		})

		It("Error should be not nil if value is in disallow list (string)", func() {
			s := Any().Disallow("id", "name", "isbn")
			Expect(s.Validate("id")).To(Equal(ErrDisallow))
			Expect(s.Validate("name")).To(Equal(ErrDisallow))
			Expect(s.Validate("isbn")).To(Equal(ErrDisallow))

			Expect(s.Validate("author")).To(BeNil())
		})

		It("Error should be not nil if value is in disallow list (bool)", func() {
			s := Any().Disallow(true)
			Expect(s.Validate(true)).To(Equal(ErrDisallow))

			Expect(s.Validate(false)).To(BeNil())
		})
	})

	Describe("Transform", func() {

		It("Should accept a PRE transform function", func() {
			f := func(value interface{}) (interface{}, error) {
				return value, nil
			}
			s := Any().Transform(joi.TransformStagePRE, f)
			Expect(s).NotTo(BeNil())
		})

		It("Should accept a POST transform function", func() {
			f := func(value interface{}) (interface{}, error) {
				return value, nil
			}
			s := Any().Transform(joi.TransformStagePOST, f)
			Expect(s).NotTo(BeNil())
		})

		It("Should allow to replace value", func() {
			// transform function that replaces a value
			f := func(value interface{}) (interface{}, error) {
				cValue, ok := value.(string)
				if !ok {
					return nil, errors.New("Failed to cast type")
				}
				if cValue == "id" {
					cValue = "name"
				}
				return cValue, nil
			}

			s := Any().
				Allow("name").
				Transform(TransformStagePRE, f)

			Expect(s.Validate("id")).To(BeNil())
		})
	})
})