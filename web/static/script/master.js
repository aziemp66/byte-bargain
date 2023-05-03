const userID = localStorage.getItem('user')

if (!userID || userID == 'undefined' || userID == "" || userID == null) {
	document.querySelectorAll('.logged-in').forEach(element => {
		element.remove();
	});
}else{
	document.querySelectorAll('.logged-out').forEach(element => {
		element.remove();
	});
}

document.getElementById('logout').addEventListener('click', async (e) => {
	e.preventDefault();
	
	localStorage.removeItem('user');

	const response = await fetch("/api/user/logout", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		}
	})

	const responseData = await response.json();

	if (!response.ok){
		alert(responseData.message);
		return
	}

	console.log(responseData);

	window.location.href = '/login';
});