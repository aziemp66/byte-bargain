const form = document.getElementById('regist-form');

console.log("here");

form.addEventListener('submit', (e) => {
		e.preventDefault();
		const formData = new FormData(form);
		const data = Object.fromEntries(formData.entries());
		data['birth_date'] = data.year + '-' + data.month + '-' + data.day;
	data.year = "";
	data.month = "";
	data.day = "";
		console.log(data);
		const url = 'api/user/register/customer';
		const options = {
				method: 'POST',
				headers: {
						'Content-Type': 'application/json'
				},
				body: JSON.stringify(data),
		};
		fetch(url, options)
				.then(response => response.json())
				.then(data => {
						if (data.status === 'success') {
								window.location.href = '/customer/login';
						}
				})
				.catch(error => console.log(error));
});