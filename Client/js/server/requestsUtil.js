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
     * @param {ServerResponseData} constructor
     * @memberof ApiRequest
     */
    constructor(path, method) {
        this.path = path;
        this.method = method;
    }

    /**
     *
     *
     * @param {ServerRequest} data
     * @param {*} onResult
     * @memberof ApiRequest
     */
    send(data, onResult, onError) {
        const requestData =
        {
            method: this.method,
            headers:
            {
                'Content-Type': 'application/json',
            },
        };

        if (this.method == POST) {
            requestData.body = JSON.stringify(data ?? {});
        }

        fetch(this.path.eval(), requestData)
            .then(async resp => {
                if (resp.headers.get('Content-Length') == 0 ||
                    !resp.headers.get('Content-Type').startsWith('application/json')) {
                    return;
                }

                const json = await resp.json();
                const result = json.result ?? {};

                onResult(json.status, result);
            }).catch(reason => {
                if (onError !== undefined) { onError(reason); }
            });
    }
}

export class ServerResponseData {
    constructor() { }
}

export class ServerRequest { constructor() { } }

export const GET = 'GET';
export const POST = 'POST';

const defaultEndpoint = '5.35.100.72'; const port = '8080'; const protocol = 'http';
// const defaultEndpoint = '192.168.0.14'; const port = '8080'; const protocol = 'http';

const requestEndpoint = `${protocol}://${defaultEndpoint}:${port}`;

export const NO_RESPONSE = -69;
export const OK = 0;
export const USER_NOT_FOUND = 1;
export const WRONG_PASSWORD = 2;
export const USER_ALREADY_EXISTS = 3;
// export const  = 0;

