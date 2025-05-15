package xpdf

import (
	"bytes"
	"fmt"
	"os"

	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/ximg"
	"github.com/jung-kurt/gofpdf"
)

// Write ...
func Write(fpath string, lines []*Line) error {
	// 初始化PDF文档, A4: 210mm*279mm
	pdf := gofpdf.New("P", "mm", "A4", "") // 纵向A4纸
	pdf.SetAutoPageBreak(true, 10)         // 自动分页, 底部留白10mm
	// 设置字体
	pdf.SetFontLocation(xfile.GetHomeDir() + "/Library/Fonts/")
	pdf.AddUTF8Font("NotoSansSC", "", "NotoSansSC.ttf")
	//pdf.AddUTF8Font("NotoSansSC", "", "./NotoSansSC.ttf")
	// 设置字体大小
	pdf.SetFont("NotoSansSC", "", 12)
	// 创建新页面
	pdf.AddPage()
	_, lineHeight := pdf.GetFontSize()
	for _, line := range lines {
		switch line.Type {
		case TypeDivider:
			currentY := pdf.GetY()
			pdf.SetLineWidth(0.5)
			pdf.SetDrawColor(200, 200, 200)
			pdf.Line(10, currentY, 200, currentY)
			pdf.Ln(10)
		case TypePage:
			pdf.AddPage()
		case TypeImage:
			imgPath := line.Text
			imgBytes, err := os.ReadFile(imgPath)
			if err != nil {
				return err
			}
			// 获取图片格式
			var imgType string
			switch {
			case ximg.IsJPEG(imgBytes):
				imgType = "JPEG"
			case ximg.IsPNG(imgBytes):
				imgType = "PNG"
			default:
				return fmt.Errorf("unsupported image format: %s", imgPath)
			}
			// 计算图片尺寸（限制宽度为页面宽度的80%）
			maxWidth := 160.0 // 210mm(A4) - 25mm*2边距
			imgOption := gofpdf.ImageOptions{ReadDpi: true, ImageType: imgType}
			info := pdf.RegisterImageOptionsReader(imgPath, imgOption, bytes.NewReader(imgBytes))
			if info == nil {
				return fmt.Errorf("invalid image: %s, err: %v", imgPath, pdf.Error())
			}
			// 自动缩放图片
			imgWidth, imgHeight := info.Width(), info.Height()
			ratio := imgHeight / imgWidth
			displayWidth := min(maxWidth, imgWidth)
			displayHeight := displayWidth * ratio
			// 居中显示
			pageWidth, _ := pdf.GetPageSize()
			x := (pageWidth - displayWidth) / 2
			// 图片前留白
			if pdf.GetY()+displayHeight > 297 {
				pdf.AddPage()
			}
			// 插入图片并更新位置
			imgOptions := gofpdf.ImageOptions{ImageType: imgType}
			pdf.ImageOptions(imgPath, x, pdf.GetY(), displayWidth, displayHeight, false, imgOptions, 0, "")
			// 图片后留白
			pdf.Ln(displayHeight + 5)
		default:
			fontSize := line.Size
			if fontSize == 0 {
				fontSize = 14
			}
			pdf.SetFontSize(fontSize)
			pdf.SetTextColor(40, 40, 40)
			pdf.SetLeftMargin(10)
			pdf.SetRightMargin(10)
			pdf.MultiCell(0, lineHeight*1.5, line.Text, "", "", false)
			pdf.Ln(lineHeight * (fontSize / 14))
		}
	}
	return pdf.OutputFileAndClose(fpath)
}
