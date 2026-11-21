package polyglot

const (
	English  string = "ENG"
	Francais string = "FRA"
	Chinese  string = "CHI"
	Russian  string = "RUS"

	Espanol string = "ESP"
	Japanes string = "JAP"
	Italian string = "ITA"
	Greek   string = "GRE"
)

type ITranslation interface {
	GetTranslation(language string) string
}

type Word struct {
	Description  string
	Translations map[string]string
}

func GetTranslation(word Word, language string) string {
	return word.Translations[language]
}

// Местоимение, отвечающее на вопрос "Кто?"
type IPronounWho interface {
	Who(language string) string
}

type PronounWho struct {
	Word
}

var Pronouns map[string]PronounWho

func init() {
	Pronouns = make(map[string]PronounWho)
}

func init() {
	var pronounsWho = []string{"I", "you(thou)", "we", "you", "he", "she", "it", "they", "ils", "elles"}
	for _, val := range pronounsWho {
		Pronouns[val] = PronounWho{Translations: map[string]string{}}
		switch val {
		case "I":
			Pronouns[val].Translations[English] = val
			Pronouns[val].Translations[Russian] = "я"
			Pronouns[val].Translations[Francais] = "je"
			Pronouns[val].Translations[Chinese] = "WŎ"

		//"ты"
		case "you(thou)":
			Pronouns[val].Translations[English] = val
			Pronouns[val].Translations[Russian] = "ты"
			Pronouns[val].Translations[Francais] = "tu"
			Pronouns[val].Translations[Chinese] = "NĬ"

		case "we":
			Pronouns[val].Translations[English] = val
			Pronouns[val].Translations[Russian] = "мы"
			Pronouns[val].Translations[Francais] = "nous"
			Pronouns[val].Translations[Chinese] = "WŎ MEN"

		//"вы"
		case "you":
			Pronouns[val].Translations[English] = val
			Pronouns[val].Translations[Russian] = "вы"
			Pronouns[val].Translations[Francais] = "vous"
			Pronouns[val].Translations[Chinese] = "NĬ MEN"

		case "he":
			Pronouns[val].Translations[English] = val
			Pronouns[val].Translations[Russian] = "он"
			Pronouns[val].Translations[Francais] = "il"
			Pronouns[val].Translations[Chinese] = "TĀ"

		case "she":
			Pronouns[val].Translations[English] = val
			Pronouns[val].Translations[Russian] = "она"
			Pronouns[val].Translations[Francais] = "elle"
			Pronouns[val].Translations[Chinese] = "TĀ"

		case "it":
			Pronouns[val].Translations[English] = val
			Pronouns[val].Translations[Russian] = "оно"
			Pronouns[val].Translations[Francais] = "il/ça"
			Pronouns[val].Translations[Chinese] = "TĀ"

		case "they":
			//Pronouns[val].Description = "они"
			Pronouns[val].Translations[English] = "they"
			Pronouns[val].Translations[Russian] = "они"
			Pronouns[val].Translations[Francais] = "ils"
			Pronouns[val].Translations[Chinese] = "TĀ MEN"

		case "ils":
			//Pronouns[val].Description = "они (М/М+Ж)"
			Pronouns[val].Translations[English] = "they"
			Pronouns[val].Translations[Russian] = "они"
			Pronouns[val].Translations[Francais] = "ils"
			Pronouns[val].Translations[Chinese] = "TĀ MEN"

		case "elles":
			//Pronouns[val].Description = "они (женщины)"
			Pronouns[val].Translations[English] = "they"
			Pronouns[val].Translations[Russian] = "они"
			Pronouns[val].Translations[Francais] = "ils"
			Pronouns[val].Translations[Chinese] = "TĀ MEN"
		}
	}

}
