function calculateBMI() {
    const weight = parseFloat(document.getElementById('weight').value);
    const height = parseFloat(document.getElementById('height').value);
    
    if (weight && height) {
        const bmi = (weight / (height * height)).toFixed(2);
        let classification = '';

        if (bmi < 18.5) {
            classification = 'You are underweight.';
        } else if (bmi >= 18.5 && bmi <= 24.9) {
            classification = 'Your weight is normal.';
        } else if (bmi >= 25 && bmi <= 29.9) {
            classification = 'You are overweight.';
        } else {
            classification = 'You are obese.';
        }

        document.getElementById('result').innerText = `Your BMI is: ${bmi}. ${classification}`;
    } else {
        document.getElementById('result').innerText = 'Please enter both weight and height.';
    }
}
