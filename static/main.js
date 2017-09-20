let b = new Board("gameScreen", 500, 500, 5);

// b.drawFromSnakeStrings("15,0,0,0,1;17,50,50,50,51,50,52");

let s = "15";

demo();

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function demo() {
    // for (let i = 0; i < 500; i++) {
    //
    //     s += ",30," + i;
    //     b.drawFromSnakeStrings(s);
    //     await sleep(50);
    // }

    // Connect to websocket
    var websocket = wsTest();

    websocket.onmessage = function (event) {
        var msg = JSON.parse(event.data);
        b.drawFromSnakeStrings(msg.data);
    }
}
