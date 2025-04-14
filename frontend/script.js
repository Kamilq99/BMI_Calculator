const backendUrl = "http://localhost:8080";

document.getElementById("calculate").addEventListener("click", async () => {
    const weight = parseFloat(document.getElementById("weight").value);
    const height = parseFloat(document.getElementById("height").value);

    if (!weight || !height) {
        alert("Please enter both weight and height.");
        return;
    }

    try {
        const dataResponse = await fetch(`${backendUrl}/data`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ weight, height }),
        });

        if (!dataResponse.ok) throw new Error("Failed to save data.");

        const bmiResponse = await fetch(`${backendUrl}/calculate`);
        const bmiData = await bmiResponse.json();

        document.getElementById("readonly").value = `Your BMI is ${bmiData.BMI.toFixed(2)}`;
    } catch (error) {
        alert("Error calculating BMI: " + error.message);
    }
});

document.getElementById("reset").addEventListener("click", () => {
    document.getElementById("weight").value = "";
    document.getElementById("height").value = "";
    document.getElementById("readonly").value = "";
});
