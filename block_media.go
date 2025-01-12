package notion

import "time"

// Media stands for a group of: audio, video, image, file and pdf blocks
//
// Reference:
// 	https://developers.notion.com/page/changelog#media-blocks-video-audio-image-file-pdf
// 	https://developers.notion.com/reference/block#image
//  https://developers.notion.com/reference/block#file
//  https://developers.notion.com/reference/block#video
//  https://developers.notion.com/reference/block#file
//  https://developers.notion.com/reference/block#pdf
//

// FileBlock is a Notion block for files
type FileBlock struct {
	BasicBlock
	File File `json:"file"`
}

// NewFileBlock creates a new FileBlock
func NewFileBlock(file File) *FileBlock {
	return &FileBlock{
		BasicBlock: NewBasicBlock(BlockTypeFile),
		File:       file,
	}
}

// PdfBlock is a Notion block for PDF files
type PdfBlock struct {
	BasicBlock
	Pdf File `json:"pdf"`
}

// NewPdfBlock creates a new PdfBlock
func NewPdfBlock(pdf File) *PdfBlock {
	return &PdfBlock{
		BasicBlock: NewBasicBlock(BlockTypePdf),
		Pdf:        pdf,
	}
}

// ImageBlock is a Notion block for images
type ImageBlock struct {
	BasicBlock
	Image File `json:"image"`
}

// NewImageBlock creates a new ImageBlock
func NewImageBlock(data File) *ImageBlock {
	return &ImageBlock{
		BasicBlock: NewBasicBlock(BlockTypeImage),
		Image:      data,
	}
}

// AudioBlock is a Notion block for audio files
type AudioBlock struct {
	BasicBlock
	Audio File `json:"audio"`
}

// NewAudioBlock creates a new AudioBlock
func NewAudioBlock(audio File) *AudioBlock {
	return &AudioBlock{
		BasicBlock: NewBasicBlock(BlockTypeAudio),
		Audio:      audio,
	}
}

// VideoBlock is a Notion block for video files
type VideoBlock struct {
	BasicBlock
	Video File `json:"video"`
}

// NewVideoBlock creates a new VideoBlock
func NewVideoBlock(video File) *VideoBlock {
	return &VideoBlock{
		BasicBlock: NewBasicBlock(BlockTypeVideo),
		Video:      video,
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

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *PdfBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *FileBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *ImageBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *AudioBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *VideoBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*PdfBlock)(nil)
	_ HierarchicalBlock = (*PdfBlock)(nil)
	_ BasicBlockHolder  = (*PdfBlock)(nil)

	_ Block             = (*FileBlock)(nil)
	_ HierarchicalBlock = (*FileBlock)(nil)
	_ BasicBlockHolder  = (*FileBlock)(nil)

	_ Block             = (*ImageBlock)(nil)
	_ HierarchicalBlock = (*ImageBlock)(nil)
	_ BasicBlockHolder  = (*ImageBlock)(nil)

	_ Block             = (*AudioBlock)(nil)
	_ HierarchicalBlock = (*AudioBlock)(nil)
	_ BasicBlockHolder  = (*AudioBlock)(nil)

	_ Block             = (*VideoBlock)(nil)
	_ HierarchicalBlock = (*VideoBlock)(nil)
	_ BasicBlockHolder  = (*VideoBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeFile, func() Block { return &FileBlock{} })
	registerBlockDecoder(BlockTypePdf, func() Block { return &PdfBlock{} })
	registerBlockDecoder(BlockTypeImage, func() Block { return &ImageBlock{} })
	registerBlockDecoder(BlockTypeAudio, func() Block { return &AudioBlock{} })
	registerBlockDecoder(BlockTypeVideo, func() Block { return &VideoBlock{} })
}
