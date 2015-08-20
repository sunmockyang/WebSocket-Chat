function WebSocketChat(username, chatWindow, url) {

	this.username = username;
	this.chatWindow = chatWindow;
	this.url = url;

	if (username == null || username == ""){
		this.displayMessage("CLIENT", "YOU NEED A USERNAME. REFRESH THE PAGE.");
		return;
	}

	this.displayMessage("CLIENT", "Connecting to: '" + this.url + "'");

	this.websocket = new WebSocket(this.url);
	this.websocket.onmessage = this.onReceive.bind(this);
	this.websocket.onopen = this.onOpen.bind(this);
	this.websocket.onclose = this.onClose.bind(this);
}

WebSocketChat.prototype.generateMessageTemplate = function(username, message) {
	var elem = document.createElement("div");
	elem.className = "message";

	var usernameElem = document.createElement("div");
	usernameElem.className = "username";
	usernameElem.style.color = generateRandomColourFromString(username);
	usernameElem.innerHTML = username;
	
	var colonElem = document.createElement("div");
	colonElem.className = "message-colon";
	colonElem.innerHTML = ":";
	
	var messageElem = document.createElement("div");
	messageElem.className = "user-message";
	messageElem.innerHTML = message;

	elem.appendChild(usernameElem);
	elem.appendChild(colonElem);
	elem.appendChild(messageElem);

	return elem;
};

WebSocketChat.prototype.onOpen = function(event) {
	this.displayMessage("SERVER", "Connection made.");
};

WebSocketChat.prototype.onClose = function(event) {
	this.displayMessage("SERVER", "Connection closed.")
};

WebSocketChat.prototype.onReceive = function(rawPacket) {
	console.log(rawPacket.data)
	var packet = JSON.parse(rawPacket.data);
	this.displayMessage(packet.username, packet.message);
};

WebSocketChat.prototype.sendMessage = function(message) {
	var packet = {
		username: this.username,
		message: message
	};

	this.displayMessage(this.username, message);

	this.websocket.send(JSON.stringify(packet));
};

WebSocketChat.prototype.displayMessage = function(username, message) {
	this.chatWindow.appendChild(this.generateMessageTemplate(username, message));
	this.chatWindow.scrollTop = this.chatWindow.scrollHeight;
};

function generateRandomColourFromString(str) {
	var len = str.length;
	var num = 1;

	for (var i = 0; i < len; i++) {
		num *= str.charCodeAt(i);
	};

	num %= 16777215;

	return "#" + num.toString(16);
}
