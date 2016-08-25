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

	Context("resizing a photo", func() {

		BeforeEach(func() {
			testImage, _ = GetPng("./../static/images/University of Houston Logo.png")
			scale = NewTransformation()
		})
		Context("that is larger than the base image", func() {

			It("should scale up with a horozontal (x) pixel scale and known measuremnt", func() {
				scale.Length = 4
				scale.Line.Start.X = 0
				scale.Line.Start.Y = 0
				scale.Line.End.X = 144
				scale.Line.End.Y = 0

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				xBounds := scaledImg.Bounds().Dx()
				Ω(xBounds).Should(Equal(386))
			})

			It("should scale up with a vertical (y) pixel scale and known measuremnt", func() {
				scale.Length = 4
				scale.Line.Start.X = 0
				scale.Line.End.Y = 0
				scale.Line.Start.Y = 0
				scale.Line.End.Y = 144
				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				xBounds := scaledImg.Bounds().Dy()
				Ω(xBounds).Should(Equal(432))
			})
		})

		Context("that is smaller than the base image", func() {

			It("should scale down with a horozontal (x) pixel scale and known measuremnt", func() {
				scale.Length = 1
				scale.Line.Start.X = 0
				scale.Line.End.X = 144
				scale.Line.Start.Y = 0
				scale.Line.End.Y = 0
				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				xBounds := scaledImg.Bounds().Dx()
				Ω(xBounds).Should(Equal(96)) // TODO: It is not clear what is determineing this rounding
			})

			It("should scale down with a vertical (y) pixel scale and known measuremnt", func() {
				scale.Length = 1
				scale.Line.Start.X = 0
				scale.Line.End.X = 0
				scale.Line.Start.Y = 0
				scale.Line.End.Y = 144
				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dy()).Should(Equal(108))

			})
		})

		Context("when scale line is not perpendicular or parellel to the X or Y axis", func() {

			It("Should work with a downward sloaping line", func() {
				scale.Length = 4.02305555556
				scale.Line.Start.X = 0
				scale.Line.Start.Y = 0
				scale.Line.End.X = 193
				scale.Line.End.Y = 216

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dx()).Should(Equal(192))
			})
			It("Should work with a downward sloaping line", func() {
				scale.Length = 8.046111
				scale.Line.Start.X = 0
				scale.Line.Start.Y = 0
				scale.Line.End.X = 193
				scale.Line.End.Y = 216

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dx()).Should(Equal(385))
			})

			It("Should work with a upward sloaping line", func() {
				scale.Length = 4.02305555556
				scale.Line.Start.X = 0
				scale.Line.Start.Y = 216
				scale.Line.End.X = 193
				scale.Line.End.Y = 0

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dy()).Should(Equal(215))
			})
		})

		Context("with bearly off measuerments", func() {
			It("should accept measurments that close to even pixels ", func() {
				scale.Length = 4
				scale.Line.Start.X = 0
				scale.Line.End.X = 143
				scale.Line.Start.Y = 0
				scale.Line.End.Y = 0

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dx()).Should(Equal(388))

			})

			It("should accept measurments that close even", func() {
				scale.Length = 4
				scale.Line.Start.X = 0
				scale.Line.End.X = 0
				scale.Line.Start.Y = 0
				scale.Line.End.Y = 143

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dy()).Should(Equal(435))

			})

			It("1should accept measurments that close but not == ", func() {
				scale.Length = 4
				scale.Line.Start.X = 49
				scale.Line.End.X = 192
				scale.Line.Start.Y = 61
				scale.Line.End.Y = 63

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dx()).Should(Equal(388))

			})
			It("4should accept measurments that close but not == ", func() {
				scale.Length = 4
				scale.Line.Start.X = 49
				scale.Line.End.X = 192
				scale.Line.Start.Y = 63
				scale.Line.End.Y = 61

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dx()).Should(Equal(388))

			})
			It("5should accept measurments that close but not == ", func() {
				scale.Length = 4
				scale.Line.Start.X = 192
				scale.Line.End.X = 49
				scale.Line.Start.Y = 61
				scale.Line.End.Y = 63

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dy()).Should(Equal(434))

			})
			It("should accept x measurments that close but not == ", func() {
				scale.Length = 4
				scale.Line.Start.X = 192
				scale.Line.End.X = 49
				scale.Line.Start.Y = 63
				scale.Line.End.Y = 61

				scaledImg, err := NewScale(testImage, scale)

				Ω(err).ShouldNot(HaveOccurred())

				Ω(scaledImg.Bounds().Dy()).Should(Equal(434))

			})

		})

		Context("with bad data", func() {

			BeforeEach(func() {
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

})
