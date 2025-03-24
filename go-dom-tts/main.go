package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

var (
	voices          []string
	voicesJs        = js.Value{}
	speechSynthesis = window.Get("speechSynthesis")
)

var (
	window   = js.Global().Get("window")
	document = js.Global().Get("document")
	body     = document.Get("body")
)

func main() {
	makeForm()

	setupVoices()
	<-make(chan (interface{}))
}

func makeForm() {

	form := document.Call("createElement", "form")
	body.Call("appendChild", form)

	voicesLabel := document.Call("createElement", "label")
	voicesLabel.Set("textContent", "voices:")
	form.Call("appendChild", voicesLabel)
	voicesSelector := document.Call("createElement", "select")
	voicesSelector.Set("id", "voices")
	voicesSelector.Set("name", "voices")
	form.Call("appendChild", voicesSelector)
	form.Call("appendChild", document.Call("createElement", "br"))

	// rate
	rateLabel := document.Call("createElement", "label")
	rateLabel.Set("textContent", "Rate:")
	form.Call("appendChild", rateLabel)

	rateInput := document.Call("createElement", "input")
	rateInput.Set("id", "rate")
	rateInput.Set("type", "range")
	rateInput.Set("name", "rate")
	rateInput.Set("min", "0.5")
	rateInput.Set("max", "2")
	rateInput.Set("step", "0.1")
	rateInput.Set("value", "1")
	form.Call("appendChild", rateInput)
	form.Call("appendChild", document.Call("createElement", "br"))

	// pitch
	pitchLabel := document.Call("createElement", "label")
	pitchLabel.Set("textContent", "Pitch:")
	form.Call("appendChild", pitchLabel)

	pitchInput := document.Call("createElement", "input")
	pitchInput.Set("id", "pitch")
	pitchInput.Set("type", "range")
	pitchInput.Set("name", "pitch")
	pitchInput.Set("min", "0")
	pitchInput.Set("max", "2")
	pitchInput.Set("step", "0.1")
	pitchInput.Set("value", "1")
	form.Call("appendChild", pitchInput)
	form.Call("appendChild", document.Call("createElement", "br"))

	// text to speak
	textLabel := document.Call("createElement", "label")
	textLabel.Set("textContent", "Text to Speak:")
	form.Call("appendChild", textLabel)

	textInput := document.Call("createElement", "input")
	textInput.Set("type", "text")
	textInput.Set("name", "text")
	textInput.Set("value", "Some text to speech")
	form.Call("appendChild", textInput)
	form.Call("appendChild", document.Call("createElement", "br"))

	// speak button
	handleSpeak := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		event.Call("preventDefault")
		text := form.Get("text").Get("value").String()
		rate, _ := strconv.ParseFloat(rateInput.Get("value").String(), 64)
		pitch, _ := strconv.ParseFloat(pitchInput.Get("value").String(), 64)
		selectedVoice := voicesSelector.Get("value").String()
		speak(selectedVoice, text, rate, pitch)
		return nil
	})

	speakButton := document.Call("createElement", "button")
	speakButton.Set("textContent", "Speak")
	speakButton.Set("onclick", handleSpeak)

	form.Call("appendChild", speakButton)
}

func updateVoiceSelector() {
	voicesSelector := document.Call("getElementById", "voices")

	for _, voiceName := range voices {
		voiceOption := document.Call("createElement", "option")
		voiceOption.Set("value", voiceName)
		voiceOption.Set("textContent", voiceName)
		voicesSelector.Call("appendChild", voiceOption)
	}
}

func setupVoices() {
	voicesChangedCallback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		voices = getVoices()
		updateVoiceSelector()
		return nil
	})

	js.Global().Get("speechSynthesis").Set("onvoiceschanged", voicesChangedCallback)
	voices = getVoices()
	updateVoiceSelector()
}

func getVoices() []string {
	voicesJs = speechSynthesis.Call("getVoices")
	voicesList := []string{}

	println(fmt.Sprintf("found %d voices", voicesJs.Length()))

	for i := 0; i < voicesJs.Length(); i++ {
		voice := voicesJs.Index(i)
		voicesList = append(voicesList, voice.Get("name").String())
	}
	return voicesList
}

func pickSelectedVoice(selectedVoiceName string) js.Value {
	for vIdx, voiceName := range voices {
		if voiceName == selectedVoiceName {
			return voicesJs.Index(vIdx)
		}
	}
	return voicesJs.Index(0)
}

func speak(selectedVoice, text string, rate, pitch float64) {
	speech := js.Global().Get("SpeechSynthesisUtterance").New(text)
	speech.Set("rate", rate)
	speech.Set("pitch", pitch)
	speech.Set("voice", pickSelectedVoice(selectedVoice))
	speechSynthesis.Call("speak", speech)
}
