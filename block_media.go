package notion

import "time"

// Media stands for a group of: audio, video, image, file and pdf blocks
//
// Ref:
// 	https://developers.notion.com/page/changelog#media-blocks-video-audio-image-file-pdf
// 	https://developers.notion.com/reference/block#image
//  https://developers.notion.com/reference/block#file
//  https://developers.notion.com/reference/block#video
//  https://developers.notion.com/reference/block#file
//  https://developers.notion.com/reference/block#pdf
//

type FileBlock struct {
	BaseBlock
	File File `json:"file"`
}

func NewFileBlock(file File) *FileBlock {
	return &FileBlock{
		BaseBlock: NewBaseBlock(BlockTypeFile),
		File:      file,
	}
}

type PdfBlock struct {
	BaseBlock
	Pdf File `json:"pdf"`
}

func NewPdfBlock(pdf File) *PdfBlock {
	return &PdfBlock{
		BaseBlock: NewBaseBlock(BlockTypePdf),
		Pdf:       pdf,
	}
}

type ImageBlock struct {
	BaseBlock
	Image File `json:"image"`
}

func NewImageBlock(data File) *ImageBlock {
	return &ImageBlock{
		BaseBlock: NewBaseBlock(BlockTypeImage),
		Image:     data,
	}
}

type AudioBlock struct {
	BaseBlock
	Audio File `json:"audio"`
}

func NewAudioBlock(audio File) *AudioBlock {
	return &AudioBlock{
		BaseBlock: NewBaseBlock(BlockTypeAudio),
		Audio:     audio,
	}
}

type VideoBlock struct {
	BaseBlock
	Video File `json:"video"`
}

func NewVideoBlock(video File) *VideoBlock {
	return &VideoBlock{
		BaseBlock: NewBaseBlock(BlockTypeVideo),
		Video:     video,
	}
}

// Media is an interface for blocks that can be downloaded
// such as Pdf, FileBlock, and Image
type Media interface {
	GetURL() string
	GetExpiryTime() *time.Time
}

// GetURL implements Media interface for PdfBlock
func (b *PdfBlock) GetURL() string { return b.Pdf.GetURL() }

// GetExpiryTime implements Media interface for PdfBlock
func (b *PdfBlock) GetExpiryTime() *time.Time { return b.Pdf.GetExpiryTime() }

// GetURL implements Media interface for FileBlock
func (b *FileBlock) GetURL() string { return b.File.GetURL() }

// GetExpiryTime implements Media interface for FileBlock
func (b *FileBlock) GetExpiryTime() *time.Time { return b.File.GetExpiryTime() }

// GetURL implements Media interface for ImageBlock
func (b *ImageBlock) GetURL() string { return b.Image.GetURL() }

// GetExpiryTime implements Media interface for ImageBlock
func (b *ImageBlock) GetExpiryTime() *time.Time { return b.Image.GetExpiryTime() }

// GetURL implements Media interface for AudioBlock
func (b *AudioBlock) GetURL() string { return b.Audio.GetURL() }

// GetExpiryTime implements Media interface for AudioBlock
func (b *AudioBlock) GetExpiryTime() *time.Time { return b.Audio.GetExpiryTime() }

// GetURL implements Media interface for VideoBlock
func (b *VideoBlock) GetURL() string { return b.Video.GetURL() }

// GetExpiryTime implements Media interface for VideoBlock
func (b *VideoBlock) GetExpiryTime() *time.Time { return b.Video.GetExpiryTime() }

// Verify that types implement Media interface
var (
	_ Media = (*PdfBlock)(nil)
	_ Media = (*FileBlock)(nil)
	_ Media = (*ImageBlock)(nil)
	_ Media = (*AudioBlock)(nil)
	_ Media = (*VideoBlock)(nil)
)

var (
	_ Block             = (*PdfBlock)(nil)
	_ HierarchicalBlock = (*PdfBlock)(nil)

	_ Block             = (*FileBlock)(nil)
	_ HierarchicalBlock = (*FileBlock)(nil)

	_ Block             = (*ImageBlock)(nil)
	_ HierarchicalBlock = (*ImageBlock)(nil)

	_ Block             = (*AudioBlock)(nil)
	_ HierarchicalBlock = (*AudioBlock)(nil)

	_ Block             = (*VideoBlock)(nil)
	_ HierarchicalBlock = (*VideoBlock)(nil)
)
