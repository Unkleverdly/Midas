import { GET, POST, RequestPath, ApiRequest, ServerResponseData, ServerRequest } from './requestsUtil.js';
import { getUserId, getUserToken, hasUserData } from '../local.js';

export class SignUpRequest extends ServerRequest {
    constructor(name, login, password) {
        super();
        this.name = name;
        this.login = login;
        this.password = password;
    }
}

export class SignInRequest extends ServerRequest {
    constructor(login, password) {
        super();
        this.login = login;
        this.password = password;
    }
}

export class UserDataRequest extends ServerRequest {
    constructor() {
        super();
        this.user =
        {
            id: getUserId(),
            token: getUserToken()
        };
        this.request = {};
    }
}

export class MainDataRequest extends UserDataRequest {
    constructor(timeStart, timeEnd) {
        super();
        this.request =
        {
            timeStart: timeStart,
            timeEnd: timeEnd
        };
    }
}

export class NewCategoryRequset extends UserDataRequest {
    constructor(name, limit, id) {
        super();
        const idToUse = id ?? -1;
        this.request =
        {
            name: name,
            limit: Number(limit),
            id: idToUse
        };
    }
}

export class TransactionRequset extends UserDataRequest {
    constructor(category, amount) {
        super();
        this.request =
        {
            id: category,
            amount: amount
        };
    }
}

export class DeleteCategoryRequset extends UserDataRequest {
    constructor(categoryId, amount) {
        super();
        this.request =
        {
            id: categoryId,
            amount: amount
        };
    }
}

const authPath = new RequestPath('auth');
const userPath = new RequestPath('user');

export const signIn = new ApiRequest(authPath.add('signIn'), POST);
export const signUp = new ApiRequest(authPath.add('signUp'), POST);

export const getMainData = new ApiRequest(userPath.add('getMainData'), POST);
export const getCategories = new ApiRequest(userPath.add('getCategories'), POST);
export const addNewCategory = new ApiRequest(userPath.add('addCategory'), POST);
export const deleteCategory = new ApiRequest(userPath.add('deleteCategory'), POST);
export const makeTransaction = new ApiRequest(userPath.add('makeTransaction'), POST);

