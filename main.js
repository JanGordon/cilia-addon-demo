import showdown from "showdown"

var converter = new showdown.Converter()

function modifier(content) {
    return converter.makeHtml(content)
}