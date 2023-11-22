import { GET, POST, RequestPath, ApiRequest, ServerResponse, ServerRequest } from './requestsUtil.js';

export class UserData extends ServerResponse {
    constructor() {
        super();
        this.id = '';
    }
}

export class SignUpRequest extends ServerRequest {
    constructor(name, login, password) {
        super();
        this.name = name;
        this.login = login;
        this.password = password;
    }
}

const authPath = new RequestPath('auth');

export const signIn = new ApiRequest(authPath.add('sign_in'), POST, UserData);
export const signUp = new ApiRequest(authPath.add('sign_up'), POST, UserData);
export const check = new ApiRequest(authPath.add('check'), GET, UserData);
