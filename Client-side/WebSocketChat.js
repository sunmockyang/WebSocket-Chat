function WebSocketChat(username, chatWindow) {
	this.username = username;
	this.chatWindow = chatWindow;
}

WebSocketChat.prototype.onReceive = function(rawPacket) {
	var packet = JSON.parse(rawPacket);
	this.displayMessage(packet.username, packet.message);
};

WebSocketChat.prototype.sendMessage = function(message) {
	var packet = {
		username: this.username,
		message: message
	};

	this.displayMessage(this.username, message);
};

WebSocketChat.prototype.displayMessage = function(username, message) {
	var elem = '<div id="message"><div id="username">' + username + '</div><div id="message-colon">:</div><div id="user-message">' + message + '</div></div>'

	this.chatWindow.innerHTML += elem;

	this.chatWindow.scrollTop = this.chatWindow.scrollHeight;
};
