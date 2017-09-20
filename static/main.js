let b = new Board("gameScreen", 500, 500, 5, 1);
let conn = new Websocket(redrawBoard);

document.addEventListener("keydown", keyDownHandler);
function keyDownHandler(e) {
    switch (e.keyCode) {
        case 38:
            conn.sendChangeDirection(UP);
            break;

        case 40:
            conn.sendChangeDirection(DOWN);
            break;

        case 37:
            conn.sendChangeDirection(LEFT);
            break;

        case 39:
            conn.sendChangeDirection(RIGHT);
            break;
    }
}

demo();

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function demo() {
}

async function redrawBoard(boardUpdateString) {
    b.drawFromSnakeStrings(boardUpdateString)
}