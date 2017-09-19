// WS test
function wsTest() {
    var loc = window.location, new_uri;
    if (loc.protocol === "https:") {
        new_uri = "wss:";
    } else {
        new_uri = "ws:";
    }
    new_uri += "//" + loc.host;
    new_uri += "/ws";

    console.log(new_uri);

    var ws = new WebSocket(new_uri);

    ws.onopen = function (event) {
        ws.send("sent message");
        ws.close();
    }
}