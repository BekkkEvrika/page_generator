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
	return &inputs.Input{Type: types[11], Name: f.getName()}
}

func submit(f *FieldType) *inputs.Input {
	return &inputs.Input{Type: types[12], Name: f.getName()}
}

func dateTime(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[2],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		Format:     globalDateFormat + " " + globalTimeFormat,
		SearchName: f.pgSearchObject,
	}
}

func dateF(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[1],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		Format:     globalDateFormat,
		SearchName: f.pgSearchObject,
	}
}

func numberView(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[4],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		SearchName: f.pgSearchObject,
	}
}

func checkBox(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[5],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		SearchName: f.pgSearchObject,
	}
}

func label(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[6],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		SearchName: f.pgSearchObject,
	}
}

func fileUploader(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:     types[10],
		Name:     f.getName(),
		FromName: f.getFromName(),
	}
}

func hidden(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[9],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		DataType:   f.pgDataType,
		SearchName: f.pgSearchObject,
	}
}

func searchView(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:     types[7],
		Name:     f.getName(),
		FromName: f.getFromName(),
		Search:   f.pgSearchSource,
		DataType: f.pgDataType,
	}
}

func textView(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[3],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		SearchName: f.pgSearchObject,
	}
}

func textField(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[8],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		SearchName: f.pgSearchObject,
	}
}

func comboBox(f *FieldType) *inputs.Input {
	return &inputs.Input{
		Type:       types[0],
		Name:       f.getName(),
		FromName:   f.getFromName(),
		ReadOnly:   f.PgReadOnly,
		SearchName: f.pgSearchObject,
	}
}
