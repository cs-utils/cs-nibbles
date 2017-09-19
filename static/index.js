// WS test
var ws;

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

    ws = new WebSocket(new_uri);
    return ws

    // ws.onmessage = function (p1) {
    //     d = JSON.parse(p1.data);
    //     console.log(d);
    // }
}

function wsClose() {
    ws.close();
}