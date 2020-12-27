package mediator

// by book implementation
// but in Go it should be done with channels

type WidType uint8

const (
	WBtnOK WidType = iota + 100
	WBtnCancel
	WListBox
	WEntryField
)

var ListVals = []string{"alpha", "betta", "gamma", "delta"}

type Widget interface {
	WType() WidType
}

type DialogDirector interface {
	WidgetChanged(w Widget)
}

type Button struct {
	Enabled  bool
	Name     string
	Type     WidType
	Director DialogDirector
}

func (b *Button) WType() WidType {
	return b.Type
}

type ListBox struct {
	List     []string
	Selected string
	Type     WidType
	Director DialogDirector
}

func NewListBox() *ListBox {
	return &ListBox{ListVals, "", WListBox, nil}
}

func (l *ListBox) WType() WidType {
	return l.Type
}

type EntryField struct {
	Text     string
	Active   bool
	Type     WidType
	Director DialogDirector
}

func (f *EntryField) WType() WidType {
	return f.Type
}

type FrontDialog struct {
	BtnOk     *Button
	BtnCancel *Button
	List      *ListBox
	Field     *EntryField
}

func (d *FrontDialog) WidgetChanged(w Widget) {
	if w.WType() == WListBox {
		d.BtnCancel.Enabled = true
		d.Field.Active = true
		d.Field.Text = d.List.Selected
	}
}

func NewDialog(ok *Button, cancel *Button, list *ListBox, field *EntryField) *FrontDialog {
	d := &FrontDialog{ok, cancel, list, field}
	ok.Director = d
	cancel.Director = d
	list.Director = d
	field.Director = d
	return d
}

func (d *FrontDialog) SetList(id int) {
	d.List.Selected = d.List.List[id]
	d.List.Director.WidgetChanged(d.List)
}
