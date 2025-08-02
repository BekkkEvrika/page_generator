package page_generator

import "github.com/BekkkEvrika/page_generator/inputs"

var types = []string{"combo-box", "date-time", "text-view", "number-view", "check-box", "label", "search-view", "text-field", "hidden", "auto-complete", "file-uploader"}

type createInput func(f *FieldType) *inputs.Input

var inputCreators map[string]createInput

func creatorsInit() {
	inputCreators = make(map[string]createInput)
	inputCreators[types[0]] = comboBox
	inputCreators[types[1]] = dateTime
	inputCreators[types[2]] = textView
	inputCreators[types[3]] = numberView
	inputCreators[types[4]] = checkBox
	inputCreators[types[5]] = label
	inputCreators[types[6]] = searchView
	inputCreators[types[7]] = textField
	inputCreators[types[8]] = hidden
	inputCreators[types[9]] = autoComplete
	inputCreators[types[10]] = fileUploader
}

func dateTime(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[1],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		Text:       f.PgText,
		Format:     globalDateFormat,
		Visible:    f.pgVisible,
		SearchName: f.pgSearchObject,
	}
	in.IsDefault, in.ValidMessage = f.getValidation()
	return &in
}

func numberView(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[3],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		Text:       f.PgText,
		MaxLength:  f.pgMax,
		Visible:    f.pgVisible,
		MinLength:  f.pgMin,
		SearchName: f.pgSearchObject,
	}
	in.IsDefault, in.ValidMessage = f.getValidation()
	return &in
}

func checkBox(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[4],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		SearchName: f.pgSearchObject,
		Visible:    f.pgVisible,
		Text:       f.PgText,
	}
	return &in
}

func label(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[5],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		Visible:    f.pgVisible,
		Template:   f.pgTemplate,
		Text:       f.PgText,
		SearchName: f.pgSearchObject}
	return &in
}

func autoComplete(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:     types[9],
		Name:     f.getName(),
		FromName: f.getFromName(),
		Visible:  f.pgVisible,
		Template: f.pgTemplate,
		Text:     f.PgText,
	}
	return &in
}

func fileUploader(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[10],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		Visible:     f.pgVisible,
		Text:        f.PgText,
		FileSource:  f.pgFileSource,
		FileMaxSize: f.pgFileMaxSize,
	}
	return &in
}

func hidden(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[8],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		DataType:   f.pgDataType,
		SearchName: f.pgSearchObject,
	}
	return &in
}

func searchView(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:     types[6],
		Name:     f.getName(),
		FromName: f.getFromName(),
		Visible:  f.pgVisible,
		Search:   f.pgSearchSource,
		Text:     f.PgText,
		DataType: f.pgDataType,
	}
	return &in
}

func textView(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[2],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		Text:       f.PgText,
		MaxLength:  f.getGormSize(),
		SearchName: f.pgSearchObject,
		Visible:    f.pgVisible,
	}
	in.IsDefault, in.ValidMessage = f.getValidation()
	return &in
}

func textField(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[7],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		Text:       f.PgText,
		SearchName: f.pgSearchObject,
		MaxLength:  f.getGormSize(),
		Visible:    f.pgVisible,
	}
	in.IsDefault, in.ValidMessage = f.getValidation()
	return &in
}

func comboBox(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[0],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		Visible:    f.pgVisible,
		Text:       f.PgText,
		SearchName: f.pgSearchObject,
	}
	in.IsDefault, in.ValidMessage = f.getValidation()
	return &in
}
