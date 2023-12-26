package airportrobot

type (
	Italian    struct{}
	Portuguese struct{}
	Greeter    interface {
		LanguageName() string
		Greet(name string) string
	}
)

func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}
