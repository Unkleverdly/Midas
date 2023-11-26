import { hasUserData } from './local.js';
import { OK, USER_NOT_FOUND } from './server/requestsUtil.js';
import { MainDataRequest, NewCategoryRequset, addNewCategory, getMainData } from './server/toServer.js';

if (!hasUserData()) {
    window.location = `../html/main.html`;
}

const newCategoryNameInput = document.getElementById('newCategoryName');
const newCategoryLimitInput = document.getElementById('newCategoryLimit');
const cats = document.getElementById('categories');

let allCategories = [];

function addCategory(cat) {
    cats.innerHTML +=
        `<li id="c_id_${cat.id}">
        <div class="uk-accordion-title text">${cat.name}</div>
        <div class="uk-accordion-content">
            <div class="text">${cat.amount}/${cat.limit}</div>
        </div>
    </li>`;
}

function removeCategory(cat) {
    const ind = allCategories.findIndex(e => e.id == cat.id)
    allCategories.splice(ind, 1);
    initCategories();
}

function initCategories() {
    cats.innerHTML = '';
    for (let i = 0; i < allCategories.length; i++) {
        const cat = allCategories[i];
        addCategory(cat);
    }
}

function loadMainData() {
    const date = new Date();
    const monthStart = new Date(date.getUTCFullYear(), date.getUTCMonth(), 1);

    const unixNow = date.getTime();
    const unixMonthStart = monthStart.getTime();

    getMainData.send(new MainDataRequest(unixNow, unixMonthStart), (status, result) => {
        if (status == USER_NOT_FOUND) {
            window.location = '../html/main.html';
            return;
        }

        document.getElementById('name').textContent = result.name;
        document.getElementById('monthSpendings').textContent = result.monthSpendings;

        allCategories = result.categories;
        initCategories();
    });
}

function newCategory() {
    const name = newCategoryNameInput.value;
    const limit = Number(newCategoryLimitInput.value);

    addNewCategory.send(new NewCategoryRequset(name, limit), (status, result) => {
        if (status == OK) {
            addCategory(result);
            return;
        }
        alert('unknown error; i dont know what exactly went wrong');
    });

}

newCategoryNameInput.value = '';
newCategoryLimitInput.value = '';

loadMainData();

document.getElementById('addNewCategorySend').addEventListener('click', newCategory);
// document.getElementById('testSend').addEventListener('click', test);

