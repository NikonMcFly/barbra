package imageScaler_test

import (
	"image"
	_ "image/png"

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

	Context("resizing a photo with bad data", func() {

		Context("with default scale values", func() {

			BeforeEach(func() {
				scale = NewTransformation()
				scale.Length = 0
				scale.Line.Start.X = 0
				scale.Line.End.X = 0
				scale.Line.Start.Y = 0
				scale.Line.End.Y = 0
			})

			It("should not change the photo in the x scale", func() {

				_, err := NewScale(testImage, scale)

				Ω(err).To(HaveOccurred())

			})
			Context("with Pixel Scale set to 0", func() {

				BeforeEach(func() {
					scale = NewTransformation()
					scale.Length = 4
					scale.Line.Start.Y = 0
					scale.Line.End.Y = 0
					scale.Line.Start.X = 0
					scale.Line.End.X = 0
				})

				It("should return an error", func() {

					_, err := NewScale(testImage, scale)

					Ω(err).To(HaveOccurred())
				})

				It("should return an error", func() {

					_, err := NewScale(testImage, scale)

					Ω(err).To(HaveOccurred())
				})

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
				Ω(xBounds).Should(Equal(97)) // TODO: this is rounding up from 96.5
			})

			It("should scale down with a vertical (y) pixel scale and known measuremnt", func() {

				scale.Line.Start.Y = 0
				scale.Line.End.Y = 144
				scale.Axis = "y"
				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dy()).Should(Equal(108))

			})
		})

		Context("when scale line is not perpendicular or parellel to the X or Y axis", func() {

			BeforeEach(func() {
				scale = NewTransformation()
				scale.Length = 4.02305555556
				scale.Axis = "xy"
			})

			It("Should work with a downward sloaping line", func() {
				scale.Line.Start.X = 0
				scale.Line.Start.Y = 0
				scale.Line.End.X = 193
				scale.Line.End.Y = 216

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dx()).Should(Equal(193))
			})

			It("Should work with a upward sloaping line", func() {
				scale.Line.Start.X = 0
				scale.Line.Start.Y = 216
				scale.Line.End.X = 193
				scale.Line.End.Y = 0

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dy()).Should(Equal(216))
			})
		})

	})
})
