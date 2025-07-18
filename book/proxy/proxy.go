package proxy

import "fmt"

type Graphic interface {
	Draw() string
	Name() string
}

type Image struct {
	Type string
	File string
}

type ImageProxy struct {
	img      *Image
	FileName string
}

func (i *Image) Draw() string {
	return fmt.Sprintf("Draw %s. Type: %s", i.File, i.Type)
}

func (p *ImageProxy) GetImage() *Image {
	if p.img == nil {
		exten := p.FileName[len(p.FileName)-3:]
		p.img = &Image{exten, p.FileName}
		return p.img
	}
	return p.img
}

func (p *ImageProxy) Draw() string {
	return p.GetImage().Draw()
}

func (p *ImageProxy) Name() string {
	return p.GetImage().File
}

func ImageGraph(fileName string) Graphic {
	p := &ImageProxy{FileName: fileName}
	return p
}
