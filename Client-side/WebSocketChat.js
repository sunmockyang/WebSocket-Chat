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
	var elem = '<div id="message"><div id="username">' + username + '</div><div id="message-colon">:</div><div id="user-message">' + message + '</div></div>'

	this.chatWindow.innerHTML += elem;

	this.chatWindow.scrollTop = this.chatWindow.scrollHeight;
};
