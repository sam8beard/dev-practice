const quoteButton = document.getElementById("quote-button")
const displayContents = document.getElementById("response-contents")

quoteButton.addEventListener("click", function(event) { 
    event.preventDefault()
    const pickupField = document.getElementById("pickup")
    const dropoffField = document.getElementById("dropoff")

    const pickupValue = pickupField.value 
    const dropoffValue = dropoffField.value 

    fetch('https://jsonplaceholder.typicode.com/posts', { 
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({pickup: pickupValue, dropoff: dropoffValue})
    })
    .then(response => { 
        displayContents.innerText = JSON.stringify({price: "$42.0"})
        console.log(response)
    })
})