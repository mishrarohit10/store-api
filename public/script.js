function addBook() {
    location.assign("/addBooks")
}

function removeBook() {
    location.assign("/removeBooks")
}

function updateBook() {
    location.assign("/updateBooks")
}

function goToRaiseissue() {
    location.assign('/resolveIssue')
}

// function showDiv() {
//     document.getElementsByClassName('book').style.display = "";
// }

const button = document.querySelector('button[name="issues"]');

button.addEventListener('click', function () {
    button.disabled = true;
});


async function getIssue() {
    const container = document.querySelector('.container');

    try {
        const res = await fetch("/getIssues", {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json();
        console.log(data);
        if (!data) {
            return;
        }
        data.forEach(item => {
            console.log(item);

            const card = document.createElement('div');
            card.classList.add('card');

            const title = document.createElement('h2');
            title.textContent = item.id;

            const body = document.createElement('p');
            body.textContent = item.RequestDate;

            card.appendChild(title);
            card.appendChild(body);
            container.appendChild(card);
            // showDiv();
        });
    } catch (err) {
        console.log(err);
    }
}

async function addBooks() {
    const form = document.querySelector('form');
    console.log("addBooks");
    // get values
    const id = form.ISBN.value;
    const libID = form.libID.value;
    const title = form.title.value;
    const authors = form.authors.value;
    const publisher = form.publisher.value;
    const version = form.libID.value;
    const totalCopies = form.libID.value;
    const availabeCopies = form.libID.value;

    try {
        const res = await fetch('/addBooks', {
            method: 'POST',
            body: JSON.stringify({ id, libID, title, authors, publisher, version, totalCopies, availabeCopies, totalCopies }),
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json();
        console.log(data);
        // if (data) {
        //     location.assign('/success');
        // }
    }
    catch (err) {
        console.log(err);
    }
}

async function updateBooks() {
    const form = document.querySelector('form');
    console.log("addBooks");
    // get values
    const id = Number(form.ISBN.value);
    console.log("updateBooks");
    const libID = form.libID.value;
    const title = form.title.value;
    const authors = form.authors.value;
    const publisher = form.publisher.value;
    const version = form.libID.value;
    const totalCopies = form.libID.value;
    const availabeCopies = form.libID.value;

    var url = `/updateBook/${id}`
    console.log(url);
    try {
        console.log(typeof (id));
        const res = await fetch(url, {
            method: 'PUT',
            body: JSON.stringify({ id, libID, title, authors, publisher, version, totalCopies, availabeCopies, totalCopies }),
            headers: { 'Content-Type': 'application/json' }
        });
        const body = JSON.stringify({ id, libID, title, authors, publisher, version, totalCopies, availabeCopies, totalCopies })
        console.log(body, "this is body", res);
    }
    catch (err) {
        console.log(" err inside");
        console.log(err);
    }
}

async function deleteBook() {
    const form = document.querySelector('form');

    // get values
    const id = Number(form.id.value);

    var url = `/deleteBook/${id}`
    console.log(url);
    try {
        console.log(typeof (id));
        const res = await fetch(url, {
            method: 'DELETE',
            // body: JSON.stringify({ id, libID, title, authors, publisher, version, totalCopies, availabeCopies, totalCopies }),
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json();
        console.log(data);
    }
    catch (err) {
        console.log(" err inside");
        console.log(err);
    }
}

async function searchByTitle() {
    const form = document.querySelector('form');

    var button = document.getElementById('searchT')
    button.disabled = true;
    // get values
    const title = form.title.value;

    var url = `/searchByTitle/${title}`
    console.log(url);
    try {
        console.log(typeof (title));
        const res = await fetch(url, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json()
        console.log(data);
        if (!data) {
            return;
        }
        document.getElementById("h").innerHTML = data.Status;
        // document.getElementById("hh").innerHTML = data.AvailableDate;
    }
    catch (err) {
        console.log(" err inside");
        console.log(err);
    }
}

async function searchByAuthor() {
    const form = document.querySelector('form');
    const container = document.querySelector('.container');

    var button = document.getElementById('searchA')
    button.disabled = true;
    // get values
    const author = form.author.value;

    var url = `/searchByAuthor/${author}`
    console.log(url);
    try {
        console.log(typeof (author));
        const res = await fetch(url, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json()
        console.log(data);
        if (!data) {
            document.getElementById("h").innerHTML = 'No books available';
        }
        data.forEach(item => {
            console.log(item.Title);
            // document.getElementById("h").innerHTML = item.Title;

            const card = document.createElement('div');
            card.classList.add('card');

            const title = document.createElement('h2');
            title.textContent = item.Title;

            const body = document.createElement('p');
            body.textContent = item.Status;

            card.appendChild(title);
            card.appendChild(body);
            container.appendChild(card);

        });
        
    }
    catch (err) {
        console.log(" err inside");
        console.log(err);
    }
}

async function searchByPublisher() {
    const form = document.querySelector('form');
    const container = document.querySelector('.container');

    var button = document.getElementById('search')
    button.disabled = true;
    // get values
    const publisher = form.publisher.value;

    var url = `/searchByPublisher/${publisher}`
    console.log(url);
    try {
        console.log(typeof(publisher));
        const res = await fetch(url, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json()
        console.log(data, 'this is data --------------');

        if (!data) {
            const h1 = document.getElementById('h')
            h1.innerHTML = 'No books available';
        }
        data.forEach(item => {
            console.log(item.Title);
            // document.getElementById("h").innerHTML = item.Title;

            const card = document.createElement('div');
            card.classList.add('card');

            const title = document.createElement('h2');
            title.textContent = item.Title;

            const body = document.createElement('p');
            body.textContent = item.Status;

            card.appendChild(title);
            card.appendChild(body);
            container.appendChild(card);

        });
    }
    catch (err) {
        console.log(" err inside");
        console.log(err);
    }
}

async function raiseIssue() {
    const form = document.querySelector('form');

    console.log('inside raiseIssue')
    // get values
    const bookID = Number(form.bookID.value);
    const email = form.email.value;

    try {
        const res = await fetch('/raiseIssue', {
            method: 'POST',
            body: JSON.stringify({ bookID, email }),
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json();
        console.log(data);
        if (data) {
            const h2 = document.getElementById('h')
            h2.innerHTML = 'Issue Raised'
        }
    } catch (err) {
        console.log(err);
    }
}

async function owner() {
    const form = document.querySelector('form');

    var button = document.getElementById('ownerLib')
    button.disabled = true;
    console.log("addBooks");
    // get values
    const name = form.library.value;
    const username = form.username.value;
    const role = form.role.value;
    
    var url = '/createLib'
    console.log(url);
    try {
        // console.log(typeof (id));
        const res = await fetch(url, {
            method: 'POST',
            body: JSON.stringify({ name, username, role }),
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json();
        console.log(data.message, "this is message");
        console.log(data.error);
        if (data.message) {
            const h1 = document.getElementById('h')
            h1.innerHTML = data.message;
        } else {
            const h1 = document.getElementById('h')
            h1.innerHTML = data.error;
        }
    }
    catch (err) {
        console.log(" err inside");
        console.log(err);
    }
}

async function resolveIssue() {
    const form = document.querySelector('form');

    // get values
    const id = Number(form.id.value);

    var url = `/resolveIssue/${id}`
    console.log(url);
    try {
        console.log(typeof (id));
        const res = await fetch(url, {
            method: 'PUT',
            body: JSON.stringify({ id, status:"accepted" }),
            headers: { 'Content-Type': 'application/json' }
        });
        const data = await res.json();
        console.log(data);
        if(data.message) {
            document.getElementById('h').innerHTML=data.message;
        } else {
            document.getElementById('h').innerHTML=data.IssueStatus
            ;
        }
    }
    catch (err) {
        console.log(" err inside");
        console.log(err);
    }
}
