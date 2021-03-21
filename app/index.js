function getScrap() {
    let input = document.getElementById('scrapInput').value
    if (input == '') {
        alert('Input is empty, please enter anything.')
    } else {
        fetch('/scrap', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(result => alert(result.status))
    }
}