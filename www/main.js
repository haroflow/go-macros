// Set-up code editor
var editor = ace.edit("editor");
// editor.setTheme("ace/theme/monokai");
editor.session.setMode("ace/mode/javascript");
// ---

setInterval(async () => {
	let pos = await getMousePosition();
	mousePosition.innerText = "mouse position: (" + pos.x + ", " + pos.y + ")";
}, 1000)

async function save() {
	error.innerText = "";
	try {
		var code = editor.getValue();
		await saveCode(code);
	} catch (ex) {
		error.innerText = ex;
	}
}

async function run() {
	error.innerText = "";
	try {
		var code = editor.getValue();
		await executeCode(code);
	} catch (ex) {
		error.innerText = ex;
	}
}

async function stop() {
	error.innerText = "";
	try {
		await stopMacros();
	} catch (ex) {
		error.innerText = ex;
	}
}

// Called from Go
function log(msg) {
	logWindow.innerHTML += msg + "<br>";
	logWindow.scrollTop = logWindow.scrollHeight;
}

// Called from Go
function logImage(img) {
	logWindow.innerHTML += "<img src='data:image/jpeg;base64," + img + "'><br>";
	logWindow.scrollTop = logWindow.scrollHeight;
}

function help() {
	var w = window.open("/help", "popup");
}

function getCode() {
	return editor.getValue();
}