function calculate() {

    const form = document.getElementById("form");
    const formData = new FormData(form);
    const firstNumber = parseFloat(formData.get("firstNumber"));
    const secondNumber = parseFloat(formData.get("secondNumber"));

    const jsonData = JSON.stringify({
        firstNumber: firstNumber,
        operation: formData.get("operation"),
        secondNumber: secondNumber
    });

    fetch("/calculate", {
        method: "POST",
        body: jsonData
    })
        .then(response => response.json())
        .then(data => {
            document.getElementById("result").innerText = `Result: ${data.result}`;
        })
        .catch(error => {
            console.error("Error:", error);
            document.getElementById("result").innerText = "Error occurred during calculation.";
        });
}