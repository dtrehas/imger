package edgedetection

import (
	"github.com/ernyoke/imger/blend"
	"github.com/ernyoke/imger/convolution"
	"github.com/ernyoke/imger/grayscale"
	"github.com/ernyoke/imger/padding"
	"image"
)

var horizontalKernel = convolution.Kernel{[][]float64{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}, 3, 3}

var verticalKernel = convolution.Kernel{[][]float64{
	{-1, -2, -1},
	{0, 0, 0},
	{1, 2, 1},
}, 3, 3}

func HorizontalSobelGray(gray *image.Gray, border padding.Border) (*image.Gray, error) {
	return convolution.ConvolveGray(gray, &horizontalKernel, image.Point{X: 1, Y: 1}, border)
}

func VerticalSobelGray(gray *image.Gray, border padding.Border) (*image.Gray, error) {
	return convolution.ConvolveGray(gray, &verticalKernel, image.Point{X: 1, Y: 1}, border)
}

func SobelGray(img *image.Gray, border padding.Border) (*image.Gray, error) {
	horizontal, error := HorizontalSobelGray(img, border)
	if error != nil {
		return nil, error
	}
	vertical, error := VerticalSobelGray(img, border)
	if error != nil {
		return nil, error
	}
	res, error := blend.AddGrayWeighted(horizontal, 0.5, vertical, 0.5)
	if error != nil {
		return nil, error
	}
	return res, nil
}

func HorizontalSobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return convolution.ConvolveGray(gray, &horizontalKernel, image.Point{X: 1, Y: 1}, border)
}

func VerticalSobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return convolution.ConvolveGray(gray, &verticalKernel, image.Point{X: 1, Y: 1}, border)
}

func SobelRGBA(img *image.RGBA, border padding.Border) (*image.Gray, error) {
	gray := grayscale.Grayscale(img)
	return SobelGray(gray, border)
}
