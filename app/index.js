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
            let list = document.getElementById('listResults')

            for (let i = 0; i < json.length; i++) {
                const projectEl = document.createElement('div')
                projectEl.innerHTML = json[i].title
                projectEl.innerHTML = json[i].url
                list.appendChild(projectEl)
            }
        })
    }
}