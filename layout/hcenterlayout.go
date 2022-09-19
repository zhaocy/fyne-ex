package layout

import "fyne.io/fyne/v2"

type HCenterLayout struct {
}

func (d *HCenterLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)

	for _, o := range objects {
		child := o.MinSize()
		w = child.Width
		h = child.Height
	}
	return fyne.NewSize(w, h)
}

func (d *HCenterLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	//初始位置
	pos := fyne.NewPos(0, 0)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		//计算每个object的x坐标
		pos.X = containerSize.Width/2 - size.Width/2
		o.Move(pos)

		// 计算每个Object的y轴偏移量 x轴不用偏移
		pos = pos.Add(fyne.NewPos(0, size.Height))
	}
}
