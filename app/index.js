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
        }).then(result => alert(result.status))
    }
}