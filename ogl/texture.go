package ogl

import (
	"image/png"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Texture struct {
	Width     uint32
	Height    uint32
	textureID uint32
}

func CreateTexture(texturePath string) (Texture, error) {
	var t Texture
	t.createTexture()
	t.SetTextureParameters()
	err := t.LoadFromFile(texturePath)
	return t, err
}
func (t *Texture) Delete() {
	gl.DeleteTextures(1, &t.textureID)
	t.textureID = 0

}
func (t *Texture) GetTextureID() uint32 {
	return t.textureID
}
func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.textureID)
}
func (t *Texture) UnBind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}
func (t *Texture) SetTextureParameters() {
	t.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
}
func (t *Texture) SetTextureData(width int32, height int32, internalFormat int32, format uint32, typ uint32, data []uint8) {
	t.Height = uint32(height)
	t.Width = uint32(width)
	t.Bind()
	gl.TexImage2D(gl.TEXTURE_2D, 0, internalFormat, width, height, 0, format, typ, gl.Ptr(data))
}
func (t *Texture) LoadFromFile(path string) error {
	t.Bind()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return err
	}
	var data []uint8
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	// for y := height - 1; y >= 0; y-- {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			app := func(r uint32, g uint32, b uint32, a uint32) {
				data = append(data, uint8(r))
				data = append(data, uint8(g))
				data = append(data, uint8(b))
				data = append(data, uint8(a))
			}
			app(img.At(x, y).RGBA())
		}
	}
	t.SetTextureData(int32(width), int32(height), gl.RGBA, gl.RGBA, gl.UNSIGNED_BYTE, data)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	t.UnBind()
	return nil
}
func (t *Texture) createTexture() {
	gl.GenTextures(1, &t.textureID)

}
