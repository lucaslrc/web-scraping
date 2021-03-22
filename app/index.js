function getScrap() {
    let input = document.getElementById('scrapInput').value
    if (input == '') {
        alert('Input is empty, please enter anything.')
    } else {
        let searchBody = { search: `${input}` }
        fetch('/scrap', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(searchBody)
        })
        .then(res => res.json())
        .then(json => {
            let table = document.getElementById('tableResults')
            removeAllChildNodes(table)

            for (let i = 0; i < json.length; i++) {
                let tr = document.createElement('tbody')
                let row = `<tr>
                                <td>${json[i].title}</td>
                            </tr>
                            <tr>
                                <th><a href="${json[i].url}" target="_blank">${json[i].url}</a></th>
                            </tr>`
                tr.innerHTML = row
                table.appendChild(tr)
            }
        })
    }
}

function removeAllChildNodes(parent) {
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild)
    }
}