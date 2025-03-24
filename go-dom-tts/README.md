# go-dom-tts

Example with HTML DOM API and text-to-speech API.

Build the Wasm files:

```bash
./build.sh
```

Start the HTTP server:

```bash
./server.sh
```

then open them in the corresponding HTML pages
e.g. [http://localhost:8080/index.html](http://localhost:8080/index_go.html),
and [http://localhost:8080/index_tinygo.html](http://localhost:8080/index_tinygo.html).

## References

Document Object Model (DOM):
- [The Core Document Object Model (DOM)](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model)
- [Introduction to the DOM](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model/Introduction)
- [The HTML DOM API](https://developer.mozilla.org/en-US/docs/Web/API/HTML_DOM_API)

Web Speech API:
- [Web Speech API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Speech_API)
- https://webaudio.github.io/web-speech-api/
- https://www.w3.org/TR/speech-synthesis/
- https://github.com/mdn/dom-examples/tree/main/web-speech-api

