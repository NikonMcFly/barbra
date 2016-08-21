package imageScale_test

import (
	"image"
	_ "image/png"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pieperz/barbra/imageScale"
	"golang.org/x/exp/shiny/unit"
)

var _ = Describe("ImageScale", func() {
	var (
		testImage   image.Image
		pixelScale  unit.Value
		knownLength unit.Value
	)

	BeforeEach(func() {
		testImage, _ = GetPng("./../static/images/University of Houston Logo.png")
	})

	Context("resizing a photo", func() {

		Context("with Pixel Scale set to 0", func() {

			BeforeEach(func() {
				pixelScale = unit.Pixels(0)
				knownLength = unit.Inches(3)
			})

			It("should resize a photo to a given width", func() {

				axis := "x"
				scaledImg, err := ScaleImage(testImage, pixelScale, knownLength, axis)

				Ω(err).ShouldNot(HaveOccurred())

				bounds := scaledImg.Bounds()
				c := unit.Converter(Default)
				got := c.Convert(knownLength, unit.Px)

				Ω(float64(bounds.Dx())).Should(Equal(got.F))

			})

			It("should resize a photo to a given height", func() {

				axis := "y"
				scaledImg, err := ScaleImage(testImage, pixelScale, knownLength, axis)

				Ω(err).ShouldNot(HaveOccurred())

				bounds := scaledImg.Bounds()
				c := unit.Converter(Default)
				got := c.Convert(knownLength, unit.Px)

				Ω(float64(bounds.Dy())).Should(Equal(got.F))
			})
		})

		Context("with a Pixel Scale", func() {

			Context("that is larger than the base image", func() {

				BeforeEach(func() {
					pixelScale = unit.Pixels(144)
					knownLength = unit.Inches(4)
				})

				It("should scale up with a horozontal (x) pixel scale and known measuremnt", func() {

					axis := "x"
					scaledImg, err := ScaleImage(testImage, pixelScale, knownLength, axis)

					Ω(err).ShouldNot(HaveOccurred())

					xBounds := scaledImg.Bounds().Dx()
					Ω(xBounds).Should(Equal(386))
				})

				It("should scale up with a vertical (y) pixel scale and known measuremnt", func() {

					axis := "y"
					scaledImg, err := ScaleImage(testImage, pixelScale, knownLength, axis)

					Ω(err).ShouldNot(HaveOccurred())

					xBounds := scaledImg.Bounds().Dy()
					Ω(xBounds).Should(Equal(432))
				})
			})
		})

		Context("that is smaller than the base image", func() {

			BeforeEach(func() {
				pixelScale = unit.Pixels(144)
				knownLength = unit.Inches(1)
			})

			It("should scale down with a horozontal (x) pixel scale and known measuremnt", func() {

				axis := "x"
				scaledImg, err := ScaleImage(testImage, pixelScale, knownLength, axis)

				Ω(err).ShouldNot(HaveOccurred())

				xBounds := scaledImg.Bounds().Dx()
				Ω(xBounds).Should(Equal(96))
			})

			It("should scale down with a vertical (y) pixel scale and known measuremnt", func() {

				axis := "y"
				scaledImg, err := ScaleImage(testImage, pixelScale, knownLength, axis)

				Ω(err).ShouldNot(HaveOccurred())

				xBounds := scaledImg.Bounds().Dy()
				Ω(xBounds).Should(Equal(108))
			})
		})
	})
})
