const fetch = require('node-fetch')

let searchBody = { search: 'maringa' }
        fetch('http://localhost:8080/scrap', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(searchBody)
        })
        .then(res => res.json())
        .then(json => console.log(json))