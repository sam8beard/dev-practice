const previewButton = document.getElementById("preview-button") 
const displayField = document.getElementById("displayTrip")

previewButton.addEventListener("click", function(event)  { 
    event.preventDefault()
    const pickupField = document.getElementById("pickup")
    const pickupValue = pickupField.value 

    const dropoffField = document.getElementById("dropoff")
    const dropoffValue = dropoffField.value

    displayField.innerText = "Pickup: " + pickupValue + "\n" + "Dropoff: " + dropoffValue
})  