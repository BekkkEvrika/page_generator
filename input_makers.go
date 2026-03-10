package page_generator

import "github.com/BekkkEvrika/page_generator/inputs"

var types = []string{"select", "date", "datetime", "text", "number", "checkbox", "label", "search", "textarea", "hidden", "file", "button", "submit"}

type createInput func(f *FieldType) *inputs.Input

var inputCreators map[string]createInput

func creatorsInit() {
	inputCreators = make(map[string]createInput)
	inputCreators[types[0]] = comboBox
	inputCreators[types[1]] = dateF
	inputCreators[types[2]] = dateTime
	inputCreators[types[3]] = textView
	inputCreators[types[4]] = numberView
	inputCreators[types[5]] = checkBox
	inputCreators[types[6]] = label
	inputCreators[types[7]] = searchView
	inputCreators[types[8]] = textField
	inputCreators[types[9]] = hidden
	inputCreators[types[10]] = fileUploader
	inputCreators[types[11]] = button
	inputCreators[types[12]] = submit
}

func button(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[11],
		Name:        f.getName(),
		Placeholder: f.PgText,
	}
	return &in
}

func submit(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[12],
		Name:        f.getName(),
		Placeholder: f.PgText,
	}
	return &in
}

func dateTime(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[2],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		ReadOnly:    f.PgReadOnly,
		Placeholder: f.PgText,
		Format:      globalDateFormat + " " + globalTimeFormat,
		SearchName:  f.pgSearchObject,
	}
	return &in
}

func dateF(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[1],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		ReadOnly:    f.PgReadOnly,
		Placeholder: f.PgText,
		Format:      globalDateFormat,
		SearchName:  f.pgSearchObject,
	}
	return &in
}

func numberView(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[4],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		ReadOnly:    f.PgReadOnly,
		Placeholder: f.PgText,
		SearchName:  f.pgSearchObject,
	}
	return &in
}

func checkBox(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[4],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		ReadOnly:    f.PgReadOnly,
		SearchName:  f.pgSearchObject,
		Placeholder: f.PgText,
	}
	return &in
}

func label(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[6],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		Placeholder: f.PgText,
		SearchName:  f.pgSearchObject}
	return &in
}

func autoComplete(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[9],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		Placeholder: f.PgText,
	}
	return &in
}

func fileUploader(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[10],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		Placeholder: f.PgText,
	}
	return &in
}

func hidden(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:       types[9],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		DataType:   f.pgDataType,
		SearchName: f.pgSearchObject,
	}
	return &in
}

func searchView(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[7],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		Search:      f.pgSearchSource,
		Placeholder: f.PgText,
		DataType:    f.pgDataType,
	}
	return &in
}

func textView(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[3],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		ReadOnly:    f.PgReadOnly,
		Placeholder: f.PgText,
		SearchName:  f.pgSearchObject,
	}
	return &in
}

func textField(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[8],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		ReadOnly:    f.PgReadOnly,
		Placeholder: f.PgText,
		SearchName:  f.pgSearchObject,
	}
	return &in
}

func comboBox(f *FieldType) *inputs.Input {
	in := inputs.Input{
		Type:        types[0],
		Name:        f.getName(),
		FromName:    f.getFromName(),
		ReadOnly:    f.PgReadOnly,
		Placeholder: f.PgText,
		SearchName:  f.pgSearchObject,
	}
	return &in
}
