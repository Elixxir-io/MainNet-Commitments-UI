package form

import (
	"github.com/dtylman/gowd"
	"github.com/dtylman/gowd/bootstrap"
	jww "github.com/spf13/jwalterweatherman"
	"path/filepath"
)

// <div class="form-group">
// <label for="exampleInputFile">File input</label>
// <input type="file" id="exampleInputFile">
// <p class="help-block">Example block-level help text here.</p>
// </div>

// FileButton is a bootstrap "form-group" input
type FileButton struct {
	*gowd.Element
	input   *gowd.Element
	txt     *gowd.Element
	txt2    *gowd.Element
	lbl2    *gowd.Element
	helpTxt *gowd.Element
	v       ValidateFunc
	value   string
	caption string
}

// NewFileButton creates a bootstrap "form-group" containing an input with a given type and caption
func NewFileButton(caption string, v ValidateFunc) *FileButton {
	i := new(FileButton)
	i.v = v
	i.Element = bootstrap.NewElement("div", "form-group")

	i.lbl2 = gowd.NewElement("label")
	i.lbl2.SetAttribute("class", "form-control fakeFileInput")
	btn := bootstrap.NewElement("button", "labelButton")
	btn.SetText("Choose File")
	i.lbl2.AddElement(btn)
	// btn :=  = bootstrap.NewElement("label", "labelButton")
	i.txt2 = gowd.NewText("No file chosen")
	i.lbl2.AddElement(i.txt2)

	lbl := gowd.NewElement("label")
	i.txt = gowd.NewText(caption)
	lbl.AddElement(i.txt)

	i.input = bootstrap.NewElement("input", "form-control fileInput")
	i.input.SetAttribute("type", "file")
	i.input.SetAttribute("style", "width:0;height:0;padding:0;border:0;")
	i.helpTxt = bootstrap.NewElement("p", "help-block")

	i.AddElement(lbl)
	i.AddElement(i.input)
	i.AddElement(i.lbl2)
	i.AddElement(i.helpTxt)
	lbl.SetAttribute("for", i.input.GetID())
	i.lbl2.SetAttribute("for", i.input.GetID())
	i.helpTxt.Hidden = true
	i.caption = caption

	btn.OnEvent(gowd.OnClick, func(sender *gowd.Element, event *gowd.EventElement) {
		gowd.ExecJSNow("document.getElementById('" + i.lbl2.GetID() + "').click();")
	})

	i.input.OnEvent(gowd.OnChange, func(sender *gowd.Element, event *gowd.EventElement) {
		i.txt2.SetText(filepath.Base(event.GetValue()))
		i.value = event.GetValue()
	})

	return i
}

// SetPlaceHolder sets the input placeholder text
func (i *FileButton) SetPlaceHolder(placeHolder string) {
	i.input.SetAttribute("placeHolder", placeHolder)
}

// SetHelpText sets the input help text
func (i *FileButton) SetHelpText(help string) {
	i.helpTxt.SetText(help)
	i.helpTxt.Hidden = false
}

// HideHelpText hides the input help text
func (i *FileButton) HideHelpText() {
	i.helpTxt.Hidden = true
}

// SetValue sets the input value
func (i *FileButton) SetValue(value string) {
	if value == "" {
		value = "No file chosen"
	}
	i.txt2.SetText(filepath.Base(value))
	i.value = value
}

// GetValue returns the input value
func (i *FileButton) GetValue() string {
	return i.value
}

// SetFile sets the input file value
func (i *FileButton) SetFile(value string) {
	i.lbl2.SetText(value)
}

// Validate checks the value of the form input against the validator function.
// If validation fails, the error is set as the help text and returns true.
// If validations succeeds, it returns true.
func (i *FileButton) Validate() (interface{}, bool) {
	validated, helpText, err := i.v(i.value)
	if err != nil {
		jww.ERROR.Printf("Failed to validate input %+v: %+v", i.caption, err)
		i.SetHelpText(helpText)
		return nil, false
	}

	i.HideHelpText()

	return validated, true
}
