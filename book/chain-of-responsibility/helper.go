package helper

type HelpTopic uint8

const (
	NO_HELP_TOPIC HelpTopic = iota
	TOPIC_APPLICATION
	TOPIC_DIALOG
	TOPIC_BUTTON
)

var helpStack = []string{
	"",                 // NO_HELP_TOPIC
	"Application Help", // TOPIC_APPLICATION
	"Dialog Help",      // TOPIC_DIALOG
	"Button Help",      // TOPIC_BUTTON
}

type HelpHandler interface {
	HandleHelp() string
	HasHelp() bool
}

type Help struct {
	topic HelpTopic
}

func (h Help) HandleHelp() string {
	return helpStack[h.topic]
}

func (h Help) HasHelp() bool {
	return h.topic != NO_HELP_TOPIC
}

type Application struct {
	Help
}

func NewApplication(topic HelpTopic) HelpHandler {
	return &Application{Help{topic}}
}

type Widget struct {
	help      Help
	successor HelpHandler
}

func (w *Widget) HandleHelp() string {
	if w.help.HasHelp() {
		return w.help.HandleHelp()
	}
	return w.successor.HandleHelp()
}

func (w *Widget) HasHelp() bool {
	return w.help.HasHelp()
}

type Dialog struct {
	Widget
}

func NewDialog(h HelpHandler, t HelpTopic) HelpHandler {
	return &Dialog{Widget{Help{t}, h}}
}

type Button struct {
	Widget
}

func NewButton(h HelpHandler, t HelpTopic) HelpHandler {
	return &Button{Widget{Help{t}, h}}
}
