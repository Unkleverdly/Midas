import { hasUserData } from './local.js';
import { OK, USER_NOT_FOUND } from './server/requestsUtil.js';
import {
    MainDataRequest, NewCategoryRequset, TransactionRequset,
    addNewCategory, getMainData, makeTransaction
} from './server/toServer.js';

if (!hasUserData()) {
    window.location = `../html/main.html`;
}

const newCategoryNameInput = document.getElementById('newCategoryName');
const newCategoryLimitInput = document.getElementById('newCategoryLimit');
const newSavingNameInput = document.getElementById('newSavingName');
const newSavingGoalInput = document.getElementById('newSavingGoal');
const categoriesList = document.getElementById('categories');
const transactionList = document.getElementById('transactions');
const savingsList = document.getElementById('savings');
const newTransactionCategoryDropdown = document.getElementById('newTransactionCategory');
const newTransactionAmountInput = document.getElementById('newTransactionAmount');
const newTransactionCategoryToggle = document.getElementById('newTransactionCategoryToggle');

let categoriesMap = new Map();
let allCategories = [];
let allTransactions = [];

let chosenTransactionCategory = null;
let useCategoryForNewTransaction = true;
let editedCategory = null;
let editedSaving = null;

function addCategory(cat) {
    const id = cat.id;
    const name = cat.name;
    const isCategory = cat.limit >= 0;

    const list = isCategory ? categoriesList : savingsList;

    const preText = isCategory ? 'Spent' : 'Saved up';

    let amountText = `${preText} in this month: ${Math.abs(cat.amount)}`;
    if (cat.limit != 0) {
        amountText += ` out of ${Math.abs(cat.limit)}`;
    }

    list.insertAdjacentHTML('beforeend',
        `
    <li>
        <div class="uk-accordion-title text">
            <div>${name}</div>
            <button class="uk-badge button bg" id="editCat${id}">Edit</button>
        </div>
        <div class="uk-accordion-content">
            <div class="text">${amountText}</div>
        </div>
    </li>`);

    // document.getElementById(`editCat${id}`).e;

    if (isCategory != useCategoryForNewTransaction) return;
    const dropdown = newTransactionCategoryDropdown;
    const buttonId = `choose${isCategory ? 'Category' : 'Saving'}${id}`;
    dropdown.insertAdjacentHTML('beforeend',
        `<li><button id="${buttonId}" class="category-button">${name}</button></li>`);

    const button = document.getElementById(buttonId);
    button.addEventListener('click', () => {
        chosenTransactionCategory = cat;
        newTransactionCategoryToggle.textContent = name;
        UIkit.toggle(newTransactionCategoryToggle).toggle();
    });
}

function removeCategory(id) {
    const ind = allCategories.findIndex(e => e.id == id)
    allCategories.splice(ind, 1);
    initCategories();
}

function initCategories() {
    categoriesList.innerHTML = '';
    savingsList.innerHTML = '';
    newTransactionCategoryDropdown.innerHTML = '';
    for (let i = 0; i < allCategories.length; i++) {
        const cat = allCategories[i];
        categoriesMap.set(cat.id, cat);
        addCategory(cat);
    }
}

function addTransaction(t) {
    const time = new Date(t.time * 1000);

    const mm = time.getMonth() + 1;
    const dd = time.getDate();
    const hh = time.getHours();
    const min = time.getMinutes();

    const strDate = [
        (dd > 9 ? '' : '0') + dd, '.',
        (mm > 9 ? '' : '0') + mm, '.',
        time.getFullYear(), ' ',
        (hh > 9 ? '' : '0') + hh, ':',
        (min > 9 ? '' : '0') + min,
    ].join('');

    const cat = categoriesMap.get(t.categoryId);

    transactionList.insertAdjacentHTML('beforeend',
        `
    <li class="no-bg">
        <div class="text">${cat.limit >= 0 ? 'Spent on' : 'Saved up for'} ${cat.name}: ${cat.amount} (${strDate})</div>
    </li>`);
}

function initTransactions() {
    transactionList.innerHTML = '';
    for (let i = allTransactions.length - 1; i >= 0; i--) {
        const t = allTransactions[i];
        addTransaction(t);
    }
}

function loadMainData() {
    const date = new Date();
    const monthStart = new Date(date.getUTCFullYear(), date.getUTCMonth(), 1);

    const unixNow = Math.floor(date.getTime() / 1000);
    const unixMonthStart = Math.floor(monthStart.getTime() / 1000);

    getMainData.send(new MainDataRequest(unixMonthStart, unixNow), (status, result) => {
        if (status == USER_NOT_FOUND) {
            window.location = '../html/main.html';
            return;
        }

        document.getElementById('name').textContent = result.name;
        document.getElementById('monthSpendings').textContent = result.monthSpendings;

        allCategories = result.categories;
        initCategories();
        allTransactions = result.transactions;
        initTransactions();
    });
}

function makeNewTransaction() {
    const amount = Number(newTransactionAmountInput.value);
    const category = chosenTransactionCategory;

    if (category == null) {
        alert(`Set ${useCategoryForNewTransaction ? 'category' : 'saving'} first`);
        return;
    }

    if (amount == Number.NaN) {
        alert('Set number in \'amount\'');
        return;
    }

    makeTransaction.send(new TransactionRequset(category.id, amount), (status, result) => {
        if (status == OK) {
            const t = {
                time: Math.floor((new Date()).getTime() / 1000),
                categoryId: category.id,
                amount: amount
            };

            allTransactions.push(t);
            initTransactions();
        }
    });
}

function newCategory() {
    const name = newCategoryNameInput.value;
    let limit = Number(newCategoryLimitInput.value);

    if (limit == Number.NaN) {
        alert('Limit must be a number');
        return;
    }

    if (limit <= 0) {
        alert('Limit must be greater than zero');
        return;
    }

    let id = -1;
    if (editedCategory != null) {
        id = editedCategory.id;
    }

    addNewCategory.send(new NewCategoryRequset(name, limit, id), (status, result) => {
        if (status == OK) {
            if (id != -1) removeCategory(id);
            const cat = {
                id: result,
                name: name,
                limit: limit,
                amount: 0
            };
            addCategory(cat);
            return;
        }
        alert('unknown error; i dont know what exactly went wrong');
    });
}

function newSaving() {
    const name = newSavingNameInput.value;
    let limit = Number(newSavingGoalInput.value);

    if (limit == Number.NaN) {
        alert('Goal must be a number');
        return;
    }

    if (limit <= 0) {
        alert('Goal must be greater than zero');
        return;
    }

    let id = -1;
    if (editedSaving != null) {
        id = editedSaving.id;
    }

    addNewCategory.send(new NewCategoryRequset(name, -limit, id), (status, result) => {
        if (status == OK) {
            if (id != -1) removeCategory(id);
            const cat = {
                id: result,
                name: name,
                limit: -limit,
                amount: 0
            };
            addCategory(cat);
            return;
        }
        alert('unknown error; i dont know what exactly went wrong');
    });
}

newCategoryNameInput.value = '';
newCategoryLimitInput.value = '';
newTransactionAmountInput.value = '';

loadMainData();

const cat = {
    id: 123,
    name: 'test',
    limit: -456,
    amount: 45
};
addCategory(cat);

document.getElementById('addNewCategorySend').addEventListener('click', newCategory);
document.getElementById('addNewSavingSend').addEventListener('click', newSaving);
document.getElementById('newTransactionSend').addEventListener('click', makeNewTransaction);
document.getElementById('newTransactionSwitch').addEventListener('click', () => {
    useCategoryForNewTransaction = !useCategoryForNewTransaction;
    const trswitch = document.getElementById('newTransactionSwitch');
    trswitch.textContent = useCategoryForNewTransaction ? 'Switch to savings' : 'Switch to categories';
    document.getElementById('newTransactionText').textContent = useCategoryForNewTransaction
        ? 'Category' : 'Saving';
    initCategories();
});

// document.getElementById('testSend').addEventListener('click', test);

