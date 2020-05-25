const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = function() {
  console.log("Websocket connected");
};

socket.onmessage = function (e) {
  const content = document.getElementById("content");
  content.innerHTML = e.data
};
