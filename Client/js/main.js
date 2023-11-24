import { setUserId, setUserToken } from './local.js';
import { OK, USER_ALREADY_EXISTS, USER_NOT_FOUND, WRONG_PASSWORD } from './server/requestsUtil.js';
import { SignInRequest, SignUpRequest, signIn, signUp } from './server/toServer.js';

const nameInput = document.getElementById('nameInput');
const loginInput = document.getElementById('loginInput');
const passwordInput = document.getElementById('passwordInput');
const passwordAgainInput = document.getElementById('passwordAgainInput');
const signInLoginInput = document.getElementById('signInLoginInput');
const signInPasswordInput = document.getElementById('signInPasswordInput');

function setStuff(result) {
    setUserId(result.id);
    setUserToken(result.token);

    window.location = '../html/user.html';
}

function signUpSend() {
    const name = nameInput.value;
    const login = loginInput.value;
    const password = passwordInput.value;
    const passwordAgain = passwordAgainInput.value;

    if (password != passwordAgain) {
        alert('Passwords do not match');
        return;
    }

    signUp.send(new SignUpRequest(name, login, password),
        (status, result) => {
            switch (status) {
                case OK:
                    setStuff(result);
                    break;
                case USER_ALREADY_EXISTS:
                    alert('User with this login already exists');
                    break;
            }
        });
}

function signInSend() {
    const login = signInLoginInput.value;
    const password = signInPasswordInput.value;

    signIn.send(new SignInRequest(login, password),
        (status, result) => {
            switch (status) {
                case OK:
                    setStuff(result);
                    break;
                case WRONG_PASSWORD:
                    alert('Wrong password');
                    break;
                case USER_NOT_FOUND:
                    alert('User with this login does not exist');
                    break;
            }
        });
}

nameInput.value = '';
loginInput.value = '';
passwordInput.value = '';
passwordAgainInput.value = '';
signInLoginInput.value = '';
signInPasswordInput.value = '';

document.getElementById('signUpSend').addEventListener('click', signUpSend);
document.getElementById('signInSend').addEventListener('click', signInSend);
