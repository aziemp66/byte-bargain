const userID = localStorage.getItem('user_id')

if (!userID || userID == 'undefined' || userID == "" || userID == null) {
	document.querySelectorAll('.logged-in').forEach(element => {
		element.remove();
	});
}else{
	document.querySelectorAll('.logged-out').forEach(element => {
		element.remove();
	});
}