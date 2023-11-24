import { hasUserData } from "./local";

const page = hasUserData ? 'user' : 'main';
window.location = `../html/${page}.html`;
