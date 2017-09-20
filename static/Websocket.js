// Constants relating to direction
const UP = 0;
const DOWN = 1;
const LEFT = 2;
const RIGHT = 3;

/*
Object representing connection to server. All communication from and to server
will be managed by this class.
 */
class Websocket {

    // Connect to websocket server on same host address.
    // Takes callbacks required as parameters
    //      boardUpdateCallback: Callback that will receive the board update string. See hub.go's broadcastBoard
    constructor(boardUpdateCallback) {

        if (!window["WebSocket"]) {
            alert("Your browser does not support websockets!!");
            return
        }

        this.socket = new WebSocket("ws://" + document.location.host + "/ws");

        this.boardUpdateCallback = boardUpdateCallback;

        // Handle all events from server
        this.socket.onmessage = function (event) {
            const parsed = JSON.parse(event.data);

            // Differentiate between message types. See WebsocketMessage's constants
            if (parsed.type == 0) {
                // Board update message
                this.boardUpdateCallback(parsed.data);
            }
        }.bind(this)
    }

    // Notify the server that the user wants to change the snake's direction.
    // Direction is an integer with the possible values:
    //  0 = Up, 1 = Down, 2 = Left, 3 = Right
    sendChangeDirection(direction) {
        const payload = {
            type: 1,
            data: direction.toString()
        };

        this.socket.send(JSON.stringify(payload))
    }
}