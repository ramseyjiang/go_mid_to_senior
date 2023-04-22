package gui

// Step  1: Define the abstract product interfaces

type Button interface {
	Click() string
}

type Checkbox interface {
	Check() string
}

// Step 2: Create concrete product implementations

type WindowsButton struct{}

func (w *WindowsButton) Click() string {
	return "Windows button clicked"
}

type MacOSButton struct{}

func (m *MacOSButton) Click() string {
	return "MacOS button clicked"
}

type WindowsCheckbox struct{}

func (w *WindowsCheckbox) Check() string {
	return "Windows checkbox checked"
}

type MacOSCheckbox struct{}

func (m *MacOSCheckbox) Check() string {
	return "MacOS checkbox checked"
}

// Step 3: Define the abstract factory interface

type FactoryGUI interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Step 4: Create concrete factory implementations

type WindowsFactory struct{}

func (w *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (w *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

type MacOSFactory struct{}

func (m *MacOSFactory) CreateButton() Button {
	return &MacOSButton{}
}

func (m *MacOSFactory) CreateCheckbox() Checkbox {
	return &MacOSCheckbox{}
}
