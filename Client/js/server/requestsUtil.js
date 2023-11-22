export class RequestPath {
    constructor(path) {
        this.path = path;
    }
    add(path) { return new RequestPath(`${this.path}/${path}`); }
    eval() { return `${requestEndpoint}/${this.path}`; }
}

export class ApiRequest {
    /**
     * Creates an instance of ApiRequest.
     * @param {RequestPath} path
     * @param {string} method
     * @param {ServerResponse} constructor
     * @memberof ApiRequest
     */
    constructor(path, method, constructor) {
        this.path = path;
        this.method = method;
        this.ctor = constructor;
    }

    /**
     *
     *
     * @param {ServerRequest} data
     * @param {*} onResult
     * @memberof ApiRequest
     */
    send(data, onResult) {
        console.log('a');
        const requestData =
        {
            'method': this.method,
            'headers':
                { 'Content-Type': 'application/json' },
            'body': JSON.stringify(data)
        };

        fetch(this.path.eval(), requestData)
            .then(async resp => {
                const json = await resp.json();
                const response = new this.ctor();
                response.readFrom(json);
                onResult(response);
            });
    }
}

export class ServerResponse {
    constructor() { }
    readFrom(response) {
        Object.assign(this, response);
    }
}

export class ServerRequest { constructor() { } }

export const GET = 'GET';
export const POST = 'POST';

const defaultEndpoint = '192.168.0.17'; const port = '8080'; const protocol = 'http';
// const defaultEndpoint = '127.0.0.1'; const port = '5656'; const protocol = 'http';

const requestEndpoint = `${protocol}://${defaultEndpoint}:${port}`;
