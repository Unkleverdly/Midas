const userIdKey = 'userId';
const userTokenKey = 'userToken';

export function hasUserData() { return getUserToken() != null && getUserId() != null; }
export function deleteUserData() {
    localStorage.removeItem(userIdKey);
    localStorage.removeItem(userTokenKey);
}

export function getUserId() { return Number(localStorage.getItem(userIdKey)); }
export function getUserToken() { return localStorage.getItem(userTokenKey); }

export function setUserId(id) { localStorage.setItem(userIdKey, id); }
export function setUserToken(token) { localStorage.setItem(userTokenKey, token); }
