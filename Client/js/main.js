import { SignUpRequest, signIn, signUp } from './server/toServer.js';

function signUpClick() {
    console.log('ab');
    signUp.send(new SignUpRequest('yakov sigma', 'sigma_login', 'sigmapassword'),
        user => {
            console.log(user);
            document.getElementById('signUp').textContent = user.id;
        });
    console.log('boba');
}

document.getElementById('signUp').addEventListener('click', signUpClick);
