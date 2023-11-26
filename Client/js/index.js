import { hasUserData } from "./local.js";

const page = hasUserData() ? 'user' : 'main';
window.location = `../html/${page}.html`;
