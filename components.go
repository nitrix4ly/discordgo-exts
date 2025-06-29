package discordgo

import (
	"encoding/json"
	"fmt"
	"time"
)

// Component types for Discord's UI system
type ComponentType uint

const (
	ActionsRowComponent            ComponentType = 1
	ButtonComponent                ComponentType = 2
	SelectMenuComponent            ComponentType = 3
	TextInputComponent             ComponentType = 4
	UserSelectMenuComponent        ComponentType = 5
	RoleSelectMenuComponent        ComponentType = 6
	MentionableSelectMenuComponent ComponentType = 7
	ChannelSelectMenuComponent     ComponentType = 8
	SectionComponent               ComponentType = 9
	TextDisplayComponent           ComponentType = 10
	ThumbnailComponent             ComponentType = 11
	MediaGalleryComponent          ComponentType = 12
	FileComponentType              ComponentType = 13
	SeparatorComponent             ComponentType = 14
	ButtonGroupComponent           ComponentType = 15  // v2 addition
	TableComponent                 ComponentType = 16  // v2 addition
	ContainerComponent             ComponentType = 17
	ModalComponent                 ComponentType = 18  // v2 addition
	TabsComponent                  ComponentType = 19  // v2 addition
	AccordionComponent             ComponentType = 20  // v2 addition
)

// Base interface that all components must implement
type MessageComponent interface {
	json.Marshaler
	Type() ComponentType
}

// ===== BUILDER FACTORY =====

// Easy way to start building components
type ComponentBuilder struct{}

func NewBuilder() *ComponentBuilder {
	return &ComponentBuilder{}
}

func (cb *ComponentBuilder) Button(label string) *ButtonBuilder {
	return &ButtonBuilder{
		button: Button{
			Label: label,
			Style: PrimaryButton,
		},
	}
}

func (cb *ComponentBuilder) SelectMenu(customID string) *SelectMenuBuilder {
	return &SelectMenuBuilder{
		menu: SelectMenu{
			CustomID: customID,
		},
	}
}

func (cb *ComponentBuilder) TextInput(customID, label string) *TextInputBuilder {
	return &TextInputBuilder{
		input: TextInput{
			CustomID: customID,
			Label:    label,
			Style:    TextInputShort,
		},
	}
}

func (cb *ComponentBuilder) ActionsRow() *ActionsRowBuilder {
	return &ActionsRowBuilder{
		row: ActionsRow{
			Components: make([]MessageComponent, 0, 5),
		},
	}
}

// v2 additions
func (cb *ComponentBuilder) Modal(customID, title string) *ModalBuilder {
	return &ModalBuilder{
		modal: Modal{
			CustomID: customID,
			Title:    title,
		},
	}
}

func (cb *ComponentBuilder) Tabs(customID string) *TabsBuilder {
	return &TabsBuilder{
		tabs: Tabs{
			CustomID: customID,
		},
	}
}

// ===== BUTTON BUILDER =====

type ButtonBuilder struct {
	button Button
}

func (bb *ButtonBuilder) Style(style ButtonStyle) *ButtonBuilder {
	bb.button.Style = style
	return bb
}

// Common button styles for quick setup
func (bb *ButtonBuilder) Primary() *ButtonBuilder   { return bb.Style(PrimaryButton) }
func (bb *ButtonBuilder) Secondary() *ButtonBuilder { return bb.Style(SecondaryButton) }
func (bb *ButtonBuilder) Success() *ButtonBuilder   { return bb.Style(SuccessButton) }
func (bb *ButtonBuilder) Danger() *ButtonBuilder    { return bb.Style(DangerButton) }
func (bb *ButtonBuilder) Premium() *ButtonBuilder   { return bb.Style(PremiumButton) }

func (bb *ButtonBuilder) Link(url string) *ButtonBuilder {
	bb.button.Style = LinkButton
	bb.button.URL = url
	return bb
}

func (bb *ButtonBuilder) CustomID(id string) *ButtonBuilder {
	bb.button.CustomID = id
	return bb
}

func (bb *ButtonBuilder) Disabled(disabled bool) *ButtonBuilder {
	bb.button.Disabled = disabled
	return bb
}

func (bb *ButtonBuilder) Emoji(name, id string, animated bool) *ButtonBuilder {
	bb.button.Emoji = &ComponentEmoji{
		Name:     name,
		ID:       id,
		Animated: animated,
	}
	return bb
}

// v2 enhancements
func (bb *ButtonBuilder) Tooltip(text string) *ButtonBuilder {
	bb.button.Tooltip = text
	return bb
}

func (bb *ButtonBuilder) Badge(count int) *ButtonBuilder {
	bb.button.Badge = &count
	return bb
}

func (bb *ButtonBuilder) Loading(loading bool) *ButtonBuilder {
	bb.button.Loading = loading
	return bb
}

func (bb *ButtonBuilder) Size(size ButtonSize) *ButtonBuilder {
	bb.button.Size = size
	return bb
}

func (bb *ButtonBuilder) Build() Button {
	return bb.button
}

// ===== SELECT MENU BUILDER =====

type SelectMenuBuilder struct {
	menu SelectMenu
}

func (smb *SelectMenuBuilder) Placeholder(text string) *SelectMenuBuilder {
	smb.menu.Placeholder = text
	return smb
}

func (smb *SelectMenuBuilder) MinValues(min int) *SelectMenuBuilder {
	smb.menu.MinValues = &min
	return smb
}

func (smb *SelectMenuBuilder) MaxValues(max int) *SelectMenuBuilder {
	smb.menu.MaxValues = max
	return smb
}

func (smb *SelectMenuBuilder) AddOption(label, value, description string) *SelectMenuBuilder {
	option := SelectMenuOption{
		Label:       label,
		Value:       value,
		Description: description,
	}
	smb.menu.Options = append(smb.menu.Options, option)
	return smb
}

func (smb *SelectMenuBuilder) AddOptionWithEmoji(label, value, description string, emoji ComponentEmoji) *SelectMenuBuilder {
	option := SelectMenuOption{
		Label:       label,
		Value:       value,
		Description: description,
		Emoji:       &emoji,
	}
	smb.menu.Options = append(smb.menu.Options, option)
	return smb
}

// Different select menu types
func (smb *SelectMenuBuilder) UserSelect() *SelectMenuBuilder {
	smb.menu.MenuType = UserSelectMenu
	return smb
}

func (smb *SelectMenuBuilder) RoleSelect() *SelectMenuBuilder {
	smb.menu.MenuType = RoleSelectMenu
	return smb
}

func (smb *SelectMenuBuilder) ChannelSelect(channelTypes ...ChannelType) *SelectMenuBuilder {
	smb.menu.MenuType = ChannelSelectMenu
	smb.menu.ChannelTypes = channelTypes
	return smb
}

// v2 enhancements
func (smb *SelectMenuBuilder) Searchable(searchable bool) *SelectMenuBuilder {
	smb.menu.Searchable = searchable
	return smb
}

func (smb *SelectMenuBuilder) Grouped(grouped bool) *SelectMenuBuilder {
	smb.menu.Grouped = grouped
	return smb
}

func (smb *SelectMenuBuilder) Build() SelectMenu {
	return smb.menu
}

// ===== TEXT INPUT BUILDER =====

type TextInputBuilder struct {
	input TextInput
}

func (tib *TextInputBuilder) Placeholder(text string) *TextInputBuilder {
	tib.input.Placeholder = text
	return tib
}

func (tib *TextInputBuilder) Value(text string) *TextInputBuilder {
	tib.input.Value = text
	return tib
}

func (tib *TextInputBuilder) Required(required bool) *TextInputBuilder {
	tib.input.Required = required
	return tib
}

func (tib *TextInputBuilder) Short() *TextInputBuilder {
	tib.input.Style = TextInputShort
	return tib
}

func (tib *TextInputBuilder) Paragraph() *TextInputBuilder {
	tib.input.Style = TextInputParagraph
	return tib
}

func (tib *TextInputBuilder) MinLength(length int) *TextInputBuilder {
	tib.input.MinLength = length
	return tib
}

func (tib *TextInputBuilder) MaxLength(length int) *TextInputBuilder {
	tib.input.MaxLength = length
	return tib
}

// v2 enhancements
func (tib *TextInputBuilder) Validation(pattern string) *TextInputBuilder {
	tib.input.ValidationPattern = pattern
	return tib
}

func (tib *TextInputBuilder) Masked(masked bool) *TextInputBuilder {
	tib.input.Masked = masked
	return tib
}

func (tib *TextInputBuilder) Build() TextInput {
	return tib.input
}

// ===== ACTIONS ROW BUILDER =====

type ActionsRowBuilder struct {
	row ActionsRow
}

func (arb *ActionsRowBuilder) AddComponent(component MessageComponent) *ActionsRowBuilder {
	if len(arb.row.Components) < 5 {
		arb.row.Components = append(arb.row.Components, component)
	}
	return arb
}

func (arb *ActionsRowBuilder) AddButton(button Button) *ActionsRowBuilder {
	return arb.AddComponent(button)
}

func (arb *ActionsRowBuilder) AddSelectMenu(menu SelectMenu) *ActionsRowBuilder {
	return arb.AddComponent(menu)
}

func (arb *ActionsRowBuilder) Build() ActionsRow {
	return arb.row
}

// ===== v2 MODAL BUILDER =====

type ModalBuilder struct {
	modal Modal
}

func (mb *ModalBuilder) AddComponent(component MessageComponent) *ModalBuilder {
	mb.modal.Components = append(mb.modal.Components, component)
	return mb
}

func (mb *ModalBuilder) AddTextInput(input TextInput) *ModalBuilder {
	return mb.AddComponent(input)
}

func (mb *ModalBuilder) Size(size ModalSize) *ModalBuilder {
	mb.modal.Size = size
	return mb
}

func (mb *ModalBuilder) Closable(closable bool) *ModalBuilder {
	mb.modal.Closable = closable
	return mb
}

func (mb *ModalBuilder) Build() Modal {
	return mb.modal
}

// ===== v2 TABS BUILDER =====

type TabsBuilder struct {
	tabs Tabs
}

func (tb *TabsBuilder) AddTab(id, label string, content MessageComponent) *TabsBuilder {
	tab := Tab{
		ID:      id,
		Label:   label,
		Content: content,
	}
	tb.tabs.TabList = append(tb.tabs.TabList, tab)
	return tb
}

func (tb *TabsBuilder) DefaultTab(id string) *TabsBuilder {
	tb.tabs.DefaultTab = id
	return tb
}

func (tb *TabsBuilder) Build() Tabs {
	return tb.tabs
}

// ===== QUICK HELPERS =====

// Create multiple buttons in one row
func QuickButtons(buttons ...Button) ActionsRow {
	components := make([]MessageComponent, len(buttons))
	for i, btn := range buttons {
		components[i] = btn
	}
	return ActionsRow{Components: components}
}

// Create a simple button
func QuickButton(label, customID string, style ButtonStyle) Button {
	return Button{
		Label:    label,
		CustomID: customID,
		Style:    style,
	}
}

// Create a select menu with options
func QuickSelectMenu(customID, placeholder string, options ...SelectMenuOption) SelectMenu {
	return SelectMenu{
		CustomID:    customID,
		Placeholder: placeholder,
		Options:     options,
		MaxValues:   1,
	}
}

// Create a select menu option
func QuickOption(label, value, description string) SelectMenuOption {
	return SelectMenuOption{
		Label:       label,
		Value:       value,
		Description: description,
	}
}

// Create a confirmation dialog with Yes/No buttons
func QuickConfirmDialog(customID string) ActionsRow {
	return QuickButtons(
		QuickButton("Yes", customID+"_yes", SuccessButton),
		QuickButton("No", customID+"_no", DangerButton),
	)
}

// Create pagination buttons
func QuickPagination(customID string, currentPage, totalPages int) ActionsRow {
	buttons := []Button{
		QuickButton("⏮️", customID+"_first", SecondaryButton),
		QuickButton("◀️", customID+"_prev", SecondaryButton),
		QuickButton(fmt.Sprintf("%d/%d", currentPage, totalPages), customID+"_current", SecondaryButton),
		QuickButton("▶️", customID+"_next", SecondaryButton),
		QuickButton("⏭️", customID+"_last", SecondaryButton),
	}
	
	// Disable buttons based on current page
	if currentPage <= 1 {
		buttons[0].Disabled = true
		buttons[1].Disabled = true
	}
	if currentPage >= totalPages {
		buttons[3].Disabled = true
		buttons[4].Disabled = true
	}
	
	return QuickButtons(buttons...)
}

// ===== VALIDATION =====

func ValidateComponent(component MessageComponent) error {
	switch c := component.(type) {
	case ActionsRow:
		if len(c.Components) > 5 {
			return fmt.Errorf("actions row can have maximum 5 components")
		}
		if len(c.Components) == 0 {
			return fmt.Errorf("actions row must have at least 1 component")
		}
	case Button:
		if c.Label == "" && c.Emoji == nil {
			return fmt.Errorf("button must have either label or emoji")
		}
		if c.Style == LinkButton && c.URL == "" {
			return fmt.Errorf("link button must have URL")
		}
		if c.Style != LinkButton && c.CustomID == "" {
			return fmt.Errorf("non-link button must have custom ID")
		}
	case SelectMenu:
		if c.CustomID == "" {
			return fmt.Errorf("select menu must have custom ID")
		}
		if c.MenuType == StringSelectMenu && len(c.Options) == 0 {
			return fmt.Errorf("string select menu must have options")
		}
	case TextInput:
		if c.CustomID == "" {
			return fmt.Errorf("text input must have custom ID")
		}
		if c.Label == "" {
			return fmt.Errorf("text input must have label")
		}
	case Modal:
		if c.CustomID == "" {
			return fmt.Errorf("modal must have custom ID")
		}
		if c.Title == "" {
			return fmt.Errorf("modal must have title")
		}
	}
	return nil
}

// ===== COMPONENT STRUCTS =====

type unmarshalableMessageComponent struct {
	MessageComponent
}

func (umc *unmarshalableMessageComponent) UnmarshalJSON(src []byte) error {
	var v struct {
		Type ComponentType `json:"type"`
	}
	err := json.Unmarshal(src, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case ActionsRowComponent:
		umc.MessageComponent = &ActionsRow{}
	case ButtonComponent:
		umc.MessageComponent = &Button{}
	case SelectMenuComponent, ChannelSelectMenuComponent, UserSelectMenuComponent,
		RoleSelectMenuComponent, MentionableSelectMenuComponent:
		umc.MessageComponent = &SelectMenu{}
	case TextInputComponent:
		umc.MessageComponent = &TextInput{}
	case SectionComponent:
		umc.MessageComponent = &Section{}
	case TextDisplayComponent:
		umc.MessageComponent = &TextDisplay{}
	case ThumbnailComponent:
		umc.MessageComponent = &Thumbnail{}
	case MediaGalleryComponent:
		umc.MessageComponent = &MediaGallery{}
	case FileComponentType:
		umc.MessageComponent = &FileComponent{}
	case SeparatorComponent:
		umc.MessageComponent = &Separator{}
	case ContainerComponent:
		umc.MessageComponent = &Container{}
	case ModalComponent:
		umc.MessageComponent = &Modal{}
	case TabsComponent:
		umc.MessageComponent = &Tabs{}
	case AccordionComponent:
		umc.MessageComponent = &Accordion{}
	default:
		return fmt.Errorf("unknown component type: %d", v.Type)
	}
	return json.Unmarshal(src, umc.MessageComponent)
}

func MessageComponentFromJSON(b []byte) (MessageComponent, error) {
	var u unmarshalableMessageComponent
	err := u.UnmarshalJSON(b)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal into MessageComponent: %w", err)
	}
	return u.MessageComponent, nil
}

// Container for other components
type ActionsRow struct {
	Components []MessageComponent `json:"components"`
	ID         int                `json:"id,omitempty"`
}

func (r ActionsRow) MarshalJSON() ([]byte, error) {
	type actionsRow ActionsRow
	return json.Marshal(struct {
		actionsRow
		Type ComponentType `json:"type"`
	}{
		actionsRow: actionsRow(r),
		Type:       r.Type(),
	})
}

func (r *ActionsRow) UnmarshalJSON(data []byte) error {
	type actionsRow ActionsRow
	var v struct {
		actionsRow
		RawComponents []unmarshalableMessageComponent `json:"components"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*r = ActionsRow(v.actionsRow)

	r.Components = make([]MessageComponent, len(v.RawComponents))
	for i, v := range v.RawComponents {
		r.Components[i] = v.MessageComponent
	}

	return err
}

func (r ActionsRow) Type() ComponentType {
	return ActionsRowComponent
}

// Button styles
type ButtonStyle uint

const (
	PrimaryButton   ButtonStyle = 1
	SecondaryButton ButtonStyle = 2
	SuccessButton   ButtonStyle = 3
	DangerButton    ButtonStyle = 4
	LinkButton      ButtonStyle = 5
	PremiumButton   ButtonStyle = 6
)

// v2 button sizes
type ButtonSize string

const (
	ButtonSizeSmall  ButtonSize = "small"
	ButtonSizeMedium ButtonSize = "medium"
	ButtonSizeLarge  ButtonSize = "large"
)

type ComponentEmoji struct {
	Name     string `json:"name,omitempty"`
	ID       string `json:"id,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}

type Button struct {
	Label    string          `json:"label"`
	Style    ButtonStyle     `json:"style"`
	Disabled bool            `json:"disabled"`
	Emoji    *ComponentEmoji `json:"emoji,omitempty"`
	URL      string          `json:"url,omitempty"`
	CustomID string          `json:"custom_id,omitempty"`
	SKUID    string          `json:"sku_id,omitempty"`
	ID       int             `json:"id,omitempty"`
	
	// v2 additions
	Tooltip string      `json:"tooltip,omitempty"`
	Badge   *int        `json:"badge,omitempty"`
	Loading bool        `json:"loading,omitempty"`
	Size    ButtonSize  `json:"size,omitempty"`
}

func (b Button) MarshalJSON() ([]byte, error) {
	type button Button
	if b.Style == 0 {
		b.Style = PrimaryButton
	}
	return json.Marshal(struct {
		button
		Type ComponentType `json:"type"`
	}{
		button: button(b),
		Type:   b.Type(),
	})
}

func (Button) Type() ComponentType {
	return ButtonComponent
}

type SelectMenuOption struct {
	Label       string          `json:"label,omitempty"`
	Value       string          `json:"value"`
	Description string          `json:"description"`
	Emoji       *ComponentEmoji `json:"emoji,omitempty"`
	Default     bool            `json:"default"`
}

type SelectMenuDefaultValueType string

const (
	SelectMenuDefaultValueUser    SelectMenuDefaultValueType = "user"
	SelectMenuDefaultValueRole    SelectMenuDefaultValueType = "role"
	SelectMenuDefaultValueChannel SelectMenuDefaultValueType = "channel"
)

type SelectMenuDefaultValue struct {
	ID   string                     `json:"id"`
	Type SelectMenuDefaultValueType `json:"type"`
}

type SelectMenuType ComponentType

const (
	StringSelectMenu      = SelectMenuType(SelectMenuComponent)
	UserSelectMenu        = SelectMenuType(UserSelectMenuComponent)
	RoleSelectMenu        = SelectMenuType(RoleSelectMenuComponent)
	MentionableSelectMenu = SelectMenuType(MentionableSelectMenuComponent)
	ChannelSelectMenu     = SelectMenuType(ChannelSelectMenuComponent)
)

type SelectMenu struct {
	MenuType      SelectMenuType           `json:"type,omitempty"`
	CustomID      string                   `json:"custom_id,omitempty"`
	Placeholder   string                   `json:"placeholder"`
	MinValues     *int                     `json:"min_values,omitempty"`
	MaxValues     int                      `json:"max_values,omitempty"`
	DefaultValues []SelectMenuDefaultValue `json:"default_values,omitempty"`
	Options       []SelectMenuOption       `json:"options,omitempty"`
	Disabled      bool                     `json:"disabled"`
	ChannelTypes  []ChannelType            `json:"channel_types,omitempty"`
	ID            int                      `json:"id,omitempty"`
	
	// v2 additions
	Searchable bool `json:"searchable,omitempty"`
	Grouped    bool `json:"grouped,omitempty"`
}

func (s SelectMenu) Type() ComponentType {
	if s.MenuType != 0 {
		return ComponentType(s.MenuType)
	}
	return SelectMenuComponent
}

func (s SelectMenu) MarshalJSON() ([]byte, error) {
	type selectMenu SelectMenu
	return json.Marshal(struct {
		selectMenu
		Type ComponentType `json:"type"`
	}{
		selectMenu: selectMenu(s),
		Type:       s.Type(),
	})
}

type TextInputStyle uint

const (
	TextInputShort     TextInputStyle = 1
	TextInputParagraph TextInputStyle = 2
)

type TextInput struct {
	CustomID    string         `json:"custom_id"`
	Label       string         `json:"label"`
	Style       TextInputStyle `json:"style"`
	Placeholder string         `json:"placeholder,omitempty"`
	Value       string         `json:"value,omitempty"`
	Required    bool           `json:"required"`
	MinLength   int            `json:"min_length,omitempty"`
	MaxLength   int            `json:"max_length,omitempty"`
	ID          int            `json:"id,omitempty"`
	
	// v2 additions
	ValidationPattern string `json:"validation_pattern,omitempty"`
	Masked           bool   `json:"masked,omitempty"`
}

func (TextInput) Type() ComponentType {
	return TextInputComponent
}

func (m TextInput) MarshalJSON() ([]byte, error) {
	type inputText TextInput
	return json.Marshal(struct {
		inputText
		Type ComponentType `json:"type"`
	}{
		inputText: inputText(m),
		Type:      m.Type(),
	})
}

// ===== v2 COMPONENTS =====

type ModalSize string

const (
	ModalSizeSmall  ModalSize = "small"
	ModalSizeMedium ModalSize = "medium"
	ModalSizeLarge  ModalSize = "large"
)

type Modal struct {
	CustomID   string             `json:"custom_id"`
	Title      string             `json:"title"`
	Components []MessageComponent `json:"components"`
	Size       ModalSize          `json:"size,omitempty"`
	Closable   bool               `json:"closable,omitempty"`
}

func (Modal) Type() ComponentType { return ModalComponent }

func (m Modal) MarshalJSON() ([]byte, error) {
	type modal Modal
	return json.Marshal(struct {
		modal
		Type ComponentType `json:"type"`
	}{
		modal: modal(m),
		Type:  m.Type(),
	})
}

type Tab struct {
	ID      string           `json:"id"`
	Label   string           `json:"label"`
	Content MessageComponent `json:"content"`
	Badge   *int             `json:"badge,omitempty"`
	Icon    *ComponentEmoji  `json:"icon,omitempty"`
}

type Tabs struct {
	CustomID   string `json:"custom_id"`
	TabList    []Tab  `json:"tabs"`
	DefaultTab string `json:"default_tab,omitempty"`
}

func (Tabs) Type() ComponentType { return TabsComponent }

func (t Tabs) MarshalJSON() ([]byte, error) {
	type tabs Tabs
	return json.Marshal(struct {
		tabs
		Type ComponentType `json:"type"`
	}{
		tabs: tabs(t),
		Type: t.Type(),
	})
}

type AccordionItem struct {
	ID      string           `json:"id"`
	Title   string           `json:"title"`
	Content MessageComponent `json:"content"`
	Open    bool             `json:"open,omitempty"`
}

type Accordion struct {
	CustomID string          `json:"custom_id"`
	Items    []AccordionItem `json:"items"`
	Multiple bool            `json:"multiple,omitempty"` // Allow multiple items open
}

func (Accordion) Type() ComponentType { return AccordionComponent }

func (a Accordion) MarshalJSON() ([]byte, error) {
	type accordion Accordion
	return json.Marshal(struct {
		accordion
		Type ComponentType `json:"type"`
	}{
		accordion: accordion(a),
		Type:      a.Type(),
	})
}

// ===== PLACEHOLDER TYPES =====

type ChannelType int
type Section struct{}
type TextDisplay struct{}
type Thumbnail struct{}
type MediaGallery struct{}
type FileComponent struct{}
type Separator struct{}
type Container struct{}

func (Section) Type() ComponentType       { return SectionComponent }
func (TextDisplay) Type() ComponentType   { return TextDisplayComponent }
func (Thumbnail) Type() ComponentType     { return ThumbnailComponent }
func (MediaGallery) Type() ComponentType  { return MediaGalleryComponent }
func (FileComponent) Type() ComponentType { return FileComponentType }
func (Separator) Type() ComponentType     { return SeparatorComponent }
func (Container) Type() ComponentType     { return ContainerComponent }

func (s Section) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
	}{Type: s.Type()})
}

func (td TextDisplay) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
	}{Type: td.Type()})
}

func (t Thumbnail) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
	}{Type: t.Type()})
}

func (mg MediaGallery) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
	}{Type: mg.Type()})
}

func (fc FileComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
	}{Type: fc.Type()})
}

func (s Separator) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
	}{Type: s.Type()})
}

func (c Container) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
	}{Type: c.Type()})
}
