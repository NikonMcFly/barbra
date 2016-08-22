package imageScaler_test

import (
	"image"
	_ "image/png"

	"golang.org/x/exp/shiny/unit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pieperz/barbra/imageScaler"
)

var _ = Describe("ImageScaler", func() {
	var (
		testImage image.Image
		scale     *Scale
	)

	BeforeEach(func() {
		testImage, _ = GetPng("./../static/images/University of Houston Logo.png")
	})

	Context("resizing a photo", func() {

		Context("with Pixel Scale set to 0", func() {

			BeforeEach(func() {
				scale = NewTransformation()
				scale.Length = 3
			})

			It("should resize a photo to a given width", func() {
				scale.Line.Start.X = 0
				scale.Line.End.X = 0
				scale.Axis = "x"
				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				bounds := scaledImg.Bounds()
				c := unit.Converter(Default)
				got := c.Convert(scale.KnownLength(), unit.Px)

				Ω(float64(bounds.Dx())).Should(Equal(got.F))

			})

			It("should resize a photo to a given height", func() {

				scale.Line.Start.Y = 0
				scale.Line.End.Y = 0
				scale.Axis = "y"
				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				bounds := scaledImg.Bounds()
				c := unit.Converter(Default)
				got := c.Convert(scale.KnownLength(), unit.Px)

				Ω(float64(bounds.Dy())).Should(Equal(got.F))
			})
		})

		Context("with a Pixel Scale", func() {

			Context("that is larger than the base image", func() {

				BeforeEach(func() {
					scale = NewTransformation()
					scale.Length = 4
				})

				It("should scale up with a horozontal (x) pixel scale and known measuremnt", func() {
					scale.Line.Start.X = 0
					scale.Line.End.X = 144
					scale.Axis = "x"
					scaledImg, err := NewScale(testImage, scale)

					Ω(err).ShouldNot(HaveOccurred())

					xBounds := scaledImg.Bounds().Dx()
					Ω(xBounds).Should(Equal(386))
				})

				It("should scale up with a vertical (y) pixel scale and known measuremnt", func() {
					scale.Line.Start.Y = 0
					scale.Line.End.Y = 144
					scale.Axis = "y"
					scaledImg, err := NewScale(testImage, scale)

					Ω(err).ShouldNot(HaveOccurred())

					xBounds := scaledImg.Bounds().Dy()
					Ω(xBounds).Should(Equal(432))
				})
			})
		})

		Context("that is smaller than the base image", func() {

			BeforeEach(func() {
				scale = NewTransformation()
				scale.Length = 1
			})

			It("should scale down with a horozontal (x) pixel scale and known measuremnt", func() {

				scale.Line.Start.X = 0
				scale.Line.End.X = 144
				scale.Axis = "x"
				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				xBounds := scaledImg.Bounds().Dx()
				Ω(xBounds).Should(Equal(96))
			})

			It("should scale down with a vertical (y) pixel scale and known measuremnt", func() {

				scale.Line.Start.Y = 0
				scale.Line.End.Y = 144
				scale.Axis = "y"
				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				xBounds := scaledImg.Bounds().Dy()
				Ω(xBounds).Should(Equal(108))
			})
		})
	})
})
