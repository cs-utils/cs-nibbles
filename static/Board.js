class Board {
    constructor(hostElementId, width, height, segmentSize, padding) {
        this.width = width;
        this.height = height;
        this.segmentSize = segmentSize; // Square
        this.padding = padding;

        this.hostElement = document.getElementById(hostElementId);

        if (this.hostElement == null) {
            console.error("Host element " + hostElementId + " Not found!");
        }
        this.hostElement.setAttribute("width", width);
        this.hostElement.setAttribute("height", height);

        this.ctx = this.hostElement.getContext("2d");
        this.ctx.fillStyle = 'red';
    }

    logicalToPixel(x, y) {
        let coords = [];

        coords[0] = x * this.segmentSize;
        coords[1] = y * this.segmentSize;

        coords[2] = this.segmentSize;
        coords[3] = this.segmentSize;

        // console.log(x + ", " + y);
        // console.log(coords);
        return coords;
    }


    drawFromSnakeStrings(snakeString) {
        this.ctx.fillStyle = 'black';
        this.ctx.fillRect(0, 0, this.width, this.height);
        this.ctx.fillStyle = 'red';

        let snakes = snakeString.split(";");

        // console.log(snakes);
        // For each snake
        for (let i = 0; i < snakes.length; i++) {
            let segs = snakes[i].split(",");

            // For each snake segment
            for (let j = 1; j < segs.length - 1; j += 2) {
                // 1, 2
                // 3, 4
                //
                let x = segs[j];
                let y = segs[j + 1];
                let coords = this.logicalToPixel(x, y);

                this.ctx.fillRect(coords[0] + this.padding, coords[1] + this.padding,
                    coords[2] - this.padding, coords[3] - this.padding);
            }
        }
    }


}
