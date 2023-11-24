import { USER_NOT_FOUND } from './server/requestsUtil.js';
import { NewCategoryRequset, UserDataRequest, addCategory, getCategories, getMainData } from './server/toServer.js';

const newCategoryNameInput = document.getElementById('newCategoryName');
const newCategoryLimitInput = document.getElementById('newCategoryLimit');

getMainData.send(new UserDataRequest(), (status, result) => {
    if (status == USER_NOT_FOUND) {
        window.location = '../html/main.html';
        return;
    }

    document.getElementById('name').textContent = 'the sigma'; //result.name;
    document.getElementById('monthSpendings').textContent = '69 69 69'; //result.monthSpendings;
});

function reloadCategories() {
    getCategories.send(new UserDataRequest(), (status, result) => {
        const cats = document.getElementById('categories');
        const rcats = result;

        cats.innerHTML = '';
        for (let i = 0; i < rcats.length; i++) {
            const cat = rcats[i];
            cats.innerHTML += `
            <li id="c_id_${cat.id}">
            <div class="uk-accordion-title text">${cat.name}</div>
            </li>`;
        }
    });
}

function newCategory() {
    const name = newCategoryNameInput.value;
    const limit = Number(newCategoryLimitInput.value);

    addCategory.send(new NewCategoryRequset(name, limit), (status, result) => {
        if (status == OK) {
            reloadCategories();
            return;
        }
        alert('unknown error; i dont know what exactly went wrong');
    });

    alert('added');
}

newCategoryNameInput.value = '';
newCategoryLimitInput.value = '';

reloadCategories();
document.getElementById('addNewCategorySend').addEventListener('click', newCategory);
