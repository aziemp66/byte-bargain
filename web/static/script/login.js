const form = document.getElementById('form');

form.addEventListener('submit', async(e) => {
	e.preventDefault();

	email = document.getElementById('email').value;
	password = document.getElementById('password').value;

	const response = await fetch("/api/user/login", {
		method: "POST",
		body: JSON.stringify({
			email: email,
			password: password
		}),
		headers: {
			"Content-Type": "application/json"
		}
	})

	const responseData = await response.json();

	if (!response.ok){
		alert(responseData.message);
		return
	}

	localStorage.setItem('user', responseData['user_id']);
});