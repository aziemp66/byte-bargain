const addCartButton = document.getElementById('cart-button')
const id = document.getElementById('id')
const qty = document.getElementById('qty')

addCartButton.addEventListener('click', async () => {
console.log(qty.value);
		const res = await fetch('/api/product/cart', {
				method: 'POST',
				headers: {
						'Content-Type': 'application/json'
				},
				body: JSON.stringify({ "product_id": id.value, "qty": Number(qty.value) })
		})
		const data = await res.json()
		if (data.status === 'success') {
				window.location.reload()
	}
	
	console.log(data);
})